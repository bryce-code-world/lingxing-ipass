package boarding

import (
	"context"
	"testing"

	"example.com/lingxing/golib/v2/sdk/lingxing"
)

// 查询仓库库存
//
// go test -v -count=1 ./init.go ./inventory_test.go -run TestGetInventoryList
func TestGetInventoryList(t *testing.T) {
	cli := newClient(t)
	_, total, raw, err := cli.Inventory.InventoryDetailsWithRawBody(
		context.Background(),
		lingxing.InventoryDetailsRequest{
			// WID 仓库 id，多个可用英文逗号分隔；不传则按领星默认范围查询。
			// WID: "33328",
			// WID: "28393",
			WID: "28393,28394,33328",
			// Offset/Length 分页参数。
			Offset: 0,
			Length: 20,
			// SKU 可选，按 SKU 过滤（文档说明支持模糊搜索）。
			SKU: "FFDD00101YW2L",
		})
	if err != nil {
		t.Fatalf("InventoryDetailsWithRawBody() err=%v", err)
	}
	// t.Logf("InventoryDetailsWithRawBody() total=%d list=%+v", total, list)
	t.Logf("InventoryDetailsWithRawBody() total=%d raw_response=%+v", total, raw)
}
