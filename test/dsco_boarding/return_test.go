package dsco_boarding

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"lingxingipass/golib/v2/sdk/dsco"
)

// 说明：
// - 这是 boarding 集成测试，会真实调用 DSCO API。
// - 配置项放在本文件全局变量中，避免与其他 boarding 测试文件命名冲突。
var (
	// returnBaseURL 不需要改：默认生产环境；如需 staging 再改。
	returnBaseURL = dsco.BaseURLProd

	// returnToken：DSCO bearer token（务必自行填写）。
	returnToken = "8b283933-2f9e-47e6-b425-ef5eb375ad54"

	// returnOrderKey：用于 GET /order/ 拉取订单明细的 orderKey。
	// 可选值：dscoOrderId / poNumber / supplierOrderNumber
	returnOrderKey = "poNumber"

	// returnReasonCode：退货原因码（onboarding 提示 Customer Regret；具体 code 以 DSCO/portal MappingSettings 为准）。
	returnReasonCode = dsco.ReturnReasonCustomerRegret

	// returnSingleOrderNumber：用于 Create Return 的测试订单（通常选择你已发货的订单）。
	returnSingleOrderNumber = "J5DRL5M541B6BLCPOM2WER"
)

func newReturnClient(t *testing.T) *dsco.Client {
	t.Helper()

	if strings.TrimSpace(returnToken) == "" {
		t.Fatalf("请先在 return_test.go 里配置 returnToken")
	}
	cli, err := dsco.New(dsco.Config{
		BaseURL: returnBaseURL,
		Token:   returnToken,
	})
	if err != nil {
		t.Fatalf("dsco.New: %v", err)
	}
	return cli
}

func genReturnNumber() string {
	// 用时间戳避免 returnNumber 冲突。
	return "TEST-RETURN-" + time.Now().UTC().Format("20060102T150405Z")
}

func mustPickReturnLineItem(t *testing.T, o *dsco.Order) dsco.ReturnLineItemCreate {
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

	out := dsco.ReturnLineItemCreate{
		Quantity:   1,
		ReasonCode: returnReasonCode,
		LineNumber: li.LineNumber,
	}

	// ReturnLineItemCreate 需要提供一个商品标识：itemId / sku / partnerSku / upc / ean / mpn。
	switch {
	case li.DscoItemID != nil && strings.TrimSpace(*li.DscoItemID) != "":
		out.ItemID = strings.TrimSpace(*li.DscoItemID)
	case li.SKU != nil && strings.TrimSpace(*li.SKU) != "":
		out.SKU = strings.TrimSpace(*li.SKU)
	case li.PartnerSKU != nil && strings.TrimSpace(*li.PartnerSKU) != "":
		out.PartnerSKU = strings.TrimSpace(*li.PartnerSKU)
	case li.UPC != nil && strings.TrimSpace(*li.UPC) != "":
		out.UPC = strings.TrimSpace(*li.UPC)
	case li.EAN != nil && strings.TrimSpace(*li.EAN) != "":
		out.EAN = strings.TrimSpace(*li.EAN)
	default:
		t.Fatalf("order.lineItems[0] 缺少可用于退货的商品标识（dscoItemId/sku/partnerSku/upc/ean）")
	}

	return out
}

func waitReturnChangeLogHasReturnNumber(t *testing.T, ctx context.Context, cli *dsco.Client, returnNumber string) *dsco.ReturnChangeLogResponse {
	t.Helper()

	start := time.Now().Add(-10 * time.Minute).UTC().Format(time.RFC3339)
	end := time.Now().Add(2 * time.Minute).UTC().Format(time.RFC3339)

	deadline := time.Now().Add(90 * time.Second)
	for time.Now().Before(deadline) {
		resp, err := cli.Return.GetChangeLog(ctx, dsco.ReturnChangeLogQuery{
			StartDate: start,
			EndDate:   end,
			Status:    "success_or_failure",
		})
		if err != nil {
			t.Fatalf("GetChangeLog: %v", err)
		}
		if resp != nil {
			needle := `"` + returnNumber + `"`
			for _, log := range resp.Logs {
				if len(log.Payload) == 0 {
					continue
				}
				if strings.Contains(string(log.Payload), needle) {
					return resp
				}
			}
		}
		time.Sleep(2 * time.Second)
	}

	t.Fatalf("等待 /return/log 出现 returnNumber 超时（returnNumber=%s）", returnNumber)
	return nil
}

