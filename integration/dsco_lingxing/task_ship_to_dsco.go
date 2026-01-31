package dsco_lingxing

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

func normalizeTrackingJoined(values []string) string {
	var tokens []string
	for _, v := range values {
		for _, t := range strings.FieldsFunc(v, func(r rune) bool {
			switch r {
			case ',', '，', ';', '；', ' ', '\t', '\n', '\r':
				return true
			default:
				return false
			}
		}) {
			if s := strings.TrimSpace(t); s != "" {
				tokens = append(tokens, s)
			}
		}
	}
	// 统一用英文逗号拼接，便于幂等与 Admin 筛选。
	return strings.Join(uniqueNonEmptyStrings(tokens), ",")
}

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
//  3. 使用领星 WmsOrderList 获取发货数据（可能多条出库单/多运单号），按 trackingNo 聚合后回传 DSCO。
//  4. shipDate：每个 trackingNo 优先取 DeliveredAt；为空则用 StockDeliveredAt 兜底（已确认）。
//  5. SKU 映射：领星 SKU 通过 mapping.sku 的反向映射得到 DSCO partnerSku（缺省同名直传）。
//  6. SID：WmsOrderList 查询需要 sid_arr；sid 从 mapping.shop 映射值解析得到（已确认）。
//  7. 本地记录：回传成功后把 trackingNo（可能多条，用英文逗号拼接）写回 shipped_tracking_no，便于幂等与筛选。
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
	var items []store.DSCOOrderSyncRow
	var err error
	if strings.TrimSpace(ctx.OnlyPONumber) != "" {
		items, err = d.orderStore.FindByStatusAndPONumber(taskCtx, 3, ctx.OnlyPONumber)
	} else {
		items, err = d.orderStore.FindByStatus(taskCtx, 3, ctx.Size)
	}
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

	// SKU 反向映射：领星 SKU -> DSCO partnerSku（mapping.sku 的反向映射；缺省同名直传）。
	reverseSKU, err := buildReverseSKUMap(ctx.Config)
	if err != nil {
		retErr = err
		return retErr
	}

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
					tracking = normalizeTrackingJoined([]string{o.Packages[0].TrackingNumber})
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

		// 4) DSCO shipMethod 口径：优先 shipMethod；没有则用 shippingServiceLevelCode（已确认）
		dscoShipMethod := getDSCOShipMethod(dscoOrder)
		if dscoShipMethod == "" {
			dscoShipMethod = getDSCOShippingServiceLevelCode(dscoOrder)
		}

		// 5) 查询领星出库单：运单号 + 发货时间 + 实发数量（可能返回多条）
		wmsOrders, _, err := lx.Warehouse.WmsOrderList(taskCtx, lingxing.WmsOrderListRequest{
			Page:               1,
			PageSize:           200,
			SIDArr:             []int{sid},
			PlatformOrderNoArr: []string{po},
		})
		if err != nil || len(wmsOrders) == 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "lingxing_wms_order_not_found",
					"err", err,
				)...,
			)
			continue
		}

		// 6) dscoItemId/lineNumber 映射：用于 shipment 行项目（尽量补齐以提升 DSCO 侧匹配成功率）
		dscoItemIDByPartner := map[string]string{}
		lineNumberByPartner := map[string]int{}
		lineNumberConflict := map[string]bool{}
		for _, li := range dscoOrder.LineItems {
			partner := ""
			if li.PartnerSKU != nil && strings.TrimSpace(*li.PartnerSKU) != "" {
				partner = strings.TrimSpace(*li.PartnerSKU)
			} else if li.SKU != nil && strings.TrimSpace(*li.SKU) != "" {
				partner = strings.TrimSpace(*li.SKU)
			}
			if partner == "" {
				continue
			}
			if li.DscoItemID != nil && strings.TrimSpace(*li.DscoItemID) != "" {
				dscoItemIDByPartner[partner] = strings.TrimSpace(*li.DscoItemID)
			}
			if li.LineNumber != nil && *li.LineNumber > 0 {
				if prev, ok := lineNumberByPartner[partner]; ok && prev != *li.LineNumber {
					lineNumberConflict[partner] = true
				}
				lineNumberByPartner[partner] = *li.LineNumber
			}
		}

		// 7) 组装 DSCO shipment：按 trackingNo 聚合（同一 PO 可能多运单号/多出库单）。
		type shipAgg struct {
			shipDate string
			items    map[string]*dsco.ShipmentLineItemForUpdate // partnerSku -> item (sum quantity)
		}
		shipAggByTracking := map[string]*shipAgg{}

		for _, w := range wmsOrders {
			tracking := strings.TrimSpace(w.TrackingNo)
			if tracking == "" {
				continue
			}
			agg := shipAggByTracking[tracking]
			if agg == nil {
				agg = &shipAgg{items: map[string]*dsco.ShipmentLineItemForUpdate{}}
				shipAggByTracking[tracking] = agg
			}

			// shipDate：优先 DeliveredAt；为空则用 StockDeliveredAt 兜底（已确认）
			if agg.shipDate == "" {
				if rawTime := strings.TrimSpace(w.DeliveredAt); rawTime != "" {
					if t, err := parseLingXingDateTimeToRFC3339UTC(rawTime); err == nil {
						agg.shipDate = t
					}
				} else if rawTime := strings.TrimSpace(w.StockDeliveredAt); rawTime != "" {
					if t, err := parseLingXingDateTimeToRFC3339UTC(rawTime); err == nil {
						agg.shipDate = t
					}
				}
			}

			for _, p := range w.ProductInfo {
				lxSKU := strings.TrimSpace(p.SKU)
				if lxSKU == "" || p.Count <= 0 {
					continue
				}
				dscoPartner := strings.TrimSpace(reverseSKU[lxSKU])
				if dscoPartner == "" {
					dscoPartner = lxSKU
				}
				if dscoPartner == "" {
					continue
				}

				if it, ok := agg.items[dscoPartner]; ok {
					it.Quantity += p.Count
					continue
				}

				li := dsco.ShipmentLineItemForUpdate{
					Quantity:   p.Count,
					PartnerSKU: dscoPartner,
				}
				if id := dscoItemIDByPartner[dscoPartner]; id != "" {
					li.DscoItemID = id
				}
				if !lineNumberConflict[dscoPartner] {
					if n := lineNumberByPartner[dscoPartner]; n > 0 {
						nn := n
						li.LineNumber = &nn
					}
				}
				agg.items[dscoPartner] = &li
			}
		}

		if len(shipAggByTracking) == 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "no_tracking_or_items_in_wms_orders",
					"dsco_raw", integration.JSONForLog(row.Payload),
					"wms_orders_raw", integration.JSONForLog(wmsOrders),
				)...,
			)
			continue
		}

		dscoWarehouseCode := getDSCOWarehouseCode(dscoOrder)
		var shipments []dsco.ShipmentForUpdate
		trackingList := make([]string, 0, len(shipAggByTracking))
		for tracking, agg := range shipAggByTracking {
			trackingList = append(trackingList, tracking)
			var lineItems []dsco.ShipmentLineItemForUpdate
			for _, it := range agg.items {
				lineItems = append(lineItems, *it)
			}
			if len(lineItems) == 0 {
				continue
			}
			shipments = append(shipments, dsco.ShipmentForUpdate{
				TrackingNumber: tracking,
				ShipDate:       agg.shipDate,
				ShipMethod:     dscoShipMethod,
				WarehouseCode:  dscoWarehouseCode,
				LineItems:      lineItems,
			})
		}
		if len(shipments) == 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "no_line_items_to_ship",
					"dsco_raw", integration.JSONForLog(row.Payload),
					"wms_orders_raw", integration.JSONForLog(wmsOrders),
				)...,
			)
			continue
		}
		trackingJoined := normalizeTrackingJoined(trackingList)

		// 8) 组装 DSCO shipment 批量请求
		ship := dsco.ShipmentsForUpdate{
			PoNumber:  po,
			Shipments: shipments,
		}
		batch = append(batch, ship)
		toUpdate = append(toUpdate, struct {
			po       string
			tracking string
		}{po: po, tracking: trackingJoined})
		logger.Info(taskCtx, "order prepared",
			append(base,
				"task", "ship_to_dsco",
				"po_number", po,
				"tracking", trackingJoined,
				"ship_request", integration.JSONForLog(ship),
				"wms_orders_raw", integration.JSONForLog(wmsOrders),
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
