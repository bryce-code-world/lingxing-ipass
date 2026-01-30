//go:build dsco_boarding

package boarding

import (
	"context"
	"strings"
	"testing"
	"time"

	"golibv2/v2/sdk/dsco"
)

// 说明：
// - 这是 boarding 集成测试，会真实调用 DSCO API。
// - 配置项放在本文件全局变量中，避免与 inventory_test.go / order_test.go / shipment_test.go / return_test.go 命名冲突。
var (
	// invoiceBaseURL 不需要改：默认生产环境；如需 staging 再改。
	invoiceBaseURL = dsco.BaseURLProd

	// invoiceToken：DSCO bearer token（务必自行填写）。
	invoiceToken = "8b283933-2f9e-47e6-b425-ef5eb375ad54"

	// invoiceOrderKey：用于 GET /order/ 拉取订单明细的 orderKey（常用 poNumber）。
	// 可选值：dscoOrderId / poNumber / supplierOrderNumber
	invoiceOrderKey = "poNumber"

	// invoiceCurrencyCode：币种（示例用 USD；若你的订单不是 USD，请改为实际币种）。
	invoiceCurrencyCode = "USD"

	// invoiceShipDateRFC3339：invoiceDate（RFC3339）。若留空，则测试运行时使用当前 UTC 时间。
	// Portal 提示 invoiceDate 为必填（custom required）。
	invoiceDateRFC3339 = ""

	// invoiceUnitPrice：测试用单价（金额口径可能受零售商/账户策略影响；如遇校验失败请按实际规则调整）。
	invoiceUnitPrice = 1.0

	// invoiceSingleOrderNumber：用于 Create Invoice 的测试订单（通常选择你已发货的订单）。
	invoiceSingleOrderNumber = "J5DRL5M541B6BLCPOM2WER"

	// invoiceBatchOrderNumbers：用于 Create Invoice Small Batch 的测试订单（至少 2 个更贴近 onboarding）。
	invoiceBatchOrderNumbers = []string{
		"YWMKWQQFHRA6KDS145N12M",
		"6ZJ1W7IDR0L77HUQ7OUTYL",
	}
)

func newInvoiceClient(t *testing.T) *dsco.Client {
	t.Helper()

	if strings.TrimSpace(invoiceToken) == "" {
		t.Fatalf("请先在 invoice_test.go 里配置 invoiceToken")
	}
	cli, err := dsco.New(dsco.Config{
		BaseURL: invoiceBaseURL,
		Token:   invoiceToken,
	})
	if err != nil {
		t.Fatalf("dsco.New: %v", err)
	}
	return cli
}

func pickInvoiceDateRFC3339() string {
	v := strings.TrimSpace(invoiceDateRFC3339)
	if v != "" {
		return v
	}
	return time.Now().UTC().Format(time.RFC3339)
}

func genInvoiceID(prefix string) string {
	// 用时间戳避免 invoiceId 冲突（invoiceId 在 DSCO 侧通常要求唯一）。
	ts := time.Now().UTC().Format("20060102T150405Z")
	return prefix + ts
}

func pickInvoiceCurrencyCode(o *dsco.Order) string {
	if strings.TrimSpace(invoiceCurrencyCode) != "" {
		return strings.TrimSpace(invoiceCurrencyCode)
	}
	if o != nil && o.CurrencyCode != nil && strings.TrimSpace(*o.CurrencyCode) != "" {
		return strings.TrimSpace(*o.CurrencyCode)
	}
	return "USD"
}

func mustBuildInvoiceLineItemsFromOrder(t *testing.T, o *dsco.Order) ([]dsco.InvoiceLineItem, float64) {
	t.Helper()

	if o == nil {
		t.Fatalf("order 不能为空")
	}
	if len(o.LineItems) == 0 {
		t.Fatalf("order.lineItems 为空")
	}

	// 零售商策略：invoice 必须匹配整个订单（不能只开其中 1 行）。
	items := make([]dsco.InvoiceLineItem, 0, len(o.LineItems))
	var total float64

	for i := range o.LineItems {
		li := o.LineItems[i]
		if li.Quantity <= 0 {
			t.Fatalf("order.lineItems[%d].quantity 非法：%d", i, li.Quantity)
		}

		out := dsco.InvoiceLineItem{
			Quantity:  li.Quantity,
			UnitPrice: invoiceUnitPrice,
		}
		if li.LineNumber != nil && *li.LineNumber > 0 {
			out.LineNumber = *li.LineNumber
		}

		// 发票行项目需要能定位到商品：dscoItemId / sku / partnerSku。
		switch {
		case li.DscoItemID != nil && strings.TrimSpace(*li.DscoItemID) != "":
			out.DscoItemID = strings.TrimSpace(*li.DscoItemID)
		case li.SKU != nil && strings.TrimSpace(*li.SKU) != "":
			out.SKU = strings.TrimSpace(*li.SKU)
		case li.PartnerSKU != nil && strings.TrimSpace(*li.PartnerSKU) != "":
			out.PartnerSKU = strings.TrimSpace(*li.PartnerSKU)
		default:
			t.Fatalf("order.lineItems[%d] 缺少可用于发票的商品标识（dscoItemId/sku/partnerSku）", i)
		}

		items = append(items, out)
		total += float64(out.Quantity) * out.UnitPrice
	}

	return items, total
}

