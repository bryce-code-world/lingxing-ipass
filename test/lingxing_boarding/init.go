package boarding

import (
	"os"
	"testing"
	"time"

	"example.com/lingxing/golib/v2/sdk/lingxing"
)

const (
	envEnable = "LINGXING_BOARDING_ENABLE"
	envAppID  = "LINGXING_APP_ID"
	envSecret = "LINGXING_APP_SECRET"
)

// newClient 用于“真实环境手动测试”：默认跳过，避免 go test ./... 时误触发真实接口调用。
func newClient(t *testing.T) *lingxing.Client {
	if os.Getenv(envEnable) != "1" {
		t.Skipf("skip boarding tests; set %s=1 to enable", envEnable)
	}
	appID := os.Getenv(envAppID)
	appSecret := os.Getenv(envSecret)
	if appID == "" || appSecret == "" {
		t.Skipf("skip boarding tests; set %s and %s", envAppID, envSecret)
	}

	cli, err := lingxing.New(lingxing.Config{
		AppID:       appID,
		AppSecret:   appSecret,
		AutoToken:   true,
		TokenLeeway: 30 * time.Second,
		Now:         time.Now,
	})
	if err != nil {
		t.Fatalf("lingxing.New() err=%v", err)
	}
	return cli
}
