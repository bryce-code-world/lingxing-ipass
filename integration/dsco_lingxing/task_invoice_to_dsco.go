package dsco_lingxing

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
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
func (d *Domain) InvoiceToDSCO(ctx integration.TaskContext) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	// 1) 取待回传发票（status=4）
	items, err := d.orderStore.FindByStatus(taskCtx, 4, ctx.Size)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}

	// 2) 初始化客户端：DSCO + 领星
	dscoCli, err := d.dscoClient()
	if err != nil {
		return err
	}
	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		return err
	}

	var invs []dsco.Invoice
	var toUpdate []struct {
		po        string
		invoiceID string
	}

	for _, row := range items {
		po := row.PONumber
		// 0) 幂等：已写入 dsco_invoice_id 的订单视为已回传过 invoice，直接跳过。
		if row.DSCOInvoiceID != "" {
			continue
		}

		// 3) 解析 DSCO 原始订单（用于 dscoItemId/币种/参考字段等）
		var dscoOrder dsco.Order
		if err := json.Unmarshal(row.Payload, &dscoOrder); err != nil {
			continue
		}

		// 3.1) sid：用于 WmsOrderList 精准查询（从 mapping.shop 解析）
		sid, ok := lingxingSIDFromMapping(ctx.Config, dscoOrder)
		if !ok {
			logger.Warn(taskCtx, "missing mapping.shop sid for wms order list", "po_number", po)
			continue
		}

		wmsOrders, _, err := lx.Warehouse.WmsOrderList(taskCtx, lingxing.WmsOrderListRequest{
			Page:               1,
			PageSize:           20,
			SIDArr:             []int{sid},
			PlatformOrderNoArr: []string{po},
		})
		if err != nil || len(wmsOrders) == 0 {
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
		invs = append(invs, dsco.Invoice{
			InvoiceID:           invoiceID,
			PoNumber:            po,
			ConsumerOrderNumber: derefString(dscoOrder.ConsumerOrderNumber),
			TrackingNumber:      tracking,
			InvoiceDate:         shipDate,
			CurrencyCode:        currency,
			TotalAmount:         total,
			LineItems:           lineItems,
		})
		toUpdate = append(toUpdate, struct {
			po        string
			invoiceID string
		}{po: po, invoiceID: invoiceID})
	}
	if len(invs) == 0 {
		return nil
	}
	// 11) 调用 DSCO invoice 批量接口
	_, err = dscoCli.Invoice.CreateSmallBatch(taskCtx, invs)
	if err != nil {
		return err
	}
	// 12) 回传成功：写回 invoiceId，并推进到 status=5
	for _, it := range toUpdate {
		if err := d.orderStore.UpdateStatusAndFields(taskCtx, it.po, 5, "", it.invoiceID); err != nil {
			logger.Warn(taskCtx, "update status after invoice failed", "err", err)
		}
	}
	return nil
}
