package store

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDscoOrderRawStore_UpsertLatest_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewDscoOrderRawStore(gdb)
	if err != nil {
		t.Fatalf("NewDscoOrderRawStore err=%v", err)
	}

	mock.ExpectExec("INSERT INTO dsco_order_raw").
		WithArgs("d1", []byte(`{"dscoOrderId":"d1"}`), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.UpsertLatest(ctx, "d1", []byte(`{"dscoOrderId":"d1"}`), time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatalf("UpsertLatest err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}
