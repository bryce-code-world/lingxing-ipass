package dsco_lingxing

import (
	"context"
	"encoding/json"
	"math"
	"strconv"
	"strings"
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

	skuRev, err := buildReverseSKUMap(ctx.Config)
	if err != nil {
		return err
	}

	var invs []dsco.Invoice
	var toUpdate []string

	for _, row := range items {
		po := row.PONumber

		var dscoOrder dsco.Order
		if err := json.Unmarshal(row.Payload, &dscoOrder); err != nil {
			continue
		}

		shipCode := getDSCOShippingServiceLevelCode(dscoOrder)
		sidStr := ""
		if shipCode != "" {
			sidStr = ctx.Config.Mapping.Shipment[shipCode]
		}
		sidStr = strings.TrimSpace(sidStr)
		if sidStr == "" {
			logger.Warn(taskCtx, "missing mapping.shipment for dsco shippingServiceLevelCode", "po_number", po, "shipping_service_level_code", shipCode)
			continue
		}
		sid, err := strconv.Atoi(sidStr)
		if err != nil || sid <= 0 {
			logger.Warn(taskCtx, "invalid mapping.shipment sid", "po_number", po, "shipping_service_level_code", shipCode, "sid", sidStr)
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
			dscoPartner := skuRev[p.SKU]
			if dscoPartner == "" {
				dscoPartner = p.SKU
			}
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
		toUpdate = append(toUpdate, po)
	}
	if len(invs) == 0 {
		return nil
	}
	_, err = dscoCli.Invoice.CreateSmallBatch(taskCtx, invs)
	if err != nil {
		return err
	}
	for _, po := range toUpdate {
		if err := d.orderStore.UpdateStatusAndFields(taskCtx, po, 5, "", po); err != nil {
			logger.Warn(taskCtx, "update status after invoice failed", "err", err)
		}
	}
	return nil
}
