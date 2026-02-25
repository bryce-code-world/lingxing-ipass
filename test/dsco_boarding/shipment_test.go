package dsco_boarding

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"lingxingipass/golib/v2/sdk/dsco"
)

// 说明：
// - 这是 boarding 集成测试，会真实调用 DSCO API。
// - 配置项放在本文件全局变量中，避免与 inventory_test.go / order_test.go 命名冲突。

var (
	// shipmentBaseURL 不需要改：默认生产环境；如需 staging 再改。
	shipmentBaseURL = dsco.BaseURLProd

	// shipmentToken：DSCO bearer token（务必自行填写）。
	shipmentToken = "8b283933-2f9e-47e6-b425-ef5eb375ad54"

	// shipmentOrderKey：用于 GET /order/ 拉取订单明细的 orderKey。
	// 可选值：dscoOrderId / poNumber / supplierOrderNumber
	shipmentOrderKey = "poNumber"

	// 1) Create Shipment：单笔订单
	shipmentSingleOrderNumber    = "HHX124S5EIPXTY0TSRR5X6"
	shipmentSingleTrackingNumber = "679820701374"

	// Allowed Ship Methods（DSCO Portal 示例）
	//
	// #  Carrier  Method            Code  Tracking Number
	// 1  FedEx    2 Day             FE2D  679820701374
	//                                  679820701385
	//                                  674643439190
	// 2  FedEx    Home Delivery     FEHD  409943471260
	//                                  778377700428
	//                                  778377701468
	// 3  USPS     First-Class Mail  USFC  9200190326965311569391
	//                                  9200190326965311570014
	//                                  9200190326965311569698

	// shipmentWarehouseCode：发货需要的仓库 code（供应商侧 warehouseCode）。
	// 该账号策略可能要求必须提供。
	shipmentWarehouseCode = "YQN-CA"

	// shipmentShipCarrier / shipmentShipMethod：部分账号策略要求必须提供 carrier+method。
	// 如不配置，会尝试从订单的 shipCarrier/shipMethod 或 requestedShipCarrier/requestedShipMethod 推断。
	shipmentShipCarrier = "USPS"
	shipmentShipMethod  = "First-Class Mail"

	// shipmentShipDateRFC3339：发货日期（RFC3339）。
	// 若留空，则测试运行时使用当前 UTC 时间。
	shipmentShipDateRFC3339 = ""

	// 2) Create Shipment Small Batch：两笔订单
	// shipmentBatchOrders：支持“一个订单多行 item”发货（multi-line order）。
	// - 1 个订单（OrderNumber）可以配置多个 LineItems（每个 lineItem 配 1 个 TrackingNumber）
	// - LineItem 可通过 lineNumber 或 sku 来定位（推荐优先填 lineNumber，更稳定）
	shipmentBatchOrders = []shipmentBatchOrderConfig{
		{
			OrderNumber: "J5DRL5M541B6BLCPOM2WER",
			LineItems: []shipmentBatchLineItemConfig{
				{
					LineNumber:     1,
					SKU:            "TEST3",
					TrackingNumber: "9200190326965311569391",
				},
				{
					LineNumber:     2,
					SKU:            "TEST2",
					TrackingNumber: "9200190326965311570014",
				},
			},
		},
	}
)

type shipmentBatchOrderConfig struct {
	OrderNumber string
	LineItems   []shipmentBatchLineItemConfig
}

type shipmentBatchLineItemConfig struct {
	// LineNumber 订单行号（可选；若填写则优先用它定位 line item）
	LineNumber int
	// SKU 商品 SKU（可选；当 LineNumber 未填写时尝试用 SKU 定位）
	SKU string
	// TrackingNumber 发货运单号（必填）
	TrackingNumber string
}

func newShipmentClient(t *testing.T) *dsco.Client {
	t.Helper()

	if strings.TrimSpace(shipmentToken) == "" {
		t.Fatalf("请先在 shipment_test.go 里配置 shipmentToken")
	}
	cli, err := dsco.New(dsco.Config{
		BaseURL: shipmentBaseURL,
		Token:   shipmentToken,
	})
	if err != nil {
		t.Fatalf("dsco.New: %v", err)
	}
	return cli
}

