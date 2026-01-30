package dsco_lingxing

import (
	"context"
	"strings"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/integration"
)

func (d *Domain) PushToLingXing(ctx integration.TaskContext) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	items, err := d.orderStore.FindByStatus(taskCtx, 1, ctx.Size)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}

	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		return err
	}

	for _, row := range items {
		order, err := decodeDSCOOrder(row.Payload)
		if err != nil {
			logger.Warn(taskCtx, "decode dsco payload failed", "err", err)
			continue
		}
		if strings.TrimSpace(order.PoNumber) == "" {
			logger.Warn(taskCtx, "dsco poNumber empty, skip")
			continue
		}
		shopKey := strings.TrimSpace(derefString(order.DscoRetailerID))
		storeID := ""
		if shopKey != "" {
			storeID = strings.TrimSpace(ctx.Config.Mapping.Shop[shopKey])
		}
		if storeID == "" {
			logger.Warn(taskCtx, "missing mapping.shop for dsco retailer id", "po_number", order.PoNumber, "shop_key", shopKey)
			continue
		}

		wid := ""
		dscoWarehouse := getDSCOWarehouseCode(order)
		if dscoWarehouse != "" {
			if v, ok := ctx.Config.Mapping.Warehouse[dscoWarehouse]; ok {
				wid = v
			}
		}

		logisticsTypeID := ""
		if wid != "" {
			shipWarehouseCode := strings.TrimSpace(derefString(order.ShipWarehouseCode))
			if shipWarehouseCode == "" {
				shipWarehouseCode = strings.TrimSpace(derefString(order.RequestedWarehouseCode))
			}
			serviceLevelCode := strings.TrimSpace(getDSCOShippingServiceLevelCode(order))
			key := shipWarehouseCode + "-" + serviceLevelCode
			logisticsTypeID = strings.TrimSpace(ctx.Config.Mapping.Shipment[key])
			if logisticsTypeID == "" {
				logger.Warn(taskCtx, "missing mapping.shipment for logistics_type_id", "po_number", order.PoNumber, "key", key)
				continue
			}
		}

		// Idempotency: if order already exists in LingXing, just advance status and continue.
		if _, err := lx.Order.GetOrderDetailV2(taskCtx, lingxing.OrderDetailV2Request{PlatformOrderNo: order.PoNumber}); err == nil {
			if err := d.orderStore.UpdateStatusAndFields(taskCtx, order.PoNumber, 2, "", ""); err != nil {
				logger.Warn(taskCtx, "update status failed", "err", err)
			}
			continue
		}

		addr := order.Shipping
		if addr == nil {
			logger.Warn(taskCtx, "missing shipping address", "po_number", order.PoNumber)
			continue
		}
		country := strings.TrimSpace(derefString(addr.Country))
		if country == "" {
			country = "US"
		}
		name := strings.TrimSpace(derefString(addr.Name))
		if name == "" {
			name = strings.TrimSpace(strings.TrimSpace(derefString(addr.FirstName)) + " " + strings.TrimSpace(derefString(addr.LastName)))
		}
		line1 := ""
		if addr.Address1 != nil && strings.TrimSpace(*addr.Address1) != "" {
			line1 = strings.TrimSpace(*addr.Address1)
		} else if len(addr.Address) > 0 && strings.TrimSpace(addr.Address[0]) != "" {
			line1 = strings.TrimSpace(addr.Address[0])
		}
		if line1 == "" || name == "" || strings.TrimSpace(addr.City) == "" {
			logger.Warn(taskCtx, "missing required address fields", "po_number", order.PoNumber)
			continue
		}

		createItems := make([]lingxing.CreateOrderItemV2, 0, len(order.LineItems))
		for _, li := range order.LineItems {
			msku := ""
			if li.SKU != nil && *li.SKU != "" {
				msku = *li.SKU
			} else if li.PartnerSKU != nil && *li.PartnerSKU != "" {
				msku = *li.PartnerSKU
			}
			if msku == "" {
				continue
			}
			unitPrice, ok := pickUnitPrice(li)
			if !ok {
				logger.Warn(taskCtx, "unit price missing", "po_number", order.PoNumber, "sku", msku)
				continue
			}
			createItems = append(createItems, lingxing.CreateOrderItemV2{
				MSKU:      msku,
				Quantity:  li.Quantity,
				UnitPrice: unitPrice,
			})
		}
		if len(createItems) == 0 {
			logger.Warn(taskCtx, "no items for order", "po_number", order.PoNumber)
			continue
		}

		req := lingxing.CreateOrdersV2Request{
			PlatformCode: lingxing.PlatformCode(d.env.Integration.LingXing.PlatformCode),
			StoreID:      storeID,
			Orders: []lingxing.CreateOrderV2{
				{
					PlatformOrderNo:     order.PoNumber,
					ReceiverCountryCode: country,
					ReceiverName:        name,
					City:                strings.TrimSpace(addr.City),
					AddressLine1:        line1,
					WID:                 wid,
					LogisticsTypeID:     logisticsTypeID,
					AmountCurrency:      "USD", // DSCO 用户都是美国用户，币种为固定值
					Items:               createItems,
				},
			},
		}

		resp, err := lx.Order.CreateOrdersV2(taskCtx, req)
		if err != nil {
			logger.Warn(taskCtx, "lingxing create order failed", "po_number", order.PoNumber, "err", err)
			continue
		}
		if len(resp.SuccessDetails) == 0 && len(resp.ErrorDetails) > 0 {
			// Keep simple: do not advance status, rely on retry.
			logger.Warn(taskCtx, "lingxing create order error detail", "po_number", order.PoNumber, "msg", resp.ErrorDetails[0].ErrorMessage)
			continue
		}

		if err := d.orderStore.UpdateStatusAndFields(taskCtx, order.PoNumber, 2, "", ""); err != nil {
			logger.Warn(taskCtx, "update status failed", "err", err)
			continue
		}
	}
	return nil
}

func pickUnitPrice(li dsco.OrderLineItem) (float64, bool) {
	if li.ConsumerPriceWithTax != nil && *li.ConsumerPriceWithTax > 0 {
		return *li.ConsumerPriceWithTax, true
	}
	if li.ConsumerPrice != nil && *li.ConsumerPrice > 0 {
		return *li.ConsumerPrice, true
	}
	if li.RetailPrice != nil && *li.RetailPrice > 0 {
		return *li.RetailPrice, true
	}
	if li.ExpectedCost != nil && *li.ExpectedCost > 0 {
		return *li.ExpectedCost, true
	}
	return 0, false
}
