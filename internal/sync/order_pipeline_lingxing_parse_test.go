package sync

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"example.com/lingxing/golib/v2/sdk/lingxing"
)

func TestOrderPipeline_QueryLingXingGlobalOrderNo_Behavior(t *testing.T) {
	t.Parallel()

	now := func() time.Time { return time.Unix(1720429074, 0) }
	lxSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method=%s want=%s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/pb/mp/order/v2/list" {
			t.Fatalf("path=%s want=%s", r.URL.Path, "/pb/mp/order/v2/list")
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"code":0,"message":"success","data":{"total":1,"list":[{"global_order_no":"g1","platform_info":[{"platform_order_no":"x"},{"platform_order_no":"d1"}]}]}}`)
	}))
	t.Cleanup(lxSrv.Close)

	lxCli, err := lingxing.New(lingxing.Config{
		BaseURL:     lxSrv.URL,
		AppID:       "1234567890abcdef",
		AccessToken: "tok",
		Now:         now,
	})
	if err != nil {
		t.Fatalf("lingxing.New err=%v", err)
	}

	p := &OrderPipeline{lingxingCli: lxCli}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	t.Cleanup(cancel)

	got, ok, err := p.queryLingXingGlobalOrderNo(ctx, 10009, "s1", "d1")
	if err != nil {
		t.Fatalf("queryLingXingGlobalOrderNo err=%v", err)
	}
	if !ok || got != "g1" {
		t.Fatalf("got=%q ok=%v want got=%q ok=true", got, ok, "g1")
	}
}