func mustPickShipmentLineItem(o *dsco.Order) (dsco.ShipmentLineItemForUpdate, error) {
	if o == nil {
		return dsco.ShipmentLineItemForUpdate{}, errors.New("order 不能为空")
	}
	if len(o.LineItems) == 0 {
		return dsco.ShipmentLineItemForUpdate{}, errors.New("order.lineItems 为空")
	}

	li := o.LineItems[0]
	if li.Quantity <= 0 {
		return dsco.ShipmentLineItemForUpdate{}, errors.New("order.lineItems[0].quantity 非法")
	}
	if li.LineNumber == nil {
		return dsco.ShipmentLineItemForUpdate{}, errors.New("order.lineItems[0].lineNumber 为空（该账号策略要求发货必须提供 lineNumber）")
	}

	out := dsco.ShipmentLineItemForUpdate{
		Quantity:   1, // onboarding 要求：针对“单个 order line”发货即可
		LineNumber: li.LineNumber,
	}

	if li.DscoItemID != nil && strings.TrimSpace(*li.DscoItemID) != "" {
		out.DscoItemID = strings.TrimSpace(*li.DscoItemID)
		return out, nil
	}
	if li.SKU != nil && strings.TrimSpace(*li.SKU) != "" {
		out.SKU = strings.TrimSpace(*li.SKU)
		return out, nil
	}
	if li.PartnerSKU != nil && strings.TrimSpace(*li.PartnerSKU) != "" {
		out.PartnerSKU = strings.TrimSpace(*li.PartnerSKU)
		return out, nil
	}
	if li.UPC != nil && strings.TrimSpace(*li.UPC) != "" {
		out.UPC = strings.TrimSpace(*li.UPC)
		return out, nil
	}
	if li.EAN != nil && strings.TrimSpace(*li.EAN) != "" {
		out.EAN = strings.TrimSpace(*li.EAN)
		return out, nil
	}
	return dsco.ShipmentLineItemForUpdate{}, errors.New("order.lineItems[0] 缺少可用于发货的商品标识（dscoItemId/sku/partnerSku/upc/ean）")
}

func mustPickShipmentLineItemByConfig(t *testing.T, o *dsco.Order, cfg shipmentBatchLineItemConfig) dsco.ShipmentLineItemForUpdate {
	t.Helper()

	if o == nil {
		t.Fatalf("order 不能为空")
	}
	if len(o.LineItems) == 0 {
		t.Fatalf("order.lineItems 为空")
	}

	// 优先用 lineNumber，其次用 sku；两者都没填则默认取第一个 line item。
	var picked *dsco.OrderLineItem
	if cfg.LineNumber > 0 {
		for i := range o.LineItems {
			li := &o.LineItems[i]
			if li.LineNumber != nil && *li.LineNumber == cfg.LineNumber {
				picked = li
				break
			}
		}
		if picked == nil {
			t.Fatalf("按 lineNumber=%d 未找到对应的 line item（poNumber=%s）", cfg.LineNumber, o.PoNumber)
		}
	} else if strings.TrimSpace(cfg.SKU) != "" {
		want := strings.TrimSpace(cfg.SKU)
		for i := range o.LineItems {
			li := &o.LineItems[i]
			if li.SKU != nil && strings.TrimSpace(*li.SKU) == want {
				picked = li
				break
			}
		}
		if picked == nil {
			t.Fatalf("按 sku=%s 未找到对应的 line item（poNumber=%s）", want, o.PoNumber)
		}
	} else {
		picked = &o.LineItems[0]
	}

	if picked.Quantity <= 0 {
		t.Fatalf("picked lineItem quantity 非法：%d（poNumber=%s）", picked.Quantity, o.PoNumber)
	}
	if picked.LineNumber == nil {
		t.Fatalf("picked lineItem.lineNumber 为空（poNumber=%s），该账号策略要求发货必须提供 lineNumber", o.PoNumber)
	}

	out := dsco.ShipmentLineItemForUpdate{
		Quantity:   picked.Quantity, // multi-line order：通常需要把该行的全部数量都发掉
		LineNumber: picked.LineNumber,
	}

	if picked.DscoItemID != nil && strings.TrimSpace(*picked.DscoItemID) != "" {
		out.DscoItemID = strings.TrimSpace(*picked.DscoItemID)
		return out
	}
	if picked.SKU != nil && strings.TrimSpace(*picked.SKU) != "" {
		out.SKU = strings.TrimSpace(*picked.SKU)
		return out
	}
	if picked.PartnerSKU != nil && strings.TrimSpace(*picked.PartnerSKU) != "" {
		out.PartnerSKU = strings.TrimSpace(*picked.PartnerSKU)
		return out
	}
	if picked.UPC != nil && strings.TrimSpace(*picked.UPC) != "" {
		out.UPC = strings.TrimSpace(*picked.UPC)
		return out
	}
	if picked.EAN != nil && strings.TrimSpace(*picked.EAN) != "" {
		out.EAN = strings.TrimSpace(*picked.EAN)
		return out
	}
	t.Fatalf("picked lineItem 缺少可用于发货的商品标识（dscoItemId/sku/partnerSku/upc/ean）（poNumber=%s）", o.PoNumber)
	return dsco.ShipmentLineItemForUpdate{}
}

