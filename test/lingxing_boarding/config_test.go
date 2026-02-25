package boarding

import (
	"context"
	"fmt"
	"testing"

	"lingxingipass/golib/v2/sdk/lingxing"
)

// 查询多平台 SKU 配对列表（原始响应）。
//
// go test -v -count=1 ./init.go ./config_test.go -run TestConfigService_GetPairListV2WithRawBody
func TestConfigService_GetPairListV2WithRawBody(t *testing.T) {
	cli := newClient(t)

	out, raw, err := cli.Config.GetPairListV2WithRawBody(context.Background(), lingxing.PairListV2Request{
		Offset: 1,
		Length: 20,
		// 可按需指定店铺与平台过滤，避免全量查询。
		PlatformCodes: []string{"10009"},
		StoreIDs:      []string{"110658143132021760"},
		// MSKU:           []string{"TEST3"},
		// SKU: []string{"FFDD00101YW2L"},
		// StartTime:      "2024-01-01 00:00:00",
		// EndTime:        "2024-01-31 23:59:59",
	})
	if err != nil {
		t.Fatalf("GetPairListV2WithRawBody() err=%v", err)
	}

	fmt.Println("GetPairListV2 resp:", out)
	fmt.Println("GetPairListV2 raw resp:", raw)
}
