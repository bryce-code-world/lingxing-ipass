package boarding

import (
	"context"
	"fmt"
	"testing"

	"gitee.com/lsy007/golibv2/v2/sdk/lingxing"
)

func TestShopService_MultiPlatformStoreListV2_RealEnv(t *testing.T) {
	cli := newClient(t)
	out, err := cli.Shop.MultiPlatformStoreListV2(context.Background(), lingxing.MultiPlatformStoreListV2Request{
		Offset:       0,
		Length:       200,
		PlatformCode: []int{10009},
	})
	if err != nil {
		t.Fatalf("MultiPlatformStoreListV2() err=%v", err)
	}

	total, _ := out.TotalInt()
	fmt.Println("platform_code:", []int{10009}, "total:", total, "list_len:", len(out.List))
	for i := 0; i < len(out.List) && i < 20; i++ {
		fmt.Printf("store[%d]: store_id=%s store_name=%s platform_code=%s platform_name=%s currency=%s is_sync=%d status=%d\n",
			i,
			out.List[i].StoreID,
			out.List[i].StoreName,
			out.List[i].PlatformCode,
			out.List[i].PlatformName,
			out.List[i].Currency,
			out.List[i].IsSync,
			out.List[i].Status,
		)
	}

	// store[0]: store_id=110655614355155968 store_name=FSI platform_code=10009 platform_name=自定义平台 currency=USD is_sync=1 status=1
	// store[1]: store_id=110658143132021760 store_name=Chewy platform_code=10009 platform_name=自定义平台 currency=USD is_sync=1 status=1
	// store[2]: store_id=110661999149898752 store_name=Petco platform_code=10009 platform_name=自定义平台 currency=USD is_sync=1 status=1
	// store[3]: store_id=110661999252100096 store_name=Tractor Supply platform_code=10009 platform_name=自定义平台 currency=USD is_sync=1 status=1
}