func waitInvoiceVisibleByInvoiceID(t *testing.T, ctx context.Context, cli *dsco.Client, invoiceID string) {
	t.Helper()

	deadline := time.Now().Add(60 * time.Second)
	for time.Now().Before(deadline) {
		resp, err := cli.Invoice.GetByID(ctx, dsco.InvoiceGetQuery{
			Key:   "invoiceId",
			Value: invoiceID,
		})
		if err == nil && resp != nil {
			for _, inv := range resp.Invoices {
				if inv.InvoiceID == invoiceID {
					return
				}
			}
		}
		time.Sleep(2 * time.Second)
	}
	t.Fatalf("等待发票可查询超时（invoiceId=%s）", invoiceID)
}

func waitInvoiceChangeLogCompleted(t *testing.T, ctx context.Context, cli *dsco.Client, requestID string) *dsco.InvoiceChangeLogResponse {
	t.Helper()

	deadline := time.Now().Add(60 * time.Second)
	for time.Now().Before(deadline) {
		resp, err := cli.Invoice.GetChangeLog(ctx, dsco.InvoiceChangeLogQuery{
			RequestID: requestID,
		})
		if err != nil {
			t.Fatalf("GetChangeLog requestId=%s: %v", requestID, err)
		}
		if resp != nil && resp.Status == "COMPLETED" {
			return resp
		}
		time.Sleep(2 * time.Second)
	}
	t.Fatalf("等待 /invoice/log COMPLETED 超时（requestId=%s）", requestID)
	return nil
}

