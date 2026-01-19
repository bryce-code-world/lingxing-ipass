package store

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestWatermarkStore_Get_NotFound(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewWatermarkStore(gdb)
	if err != nil {
		t.Fatalf("NewWatermarkStore err=%v", err)
	}

	mock.ExpectQuery("SELECT watermark FROM job_watermark").
		WithArgs("pull_dsco_orders").
		WillReturnRows(sqlmock.NewRows([]string{"watermark"}))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	raw, ok, err := s.Get(ctx, "pull_dsco_orders")
	if err != nil {
		t.Fatalf("Get err=%v", err)
	}
	if ok || raw != nil {
		t.Fatalf("ok=%v raw=%s want ok=false raw=nil", ok, string(raw))
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestWatermarkStore_Set_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewWatermarkStore(gdb)
	if err != nil {
		t.Fatalf("NewWatermarkStore err=%v", err)
	}

	mock.ExpectExec("INSERT INTO job_watermark").
		WithArgs("pull_dsco_orders", []byte(`{"since":"a"}`)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.Set(ctx, "pull_dsco_orders", []byte(`{"since":"a"}`)); err != nil {
		t.Fatalf("Set err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestWatermarkStore_ListAll_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewWatermarkStore(gdb)
	if err != nil {
		t.Fatalf("NewWatermarkStore err=%v", err)
	}

	now := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	mock.ExpectQuery("SELECT job_name, watermark, updated_at FROM job_watermark").
		WillReturnRows(sqlmock.NewRows([]string{"job_name", "watermark", "updated_at"}).
			AddRow("ack_to_dsco", []byte(`{"mode":"update_time","since":0}`), now).
			AddRow("pull_dsco_orders", []byte(`{"mode":"updatedSince","since":"1970-01-01T00:00:00Z"}`), now))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	rows, err := s.ListAll(ctx)
	if err != nil {
		t.Fatalf("ListAll err=%v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("len=%d want=2", len(rows))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}
