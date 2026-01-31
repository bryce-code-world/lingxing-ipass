package dsco_lingxing

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/integration"
)

// ShipToDSCO 将领星侧的发货信息回传给 DSCO（Shipment CreateSmallBatch）。
//
// 状态机（一期口径）：
// - 本任务处理 dsco_order_sync.status = 3 的订单（已 ACK，待回传发货信息）。
// - 回传 shipment 成功后将状态推进到 4（待回传发票）。
//
// 核心逻辑（你最新确认）：
//  1. 先查 DSCO 订单状态（dscoStatus）：
//     - dscoStatus == shipped：表示 DSCO 已发货，直接推进本地状态到 4（并尽量补齐 tracking）。
//     - dscoStatus == shipment_pending：表示 DSCO 等待发货回传，本任务可以继续调用 DSCO shipment 接口。
//     - 其他状态：本轮跳过（避免乱回传）。
//  2. 幂等：若本地 shipped_tracking_no 已有值，则认为已回传过 shipment，直接跳过。
//  3. 批量查询领星订单详情（trackingNo），避免逐单调用触发限流。
//  4. shipDate：优先使用 WmsOrder.DeliveredAt；若为空则用 StockDeliveredAt 兜底（已确认）。
//  5. 多出库单：WmsOrderList 返回多条时，按“取第一条”回传（已确认）。
//  6. SID：WmsOrderList 查询需要 sid_arr；sid 从 mapping.shop 映射值解析得到（已确认）。
func (d *Domain) ShipToDSCO(ctx integration.TaskContext) (retErr error) {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	startedAt := time.Now().UTC()
	base := ctx.BaseLogFields()
	logger.Info(taskCtx, "task begin", append(base, "task", "ship_to_dsco")...)

	var (
		total    int
		okCount  int
		skip     int
		fail     int
		advanced int
	)
	defer func() {
		fields := append(base,
			"task", "ship_to_dsco",
			"duration_ms", time.Since(startedAt).Milliseconds(),
			"total", total,
			"ok", okCount,
			"skip", skip,
			"fail", fail,
			"advanced", advanced,
		)
		if retErr != nil {
			logger.Error(taskCtx, "task end", append(fields, "result", "failed", "err", retErr)...)
			return
		}
		logger.Info(taskCtx, "task end", append(fields, "result", "ok")...)
	}()

	// 1) 取待回传发货（status=3）
	items, err := d.orderStore.FindByStatus(taskCtx, 3, ctx.Size)
	if err != nil {
		retErr = err
		return retErr
	}
	if len(items) == 0 {
		return nil
	}
	total = len(items)

	// 2) 初始化客户端：DSCO + 领星
	dscoCli, err := d.dscoClient()
	if err != nil {
		retErr = err
		return retErr
	}
	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		retErr = err
		return retErr
	}

	// 3) 先查 DSCO 状态（避免重复回传 shipment）：
	poNumbers := make([]string, 0, len(items))
	for _, it := range items {
		if strings.TrimSpace(it.PONumber) == "" || it.ShippedTrackingNo != "" {
			continue
		}
		poNumbers = append(poNumbers, it.PONumber)
	}
	dscoByPO := fetchDSCOOrdersByPONumbers(taskCtx, dscoCli, poNumbers, 5)
	logger.Info(taskCtx, "dsco orders fetched",
		append(base,
			"task", "ship_to_dsco",
			"po_count", len(uniqueNonEmptyStrings(poNumbers)),
			"fetched", len(dscoByPO),
		)...,
	)

	// 4) 批量查询领星订单详情（主要为了 trackingNo），避免逐单 GetOrderDetailV2 触发限流。
	poNumbers = uniqueNonEmptyStrings(poNumbers)
	detailsByPO := make(map[string]lingxing.OrderDetailV2, len(poNumbers))
	includeDelete := true
	const maxBatch = 50
	for _, chunk := range chunkStrings(poNumbers, maxBatch) {
		out, err := lx.Order.ListOrdersV2(taskCtx, lingxing.OrderListV2Request{
			Offset:           0,
			Length:           len(chunk),
			PlatformOrderNos: chunk,
			IncludeDelete:    &includeDelete,
		})
		if err != nil {
			logger.Warn(taskCtx, "lingxing list orders failed, fallback to per-order check", "err", err)
			for _, po := range chunk {
				detail, derr := lx.Order.GetOrderDetailV2(taskCtx, lingxing.OrderDetailV2Request{PlatformOrderNo: po})
				if derr == nil {
					detailsByPO[po] = detail
				}
			}
			continue
		}
		for _, detail := range out.List {
			if po := poNumberFromLingXingOrderDetail(detail); po != "" {
				detailsByPO[po] = detail
			}
		}
	}
	logger.Info(taskCtx, "lingxing orders fetched",
		append(base,
			"task", "ship_to_dsco",
			"po_count", len(poNumbers),
			"fetched", len(detailsByPO),
		)...,
	)

	var batch []dsco.ShipmentsForUpdate
	var toUpdate []struct {
		po       string
		tracking string
	}

	for _, row := range items {
		po := strings.TrimSpace(row.PONumber)
		if po == "" {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", "",
					"result", "skip",
					"reason", "po_number_empty",
				)...,
			)
			continue
		}
		// 2) 幂等：已写入 shipped_tracking_no 的订单视为已回传过 shipment，直接跳过。
		if row.ShippedTrackingNo != "" {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "already_has_shipped_tracking_no",
					"tracking", row.ShippedTrackingNo,
				)...,
			)
			continue
		}

		// 1) DSCO 状态判断：shipped -> 直接推进；shipment_pending -> 允许回传。
		if o, ok := dscoByPO[po]; ok {
			switch strings.TrimSpace(o.DscoStatus) {
			case "shipped", "completed":
				tracking := ""
				if len(o.Packages) > 0 {
					tracking = strings.TrimSpace(o.Packages[0].TrackingNumber)
				}
				if uerr := d.orderStore.UpdateStatusAndFields(taskCtx, po, 4, tracking, ""); uerr != nil {
					fail++
					logger.Warn(taskCtx, "order done",
						append(base,
							"task", "ship_to_dsco",
							"po_number", po,
							"result", "fail",
							"reason", "update_local_status_failed_for_dsco_already_shipped",
							"dsco_status", strings.TrimSpace(o.DscoStatus),
							"dsco_raw", integration.JSONForLog(o),
							"tracking", tracking,
							"err", uerr,
						)...,
					)
					continue
				}
				skip++
				advanced++
				logger.Info(taskCtx, "order done",
					append(base,
						"task", "ship_to_dsco",
						"po_number", po,
						"result", "skip",
						"reason", "dsco_already_shipped_or_completed",
						"dsco_status", strings.TrimSpace(o.DscoStatus),
						"dsco_raw", integration.JSONForLog(o),
						"tracking", tracking,
						"new_status", 4,
					)...,
				)
				continue
			case "shipment_pending":
				// ok: continue
			default:
				// 其他状态不做回传，避免乱推进
				skip++
				logger.Info(taskCtx, "order done",
					append(base,
						"task", "ship_to_dsco",
						"po_number", po,
						"result", "skip",
						"reason", "dsco_status_not_allowed_for_ship",
						"dsco_status", strings.TrimSpace(o.DscoStatus),
						"dsco_raw", integration.JSONForLog(o),
					)...,
				)
				continue
			}
		} else {
			// 获取不到 DSCO 订单对象时，无法使用 dscoStatus 做幂等/状态门控。
			// 一期策略：仍然基于本地 payload + 领星出库信息继续组装回传，避免因单次查询失败导致卡住。
			logger.Warn(taskCtx, "dsco order not fetched, continue by local payload",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
				)...,
			)
		}

		// 3) 解析 DSCO 原始订单（用于 dscoItemId/shipMethod 等字段拼装）
		var dscoOrder dsco.Order
		if err := json.Unmarshal(row.Payload, &dscoOrder); err != nil {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
					"result", "fail",
					"reason", "decode_dsco_payload_failed",
					"dsco_raw", integration.JSONForLog(row.Payload),
					"err", err,
				)...,
			)
			continue
		}
		// 3.1) sid：用于 WmsOrderList 精准查询（从 mapping.shop 解析）
		sid, ok := lingxingSIDFromMapping(ctx.Config, dscoOrder)
		if !ok {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "missing_mapping_shop_sid",
					"dsco_raw", integration.JSONForLog(row.Payload),
				)...,
			)
			continue
		}

		// 4.1) 从批量查询结果中提取 trackingNo（无运单号则暂不回传）
		detail, ok := detailsByPO[po]
		if !ok {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "lingxing_order_not_fetched",
				)...,
			)
			continue
		}
		tracking := strings.TrimSpace(detail.LogisticsInfo.TrackingNo)
		if tracking == "" {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "lingxing_tracking_empty",
					"lingxing_status", int(detail.Status),
					"lingxing_raw", integration.JSONForLog(detail),
				)...,
			)
			continue
		}

		// 5) DSCO shipMethod 口径：优先 shipMethod；没有则用 shippingServiceLevelCode（已确认）
		dscoShipMethod := getDSCOShipMethod(dscoOrder)
		if dscoShipMethod == "" {
			dscoShipMethod = getDSCOShippingServiceLevelCode(dscoOrder)
		}

		// 6) shipDate 口径：优先 DeliveredAt；为空则用 StockDeliveredAt 兜底（已确认）
		shipDateRFC3339 := ""
		orders, _, err := lx.Warehouse.WmsOrderList(taskCtx, lingxing.WmsOrderListRequest{
			Page:               1,
			PageSize:           20,
			SIDArr:             []int{sid},
			PlatformOrderNoArr: []string{po},
		})
		if err == nil && len(orders) > 0 {
			rawTime := orders[0].DeliveredAt
			if rawTime != "" {
				if t, err := parseLingXingDateTimeToRFC3339UTC(rawTime); err == nil {
					shipDateRFC3339 = t
				}
			} else if rawTime := orders[0].StockDeliveredAt; rawTime != "" {
				if t, err := parseLingXingDateTimeToRFC3339UTC(rawTime); err == nil {
					shipDateRFC3339 = t
				}
			}
		}

		// 7) dscoItemId 映射：用于 shipment 行项目（尽量补齐 dscoItemId）
		dscoItemIDByPartner := map[string]string{}
		for _, li := range dscoOrder.LineItems {
			p := ""
			if li.PartnerSKU != nil {
				p = *li.PartnerSKU
			}
			if p != "" && li.DscoItemID != nil {
				dscoItemIDByPartner[p] = *li.DscoItemID
			}
		}

		lineItems := []dsco.ShipmentLineItemForUpdate{}

		// 8) 行项目数量口径：
		// - 优先使用 WmsOrderList 的 ProductInfo（更接近真实出库数量）。
		// - 若查不到，则退回使用领星订单详情 ItemInfo 数量。
		wmsOrders, _, err := lx.Warehouse.WmsOrderList(taskCtx, lingxing.WmsOrderListRequest{
			Page:               1,
			PageSize:           20,
			SIDArr:             []int{sid},
			PlatformOrderNoArr: []string{po},
		})
		if err == nil && len(wmsOrders) > 0 && len(wmsOrders[0].ProductInfo) > 0 {
			for _, p := range wmsOrders[0].ProductInfo {
				dscoPartner := p.SKU
				li := dsco.ShipmentLineItemForUpdate{
					Quantity:   p.Count,
					PartnerSKU: dscoPartner,
				}
				if id := dscoItemIDByPartner[dscoPartner]; id != "" {
					li.DscoItemID = id
				}
				lineItems = append(lineItems, li)
			}
		} else {
			for _, it := range detail.ItemInfo {
				dscoPartner := it.MSKU
				li := dsco.ShipmentLineItemForUpdate{
					Quantity:   it.Quantity,
					PartnerSKU: dscoPartner,
				}
				if id := dscoItemIDByPartner[dscoPartner]; id != "" {
					li.DscoItemID = id
				}
				lineItems = append(lineItems, li)
			}
		}
		if len(lineItems) == 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "no_line_items_to_ship",
					"dsco_raw", integration.JSONForLog(row.Payload),
					"lingxing_raw", integration.JSONForLog(detail),
				)...,
			)
			continue
		}

		// 9) 组装 DSCO shipment 批量请求
		ship := dsco.ShipmentsForUpdate{
			PoNumber: po,
			Shipments: []dsco.ShipmentForUpdate{
				{
					TrackingNumber: tracking,
					ShipDate:       shipDateRFC3339,
					ShipMethod:     dscoShipMethod,
					LineItems:      lineItems,
				},
			},
		}
		batch = append(batch, ship)
		toUpdate = append(toUpdate, struct {
			po       string
			tracking string
		}{po: po, tracking: tracking})

		wmsRaw := ""
		if len(wmsOrders) > 0 {
			wmsRaw = integration.JSONForLog(wmsOrders[0])
		}
		logger.Info(taskCtx, "order prepared",
			append(base,
				"task", "ship_to_dsco",
				"po_number", po,
				"tracking", tracking,
				"ship_request", integration.JSONForLog(ship),
				"lingxing_raw", integration.JSONForLog(detail),
				"wms_order_raw", wmsRaw,
			)...,
		)
	}

	if len(batch) == 0 {
		return nil
	}

	// 10) 调用 DSCO shipment 批量接口
	if _, err := dscoCli.Shipment.CreateSmallBatch(taskCtx, batch); err != nil {
		retErr = err
		logger.Error(taskCtx, "dsco shipment createSmallBatch failed",
			append(base,
				"task", "ship_to_dsco",
				"batch", integration.JSONForLog(batch),
				"err", err,
			)...,
		)
		return retErr
	}

	// 11) 回传成功：写回 tracking，并推进到 status=4
	for _, it := range toUpdate {
		if err := d.orderStore.UpdateStatusAndFields(taskCtx, it.po, 4, it.tracking, ""); err != nil {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", it.po,
					"result", "fail",
					"reason", "update_local_status_failed_after_ship",
					"tracking", it.tracking,
					"err", err,
				)...,
			)
			continue
		}
		okCount++
		logger.Info(taskCtx, "order done",
			append(base,
				"task", "ship_to_dsco",
				"po_number", it.po,
				"result", "ok",
				"reason", "dsco_shipment_sent",
				"tracking", it.tracking,
				"new_status", 4,
			)...,
		)
	}
	return nil
}
