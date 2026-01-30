package boarding

import (
	"context"
	"testing"

	"example.com/lingxing/golib/v2/sdk/lingxing"
)

func TestGetWarehouseLists(t *testing.T) {
	cli := newClient(t)
	list, total, raw, err := cli.Warehouse.WarehouseListsWithRawBody(
		context.Background(),
		lingxing.WarehouseListsRequest{
			Type: 3,
			// Offset/Length 分页参数。
			Offset: 0,
			Length: 20,
		})
	if err != nil {
		t.Fatalf("WarehouseListWithRawBody() err=%v", err)
	}
	t.Logf("WarehouseListWithRawBody() total=%d list=%+v", total, list)
	t.Logf("WarehouseListWithRawBody() raw_response=%+v", raw)
}
