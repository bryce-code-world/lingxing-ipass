package dsco_boarding

import (
	"context"
	"testing"
	"time"

	"golibv2/v2/sdk/dsco"
)

// 说明：
// - 这是 boarding 集成测试，会真实调用 DSCO API。
// - 你只需要改下面 3 个全局变量：token / warehouseCode / skus。
// - 为避免默认 go test 误触发，本文件使用 build tag：dsco_boarding。

var (
	// baseURL 不需要改：默认生产环境；如需 staging 再改。
	baseURL = dsco.BaseURLProd

	// token：DSCO bearer token（务必自行填写）。
	token = "8b283933-2f9e-47e6-b425-ef5eb375ad54"

	// warehouseCode：图 2 中 Warehouse Manager 的 Code 列（例如：YQN-CA2 / COOL-INTELOGICS / YQN-CA）。
	warehouseCode = "YQN-CA"

	// skus：boarding 需要的 3 个 SKU（务必先确保这些 SKU 在 DSCO 侧是有效商品；否则库存接口可能会被过滤/失败）。
	skus = []string{"TEST1", "TEST2", "TEST3"}

	// gtins：与 skus 一一对应。
	//
	// 说明：当 DSCO 认为某个 sku 是“新 SKU”时，会校验必须提供 upc/ean/gtin 至少一个；
	// 本测试统一用 gtin 来满足该校验。
	gtins = []string{"000000000001", "000000000002", "000000000003"}
)

func newClient(t *testing.T) *dsco.Client {
	t.Helper()

	if token == "" {
		t.Fatalf("请先在测试文件里配置 token")
	}
	if warehouseCode == "" {
		t.Fatalf("请先在测试文件里配置 warehouseCode")
	}
	if len(skus) != 3 || skus[0] == "" || skus[1] == "" || skus[2] == "" {
		t.Fatalf("请先在测试文件里配置 3 个 skus")
	}
	if len(gtins) != 3 || gtins[0] == "" || gtins[1] == "" || gtins[2] == "" {
		t.Fatalf("请先在测试文件里配置 3 个 gtins（与 skus 一一对应）")
	}

	cli, err := dsco.New(dsco.Config{
		BaseURL: baseURL,
		Token:   token,
	})
	if err != nil {
		t.Fatalf("dsco.New: %v", err)
	}
	return cli
}

// TestBoarding_Method1_SingleItem_Create3SKUs
//
// 目的：满足 onboarding「Load Items & Inventory」步骤：
// - 使用 POST /inventory/singleItem 创建/更新 3 个 SKU
// - 库存数量设置为 30（>= 20）
func TestBoarding_Method1_SingleItem_Create3SKUs(t *testing.T) {
	cli := newClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	qty := 30
	for i, sku := range skus {
		gtin := gtins[i]
		inv := &dsco.ItemInventory{
			Item: dsco.Item{
				SKU:  sku,
				GTIN: &gtin,
			},
			Status:            "in-stock",
			QuantityAvailable: &qty,
			Warehouses: []dsco.ItemWarehouse{
				{
					Code:     warehouseCode,
					Quantity: &qty,
				},
			},
		}

		resp, err := cli.Inventory.UpsertSingle(ctx, inv)
		if err != nil {
			t.Fatalf("UpsertSingle sku=%s: %v", sku, err)
		}
		if resp == nil || !resp.Success {
			t.Fatalf("UpsertSingle sku=%s: success=%v", sku, resp != nil && resp.Success)
		}
		t.Logf("singleItem OK: sku=%s qty=%d warehouse=%s", sku, qty, warehouseCode)
	}
}

// TestBoarding_Method2_SmallBatch_UpdateTo262728
//
// 目的：满足 onboarding「Update Inventory」步骤：
// - 使用 POST /inventory/batch/small 将上面 3 个 SKU 的库存更新为 50/50/50
// - 该接口为异步：只校验 requestId 返回，具体处理结果需后续通过 inventory change log / streams 追踪
func TestBoarding_Method2_SmallBatch_UpdateTo50(t *testing.T) {
	cli := newClient(t)
	qtys := []int{50, 50, 50}

	items := make([]dsco.ItemInventory, 0, 3)
	for i, sku := range skus {
		qty := qtys[i]
		gtin := gtins[i]
		items = append(items, dsco.ItemInventory{
			Item: dsco.Item{
				SKU:  sku,
				GTIN: &gtin,
			},
			Status:            "in-stock",
			QuantityAvailable: &qty,
			Warehouses: []dsco.ItemWarehouse{
				{
					Code:     warehouseCode,
					Quantity: &qty,
				},
			},
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := cli.Inventory.UpdateSmallBatch(ctx, items, dsco.InventoryUpdateSmallBatchQuery{
		SkipItemsThatDontExist: nil,
	})
	if err != nil {
		t.Fatalf("UpdateSmallBatch: %v", err)
	}
	if resp == nil || resp.RequestID == "" {
		t.Fatalf("UpdateSmallBatch: requestId 为空，resp=%+v", resp)
	}
	t.Logf("small batch accepted: status=%s requestId=%s", resp.Status, resp.RequestID)
}
