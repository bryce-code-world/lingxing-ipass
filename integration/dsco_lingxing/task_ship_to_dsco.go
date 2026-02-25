package dsco_lingxing

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"lingxingipass/golib/v2/sdk/dsco"
	"lingxingipass/golib/v2/sdk/lingxing"
	"lingxingipass/golib/v2/tool/logger"

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

func parseTrackingSet(raw string) map[string]struct{} {
	normalized := normalizeTrackingJoined([]string{raw})
	if normalized == "" {
		return map[string]struct{}{}
	}
	out := make(map[string]struct{}, 8)
	for _, t := range strings.Split(normalized, ",") {
		if s := strings.TrimSpace(t); s != "" {
			out[s] = struct{}{}
		}
	}
	return out
}

func joinTrackingSet(set map[string]struct{}) string {
	if len(set) == 0 {
		return ""
	}
	items := make([]string, 0, len(set))
	for k := range set {
		if s := strings.TrimSpace(k); s != "" {
			items = append(items, s)
		}
	}
	return normalizeTrackingJoined(items)
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
//  2. 幂等：不以本地 shipped_tracking_no 判断是否回传；是否跳过以 DSCO 实际 dscoStatus 为准。
//  3. 使用领星 WmsOrderList 获取发货数据（可能多条出库单/多运单号），按 trackingNo 聚合后回传 DSCO。
//  4. shipDate：每个 trackingNo 优先取 DeliveredAt；为空则用 StockDeliveredAt 兜底（已确认）。
//  5. SKU 映射：领星 SKU 通过 mapping.sku 的反向映射得到 DSCO sku（缺省同名直传；必要时兜底 partnerSku）。
//  6. SID：WmsOrderList 查询需要 sid_arr；sid 从 mapping.shop 映射值解析得到（已确认）。
//  7. 本地记录：回传成功后把 trackingNo（可能多条，用英文逗号拼接）写回 shipped_tracking_no，便于展示与筛选。
//  8. 若同一 partnerSku 跨多个 tracking，需对该行设置 packageSpanFlag=true（否则部分零售商会拒绝 partial shipment）。
//  9. CreateSmallBatch 为异步：仅在 Order Change Log 确认 success 后，才推进本地 status / 写回 tracking（避免“HTTP 成功但最终失败”）。
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

	// 1) 取待回传发货（status=3）。
	// 手动指定 po_number 时，允许取出该单据并用于补偿回传（即使其 status 已被错误推进到 4）。
	var items []store.DSCOOrderSyncRow
	var err error
	if strings.TrimSpace(ctx.OnlyPONumber) != "" {
		row, ok, gerr := d.orderStore.GetByPONumber(taskCtx, ctx.OnlyPONumber)
		if gerr != nil {
			retErr = gerr
			return retErr
		}
		if !ok {
			return nil
		}
		if row.Status < 3 {
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", strings.TrimSpace(ctx.OnlyPONumber),
					"result", "skip",
					"reason", "status_not_allowed_for_ship",
					"status", row.Status,
				)...,
			)
			return nil
		}
		items = []store.DSCOOrderSyncRow{row}
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
		if strings.TrimSpace(it.PONumber) == "" {
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

	// SKU 反向映射：领星 SKU -> DSCO sku（mapping.sku 的反向映射；缺省同名直传）。
	reverseSKU, err := buildReverseSKUMap(ctx.Config)
	if err != nil {
		retErr = err
		return retErr
	}

	multiBan := false
	if jc, ok := ctx.Config.Jobs[ctx.Job]; ok {
		multiBan = jc.MultiBan
	}
	// 手动单号测试（OnlyPONumber 非空）允许通过：此时是人工验证单个订单，不做 multi_ban 限制。
	if strings.TrimSpace(ctx.OnlyPONumber) != "" {
		multiBan = false
	}

	var batch []dsco.ShipmentsForUpdate
	var toUpdate []struct {
		po       string
		tracking string
		status   int16
	}

	for _, row := range items {
		po := strings.TrimSpace(row.PONumber)
		existingTracking := normalizeTrackingJoined([]string{row.ShippedTrackingNo})
		existingTrackingSet := parseTrackingSet(existingTracking)
		// DSCO 侧已记录的 tracking（以 DSCO packages 为准，用于幂等与补偿）。
		// 注意：Shipment CreateSmallBatch 为异步接口，仅 HTTP 成功不代表每条都成功；
		// 若仅凭本地 shipped_tracking_no 判断“已回传”，可能导致失败单据被错误跳过。
		dscoTrackingSet := map[string]struct{}{}
		hasDSCOOrder := false
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
		// 1) DSCO 状态判断：shipped -> 直接推进；shipment_pending -> 允许回传。
		if o, ok := dscoByPO[po]; ok {
			hasDSCOOrder = true
			var pkgTrackings []string
			for _, p := range o.Packages {
				if s := strings.TrimSpace(p.TrackingNumber); s != "" {
					pkgTrackings = append(pkgTrackings, s)
				}
			}
			trackingFromDSCO := normalizeTrackingJoined(pkgTrackings)
			dscoTrackingSet = parseTrackingSet(trackingFromDSCO)
			// 用 DSCO packages 补齐本地 tracking 展示字段（DSCO 为准）。
			if trackingFromDSCO != "" {
				for t := range dscoTrackingSet {
					existingTrackingSet[t] = struct{}{}
				}
				existingTracking = joinTrackingSet(existingTrackingSet)
			}
			switch strings.TrimSpace(o.DscoStatus) {
			case "shipped", "completed":
				// 重要：DSCO 侧订单状态变为 shipped 并不一定代表“所有行都可开票/可回传发票”。
				// 例如：同一 PO 在领星拆单，先回传了一部分 shipment，DSCO 仍可能将订单标为 shipped；
				// 若此处直接跳过，会导致缺失的 tracking/行项目无法补回，后续 invoice 会因为 available-to-invoice=0 失败。
				logger.Warn(taskCtx, "order note",
					append(base,
						"task", "ship_to_dsco",
						"po_number", po,
						"reason", "dsco_status_shipped_but_continue",
						"dsco_status", strings.TrimSpace(o.DscoStatus),
						"tracking_from_dsco", trackingFromDSCO,
						"tracking_before", strings.TrimSpace(row.ShippedTrackingNo),
					)...,
				)
				// continue: still try to补回缺失 shipment
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

		if multiBan {
			isMulti, mi := detectMultiOrderByMSKUs([]string(row.MSKUs))
			if isMulti {
				skip++
				logger.Info(taskCtx, "order done",
					append(base,
						"task", "ship_to_dsco",
						"po_number", po,
						"result", "skip",
						"reason", "multi_banned",
						"multi_ban", true,
						"mskus", []string(row.MSKUs),
						"multi_info", integration.JSONForLog(mi),
						"dsco_raw", integration.JSONForLog(row.Payload),
					)...,
				)
				continue
			}
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

		// 4) DSCO 物流信息 口径
		dscoShipMethod := getDSCOShipMethod(dscoOrder)
		dscoShipCarrier := getDSCOShipCarrier(dscoOrder)
		dscoShipServiceLevelCode := getDSCOShippingServiceLevelCode(dscoOrder)

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

		// 6) 订单行映射：统一按 DSCO sku 主键（sku 缺失时兜底 partnerSku）
		// - dscoItemId/lineNumber：用于 shipment 行项目增强匹配
		// - sku/partnerSku：用于回传时双字段兼容
		dscoItemIDByKey := map[string]string{}
		lineNumberByKey := map[string]int{}
		lineNumberConflict := map[string]bool{}
		skuByKey := map[string]string{}
		partnerByKey := map[string]string{}
		aliasToKey := map[string]string{}
		for _, li := range dscoOrder.LineItems {
			key := dscoLineKey(li)
			if key == "" {
				continue
			}
			sku := dscoLineSKU(li)
			partner := dscoLinePartnerSKU(li)
			if sku != "" {
				skuByKey[key] = sku
				aliasToKey[sku] = key
			}
			if partner != "" {
				partnerByKey[key] = partner
				aliasToKey[partner] = key
			}
			aliasToKey[key] = key
			if li.DscoItemID != nil && strings.TrimSpace(*li.DscoItemID) != "" {
				dscoItemIDByKey[key] = strings.TrimSpace(*li.DscoItemID)
			}
			if li.LineNumber != nil && *li.LineNumber > 0 {
				if prev, ok := lineNumberByKey[key]; ok && prev != *li.LineNumber {
					lineNumberConflict[key] = true
				}
				lineNumberByKey[key] = *li.LineNumber
			}
		}

		// 7) 组装 DSCO shipment：按 trackingNo 聚合（同一 PO 可能多运单号/多出库单）。
		type shipAgg struct {
			shipDate string
			items    map[string]*dsco.ShipmentLineItemForUpdate // dscoKey -> item (sum quantity)
		}
		shipAggByTracking := map[string]*shipAgg{}

		for _, w := range wmsOrders {
			// 只在领星 WmsOrder.status=3（已发货）时才回传物流信息；否则发货时间往往为空值，且不应回传。
			if w.Status != 3 {
				continue
			}
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
				dscoKey := strings.TrimSpace(reverseSKU[lxSKU])
				if dscoKey == "" {
					dscoKey = lxSKU
				}
				if alias, ok := aliasToKey[dscoKey]; ok && alias != "" {
					dscoKey = alias
				}
				if dscoKey == "" {
					continue
				}

				if it, ok := agg.items[dscoKey]; ok {
					it.Quantity += p.Count
					continue
				}

				li := dsco.ShipmentLineItemForUpdate{
					Quantity: p.Count,
				}
				if sku := strings.TrimSpace(skuByKey[dscoKey]); sku != "" {
					li.SKU = sku
				} else {
					li.SKU = dscoKey
				}
				if partner := strings.TrimSpace(partnerByKey[dscoKey]); partner != "" {
					li.PartnerSKU = partner
				}
				if id := dscoItemIDByKey[dscoKey]; id != "" {
					li.DscoItemID = id
				}
				if !lineNumberConflict[dscoKey] {
					if n := lineNumberByKey[dscoKey]; n > 0 {
						nn := n
						li.LineNumber = &nn
					}
				}
				agg.items[dscoKey] = &li
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

		// 7.1) “是否已全量发货”判断（以 DSCO 订单行数量为准）：
		// - 只要有任一 dscoKey（sku 主口径）的已发货数量 < DSCO 订单行期望数量，则认为未全量发货。
		// - 未全量发货时一律跳过 shipment 回传，等待下次任务重试。
		expectedQtyByPartner := map[string]int{}
		for _, li := range dscoOrder.LineItems {
			key := dscoLineKey(li)
			if key == "" || li.Quantity <= 0 {
				continue
			}
			expectedQtyByPartner[key] += li.Quantity
		}
		shippedQtyByPartner := map[string]int{}
		for _, agg := range shipAggByTracking {
			for partner, it := range agg.items {
				if it == nil || it.Quantity <= 0 {
					continue
				}
				shippedQtyByPartner[partner] += it.Quantity
			}
		}
		notReady := false
		for partner, expected := range expectedQtyByPartner {
			if expected <= 0 {
				continue
			}
			if shippedQtyByPartner[partner] < expected {
				notReady = true
				break
			}
		}
		if notReady {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "not_fully_shipped_wait_all",
					"expected_qty", integration.JSONForLog(expectedQtyByPartner),
					"shipped_qty", integration.JSONForLog(shippedQtyByPartner),
					"tracking_sent", existingTracking,
					"wms_orders_raw", integration.JSONForLog(wmsOrders),
				)...,
			)
			continue
		}

		// 7.2) 增量回传：优先以 DSCO 侧 packages 做幂等判断；仅在无法获取 DSCO 订单时，才回退到本地 shipped_tracking_no。
		alreadyInDSCOSet := existingTrackingSet
		if hasDSCOOrder {
			alreadyInDSCOSet = dscoTrackingSet
		}
		shipAggToSend := map[string]*shipAgg{}
		for tracking, agg := range shipAggByTracking {
			if _, ok := alreadyInDSCOSet[tracking]; ok {
				continue
			}
			shipAggToSend[tracking] = agg
		}
		if len(shipAggToSend) == 0 {
			// 已全量发货，但没有新增 tracking 需要回传：说明之前已经回传过，推进到 status=4 等待回传发票。
			if existingTracking == "" {
				skip++
				logger.Info(taskCtx, "order done",
					append(base,
						"task", "ship_to_dsco",
						"po_number", po,
						"result", "skip",
						"reason", "fully_shipped_but_no_tracking_sent",
						"expected_qty", integration.JSONForLog(expectedQtyByPartner),
						"shipped_qty", integration.JSONForLog(shippedQtyByPartner),
						"wms_orders_raw", integration.JSONForLog(wmsOrders),
					)...,
				)
				continue
			}
			if uerr := d.orderStore.UpdateStatus(taskCtx, po, 4); uerr != nil {
				fail++
				logger.Warn(taskCtx, "order done",
					append(base,
						"task", "ship_to_dsco",
						"po_number", po,
						"result", "fail",
						"reason", "update_local_status_failed_for_fully_shipped_no_new_tracking",
						"tracking", existingTracking,
						"err", uerr,
					)...,
				)
				continue
			}
			okCount++
			advanced++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", po,
					"result", "ok",
					"reason", "fully_shipped_no_new_tracking_advance",
					"tracking", existingTracking,
					"new_status", 4,
				)...,
			)
			continue
		}

		// 7.3) DSCO 对“跨多个 tracking 的同一行商品”需要显式标记 packageSpanFlag，否则会被判定为 partial ship 并拒绝。
		spanPartner := map[string]bool{}
		partnerToTracking := make(map[string]map[string]struct{}, 8)
		for tracking, agg := range shipAggByTracking {
			for partner := range agg.items {
				p := strings.TrimSpace(partner)
				if p == "" {
					continue
				}
				set := partnerToTracking[p]
				if set == nil {
					set = map[string]struct{}{}
					partnerToTracking[p] = set
				}
				set[tracking] = struct{}{}
			}
		}
		for partner, set := range partnerToTracking {
			if len(set) > 1 {
				spanPartner[partner] = true
			}
		}

		dscoWarehouseCode := getDSCOWarehouseCode(dscoOrder)
		var shipments []dsco.ShipmentForUpdate
		trackingList := make([]string, 0, len(shipAggToSend))
		for tracking, agg := range shipAggToSend {
			trackingList = append(trackingList, tracking)
			var lineItems []dsco.ShipmentLineItemForUpdate
			for partner, it := range agg.items {
				if it == nil {
					continue
				}
				li := dsco.ShipmentLineItemForUpdate{
					Quantity:   it.Quantity,
					LineNumber: it.LineNumber,
					DscoItemID: it.DscoItemID,
					SKU:        it.SKU,
					PartnerSKU: it.PartnerSKU,
					UPC:        it.UPC,
					EAN:        it.EAN,
					GTIN:       it.GTIN,
					ISBN:       it.ISBN,
					MPN:        it.MPN,
				}
				if spanPartner[strings.TrimSpace(partner)] {
					v := true
					li.PackageSpanFlag = &v
				}
				lineItems = append(lineItems, li)
			}
			if len(lineItems) == 0 {
				continue
			}
			shipments = append(shipments, dsco.ShipmentForUpdate{
				TrackingNumber:           tracking,
				ShipDate:                 agg.shipDate,
				ShipCarrier:              dscoShipCarrier,
				ShipMethod:               dscoShipMethod,
				ShippingServiceLevelCode: dscoShipServiceLevelCode,
				WarehouseCode:            dscoWarehouseCode,
				LineItems:                lineItems,
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
		for _, t := range trackingList {
			if s := strings.TrimSpace(t); s != "" {
				existingTrackingSet[s] = struct{}{}
			}
		}
		nextStatus := int16(4)
		toUpdate = append(toUpdate, struct {
			po       string
			tracking string
			status   int16
		}{po: po, tracking: joinTrackingSet(existingTrackingSet), status: nextStatus})
		logger.Info(taskCtx, "order prepared",
			append(base,
				"task", "ship_to_dsco",
				"po_number", po,
				"tracking", trackingJoined,
				"tracking_sent_before", existingTracking,
				"not_fully_shipped", notReady,
				"expected_qty", integration.JSONForLog(expectedQtyByPartner),
				"shipped_qty", integration.JSONForLog(shippedQtyByPartner),
				"ship_request", integration.JSONForLog(ship),
				"wms_orders_raw", integration.JSONForLog(wmsOrders),
			)...,
		)
	}

	if len(batch) == 0 {
		return nil
	}

	// 10) 调用 DSCO shipment 批量接口（CreateSmallBatch 为异步接口）
	resp, raw, err := dscoCli.Shipment.CreateSmallBatchWithRawBody(taskCtx, batch)
	if err != nil {
		retErr = err
		logger.Error(taskCtx, "dsco shipment createSmallBatch failed",
			append(base,
				"task", "ship_to_dsco",
				"batch", integration.JSONForLog(batch),
				"resp_raw", raw,
				"err", err,
			)...,
		)
		return retErr
	}

	logger.Info(taskCtx, "dsco shipment createSmallBatch ok",
		append(base,
			"task", "ship_to_dsco",
			"resp", integration.JSONForLog(resp),
			"resp_raw", raw,
		)...,
	)

	// 11) 校验 DSCO 侧是否真正成功（CreateSmallBatch 为异步接口；仅 HTTP 成功不代表每条都成功）。
	verifyByPO := make(map[string]bool, len(toUpdate))
	verifyMsgByPO := make(map[string]any, len(toUpdate))
	want := make(map[string]struct{}, len(toUpdate))
	for _, it := range toUpdate {
		if s := strings.TrimSpace(it.po); s != "" {
			want[s] = struct{}{}
		}
	}

	completed := false
	requestID := ""
	if resp != nil {
		requestID = strings.TrimSpace(resp.RequestID)
	}
	if requestID != "" && len(want) > 0 {
		m, msg, ok := pollDSCOShipmentChangeLog(taskCtx, dscoCli, requestID, want, 60*time.Second, 3*time.Second)
		for k, v := range m {
			verifyByPO[k] = v
			if mm, ok := msg[k]; ok {
				verifyMsgByPO[k] = mm
			}
		}
		completed = ok
	}

	var confirmedOK int
	var confirmedFail int
	var confirmedPending int
	for po := range want {
		v, ok := verifyByPO[po]
		if !ok {
			confirmedPending++
			continue
		}
		if v {
			confirmedOK++
		} else {
			confirmedFail++
		}
	}
	logger.Info(taskCtx, "dsco shipment verify by requestId",
		append(base,
			"task", "ship_to_dsco",
			"request_id", requestID,
			"completed", completed,
			"confirmed_ok", confirmedOK,
			"confirmed_fail", confirmedFail,
			"confirmed_pending", confirmedPending,
		)...,
	)

	// 12) 仅在确认 DSCO success 后，才写回 tracking/status（避免「异步失败但本地推进」）。
	for _, it := range toUpdate {
		po := strings.TrimSpace(it.po)
		ok, has := verifyByPO[po]
		if !has {
			skip++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", it.po,
					"result", "skip",
					"reason", "dsco_change_log_pending_or_missing",
					"tracking", it.tracking,
					"new_status", it.status,
					"request_id", requestID,
				)...,
			)
			continue
		}
		if !ok {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", it.po,
					"result", "fail",
					"reason", "dsco_shipment_failed",
					"tracking", it.tracking,
					"new_status", it.status,
					"request_id", requestID,
					"dsco_change_log", integration.JSONForLog(verifyMsgByPO[po]),
				)...,
			)
			continue
		}
		if err := d.orderStore.UpdateStatusAndTrackingNo(taskCtx, it.po, it.status, it.tracking); err != nil {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "ship_to_dsco",
					"po_number", it.po,
					"result", "fail",
					"reason", "update_local_status_failed_after_ship",
					"tracking", it.tracking,
					"new_status", it.status,
					"err", err,
				)...,
			)
			continue
		}
		okCount++
		if it.status == 4 {
			advanced++
		}
		logger.Info(taskCtx, "order done",
			append(base,
				"task", "ship_to_dsco",
				"po_number", it.po,
				"result", "ok",
				"reason", "dsco_shipment_sent",
				"tracking", it.tracking,
				"new_status", it.status,
			)...,
		)
	}
	return nil
}
