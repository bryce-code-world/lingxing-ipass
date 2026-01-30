package boarding

import (
	"context"
	"testing"

	"example.com/lingxing/golib/v2/sdk/lingxing"
)

func TestGetInventoryList(t *testing.T) {
	cli := newClient(t)
	_, _, raw, err := cli.Inventory.InventoryDetailsWithRawBody(
		context.Background(),
		lingxing.InventoryDetailsRequest{
			// WID 仓库 id，多个可用英文逗号分隔；不传则按领星默认范围查询。
			WID: "33328",
			// Offset/Length 分页参数。
			Offset: 0,
			Length: 5,
			// SKU 可选，按 SKU 过滤（文档说明支持模糊搜索）。
			SKU: "",
		})
	if err != nil {
		t.Fatalf("InventoryDetailsWithRawBody() err=%v", err)
	}
	// t.Logf("InventoryDetailsWithRawBody() total=%d list=%+v", total, list)
	t.Logf("InventoryDetailsWithRawBody() raw_response=%+v", raw)
}
