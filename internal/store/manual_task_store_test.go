package store

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestManualTaskStore_Create_Behavior(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	s, err := NewManualTaskStore(db)
	if err != nil {
		t.Fatalf("NewManualTaskStore err=%v", err)
	}

	mock.ExpectExec("INSERT INTO manual_task").
		WithArgs("bad_payload", sqlmock.AnyArg(), []byte(`{"k":"v"}`)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.Create(ctx, ManualTask{TaskType: "bad_payload", DscoOrderID: "d1", Payload: []byte(`{"k":"v"}`)}); err != nil {
		t.Fatalf("Create err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestManualTaskStore_ListByStatus_Behavior(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New err=%v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	s, err := NewManualTaskStore(db)
	if err != nil {
		t.Fatalf("NewManualTaskStore err=%v", err)
	}

	now := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	mock.ExpectQuery("SELECT id, task_type, dsco_order_id, payload, status, created_at, updated_at").
		WithArgs(0, 50, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "task_type", "dsco_order_id", "payload", "status", "created_at", "updated_at"}).
			AddRow(int64(1), "bad_payload", sql.NullString{String: "d1", Valid: true}, []byte(`{"a":1}`), 0, now, now))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	rows, err := s.ListByStatus(ctx, 0, 50, 0)
	if err != nil {
		t.Fatalf("ListByStatus err=%v", err)
	}
	if len(rows) != 1 || rows[0].ID != 1 || rows[0].TaskType != "bad_payload" || rows[0].DscoOrderID.String != "d1" {
		t.Fatalf("rows=%+v", rows)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}
