package store

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestWatermarkStore_Get_NotFound(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	s, err := NewWatermarkStore(db)
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

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	s, err := NewWatermarkStore(db)
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