func pickShipmentCarrierMethod(o *dsco.Order) (carrier string, method string, err error) {
	if o != nil {
		if o.ShipCarrier != nil && o.ShipMethod != nil {
			c := strings.TrimSpace(*o.ShipCarrier)
			m := strings.TrimSpace(*o.ShipMethod)
			if c != "" && m != "" {
				return c, m, nil
			}
		}
		if o.RequestedShipCarrier != nil && o.RequestedShipMethod != nil {
			c := strings.TrimSpace(*o.RequestedShipCarrier)
			m := strings.TrimSpace(*o.RequestedShipMethod)
			if c != "" && m != "" {
				return c, m, nil
			}
		}
	}

	c := strings.TrimSpace(shipmentShipCarrier)
	m := strings.TrimSpace(shipmentShipMethod)
	if c == "" || m == "" {
		return "", "", errors.New("缺少 shipCarrier/shipMethod（该账号策略要求发货必须提供 carrier+method）")
	}
	return c, m, nil
}

func pickShipDateRFC3339() string {
	v := strings.TrimSpace(shipmentShipDateRFC3339)
	if v != "" {
		return v
	}
	return time.Now().UTC().Format(time.RFC3339)
}

// TestBoarding_Shipment_Method1_CreateShipment
//
// 对应 onboarding：POST Create Shipment（POST /order/singleShipment）。
func TestBoarding_Shipment_Method1_CreateShipment(t *testing.T) {
	cli := newShipmentClient(t)

	if strings.TrimSpace(shipmentSingleTrackingNumber) == "" {
		t.Fatalf("请先在 shipment_test.go 里配置 shipmentSingleTrackingNumber（按 onboarding 页 Allowed Ship Methods 对应的 tracking number）")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	o, err := cli.Order.GetByKey(ctx, shipmentOrderKey, shipmentSingleOrderNumber, nil)
	if err != nil {
		t.Fatalf("GetByKey orderKey=%s value=%s: %v", shipmentOrderKey, shipmentSingleOrderNumber, err)
	}

	lineItem, err := mustPickShipmentLineItem(o)
	if err != nil {
		t.Fatalf("pick line item: %v", err)
	}
	carrier, method, err := pickShipmentCarrierMethod(o)
	if err != nil {
		t.Fatalf("pick carrier/method: %v", err)
	}
	if strings.TrimSpace(shipmentWarehouseCode) == "" {
		t.Fatalf("请先在 shipment_test.go 里配置 shipmentWarehouseCode（该账号策略要求发货必须提供 warehouseCode）")
	}

	resp, err := cli.Shipment.CreateSingle(ctx, dsco.ShipmentsForUpdate{
		PoNumber: shipmentSingleOrderNumber,
		Shipments: []dsco.ShipmentForUpdate{
			{
				TrackingNumber: shipmentSingleTrackingNumber,
				ShipDate:       pickShipDateRFC3339(),
				ShipCarrier:    carrier,
				ShipMethod:     method,
				WarehouseCode:  shipmentWarehouseCode,
				LineItems:      []dsco.ShipmentLineItemForUpdate{lineItem},
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateSingle: %v", err)
	}
	if resp == nil || !resp.Success {
		t.Fatalf("CreateSingle: success=%v", resp != nil && resp.Success)
	}
	t.Logf("create shipment ok: order=%s tracking=%s", shipmentSingleOrderNumber, shipmentSingleTrackingNumber)
}

// TestBoarding_Shipment_Method2_CreateShipmentSmallBatch
//
// 对应 onboarding：POST Create Shipment Small Batch（POST /order/shipment/batch/small）。
func TestBoarding_Shipment_Method2_CreateShipmentSmallBatch(t *testing.T) {
	cli := newShipmentClient(t)

	if len(shipmentBatchOrders) == 0 {
		t.Fatalf("请先在 shipment_test.go 里配置 shipmentBatchOrders")
	}
	for _, cfg := range shipmentBatchOrders {
		if strings.TrimSpace(cfg.OrderNumber) == "" {
			t.Fatalf("shipmentBatchOrders: orderNumber 不能为空")
		}
		if len(cfg.LineItems) == 0 {
			t.Fatalf("shipmentBatchOrders: order=%s lineItems 不能为空", cfg.OrderNumber)
		}
		for _, li := range cfg.LineItems {
			if strings.TrimSpace(li.TrackingNumber) == "" {
				t.Fatalf("shipmentBatchOrders: order=%s 存在 trackingNumber 为空的 lineItem 配置", cfg.OrderNumber)
			}
		}
	}
	if strings.TrimSpace(shipmentWarehouseCode) == "" {
		t.Fatalf("请先在 shipment_test.go 里配置 shipmentWarehouseCode（该账号策略要求发货必须提供 warehouseCode）")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	reqs := make([]dsco.ShipmentsForUpdate, 0, len(shipmentBatchOrders))
	for _, cfg := range shipmentBatchOrders {
		o, err := cli.Order.GetByKey(ctx, shipmentOrderKey, cfg.OrderNumber, nil)
		if err != nil {
			t.Fatalf("GetByKey orderKey=%s value=%s: %v", shipmentOrderKey, cfg.OrderNumber, err)
		}
		carrier, method, err := pickShipmentCarrierMethod(o)
		if err != nil {
			t.Fatalf("pick carrier/method for %s: %v", cfg.OrderNumber, err)
		}

		shipments := make([]dsco.ShipmentForUpdate, 0, len(cfg.LineItems))
		for _, liCfg := range cfg.LineItems {
			lineItem := mustPickShipmentLineItemByConfig(t, o, liCfg)
			shipments = append(shipments, dsco.ShipmentForUpdate{
				TrackingNumber: liCfg.TrackingNumber,
				ShipDate:       pickShipDateRFC3339(),
				ShipCarrier:    carrier,
				ShipMethod:     method,
				WarehouseCode:  shipmentWarehouseCode,
				LineItems:      []dsco.ShipmentLineItemForUpdate{lineItem},
			})
		}

		reqs = append(reqs, dsco.ShipmentsForUpdate{
			PoNumber:  cfg.OrderNumber,
			Shipments: shipments,
		})
	}

	resp, err := cli.Shipment.CreateSmallBatch(ctx, reqs)
	if err != nil {
		t.Fatalf("CreateSmallBatch: %v", err)
	}
	if resp == nil || strings.TrimSpace(resp.RequestID) == "" {
		t.Fatalf("CreateSmallBatch: requestId 为空，resp=%+v", resp)
	}
	t.Logf("create shipment small batch accepted: status=%s requestId=%s", resp.Status, resp.RequestID)
}
