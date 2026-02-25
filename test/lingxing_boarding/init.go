package boarding

import (
	"testing"
	"time"

	"lingxingipass/golib/v2/sdk/lingxing"
)

const (
	envEnable = 1
	envAppID  = "ak_Grz0cgq6BIM1G"
	envSecret = "Amn/Cms7r8Q8W+9PkE8N0A=="
)

// newClient 用于“真实环境手动测试”：默认跳过，避免 go test ./... 时误触发真实接口调用。
func newClient(t *testing.T) *lingxing.Client {
	if envEnable != 1 {
		t.Skipf("skip boarding tests; set %d=1 to enable", envEnable)
	}
	if envAppID == "" || envSecret == "" {
		t.Skipf("skip boarding tests; set %s and %s", envAppID, envSecret)
	}

	cli, err := lingxing.New(lingxing.Config{
		AppID:       envAppID,
		AppSecret:   envSecret,
		AutoToken:   true,
		TokenLeeway: 30 * time.Second,
		Now:         time.Now,
	})
	if err != nil {
		t.Fatalf("lingxing.New() err=%v", err)
	}
	return cli
}
