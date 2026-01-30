package dsco_lingxing

import (
	"context"
	"encoding/json"
	"strings"

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

	var batch []dsco.ShipmentsForUpdate
	var toUpdate []string

	for _, row := range items {
		po := row.PONumber
		if row.ShippedTrackingNo != "" {
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

		detail, err := lx.Order.GetOrderDetailV2(taskCtx, lingxing.OrderDetailV2Request{PlatformOrderNo: po})
		if err != nil {
			continue
		}
		tracking := detail.LogisticsInfo.TrackingNo
		if tracking == "" {
			continue
		}

		dscoShipMethod := getDSCOShipMethod(dscoOrder)
		if dscoShipMethod == "" {
			dscoShipMethod = getDSCOShippingServiceLevelCode(dscoOrder)
		}

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
			// Fallback: order detail item quantities.
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
		// shipped_tracking_no: use tracking returned by lingxing order detail.
		// For MVP: re-query detail to get tracking (avoid carrying extra state in toUpdate).
		detail, derr := lx.Order.GetOrderDetailV2(taskCtx, lingxing.OrderDetailV2Request{PlatformOrderNo: po})
		tracking := ""
		if derr == nil {
			tracking = strings.TrimSpace(detail.LogisticsInfo.TrackingNo)
		}
		if err := d.orderStore.UpdateStatusAndFields(taskCtx, po, 4, tracking, ""); err != nil {
			logger.Warn(taskCtx, "update status after ship failed", "err", err)
		}
	}
	return nil
}
