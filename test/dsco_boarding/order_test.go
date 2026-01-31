package dsco_boarding

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golibv2/v2/sdk/dsco"
)

// 说明：
// - 这是 boarding 集成测试，会真实调用 DSCO API。
// - 本文件的配置项使用独立的全局变量命名，避免与 1.inventory_test.go 冲突。
// - 你只需要改下面的全局变量：orderToken / orderKeyForGetOrderObject / orderTestOrderNumbers。

var (
	// orderBaseURL 不需要改：默认生产环境；如需 staging 再改。
	orderBaseURL = dsco.BaseURLProd

	// orderToken：DSCO bearer token（务必自行填写）。
	orderToken = "8b283933-2f9e-47e6-b425-ef5eb375ad54"

	// orderKeyForGetOrderObject：用于 GET /order/ 的 orderKey（常见是 poNumber）。
	// 可选值：dscoOrderId / poNumber / supplierOrderNumber
	orderKeyForGetOrderObject = "poNumber"

	// orderTestOrderNumbers：boarding 生成的 3 个测试订单号（通常是 poNumber）。
	orderTestOrderNumbers = []string{
		// "HHX124S5EIPXTY0TSRR5X6",
		// "YWMKWQQFHRA6KDS145N12M",
		// "6ZJ1W7IDR0L77HUQ7OUTYL",
		// "U23FJHNUS4AOVAETQVW4I2",
		// "J5DRL5M541B6BLCPOM2WER",
		// "T77OMD7909R5CD4DZK38O4",
		// "YNMLAX1Z44KBBYQ9Q5C3",
		// "5J7PMN8NH5DSBCXL2M1Y",
		// "DJLLQWTZX5CA85ZWN1NL",
		// "S5S71DCU8X8QMVEMFRSE",   // shipped
		// "T77OMD7909R5CD4DZK38O4", // shipment_pending
		"P5E9VZQYS8A7MVWZMRZC", // created
		"XGKD4ZBKQKZLQ1K5E5FC", // cancelled
		"6ECM5T57UBS9CV6WDWNX",
	}
)

func newOrderClient(t *testing.T) *dsco.Client {
	t.Helper()

	if orderToken == "" {
		t.Fatalf("请先在 2.order_test.go 里配置 orderToken")
	}
	if len(orderTestOrderNumbers) != 3 || orderTestOrderNumbers[0] == "" || orderTestOrderNumbers[1] == "" || orderTestOrderNumbers[2] == "" {
		t.Fatalf("请先在 2.order_test.go 里配置 3 个 orderTestOrderNumbers")
	}

	cli, err := dsco.New(dsco.Config{
		BaseURL: orderBaseURL,
		Token:   orderToken,
	})
	if err != nil {
		t.Fatalf("dsco.New: %v", err)
	}
	return cli
}

// TestBoarding_Order_Method1_GetOrderObject
//
// 对应 onboarding：GET Get Order Object（GET /order/）。
//
// go test -v -count=1 ./order_test.go -run TestBoarding_Order_Method1_GetOrderObject
func TestBoarding_Order_Method1_GetOrderObject(t *testing.T) {
	cli := newOrderClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, no := range orderTestOrderNumbers {
		o, res, err := cli.Order.GetByKeyWithRawBody(ctx, orderKeyForGetOrderObject, no, nil)
		if err != nil {
			t.Fatalf("GetByKey orderKey=%s value=%s: %v", orderKeyForGetOrderObject, no, err)
		}
		if o == nil {
			t.Fatalf("GetByKey orderKey=%s value=%s: order 为空", orderKeyForGetOrderObject, no)
		}
		t.Logf("get order ok: value=%s dscoOrderId=%s poNumber=%s lineItems=%d", no, o.DscoOrderID, o.PoNumber, len(o.LineItems))
		// fmt.Printf("order: %+v\n", o.LineItems)
		// fmt.Println(*o.Shipping.Country, *o.Shipping.State, o.Shipping.City, o.Shipping.Address, *o.Shipping.Name)
		// fmt.Println(*o.DscoRetailerID, o.PoNumber, *o.RequestedWarehouseCode, *o.ShipWarehouseCode)
		fmt.Println("Raw response:", res)
		fmt.Println(o.PoNumber, o.DscoStatus, o.Packages,
			*o.Shipping.Phone, *o.Shipping.Name, o.Shipping.Postal, *o.BuyerMessage)
		// for _, li := range o.LineItems {
		// 	t.Logf("lineItem: %+v", li)
		// }
	}
}

