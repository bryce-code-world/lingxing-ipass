package dsco_lingxing

import (
	"context"
	"encoding/json"
	"math"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

// InvoiceToDSCO 将领星侧“实际出库”的发票信息回传给 DSCO（Invoice CreateSmallBatch）。
//
// 状态机（一期口径）：
// - 本任务处理 dsco_order_sync.status = 4 的订单（已回传发货，待回传发票）。
// - 回传 invoice 成功后将状态推进到 5（完成态）。
//
// 关键点：
// - 幂等：若本地 dsco_invoice_id 已有值，则认为该 poNumber 已回传过发票，直接跳过。
// - 发票数据来源：
//   - 数量：优先使用 WmsOrderList 的 ProductInfo（以实际出库为准）。
//   - 单价：优先从 DSCO 原始订单行项目中挑选（见 pickUnitPrice），以 DSCO 订单字段为准。
//   - 订单总金额：按“数量 * 单价”汇总，并做 2 位小数 round。
//
// - invoiceDate：优先 DeliveredAt；为空则用 StockDeliveredAt 兜底（已确认）。
// - 发票回传口径（一期最新确认）：
//   - 以 poNumber 为维度：同一 poNumber 即使对应领星拆单/多行，也只回传“1 张汇总发票”给 DSCO。
//   - 前置条件：poNumber 下所有 SKU 均已发货（以 WmsOrderList 聚合数量 >= DSCO 原始订单数量 为准）。
//   - invoiceId：固定 invoiceId = poNumber。
//   - tracking：DSCO 发票接口不接受顶层 trackingNumber 字段；本任务仅在日志中记录运单号（若多个用逗号拼接），tracking 由 `ship_to_dsco` 回传。
func (d *Domain) InvoiceToDSCO(ctx integration.TaskContext) (retErr error) {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	startedAt := time.Now().UTC()
	base := ctx.BaseLogFields()
	logger.Info(taskCtx, "task begin", append(base, "task", "invoice_to_dsco")...)

	var (
		total   int
		okCount int
		skip    int
		fail    int
	)
	defer func() {
		fields := append(base,
			"task", "invoice_to_dsco",
			"duration_ms", time.Since(startedAt).Milliseconds(),
			"total", total,
			"ok", okCount,
			"skip", skip,
			"fail", fail,
		)
		if retErr != nil {
			logger.Error(taskCtx, "task end", append(fields, "result", "failed", "err", retErr)...)
			return
		}
		logger.Info(taskCtx, "task end", append(fields, "result", "ok")...)
	}()

	// 1) 取待回传发票（status=4）
	var items []store.DSCOOrderSyncRow
	var err error
	if strings.TrimSpace(ctx.OnlyPONumber) != "" {
		items, err = d.orderStore.FindByStatusAndPONumber(taskCtx, 4, ctx.OnlyPONumber)
	} else {
		items, err = d.orderStore.FindByStatus(taskCtx, 4, ctx.Size)
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

	// SKU 反向映射：领星 SKU -> DSCO partnerSku（mapping.sku 的反向映射；缺省同名直传）。
	reverseSKU, err := buildReverseSKUMap(ctx.Config)
	if err != nil {
		retErr = err
		return retErr
	}

	multiBan := false
	if jc, ok := ctx.Config.Jobs[ctx.Job]; ok {
		multiBan = jc.MultiBan
	}

	var invs []dsco.Invoice
	var toUpdate []struct {
		po        string
		invoiceID string
	}

	for _, row := range items {
		po := strings.TrimSpace(row.PONumber)
		if po == "" {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", "",
					"result", "skip",
					"reason", "po_number_empty",
				)...,
			)
			continue
		}

		// 0) 幂等：已写入 dsco_invoice_id 的订单视为已回传过发票，直接跳过。
		if strings.TrimSpace(row.DSCOInvoiceID) != "" {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "already_has_invoice_id",
					"invoice_id", row.DSCOInvoiceID,
				)...,
			)
			continue
		}

		if multiBan {
			isMulti, mi := detectMultiOrderByMSKUs([]string(row.MSKUs))
			if isMulti {
				skip++
				logger.Info(taskCtx, "order done",
					append(base,
						"task", "invoice_to_dsco",
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

		// 3) 解析 DSCO 原始订单（用于 dscoItemId/币种/参考字段等）
		var dscoOrder dsco.Order
		if err := json.Unmarshal(row.Payload, &dscoOrder); err != nil {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
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
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "missing_mapping_shop_sid",
					"dsco_raw", integration.JSONForLog(row.Payload),
				)...,
			)
			continue
		}

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
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "lingxing_wms_order_not_found",
					"err", err,
				)...,
			)
			continue
		}

		// 4) DSCO 行项目口径（以 partnerSku 为主键）：
		// - expectedQty：用于判断“是否全部已发货”
		// - dscoItemId/lineNumber：用于发票行项目（尽量补齐以提升 DSCO 侧匹配成功率）
		expectedQtyByPartner := map[string]int{}
		dscoItemIDByPartner := map[string]string{}
		dscoItemIDConflict := map[string]bool{}
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
			if li.Quantity > 0 {
				expectedQtyByPartner[partner] += li.Quantity
			}
			if li.DscoItemID != nil && strings.TrimSpace(*li.DscoItemID) != "" {
				id := strings.TrimSpace(*li.DscoItemID)
				if prev, ok := dscoItemIDByPartner[partner]; ok && prev != id {
					dscoItemIDConflict[partner] = true
				} else {
					dscoItemIDByPartner[partner] = id
				}
			}
			if li.LineNumber != nil && *li.LineNumber > 0 {
				if prev, ok := lineNumberByPartner[partner]; ok && prev != *li.LineNumber {
					lineNumberConflict[partner] = true
				}
				lineNumberByPartner[partner] = *li.LineNumber
			}
		}

		priceByPartner := map[string]float64{}
		for _, li := range dscoOrder.LineItems {
			p := ""
			if li.PartnerSKU != nil {
				p = *li.PartnerSKU
			} else if li.SKU != nil {
				p = *li.SKU
			}
			if p == "" {
				continue
			}
			if price, ok := pickUnitPrice(li); ok {
				priceByPartner[p] = price
			}
		}

		// 5) 币种：默认 USD；若 DSCO 原始订单带 currencyCode 则使用它
		currency := "USD"
		if dscoOrder.CurrencyCode != nil && *dscoOrder.CurrencyCode != "" {
			currency = *dscoOrder.CurrencyCode
		}

		// 6) 计算“实际已发货数量”（WMS 聚合）与运单号集合：
		// - shippedQty：用于判断“是否全部已发货”
		// - tracking：仅用于日志排查；DSCO 发票接口不接受顶层 trackingNumber 字段
		shippedQtyByPartner := map[string]int{}
		trackingList := make([]string, 0, len(wmsOrders))
		var latestInvoiceDate string
		for _, w := range wmsOrders {
			if t := strings.TrimSpace(w.TrackingNo); t != "" {
				trackingList = append(trackingList, t)
			}
			// invoiceDate：取“最晚一次发货时间”（优先 DeliveredAt，否则 StockDeliveredAt）
			var chosen string
			if rawTime := strings.TrimSpace(w.DeliveredAt); rawTime != "" {
				if t, err := parseLingXingDateTimeToRFC3339UTC(rawTime); err == nil {
					chosen = t
				}
			} else if rawTime := strings.TrimSpace(w.StockDeliveredAt); rawTime != "" {
				if t, err := parseLingXingDateTimeToRFC3339UTC(rawTime); err == nil {
					chosen = t
				}
			}
			if chosen != "" {
				// RFC3339 字符串同一时区下可按字典序比较，但这里仍以字符串比较做最小实现。
				if latestInvoiceDate == "" || chosen > latestInvoiceDate {
					latestInvoiceDate = chosen
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
				shippedQtyByPartner[dscoPartner] += p.Count
			}
		}
		if latestInvoiceDate == "" {
			latestInvoiceDate = time.Now().UTC().Format(time.RFC3339)
		}
		trackingJoined := strings.Join(uniqueNonEmptyStrings(trackingList), ",")

		// 7) 全量发货判断：poNumber 下所有 SKU 均已发货才允许回传发票
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
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "not_fully_shipped",
					"expected_qty", integration.JSONForLog(expectedQtyByPartner),
					"shipped_qty", integration.JSONForLog(shippedQtyByPartner),
					"wms_orders_raw", integration.JSONForLog(wmsOrders),
				)...,
			)
			continue
		}

		// 8) 组装“汇总发票”：
		// - 发票行数量以 DSCO 订单数量为准（已确认“必须全量发货”后再回传发票）
		// - 单价以 DSCO 原始订单为准（pickUnitPrice）
		// - 如缺少单价导致无法计算总额，则跳过（避免回传不完整发票）
		var missingPrice []string
		var lineItems []dsco.InvoiceLineItem
		var totalAmount float64
		for partner, expected := range expectedQtyByPartner {
			if expected <= 0 {
				continue
			}
			unit := priceByPartner[partner]
			if unit <= 0 {
				missingPrice = append(missingPrice, partner)
				continue
			}
			line := dsco.InvoiceLineItem{
				PartnerSKU: partner,
				Quantity:   expected,
				UnitPrice:  unit,
			}
			if !dscoItemIDConflict[partner] {
				if id := dscoItemIDByPartner[partner]; id != "" {
					line.DscoItemID = id
				}
			}
			if !lineNumberConflict[partner] {
				if n := lineNumberByPartner[partner]; n > 0 {
					line.LineNumber = n
				}
			}
			lineItems = append(lineItems, line)
			totalAmount += float64(expected) * unit
		}
		if len(missingPrice) > 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "missing_unit_price",
					"missing_partner_sku", strings.Join(uniqueNonEmptyStrings(missingPrice), ","),
					"dsco_raw", integration.JSONForLog(row.Payload),
				)...,
			)
			continue
		}
		if len(lineItems) == 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "no_invoice_line_items",
					"dsco_raw", integration.JSONForLog(row.Payload),
					"wms_orders_raw", integration.JSONForLog(wmsOrders),
				)...,
			)
			continue
		}

		totalAmount = math.Round(totalAmount*100) / 100
		invoiceID := po
		inv := dsco.Invoice{
			InvoiceID:           invoiceID,
			PoNumber:            po,
			ConsumerOrderNumber: derefString(dscoOrder.ConsumerOrderNumber),
			InvoiceDate:         latestInvoiceDate,
			CurrencyCode:        currency,
			TotalAmount:         totalAmount,
			LineItems:           lineItems,
		}
		invs = append(invs, inv)
		toUpdate = append(toUpdate, struct {
			po        string
			invoiceID string
		}{po: po, invoiceID: invoiceID})
		logger.Info(taskCtx, "order prepared",
			append(base,
				"task", "invoice_to_dsco",
				"po_number", po,
				"invoice_id", invoiceID,
				"tracking", trackingJoined,
				"invoice_request", integration.JSONForLog(inv),
				"wms_orders_raw", integration.JSONForLog(wmsOrders),
			)...,
		)
	}
	if len(invs) == 0 {
		return nil
	}
	// 11) 调用 DSCO invoice 批量接口
	_, err = dscoCli.Invoice.CreateSmallBatch(taskCtx, invs)
	if err != nil {
		retErr = err
		logger.Error(taskCtx, "dsco invoice createSmallBatch failed",
			append(base,
				"task", "invoice_to_dsco",
				"batch", integration.JSONForLog(invs),
				"err", err,
			)...,
		)
		return retErr
	}
	// 12) 回传成功：写回 invoiceId，并推进到 status=5
	for _, it := range toUpdate {
		if err := d.orderStore.UpdateStatusAndFields(taskCtx, it.po, 5, "", it.invoiceID); err != nil {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", it.po,
					"result", "fail",
					"reason", "update_local_status_failed_after_invoice",
					"invoice_id", it.invoiceID,
					"err", err,
				)...,
			)
			continue
		}
		okCount++
		logger.Info(taskCtx, "order done",
			append(base,
				"task", "invoice_to_dsco",
				"po_number", it.po,
				"result", "ok",
				"reason", "dsco_invoice_sent",
				"invoice_id", it.invoiceID,
				"new_status", 5,
			)...,
		)
	}
	return nil
}
