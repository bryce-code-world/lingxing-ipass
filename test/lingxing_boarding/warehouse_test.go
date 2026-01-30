package boarding

import (
	"context"
	"testing"

	"example.com/lingxing/golib/v2/sdk/lingxing"
)

// 查询仓库列表
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

// 查询仓库关联的物流方式列表
func TestGetWarehouseLogisticsMethods(t *testing.T) {
	cli := newClient(t)
	list, total, raw, err := cli.Warehouse.ListUsedLogisticsTypeWithRawBody(
		context.Background(),
		lingxing.ListUsedLogisticsTypeRequest{
			Param: lingxing.ListUsedLogisticsTypeParam{
				ProviderType: 2, // 海外仓物流
				Page:         1,
				Length:       50,
			},
		})
	if err != nil {
		t.Fatalf("WarehouseLogisticsMethodsWithRawBody() err=%v", err)
	}
	t.Logf("WarehouseLogisticsMethodsWithRawBody() total=%d list=%+v", total, list)
	t.Logf("WarehouseLogisticsMethodsWithRawBody() raw_response=%+v", raw)
}