// TestBoarding_Order_Method2_GetOrders
//
// 对应 onboarding：GET Get Orders（GET /order/page）。
//
// 说明：该接口有严格的时间窗口要求，until 必须至少比当前时间早 5 秒。
func TestBoarding_Order_Method2_GetOrders(t *testing.T) {
	cli := newOrderClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	includeTest := true
	until := time.Now().Add(-10 * time.Second).UTC().Format(time.RFC3339)
	createdSince := time.Now().Add(-7 * 24 * time.Hour).UTC().Format(time.RFC3339)

	page, err := cli.Order.GetPage(ctx, dsco.OrderPageQuery{
		OrdersCreatedSince: createdSince,
		Until:              until,
		IncludeTestOrders:  &includeTest,
		OrdersPerPage:      1000,
	})
	if err != nil {
		t.Fatalf("GetPage: %v", err)
	}
	if page == nil || len(page.Orders) == 0 {
		t.Fatalf("GetPage: 未返回任何订单（可能需要扩大时间窗口或确认 includeTestOrders）")
	}

	need := map[string]bool{}
	for _, v := range orderTestOrderNumbers {
		need[v] = false
	}
	for _, o := range page.Orders {
		if _, ok := need[o.PoNumber]; ok {
			need[o.PoNumber] = true
		}
	}
	for v, found := range need {
		if !found {
			t.Fatalf("GetPage: 未找到测试订单（poNumber=%s），建议确认 orderTestOrderNumbers 是否为 poNumber 或扩大时间窗口", v)
		}
	}
	t.Logf("get orders ok: returned=%d scrollId=%s", len(page.Orders), page.ScrollID)
}

// TestBoarding_Order_Method3_GetOrderChangeLog
//
// 对应 onboarding：GET Get Order Change Log（GET /order/log）。
//
// 说明：change log 是异步反馈；本测试会先发起一次 acknowledge，再用返回的 requestId 轮询 /order/log 直到 COMPLETED。
func TestBoarding_Order_Method3_GetOrderChangeLog(t *testing.T) {
	cli := newOrderClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	reqID := acknowledgeOrdersForBoarding(t, ctx, cli)
	resp := waitOrderChangeLogCompleted(t, ctx, cli, reqID)
	t.Logf("order change log completed: requestId=%s logs=%d status=%s", reqID, len(resp.Logs), resp.Status)
}

// TestBoarding_Order_Method4_AcknowledgeOrders
//
// 对应 onboarding：POST Acknowledge Orders（POST /order/acknowledge）。
func TestBoarding_Order_Method4_AcknowledgeOrders(t *testing.T) {
	cli := newOrderClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	reqID := acknowledgeOrdersForBoarding(t, ctx, cli)
	t.Logf("ack accepted: requestId=%s（可用于 /order/log 查询）", reqID)
}

func acknowledgeOrdersForBoarding(t *testing.T, ctx context.Context, cli *dsco.Client) string {
	t.Helper()

	items := make([]dsco.OrderAcknowledgeRequest, 0, len(orderTestOrderNumbers))
	for _, no := range orderTestOrderNumbers {
		items = append(items, dsco.OrderAcknowledgeRequest{
			ID:   no,
			Type: dsco.OrderAcknowledgeIDTypePoNumber,
		})
	}
	resp, err := cli.Order.Acknowledge(ctx, items)
	if err != nil {
		t.Fatalf("Acknowledge: %v", err)
	}
	if resp == nil || resp.RequestID == "" {
		t.Fatalf("Acknowledge: requestId 为空，resp=%+v", resp)
	}
	return resp.RequestID
}

func waitOrderChangeLogCompleted(t *testing.T, ctx context.Context, cli *dsco.Client, requestID string) *dsco.OrderChangeLogResponse {
	t.Helper()

	deadline := time.Now().Add(45 * time.Second)
	for time.Now().Before(deadline) {
		resp, err := cli.Order.GetChangeLog(ctx, dsco.OrderChangeLogQuery{
			RequestID: requestID,
		})
		if err != nil {
			t.Fatalf("GetChangeLog requestId=%s: %v", requestID, err)
		}
		if resp != nil && resp.Status == "COMPLETED" {
			return resp
		}
		time.Sleep(1 * time.Second)
	}
	t.Fatalf("等待 /order/log COMPLETED 超时：requestId=%s", requestID)
	return nil
}