// TestBoarding_Return_Method1_CreateReturn
//
// 对应 onboarding：POST Create Return（POST /return/）。
func TestBoarding_Return_Method1_CreateReturn(t *testing.T) {
	cli := newReturnClient(t)

	if strings.TrimSpace(returnSingleOrderNumber) == "" {
		t.Fatalf("请先在 return_test.go 里配置 returnSingleOrderNumber")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	o, err := cli.Order.GetByKey(ctx, returnOrderKey, returnSingleOrderNumber, nil)
	if err != nil {
		t.Fatalf("GetByKey orderKey=%s value=%s: %v", returnOrderKey, returnSingleOrderNumber, err)
	}
	if o == nil || strings.TrimSpace(o.DscoOrderID) == "" {
		t.Fatalf("GetByKey 未返回 dscoOrderId（poNumber=%s）", returnSingleOrderNumber)
	}

	retNo := genReturnNumber()
	lineItem := mustPickReturnLineItem(t, o)

	resp, err := cli.Return.Create(ctx, &dsco.ReturnCreateRequest{
		ReturnNumber: retNo,
		DscoOrderID:  o.DscoOrderID,
		LineItems:    []dsco.ReturnLineItemCreate{lineItem},
	})
	if err != nil {
		t.Fatalf("CreateReturn: %v", err)
	}
	if resp == nil || resp.Status != "success" || resp.Return == nil || resp.Return.ReturnNumber != retNo {
		t.Fatalf("CreateReturn: resp=%+v", resp)
	}
	t.Logf("create return ok: poNumber=%s dscoOrderId=%s returnNumber=%s", returnSingleOrderNumber, o.DscoOrderID, retNo)
}

// TestBoarding_Return_Method2_GetReturnChangeLog
//
// 对应 onboarding：GET Return Changelog（GET /return/log）。
//
// 说明：先创建一张 return，再在 /return/log 中轮询直到出现该 returnNumber 的记录。
func TestBoarding_Return_Method2_GetReturnChangeLog(t *testing.T) {
	cli := newReturnClient(t)

	if strings.TrimSpace(returnSingleOrderNumber) == "" {
		t.Fatalf("请先在 return_test.go 里配置 returnSingleOrderNumber")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	o, err := cli.Order.GetByKey(ctx, returnOrderKey, returnSingleOrderNumber, nil)
	if err != nil {
		t.Fatalf("GetByKey orderKey=%s value=%s: %v", returnOrderKey, returnSingleOrderNumber, err)
	}
	if o == nil || strings.TrimSpace(o.DscoOrderID) == "" {
		t.Fatalf("GetByKey 未返回 dscoOrderId（poNumber=%s）", returnSingleOrderNumber)
	}

	retNo := genReturnNumber()
	lineItem := mustPickReturnLineItem(t, o)
	resp, err := cli.Return.Create(ctx, &dsco.ReturnCreateRequest{
		ReturnNumber: retNo,
		DscoOrderID:  o.DscoOrderID,
		LineItems:    []dsco.ReturnLineItemCreate{lineItem},
	})
	if err != nil {
		t.Fatalf("CreateReturn: %v", err)
	}
	if resp == nil || resp.Status != "success" {
		t.Fatalf("CreateReturn: resp=%+v", resp)
	}

	logResp := waitReturnChangeLogHasReturnNumber(t, ctx, cli, retNo)
	raw, _ := json.Marshal(logResp)
	t.Logf("return changelog found: returnNumber=%s resp=%s", retNo, string(raw))
}
