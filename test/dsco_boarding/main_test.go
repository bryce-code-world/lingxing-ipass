package dsco_boarding

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 说明：
	// - test/dsco_boarding 属于“真实接口手动测试”，会真实调用 DSCO API。
	// - 为避免 go test ./... 误触发，默认跳过；需要手动显式开启。
	if os.Getenv("DSCO_BOARDING") != "1" {
		fmt.Fprintln(os.Stderr, "skip DSCO boarding tests (set DSCO_BOARDING=1 to enable)")
		os.Exit(0)
	}
	os.Exit(m.Run())
}
