package store

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestOrderStateStore_UpsertOrderIDs_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)
	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("INSERT INTO sync_order_state").
		WithArgs("d1", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO sync_order_state").
		WithArgs("d2", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.UpsertOrderIDs(ctx, []string{"d1", "d2"}); err != nil {
		t.Fatalf("UpsertOrderIDs err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_ClaimForPush_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)
	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT dsco_order_id").
		WithArgs(sqlmock.AnyArg(), 2).
		WillReturnRows(sqlmock.NewRows([]string{"dsco_order_id"}).AddRow("d1").AddRow("d2"))
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1", "d2").
		WillReturnResult(sqlmock.NewResult(0, 2))
	mock.ExpectCommit()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	ids, err := s.ClaimForPush(ctx, 2)
	if err != nil {
		t.Fatalf("ClaimForPush err=%v", err)
	}
	if len(ids) != 2 || ids[0] != "d1" || ids[1] != "d2" {
		t.Fatalf("ids=%v", ids)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkPushSuccess_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)
	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), "g1", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkPushSuccess(ctx, "d1", "g1"); err != nil {
		t.Fatalf("MarkPushSuccess err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkPushFailure_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)
	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs("boom", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkPushFailure(ctx, "d1", "boom"); err != nil {
		t.Fatalf("MarkPushFailure err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkPushManual_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)
	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs("bad_payload", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkPushManual(ctx, "d1", "bad_payload"); err != nil {
		t.Fatalf("MarkPushManual err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_TryClaimAck_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)
	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	ok, err := s.TryClaimAck(ctx, "d1")
	if err != nil {
		t.Fatalf("TryClaimAck err=%v", err)
	}
	if !ok {
		t.Fatalf("ok=%v want true", ok)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkAckSuccess_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)
	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkAckSuccess(ctx, "d1"); err != nil {
		t.Fatalf("MarkAckSuccess err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkAckFailure_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)
	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs("boom", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkAckFailure(ctx, "d1", "boom"); err != nil {
		t.Fatalf("MarkAckFailure err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkAckManual_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs("bad_payload", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkAckManual(ctx, "d1", "bad_payload"); err != nil {
		t.Fatalf("MarkAckManual err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_TryClaimShipment_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	ok, err := s.TryClaimShipment(ctx, "d1")
	if err != nil {
		t.Fatalf("TryClaimShipment err=%v", err)
	}
	if !ok {
		t.Fatalf("ok=%v want true", ok)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkShipmentSuccess_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), "TN1", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkShipmentSuccess(ctx, "d1", "TN1"); err != nil {
		t.Fatalf("MarkShipmentSuccess err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkShipmentFailure_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs("boom", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkShipmentFailure(ctx, "d1", "boom"); err != nil {
		t.Fatalf("MarkShipmentFailure err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkShipmentManual_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs("multi_shipment", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkShipmentManual(ctx, "d1", "multi_shipment"); err != nil {
		t.Fatalf("MarkShipmentManual err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_ClaimForInvoice_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT dsco_order_id").
		WithArgs(sqlmock.AnyArg(), 2).
		WillReturnRows(sqlmock.NewRows([]string{"dsco_order_id"}).AddRow("d1").AddRow("d2"))
	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "d1", "d2").
		WillReturnResult(sqlmock.NewResult(0, 2))
	mock.ExpectCommit()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	ids, err := s.ClaimForInvoice(ctx, 2)
	if err != nil {
		t.Fatalf("ClaimForInvoice err=%v", err)
	}
	if len(ids) != 2 || ids[0] != "d1" || ids[1] != "d2" {
		t.Fatalf("ids=%v", ids)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkInvoiceSuccess_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs(sqlmock.AnyArg(), "INV-d1", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkInvoiceSuccess(ctx, "d1", "INV-d1"); err != nil {
		t.Fatalf("MarkInvoiceSuccess err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkInvoiceFailure_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs("boom", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkInvoiceFailure(ctx, "d1", "boom"); err != nil {
		t.Fatalf("MarkInvoiceFailure err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}

func TestOrderStateStore_MarkInvoiceManual_Behavior(t *testing.T) {
	t.Parallel()

	gdb, mock := newMockGormDB(t)

	s, err := NewOrderStateStore(gdb)
	if err != nil {
		t.Fatalf("NewOrderStateStore err=%v", err)
	}

	mock.ExpectExec("UPDATE sync_order_state").
		WithArgs("bad_payload", sqlmock.AnyArg(), "d1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(cancel)

	if err := s.MarkInvoiceManual(ctx, "d1", "bad_payload"); err != nil {
		t.Fatalf("MarkInvoiceManual err=%v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations err=%v", err)
	}
}
