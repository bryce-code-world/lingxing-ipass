package order

import (
	"testing"

	"example.com/lingxing/golib/v2/sdk/dsco"
)

func TestMapCreateOrderV2FromDSCO_Behavior(t *testing.T) {
	t.Parallel()

	price := 9.99
	retail := 19.99
	cases := []struct {
		name    string
		order   *dsco.Order
		wantErr bool
	}{
		{
			name: "正常映射_shipping.name+address数组+consumerPrice",
			order: &dsco.Order{
				DscoOrderID:  "d1",
				CurrencyCode: ptr("USD"),
				Shipping: &dsco.OrderShipping{
					Country: ptr("US"),
					City:    "New York",
					Name:    ptr("Tom"),
					Address: []string{"Street 1", "Street 2"},
				},
				LineItems: []dsco.OrderLineItem{
					{Quantity: 1, SKU: ptr("SKU-1"), ConsumerPrice: &price},
				},
			},
			wantErr: false,
		},
		{
			name: "name兜底_firstName+lastName",
			order: &dsco.Order{
				DscoOrderID: "d1",
				Shipping: &dsco.OrderShipping{
					Country:   ptr("US"),
					City:      "New York",
					FirstName: ptr("Tom"),
					LastName:  ptr("Lee"),
					Address1:  ptr("Street 1"),
				},
				LineItems: []dsco.OrderLineItem{
					{Quantity: 1, SKU: ptr("SKU-1"), ConsumerPrice: &price},
				},
			},
			wantErr: false,
		},
		{
			name: "unit_price兜底_retailPrice",
			order: &dsco.Order{
				DscoOrderID: "d1",
				Shipping: &dsco.OrderShipping{
					Country:  ptr("US"),
					City:     "New York",
					Name:     ptr("Tom"),
					Address1: ptr("Street 1"),
				},
				LineItems: []dsco.OrderLineItem{
					{Quantity: 1, SKU: ptr("SKU-1"), RetailPrice: &retail},
				},
			},
			wantErr: false,
		},
		{
			name: "缺少country报错",
			order: &dsco.Order{
				DscoOrderID: "d1",
				Shipping: &dsco.OrderShipping{
					City:     "New York",
					Name:     ptr("Tom"),
					Address1: ptr("Street 1"),
				},
				LineItems: []dsco.OrderLineItem{
					{Quantity: 1, SKU: ptr("SKU-1"), ConsumerPrice: &price},
				},
			},
			wantErr: true,
		},
		{
			name: "缺少unit_price报错",
			order: &dsco.Order{
				DscoOrderID: "d1",
				Shipping: &dsco.OrderShipping{
					Country:  ptr("US"),
					City:     "New York",
					Name:     ptr("Tom"),
					Address1: ptr("Street 1"),
				},
				LineItems: []dsco.OrderLineItem{
					{Quantity: 1, SKU: ptr("SKU-1")},
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := MapCreateOrderV2FromDSCO(tc.order)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("want err, got=nil, got=%+v", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("err=%v", err)
			}
			if got.PlatformOrderNo != "d1" {
				t.Fatalf("PlatformOrderNo=%q", got.PlatformOrderNo)
			}
			if got.ReceiverCountryCode == "" || got.ReceiverName == "" || got.City == "" || got.AddressLine1 == "" {
				t.Fatalf("缺少必填字段: %+v", got)
			}
			if len(got.Items) != 1 || got.Items[0].Quantity != 1 || got.Items[0].UnitPrice <= 0 {
				t.Fatalf("items=%+v", got.Items)
			}
		})
	}
}

func ptr(s string) *string { return &s }
