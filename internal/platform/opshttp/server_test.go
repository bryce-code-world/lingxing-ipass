package opshttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWrap_Behavior(t *testing.T) {
	t.Parallel()

	base := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	})

	h := Wrap(base, "p1")

	// healthz 放行
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}

	// 未带密码拒绝
	req = httptest.NewRequest(http.MethodPost, "/admin/run?job=heartbeat", nil)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}

	// 正确密码放行
	req = httptest.NewRequest(http.MethodPost, "/admin/run?job=heartbeat", nil)
	req.Header.Set("X-Ops-Password", "p1")
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}
}
