package dsco_lingxing

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/integration"
)

// InvoiceToDSCO 将领星侧“实际出库”的发票信息回传给 DSCO（Invoice CreateSmallBatch）。
//
// 状态机（一期口径）：
// - 本任务处理 dsco_order_sync.status = 4 的订单（已回传发货，待回传发票）。
// - 回传 invoice 成功后将状态推进到 5（完成态）。
//
// 关键点：
// - 幂等：若本地 dsco_invoice_id 已有值，则认为已回传过 invoice，直接跳过。
// - 发票数据来源：
//   - 数量：优先使用 WmsOrderList 的 ProductInfo（以实际出库为准）。
//   - 单价：优先从 DSCO 原始订单行项目中挑选（见 pickUnitPrice），以 DSCO 订单字段为准。
//   - 订单总金额：按“数量 * 单价”汇总，并做 2 位小数 round。
//
// - invoiceDate：优先 DeliveredAt；为空则用 StockDeliveredAt 兜底（已确认）。
// - invoiceId：一期默认 invoiceId = poNumber；若未来一单多出库单，可切换为 poNumber-序号（已确认）。
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
	items, err := d.orderStore.FindByStatus(taskCtx, 4, ctx.Size)
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
		// 0) 幂等：已写入 dsco_invoice_id 的订单视为已回传过 invoice，直接跳过。
		if row.DSCOInvoiceID != "" {
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
			PageSize:           20,
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

		// 4) invoiceDate 口径：优先 DeliveredAt；为空则用 StockDeliveredAt 兜底（已确认）
		shipDate := time.Now().UTC().Format(time.RFC3339)
		if rawTime := wmsOrders[0].DeliveredAt; rawTime != "" {
			if t, err := parseLingXingDateTimeToRFC3339UTC(rawTime); err == nil {
				shipDate = t
			}
		} else if rawTime := wmsOrders[0].StockDeliveredAt; rawTime != "" {
			if t, err := parseLingXingDateTimeToRFC3339UTC(rawTime); err == nil {
				shipDate = t
			}
		}

		// 5) tracking：发票回传也带 trackingNumber（从 WMS 订单取）
		tracking := wmsOrders[0].TrackingNo

		// 6) dscoItemId 映射：用于发票行项目（尽量补齐 dscoItemId）
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

		// 7) 生成发票行项目与总金额：
		// - 数量以 WMS 出库为准
		// - 单价以 DSCO 原始订单为准（pickUnitPrice）
		var lineItems []dsco.InvoiceLineItem
		var total float64
		for _, p := range wmsOrders[0].ProductInfo {
			dscoPartner := p.SKU
			unit := priceByPartner[dscoPartner]
			if unit <= 0 {
				continue
			}
			line := dsco.InvoiceLineItem{
				PartnerSKU: dscoPartner,
				Quantity:   p.Count,
				UnitPrice:  unit,
			}
			if id := dscoItemIDByPartner[dscoPartner]; id != "" {
				line.DscoItemID = id
			}
			lineItems = append(lineItems, line)
			total += float64(p.Count) * unit
		}
		if len(lineItems) == 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "no_line_items_to_invoice",
					"wms_order_raw", integration.JSONForLog(wmsOrders[0]),
					"dsco_raw", integration.JSONForLog(row.Payload),
				)...,
			)
			continue
		}

		// 8) 币种：默认 USD；若 DSCO 原始订单带 currencyCode 则使用它
		currency := "USD"
		if dscoOrder.CurrencyCode != nil && *dscoOrder.CurrencyCode != "" {
			currency = *dscoOrder.CurrencyCode
		}

		// Round total to 2 decimals to be safe.
		total = math.Round(total*100) / 100

		// 9) invoiceId：一期默认 invoiceId = poNumber；多出库单时可切换为 poNumber-序号（MVP 先写 1）
		invoiceID := po
		if len(wmsOrders) > 1 {
			// Future-proof: when a PO has multiple WMS orders, use a stable 1-based sequence.
			invoiceID = fmt.Sprintf("%s-%d", po, 1)
		}
		// 10) 组装 DSCO 发票对象
		inv := dsco.Invoice{
			InvoiceID:           invoiceID,
			PoNumber:            po,
			ConsumerOrderNumber: derefString(dscoOrder.ConsumerOrderNumber),
			TrackingNumber:      tracking,
			InvoiceDate:         shipDate,
			CurrencyCode:        currency,
			TotalAmount:         total,
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
				"invoice_request", integration.JSONForLog(inv),
				"wms_order_raw", integration.JSONForLog(wmsOrders[0]),
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
