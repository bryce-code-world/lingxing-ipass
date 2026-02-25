package dsco_boarding

import (
	"context"
	"strings"
	"testing"
	"time"

	"lingxingipass/golib/v2/sdk/dsco"
)

// 说明：
// - 这是 boarding 集成测试，会真实调用 DSCO API。
// - 配置项放在本文件全局变量中，避免与其他 boarding 测试文件命名冲突。
var (
	// cancelBaseURL 不需要改：默认生产环境；如需 staging 再改。
	cancelBaseURL = dsco.BaseURLProd

	// cancelToken：DSCO bearer token（务必自行填写）。
	cancelToken = "8b283933-2f9e-47e6-b425-ef5eb375ad54"

	// cancelOrderKey：用于 GET /order/ 拉取订单明细的 orderKey。
	// 可选值：dscoOrderId / poNumber / supplierOrderNumber
	cancelOrderKey = "poNumber"

	// cancelIDType：Cancel Order Item 接口里 OrderForCancel.type（与 Acknowledge 的枚举一致）。
	cancelIDType = dsco.OrderAcknowledgeIDTypePoNumber

	// cancelCode：取消原因码（文档为 string；具体业务枚举以 DSCO/零售商策略为准）。
	cancelCode dsco.CancelReasonCode = dsco.CancelReasonOutOfStock

	// 1) Cancel Order Item：单个测试订单（通常填 poNumber）。
	cancelSingleOrderNumber = "HHX124S5EIPXTY0TSRR5X6"

	// 2) Cancel Order Item Small Batch：两个测试订单（通常填 poNumber）。
	cancelBatchOrderNumbers = []string{
		"YWMKWQQFHRA6KDS145N12M",
		"6ZJ1W7IDR0L77HUQ7OUTYL",
	}
)

func newCancelClient(t *testing.T) *dsco.Client {
	t.Helper()

	if strings.TrimSpace(cancelToken) == "" {
		t.Fatalf("请先在 cancel_test.go 里配置 cancelToken")
	}
	cli, err := dsco.New(dsco.Config{
		BaseURL: cancelBaseURL,
		Token:   cancelToken,
	})
	if err != nil {
		t.Fatalf("dsco.New: %v", err)
	}
	return cli
}

func mustPickCancelLineItem(t *testing.T, o *dsco.Order) dsco.OrderLineItemForCancel {
	t.Helper()

	if o == nil {
		t.Fatalf("order 不能为空")
	}
	if len(o.LineItems) == 0 {
		t.Fatalf("order.lineItems 为空")
	}

	li := o.LineItems[0]
	if li.Quantity <= 0 {
		t.Fatalf("order.lineItems[0].quantity 非法：%d", li.Quantity)
	}

	out := dsco.OrderLineItemForCancel{
		CancelledQuantity: 1, // onboarding：取消一个订单行即可
		CancelCode:        cancelCode,
		LineNumber:        li.LineNumber,
	}

	// 订单行项目需要提供一个商品标识：dscoItemId / sku / partnerSku / upc / ean。
	switch {
	case li.DscoItemID != nil && strings.TrimSpace(*li.DscoItemID) != "":
		out.DscoItemID = li.DscoItemID
	case li.SKU != nil && strings.TrimSpace(*li.SKU) != "":
		out.SKU = li.SKU
	case li.PartnerSKU != nil && strings.TrimSpace(*li.PartnerSKU) != "":
		out.PartnerSKU = li.PartnerSKU
	case li.UPC != nil && strings.TrimSpace(*li.UPC) != "":
		out.UPC = li.UPC
	case li.EAN != nil && strings.TrimSpace(*li.EAN) != "":
		out.EAN = li.EAN
	default:
		t.Fatalf("order.lineItems[0] 缺少可用于取消的商品标识（dscoItemId/sku/partnerSku/upc/ean）")
	}

	return out
}

// TestBoarding_Cancel_Method1_CancelOrderItem
//
// 对应 onboarding：POST Cancel Order Item（POST /order/item/cancel）。
func TestBoarding_Cancel_Method1_CancelOrderItem(t *testing.T) {
	cli := newCancelClient(t)

	if strings.TrimSpace(cancelSingleOrderNumber) == "" {
		t.Fatalf("请先在 cancel_test.go 里配置 cancelSingleOrderNumber")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	o, err := cli.Order.GetByKey(ctx, cancelOrderKey, cancelSingleOrderNumber, nil)
	if err != nil {
		t.Fatalf("GetByKey orderKey=%s value=%s: %v", cancelOrderKey, cancelSingleOrderNumber, err)
	}

	lineItem := mustPickCancelLineItem(t, o)
	resp, err := cli.Cancel.OrderItem(ctx, &dsco.OrderForCancel{
		ID:        cancelSingleOrderNumber,
		Type:      cancelIDType,
		LineItems: []dsco.OrderLineItemForCancel{lineItem},
	})
	if err != nil {
		t.Fatalf("OrderItem: %v", err)
	}
	if resp == nil || strings.TrimSpace(resp.RequestID) == "" {
		t.Fatalf("OrderItem: requestId 为空，resp=%+v", resp)
	}
	t.Logf("cancel order item accepted: status=%s requestId=%s order=%s", resp.Status, resp.RequestID, cancelSingleOrderNumber)
}

// TestBoarding_Cancel_Method2_CancelOrderItemSmallBatch
//
// 对应 onboarding：POST Cancel Order Item Small Batch（POST /order/item/cancel/batch/small）。
func TestBoarding_Cancel_Method2_CancelOrderItemSmallBatch(t *testing.T) {
	cli := newCancelClient(t)

	if len(cancelBatchOrderNumbers) != 2 || strings.TrimSpace(cancelBatchOrderNumbers[0]) == "" || strings.TrimSpace(cancelBatchOrderNumbers[1]) == "" {
		t.Fatalf("请先在 cancel_test.go 里配置 2 个 cancelBatchOrderNumbers")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	reqs := make([]dsco.OrderForCancel, 0, 2)
	for _, orderNo := range cancelBatchOrderNumbers {
		o, err := cli.Order.GetByKey(ctx, cancelOrderKey, orderNo, nil)
		if err != nil {
			t.Fatalf("GetByKey orderKey=%s value=%s: %v", cancelOrderKey, orderNo, err)
		}
		lineItem := mustPickCancelLineItem(t, o)
		reqs = append(reqs, dsco.OrderForCancel{
			ID:        orderNo,
			Type:      cancelIDType,
			LineItems: []dsco.OrderLineItemForCancel{lineItem},
		})
	}

	resp, err := cli.Cancel.OrderItemSmallBatch(ctx, reqs)
	if err != nil {
		t.Fatalf("OrderItemSmallBatch: %v", err)
	}
	if resp == nil || strings.TrimSpace(resp.RequestID) == "" {
		t.Fatalf("OrderItemSmallBatch: requestId 为空，resp=%+v", resp)
	}
	t.Logf("cancel order item small batch accepted: status=%s requestId=%s", resp.Status, resp.RequestID)
}
