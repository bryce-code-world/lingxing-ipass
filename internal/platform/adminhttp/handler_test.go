package adminhttp

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"lingxingipass/internal/store"
)

func TestHandler_WatermarkSetAndGet_Behavior(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	wm, err := store.NewWatermarkStore(db)
	if err != nil {
		t.Fatalf("NewWatermarkStore err=%v", err)
	}
	manual, err := store.NewManualTaskStore(db)
	if err != nil {
		t.Fatalf("NewManualTaskStore err=%v", err)
	}

	h, err := NewHandler(wm, manual, nil)
	if err != nil {
		t.Fatalf("NewHandler err=%v", err)
	}

	mock.ExpectExec("INSERT INTO job_watermark").
		WithArgs("ack_to_dsco", []byte(`{"mode":"update_time","since":0}`)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := httptest.NewRequest(http.MethodPost, "/admin/watermark/set?job=ack_to_dsco", bytes.NewBufferString(`{"mode":"update_time","since":0}`))
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}

	mock.ExpectQuery("SELECT watermark FROM job_watermark").
		WithArgs("ack_to_dsco").
		WillReturnRows(sqlmock.NewRows([]string{"watermark"}).AddRow([]byte(`{"mode":"update_time","since":0}`)))

	req = httptest.NewRequest(http.MethodGet, "/admin/watermark/get?job=ack_to_dsco", nil)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}
	if got := rr.Body.String(); got != `{"mode":"update_time","since":0}` {
		t.Fatalf("got=%s", got)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestHandler_RunJob_Behavior(t *testing.T) {
	t.Parallel()

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	wm, _ := store.NewWatermarkStore(db)
	manual, _ := store.NewManualTaskStore(db)

	var called bool
	h, err := NewHandler(wm, manual, map[string]JobRunner{
		"heartbeat": func(ctx context.Context) error {
			called = true
			return nil
		},
	})
	if err != nil {
		t.Fatalf("NewHandler err=%v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/admin/run?job=heartbeat", nil)
	req = req.WithContext(context.Background())
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}
	if !called {
		t.Fatalf("called=%v want true", called)
	}
}

func TestHandler_ManualTasks_Behavior(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	wm, _ := store.NewWatermarkStore(db)
	manual, _ := store.NewManualTaskStore(db)

	h, err := NewHandler(wm, manual, nil)
	if err != nil {
		t.Fatalf("NewHandler err=%v", err)
	}

	now := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	mock.ExpectQuery("SELECT id, task_type, dsco_order_id, payload, status, created_at, updated_at").
		WithArgs(0, 50, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "task_type", "dsco_order_id", "payload", "status", "created_at", "updated_at"}).
			AddRow(int64(1), "bad_payload", "d1", []byte(`{"a":1}`), 0, now, now))

	req := httptest.NewRequest(http.MethodGet, "/admin/manual_tasks?status=0&limit=50&offset=0", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("code=%d body=%s", rr.Code, rr.Body.String())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}
