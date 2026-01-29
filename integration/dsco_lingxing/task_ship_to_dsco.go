package dsco_lingxing

import (
	"context"
	"encoding/json"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/integration"
)

func (d *Domain) ShipToDSCO(ctx integration.TaskContext) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	items, err := d.orderStore.FindByStatus(taskCtx, 3, ctx.Size)
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

	_, skuRev, shipRev, err := buildReverseMaps(ctx.Config)
	if err != nil {
		return err
	}

	var batch []dsco.ShipmentsForUpdate
	var toUpdate []string

	for _, row := range items {
		po := row.PONumber
		detail, err := lx.Order.GetOrderDetailV2(taskCtx, lingxing.OrderDetailV2Request{PlatformOrderNo: po})
		if err != nil {
			continue
		}
		tracking := detail.LogisticsInfo.TrackingNo
		if tracking == "" {
			continue
		}
		dscoShipMethod := shipRev[detail.LogisticsInfo.LogisticsTypeName]

		shipDateRFC3339 := ""
		switch d.env.Integration.Shipment.ShipDateSource {
		case "none":
		case "delivered_at", "stock_delivered_at":
			orders, _, err := lx.Warehouse.WmsOrderList(taskCtx, lingxing.WmsOrderListRequest{
				Page:               1,
				PageSize:           20,
				SIDArr:             []int{d.env.Integration.LingXing.SID},
				PlatformOrderNoArr: []string{po},
			})
			if err == nil && len(orders) > 0 {
				rawTime := orders[0].DeliveredAt
				if d.env.Integration.Shipment.ShipDateSource == "stock_delivered_at" {
					rawTime = orders[0].StockDeliveredAt
				}
				if rawTime != "" {
					if t, err := parseLingXingDateTimeToRFC3339UTC(rawTime); err == nil {
						shipDateRFC3339 = t
					}
				}
			}
			if shipDateRFC3339 == "" && d.env.Integration.Shipment.ShipDateSource == "delivered_at" {
				// Fallback to order detail timestamp seconds.
				if sec, err := parseInt64(detail.GlobalDeliveryTime); err == nil && sec > 0 {
					shipDateRFC3339 = time.Unix(sec, 0).UTC().Format(time.RFC3339)
				}
			}
		}

		// Build dscoItemId map from payload if possible.
		var dscoOrder dsco.Order
		_ = json.Unmarshal(row.Payload, &dscoOrder)
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

		// Prefer WMS shipped quantities.
		wmsOrders, _, err := lx.Warehouse.WmsOrderList(taskCtx, lingxing.WmsOrderListRequest{
			Page:               1,
			PageSize:           20,
			SIDArr:             []int{d.env.Integration.LingXing.SID},
			PlatformOrderNoArr: []string{po},
		})
		if err == nil && len(wmsOrders) > 0 && len(wmsOrders[0].ProductInfo) > 0 {
			for _, p := range wmsOrders[0].ProductInfo {
				dscoPartner := skuRev[p.SKU]
				if dscoPartner == "" {
					dscoPartner = p.SKU
				}
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
			// Fallback: order detail item quantities.
			for _, it := range detail.ItemInfo {
				dscoPartner := skuRev[it.MSKU]
				if dscoPartner == "" {
					dscoPartner = it.MSKU
				}
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
			continue
		}

		batch = append(batch, dsco.ShipmentsForUpdate{
			PoNumber: po,
			Shipments: []dsco.ShipmentForUpdate{
				{
					TrackingNumber: tracking,
					ShipDate:       shipDateRFC3339,
					ShipMethod:     dscoShipMethod,
					LineItems:      lineItems,
				},
			},
		})
		toUpdate = append(toUpdate, po)
	}
	if len(batch) == 0 {
		return nil
	}
	_, err = dscoCli.Shipment.CreateSmallBatch(taskCtx, batch)
	if err != nil {
		return err
	}
	for _, po := range toUpdate {
		if err := d.orderStore.UpdateStatusAndFields(taskCtx, po, 4, "", ""); err != nil {
			logger.Warn(taskCtx, "update status after ship failed", "err", err)
		}
	}
	return nil
}
