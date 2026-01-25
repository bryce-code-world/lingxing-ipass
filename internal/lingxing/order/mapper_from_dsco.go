package order

import (
	"errors"
	"fmt"
	"strings"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
)

// MapCreateOrderV2FromDSCO 将 DSCO Order 映射为领星 CreateOrderV2（一期最小必填集）。
//
// 一期原则：
// - 缺必填字段直接失败，交给上层写 manual_task 转人工
// - unit_price 优先取 consumerPrice，缺失时用 retailPrice 兜底（口径以 DSCO API spec 为准）
func MapCreateOrderV2FromDSCO(o *dsco.Order) (lingxing.CreateOrderV2, error) {
	if o == nil {
		return lingxing.CreateOrderV2{}, errors.New("dscoOrder 不能为空")
	}
	if strings.TrimSpace(o.DscoOrderID) == "" {
		return lingxing.CreateOrderV2{}, errors.New("缺少 dscoOrderId")
	}

	if o.Shipping == nil {
		return lingxing.CreateOrderV2{}, errors.New("缺少 shipping")
	}
	ship := o.Shipping

	country := strings.TrimSpace(ptrString(ship.Country))
	if country == "" {
		return lingxing.CreateOrderV2{}, errors.New("缺少 shipping.country")
	}
	city := strings.TrimSpace(ship.City)
	if city == "" {
		return lingxing.CreateOrderV2{}, errors.New("缺少 shipping.city")
	}

	name := strings.TrimSpace(ptrString(ship.Name))
	if name == "" {
		first := strings.TrimSpace(ptrString(ship.FirstName))
		last := strings.TrimSpace(ptrString(ship.LastName))
		name = strings.TrimSpace(strings.Join([]string{first, last}, " "))
	}
	if name == "" {
		return lingxing.CreateOrderV2{}, errors.New("缺少 shipping.name")
	}

	addressLine1 := ""
	if len(ship.Address) > 0 {
		addressLine1 = strings.TrimSpace(ship.Address[0])
	}
	if addressLine1 == "" {
		addressLine1 = strings.TrimSpace(ptrString(ship.Address1))
	}
	if addressLine1 == "" {
		return lingxing.CreateOrderV2{}, errors.New("缺少 shipping.address/address1")
	}
	_ = strings.TrimSpace(ptrString(ship.Address2)) // 一期暂不单独落字段，避免“过度翻译”

	if len(o.LineItems) == 0 {
		return lingxing.CreateOrderV2{}, errors.New("缺少 lineItems")
	}

	items := make([]lingxing.CreateOrderItemV2, 0, len(o.LineItems))
	for i, li := range o.LineItems {
		if li.Quantity <= 0 {
			return lingxing.CreateOrderV2{}, fmt.Errorf("lineItems[%d] quantity 非法", i)
		}
		sku := strings.TrimSpace(ptrString(li.SKU))
		if sku == "" {
			sku = strings.TrimSpace(ptrString(li.PartnerSKU))
		}
		if sku == "" {
			return lingxing.CreateOrderV2{}, fmt.Errorf("lineItems[%d] 缺少 sku/partnerSku", i)
		}

		var unitPrice *float64
		if li.ConsumerPrice != nil {
			unitPrice = li.ConsumerPrice
		} else if li.RetailPrice != nil {
			unitPrice = li.RetailPrice
		}
		if unitPrice == nil {
			return lingxing.CreateOrderV2{}, fmt.Errorf("lineItems[%d] 缺少 consumerPrice/retailPrice", i)
		}

		items = append(items, lingxing.CreateOrderItemV2{
			SKU:       sku,
			Quantity:  li.Quantity,
			UnitPrice: *unitPrice,
		})
	}

	return lingxing.CreateOrderV2{
		PlatformOrderNo:     strings.TrimSpace(o.DscoOrderID),
		ReceiverCountryCode: country,
		ReceiverName:        name,
		City:                city,
		AddressLine1:        addressLine1,
		AmountCurrency:      strings.TrimSpace(ptrString(o.CurrencyCode)),
		Items:               items,
	}, nil
}

func ptrString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
