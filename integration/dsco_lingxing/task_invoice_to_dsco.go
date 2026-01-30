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

func (d *Domain) InvoiceToDSCO(ctx integration.TaskContext) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	items, err := d.orderStore.FindByStatus(taskCtx, 4, ctx.Size)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}

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
		if row.DSCOInvoiceID != "" {
			continue
		}

		var dscoOrder dsco.Order
		if err := json.Unmarshal(row.Payload, &dscoOrder); err != nil {
			continue
		}

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

		tracking := wmsOrders[0].TrackingNo

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

		currency := "USD"
		if dscoOrder.CurrencyCode != nil && *dscoOrder.CurrencyCode != "" {
			currency = *dscoOrder.CurrencyCode
		}

		// Round total to 2 decimals to be safe.
		total = math.Round(total*100) / 100

		invoiceID := po
		if len(wmsOrders) > 1 {
			// Future-proof: when a PO has multiple WMS orders, use a stable 1-based sequence.
			invoiceID = fmt.Sprintf("%s-%d", po, 1)
		}
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
	_, err = dscoCli.Invoice.CreateSmallBatch(taskCtx, invs)
	if err != nil {
		return err
	}
	for _, it := range toUpdate {
		if err := d.orderStore.UpdateStatusAndFields(taskCtx, it.po, 5, "", it.invoiceID); err != nil {
			logger.Warn(taskCtx, "update status after invoice failed", "err", err)
		}
	}
	return nil
}
