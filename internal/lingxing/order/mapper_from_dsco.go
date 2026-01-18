package order

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
)

type dscoShipping struct {
	City    string `json:"city"`
	Country string `json:"country"`

	Name      string `json:"name"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	Address1 string   `json:"address1"`
	Address2 string   `json:"address2"`
	Address  []string `json:"address"`
}

// MapCreateOrderV2FromDSCO 将 DSCO Order 映射为领星 CreateOrderV2（一期最小必填集）。
//
// 一期原则：
// - 缺必填字段直接失败，交给上层写 manual_task 转人工
// - unit_price 优先取 consumerPrice，缺失时用 expectedCost 兜底
func MapCreateOrderV2FromDSCO(o *dsco.Order) (lingxing.CreateOrderV2, error) {
	if o == nil {
		return lingxing.CreateOrderV2{}, errors.New("dscoOrder 不能为空")
	}
	if strings.TrimSpace(o.DscoOrderID) == "" {
		return lingxing.CreateOrderV2{}, errors.New("缺少 dscoOrderId")
	}

	var ship dscoShipping
	if len(o.Shipping) > 0 {
		if err := json.Unmarshal(o.Shipping, &ship); err != nil {
			return lingxing.CreateOrderV2{}, fmt.Errorf("解析 shipping 失败: %w", err)
		}
	}

	country := strings.TrimSpace(ship.Country)
	if country == "" {
		return lingxing.CreateOrderV2{}, errors.New("缺少 shipping.country")
	}
	city := strings.TrimSpace(ship.City)
	if city == "" {
		return lingxing.CreateOrderV2{}, errors.New("缺少 shipping.city")
	}

	name := strings.TrimSpace(ship.Name)
	if name == "" {
		first := strings.TrimSpace(ship.FirstName)
		last := strings.TrimSpace(ship.LastName)
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
		addressLine1 = strings.TrimSpace(ship.Address1)
	}
	if addressLine1 == "" {
		return lingxing.CreateOrderV2{}, errors.New("缺少 shipping.address/address1")
	}
	_ = strings.TrimSpace(ship.Address2) // 一期暂不单独落字段，避免“过度翻译”

	if len(o.LineItems) == 0 {
		return lingxing.CreateOrderV2{}, errors.New("缺少 lineItems")
	}

	items := make([]lingxing.CreateOrderItemV2, 0, len(o.LineItems))
	for i, li := range o.LineItems {
		if li.Quantity <= 0 {
			return lingxing.CreateOrderV2{}, fmt.Errorf("lineItems[%d] quantity 非法", i)
		}
		sku := strings.TrimSpace(li.SKU)
		if sku == "" {
			sku = strings.TrimSpace(li.PartnerSKU)
		}
		if sku == "" {
			return lingxing.CreateOrderV2{}, fmt.Errorf("lineItems[%d] 缺少 sku/partnerSku", i)
		}

		var unitPrice *float64
		if li.ConsumerPrice != nil {
			unitPrice = li.ConsumerPrice
		} else if li.ExpectedCost != nil {
			unitPrice = li.ExpectedCost
		}
		if unitPrice == nil {
			return lingxing.CreateOrderV2{}, fmt.Errorf("lineItems[%d] 缺少 consumerPrice/expectedCost", i)
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
		AmountCurrency:      strings.TrimSpace(o.CurrencyCode),
		Items:               items,
	}, nil
}
