package adminweb

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"lingxingipass/internal/store"
)

func TestAdminWeb_UI_LoginAndSessionCookie_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	wm, _ := store.NewWatermarkStore(gdb)
	manual, _ := store.NewManualTaskStore(gdb)
	order, _ := store.NewOrderStateStore(gdb)

	// dashboard 会查询水位列表（空也要返回 200）。
	mock.ExpectQuery("SELECT job_name, watermark, updated_at FROM job_watermark").
		WillReturnRows(sqlmock.NewRows([]string{"job_name", "watermark", "updated_at"}))

	h := NewServer(Options{
		AdminPassword: "p1",
		Watermark:     wm,
		Manual:        manual,
		Order:         order,
		Runners:       map[string]JobRunner{},
		Now:           func() time.Time { return time.Unix(1720429074, 0).UTC() },
	})

	// 未登录访问后台，应跳转到 login。
	req := httptest.NewRequest(http.MethodGet, "/admin/ui/", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusFound {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}

	// 登录成功，设置 session cookie。
	req = httptest.NewRequest(http.MethodPost, "/admin/ui/login", bytes.NewBufferString("password=p1"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusFound {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}
	setCookie := rr.Header().Get("Set-Cookie")
	if setCookie == "" {
		t.Fatalf("missing Set-Cookie")
	}

	// 带 cookie 再访问 dashboard，应返回 200。
	req = httptest.NewRequest(http.MethodGet, "/admin/ui/", nil)
	req.Header.Set("Cookie", setCookie)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestAdminWeb_API_AuthByHeader_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	wm, _ := store.NewWatermarkStore(gdb)
	manual, _ := store.NewManualTaskStore(gdb)
	order, _ := store.NewOrderStateStore(gdb)

	h := NewServer(Options{
		AdminPassword: "p1",
		Watermark:     wm,
		Manual:        manual,
		Order:         order,
		Runners:       map[string]JobRunner{},
		Now:           func() time.Time { return time.Unix(1720429074, 0).UTC() },
	})

	// 未认证调用 API，应 401。
	req := httptest.NewRequest(http.MethodGet, "/admin/watermark/get?job=ack_to_dsco", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}

	// 通过 header 密码访问 API。
	mock.ExpectQuery("SELECT watermark FROM job_watermark").
		WithArgs("ack_to_dsco").
		WillReturnRows(sqlmock.NewRows([]string{"watermark"}).AddRow([]byte(`{"mode":"update_time","since":0}`)))

	req = httptest.NewRequest(http.MethodGet, "/admin/watermark/get?job=ack_to_dsco", nil)
	req.Header.Set("X-Admin-Password", "p1")
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

