package boarding

import (
	"context"
	"fmt"
	"testing"

	"example.com/lingxing/golib/v2/sdk/lingxing"
)

// 说明：
// - 这些是“真实环境手动测试”（boarding/e2e），会真实调用领星接口并可能创建/修改订单。
// - 默认 go test ./... 会跳过（见 init.go 的 newClient）。
func TestOrderService_CreateOrdersV2(t *testing.T) {
	cli := newClient(t)
	out, err := cli.Order.CreateOrdersV2(context.Background(), lingxing.CreateOrdersV2Request{
		PlatformCode: 10009,                // 自定义平台
		StoreID:      "110658143132021760", // Chewy 店铺ID
		Orders: []lingxing.CreateOrderV2{
			{
				PlatformOrderNo:     "HHX124S5EIPXTY0TSRR5X6",
				ReceiverCountryCode: "US",
				ReceiverName:        "SHIP_TO_FIRST_NAME SHIP_TO_LAST_NAME",
				City:                "Albany",
				AddressLine1:        "123 Abc St",
				AmountCurrency:      "USD",
				WID:                 "33328",
				LogisticsTypeID:     "203662740883276800",
				Items: []lingxing.CreateOrderItemV2{
					{MSKU: "TEST3", Quantity: 1, UnitPrice: 21.34, StockDeductionType: 1},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateOrdersV2() err=%v", err)
	}
	fmt.Println("CreateOrdersV2 resp:", out)
}

// 获取订单列表
//
// go test -v -count=1 .\\init.go .\\order_test.go -run TestOrderService_ListOrdersV2
func TestOrderService_ListOrdersV2(t *testing.T) {
	cli := newClient(t)
	out, res, err := cli.Order.ListOrdersV2WithRawBody(context.Background(), lingxing.OrderListV2Request{
		PlatformCode:     []lingxing.PlatformCode{lingxing.PlatformCodeCustom}, // 自定义平台
		StoreID:          []string{"110658143132021760"},                       // Chewy 店铺ID
		Offset:           0,
		Length:           20,
		DateType:         lingxing.MultiPlatformOrderDateTypeUpdateTime,
		StartTime:        1769304406,
		EndTime:          1769650006,
		PlatformOrderNos: []string{"HHX124S5EIPXTY0TSRR5X6"},
	})
	if err != nil {
		t.Fatalf("ListOrdersV2() err=%v", err)
	}
	fmt.Println("ListOrdersV2 resp:", out)
	fmt.Println("ListOrdersV2 raw resp:", res)
}

// 编辑订单
func TestOrderService_EditOrderV2(t *testing.T) {
	cli := newClient(t)
	_, raw, err := cli.Order.EditOrderWithRawBody(context.Background(), lingxing.EditOrderRequest{
		OrderList: []lingxing.EditOrderItem{
			{
				GlobalOrderNo: 103662123281304064,
				Logistics: lingxing.EditOrderLogistics{
					LogisticsTypeID: 203662740883276800,
					SysWID:          33328,
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("EditOrderV2() err=%v", err)
	}
	fmt.Println("EditOrderV2 resp:", raw)
}

// 更新订单
func TestOrderService_UpdateOrderV2(t *testing.T) {
	cli := newClient(t)
	out, err := cli.Order.UpdateOrderV2(context.Background(), lingxing.UpdateOrderV2Request{})
	if err != nil {
		t.Fatalf("UpdateOrderV2() err=%v", err)
	}
	fmt.Println("UpdateOrderV2 resp:", out)
}