// TestBoarding_Invoice_Method1_CreateInvoice
//
// 对应 onboarding：POST Create Invoice（POST /invoice）。
func TestBoarding_Invoice_Method1_CreateInvoice(t *testing.T) {
	cli := newInvoiceClient(t)

	if strings.TrimSpace(invoiceSingleOrderNumber) == "" {
		t.Fatalf("请先在 invoice_test.go 里配置 invoiceSingleOrderNumber")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	o, err := cli.Order.GetByKey(ctx, invoiceOrderKey, invoiceSingleOrderNumber, nil)
	if err != nil {
		t.Fatalf("GetByKey orderKey=%s value=%s: %v", invoiceOrderKey, invoiceSingleOrderNumber, err)
	}

	lineItems, total := mustBuildInvoiceLineItemsFromOrder(t, o)
	invoiceID := genInvoiceID("TEST-INVOICE-")

	resp, err := cli.Invoice.CreateSingle(ctx, &dsco.Invoice{
		InvoiceID:    invoiceID,
		PoNumber:     invoiceSingleOrderNumber,
		InvoiceDate:  pickInvoiceDateRFC3339(),
		CurrencyCode: pickInvoiceCurrencyCode(o),
		TotalAmount:  total,
		LineItems:    lineItems,
	})
	if err != nil {
		t.Fatalf("CreateSingle: %v", err)
	}
	if resp == nil || !resp.Success {
		t.Fatalf("CreateSingle: success=%v", resp != nil && resp.Success)
	}

	// 文档说明：创建后需要等待 post-processing，发票才可被查询到。
	waitInvoiceVisibleByInvoiceID(t, ctx, cli, invoiceID)
	t.Logf("create invoice ok: poNumber=%s invoiceId=%s", invoiceSingleOrderNumber, invoiceID)
}

// TestBoarding_Invoice_Method2_CreateInvoiceSmallBatch
//
// 对应 onboarding：POST Create Invoice Small Batch（POST /invoice/batch/small）。
func TestBoarding_Invoice_Method2_CreateInvoiceSmallBatch(t *testing.T) {
	cli := newInvoiceClient(t)

	if len(invoiceBatchOrderNumbers) == 0 {
		t.Fatalf("请先在 invoice_test.go 里配置 invoiceBatchOrderNumbers")
	}
	for _, v := range invoiceBatchOrderNumbers {
		if strings.TrimSpace(v) == "" {
			t.Fatalf("invoiceBatchOrderNumbers 存在空值")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	invs := make([]dsco.Invoice, 0, len(invoiceBatchOrderNumbers))
	for _, orderNo := range invoiceBatchOrderNumbers {
		o, err := cli.Order.GetByKey(ctx, invoiceOrderKey, orderNo, nil)
		if err != nil {
			t.Fatalf("GetByKey orderKey=%s value=%s: %v", invoiceOrderKey, orderNo, err)
		}
		lineItems, total := mustBuildInvoiceLineItemsFromOrder(t, o)
		invs = append(invs, dsco.Invoice{
			InvoiceID:    genInvoiceID("TEST-INVOICE-BATCH-"),
			PoNumber:     orderNo,
			InvoiceDate:  pickInvoiceDateRFC3339(),
			CurrencyCode: pickInvoiceCurrencyCode(o),
			TotalAmount:  total,
			LineItems:    lineItems,
		})
	}

	resp, err := cli.Invoice.CreateSmallBatch(ctx, invs)
	if err != nil {
		t.Fatalf("CreateSmallBatch: %v", err)
	}
	if resp == nil || strings.TrimSpace(resp.RequestID) == "" {
		t.Fatalf("CreateSmallBatch: requestId 为空，resp=%+v", resp)
	}
	t.Logf("create invoice small batch accepted: status=%s requestId=%s", resp.Status, resp.RequestID)
}

// TestBoarding_Invoice_Method3_GetInvoicesByID
//
// 对应 onboarding：GET Get Invoices by ID（GET /invoice）。
//
// 说明：为避免受“发票创建后需要 post-processing 才可查询”的影响，本测试会先创建一张发票，再按 invoiceId 查询并轮询直至可见。
func TestBoarding_Invoice_Method3_GetInvoicesByID(t *testing.T) {
	cli := newInvoiceClient(t)

	if strings.TrimSpace(invoiceSingleOrderNumber) == "" {
		t.Fatalf("请先在 invoice_test.go 里配置 invoiceSingleOrderNumber")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	o, err := cli.Order.GetByKey(ctx, invoiceOrderKey, invoiceSingleOrderNumber, nil)
	if err != nil {
		t.Fatalf("GetByKey orderKey=%s value=%s: %v", invoiceOrderKey, invoiceSingleOrderNumber, err)
	}

	lineItems, total := mustBuildInvoiceLineItemsFromOrder(t, o)
	invoiceID := genInvoiceID("TEST-INVOICE-GET-")
	resp, err := cli.Invoice.CreateSingle(ctx, &dsco.Invoice{
		InvoiceID:    invoiceID,
		PoNumber:     invoiceSingleOrderNumber,
		InvoiceDate:  pickInvoiceDateRFC3339(),
		CurrencyCode: pickInvoiceCurrencyCode(o),
		TotalAmount:  total,
		LineItems:    lineItems,
	})
	if err != nil {
		t.Fatalf("CreateSingle: %v", err)
	}
	if resp == nil || !resp.Success {
		t.Fatalf("CreateSingle: success=%v", resp != nil && resp.Success)
	}

	waitInvoiceVisibleByInvoiceID(t, ctx, cli, invoiceID)

	got, err := cli.Invoice.GetByID(ctx, dsco.InvoiceGetQuery{
		Key:   "invoiceId",
		Value: invoiceID,
	})
	if err != nil {
		t.Fatalf("GetByID: %v", err)
	}
	if got == nil || len(got.Invoices) == 0 {
		t.Fatalf("GetByID: invoices 为空，resp=%+v", got)
	}
	found := false
	for _, inv := range got.Invoices {
		if inv.InvoiceID == invoiceID {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("GetByID: 未找到 invoiceId=%s，resp=%+v", invoiceID, got)
	}
	t.Logf("get invoices ok: invoiceId=%s count=%d", invoiceID, len(got.Invoices))
}

// TestBoarding_Invoice_Method4_GetInvoiceChangeLog
//
// 对应 onboarding：GET Get Invoice Change Log（GET /invoice/log）。
//
// 说明：Change log 用于反馈异步 batch 的处理结果；本测试会先调用 CreateSmallBatch，再用 requestId 轮询直至 COMPLETED。
func TestBoarding_Invoice_Method4_GetInvoiceChangeLog(t *testing.T) {
	cli := newInvoiceClient(t)

	if len(invoiceBatchOrderNumbers) == 0 {
		t.Fatalf("请先在 invoice_test.go 里配置 invoiceBatchOrderNumbers")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()

	invs := make([]dsco.Invoice, 0, len(invoiceBatchOrderNumbers))
	for _, orderNo := range invoiceBatchOrderNumbers {
		o, err := cli.Order.GetByKey(ctx, invoiceOrderKey, orderNo, nil)
		if err != nil {
			t.Fatalf("GetByKey orderKey=%s value=%s: %v", invoiceOrderKey, orderNo, err)
		}
		lineItems, total := mustBuildInvoiceLineItemsFromOrder(t, o)
		invs = append(invs, dsco.Invoice{
			InvoiceID:    genInvoiceID("TEST-INVOICE-LOG-"),
			PoNumber:     orderNo,
			InvoiceDate:  pickInvoiceDateRFC3339(),
			CurrencyCode: pickInvoiceCurrencyCode(o),
			TotalAmount:  total,
			LineItems:    lineItems,
		})
	}

	createResp, err := cli.Invoice.CreateSmallBatch(ctx, invs)
	if err != nil {
		t.Fatalf("CreateSmallBatch: %v", err)
	}
	if createResp == nil || strings.TrimSpace(createResp.RequestID) == "" {
		t.Fatalf("CreateSmallBatch: requestId 为空，resp=%+v", createResp)
	}

	logResp := waitInvoiceChangeLogCompleted(t, ctx, cli, createResp.RequestID)
	if logResp == nil || len(logResp.Logs) == 0 {
		t.Fatalf("GetChangeLog: logs 为空，resp=%+v", logResp)
	}
	t.Logf("invoice change log completed: requestId=%s logs=%d status=%s", createResp.RequestID, len(logResp.Logs), logResp.Status)
}
