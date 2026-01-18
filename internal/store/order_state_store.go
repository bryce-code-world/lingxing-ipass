package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

// OrderStateStore 负责 sync_order_state 的读写。
type OrderStateStore struct {
	db *sql.DB
}

func NewOrderStateStore(db *sql.DB) (*OrderStateStore, error) {
	if db == nil {
		return nil, errors.New("db 不能为空")
	}
	return &OrderStateStore{db: db}, nil
}

// UpsertOrderIDs 将 dscoOrderId 写入状态表（一期只落最小信息）。
func (s *OrderStateStore) UpsertOrderIDs(ctx context.Context, dscoOrderIDs []string) error {
	if len(dscoOrderIDs) == 0 {
		return nil
	}

	now := time.Now().UTC()
	for _, id := range dscoOrderIDs {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		// 只插入最小记录；若已存在，不覆盖状态字段。
		_, err := s.db.ExecContext(ctx, `
INSERT INTO sync_order_state (dsco_order_id, created_at, updated_at)
VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at)
`, id, now, now)
		if err != nil {
			return err
		}
	}
	return nil
}

// ClaimForPush 抢占一批“待推单”的 dscoOrderId，避免多实例重复处理。
//
// pushed_to_lx_status 约定：
// - 0/2：候选（未处理/失败可重试）
// - 9：处理中（抢占成功后置为 9）
func (s *OrderStateStore) ClaimForPush(ctx context.Context, limit int) ([]string, error) {
	if limit <= 0 {
		return nil, errors.New("limit 必须大于 0")
	}

	now := time.Now().UTC()
	// 进程崩溃/重启可能留下“处理中”的脏锁，这里按时间兜底回收。
	staleBefore := now.Add(-30 * time.Minute)

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	rows, err := tx.QueryContext(ctx, `
SELECT dsco_order_id
FROM sync_order_state
WHERE pushed_to_lx_status IN (0, 2)
   OR (pushed_to_lx_status = 9 AND last_attempt_at < ?)
ORDER BY updated_at ASC
LIMIT ?
FOR UPDATE
`, staleBefore, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		if err := tx.Commit(); err != nil {
			return nil, err
		}
		return nil, nil
	}

	placeholders := make([]string, 0, len(ids))
	args := make([]any, 0, 2+len(ids))
	args = append(args, now, now)
	for _, id := range ids {
		placeholders = append(placeholders, "?")
		args = append(args, id)
	}

	// 注意：只把当前选中的这些记录置为处理中；retry_count 用于监控与兜底策略。
	q := fmt.Sprintf(`
UPDATE sync_order_state
SET pushed_to_lx_status = 9,
    last_attempt_at = ?,
    retry_count = retry_count + 1,
    updated_at = ?
WHERE dsco_order_id IN (%s)
`, strings.Join(placeholders, ","))

	if _, err := tx.ExecContext(ctx, q, args...); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return ids, nil
}

// MarkPushSuccess 标记推单成功，并写入领星 global_order_no。
func (s *OrderStateStore) MarkPushSuccess(ctx context.Context, dscoOrderID, globalOrderNo string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	globalOrderNo = strings.TrimSpace(globalOrderNo)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if globalOrderNo == "" {
		return errors.New("globalOrderNo 不能为空")
	}

	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET pushed_to_lx_status = 1,
    pushed_to_lx_at = ?,
    lingxing_global_order_no = ?,
    last_error = NULL,
    updated_at = ?
WHERE dsco_order_id = ?
`, now, globalOrderNo, now, dscoOrderID)
	return err
}

// MarkPushFailure 标记推单失败（可重试）。
func (s *OrderStateStore) MarkPushFailure(ctx context.Context, dscoOrderID, errMsg string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	errMsg = strings.TrimSpace(errMsg)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if errMsg == "" {
		errMsg = "unknown error"
	}
	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET pushed_to_lx_status = 2,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, errMsg, now, dscoOrderID)
	return err
}

// MarkPushManual 标记为人工处理（不再自动重试）。
func (s *OrderStateStore) MarkPushManual(ctx context.Context, dscoOrderID, reason string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	reason = strings.TrimSpace(reason)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if reason == "" {
		reason = "manual"
	}
	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET pushed_to_lx_status = 3,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, reason, now, dscoOrderID)
	return err
}

// TryClaimAck 抢占“待 ACK 回传”的订单，避免多实例重复处理。
//
// acked_to_dsco_status 约定：
// - 0/2：候选（未处理/失败可重试）
// - 9：处理中（抢占成功后置为 9）
func (s *OrderStateStore) TryClaimAck(ctx context.Context, dscoOrderID string) (bool, error) {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	if dscoOrderID == "" {
		return false, errors.New("dscoOrderID 不能为空")
	}
	now := time.Now().UTC()
	res, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET acked_to_dsco_status = 9,
    last_attempt_at = ?,
    retry_count = retry_count + 1,
    updated_at = ?
WHERE dsco_order_id = ?
  AND acked_to_dsco_status IN (0, 2)
`, now, now, dscoOrderID)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}

func (s *OrderStateStore) MarkAckSuccess(ctx context.Context, dscoOrderID string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET acked_to_dsco_status = 1,
    acked_to_dsco_at = ?,
    last_error = NULL,
    updated_at = ?
WHERE dsco_order_id = ?
`, now, now, dscoOrderID)
	return err
}

func (s *OrderStateStore) MarkAckFailure(ctx context.Context, dscoOrderID, errMsg string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	errMsg = strings.TrimSpace(errMsg)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if errMsg == "" {
		errMsg = "unknown error"
	}
	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET acked_to_dsco_status = 2,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, errMsg, now, dscoOrderID)
	return err
}

func (s *OrderStateStore) MarkAckManual(ctx context.Context, dscoOrderID, reason string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	reason = strings.TrimSpace(reason)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if reason == "" {
		reason = "manual"
	}
	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET acked_to_dsco_status = 3,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, reason, now, dscoOrderID)
	return err
}

// TryClaimShipment 抢占“待发货回传”的订单，避免多实例重复处理。
//
// shipped_to_dsco_status 约定：
// - 0/2：候选（未处理/失败可重试）
// - 9：处理中（抢占成功后置为 9）
func (s *OrderStateStore) TryClaimShipment(ctx context.Context, dscoOrderID string) (bool, error) {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	if dscoOrderID == "" {
		return false, errors.New("dscoOrderID 不能为空")
	}
	now := time.Now().UTC()
	res, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET shipped_to_dsco_status = 9,
    last_attempt_at = ?,
    retry_count = retry_count + 1,
    updated_at = ?
WHERE dsco_order_id = ?
  AND shipped_to_dsco_status IN (0, 2)
`, now, now, dscoOrderID)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}

func (s *OrderStateStore) MarkShipmentSuccess(ctx context.Context, dscoOrderID, trackingNo string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	trackingNo = strings.TrimSpace(trackingNo)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if trackingNo == "" {
		return errors.New("trackingNo 不能为空")
	}
	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET shipped_to_dsco_status = 1,
    shipped_to_dsco_at = ?,
    shipped_tracking_no = ?,
    last_error = NULL,
    updated_at = ?
WHERE dsco_order_id = ?
`, now, trackingNo, now, dscoOrderID)
	return err
}

func (s *OrderStateStore) MarkShipmentFailure(ctx context.Context, dscoOrderID, errMsg string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	errMsg = strings.TrimSpace(errMsg)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if errMsg == "" {
		errMsg = "unknown error"
	}
	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET shipped_to_dsco_status = 2,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, errMsg, now, dscoOrderID)
	return err
}

func (s *OrderStateStore) MarkShipmentManual(ctx context.Context, dscoOrderID, reason string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	reason = strings.TrimSpace(reason)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if reason == "" {
		reason = "manual"
	}
	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET shipped_to_dsco_status = 3,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, reason, now, dscoOrderID)
	return err
}

// ClaimForInvoice 抢占一批“待回传发票”的 dscoOrderId，避免多实例重复处理。
//
// invoiced_to_dsco_status 约定：
// - 0/2：候选（未处理/失败可重试）
// - 9：处理中（抢占成功后置为 9）
func (s *OrderStateStore) ClaimForInvoice(ctx context.Context, limit int) ([]string, error) {
	if limit <= 0 {
		return nil, errors.New("limit 必须大于 0")
	}

	now := time.Now().UTC()
	staleBefore := now.Add(-30 * time.Minute)

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	rows, err := tx.QueryContext(ctx, `
SELECT dsco_order_id
FROM sync_order_state
WHERE invoiced_to_dsco_status IN (0, 2)
   OR (invoiced_to_dsco_status = 9 AND last_attempt_at < ?)
ORDER BY updated_at ASC
LIMIT ?
FOR UPDATE
`, staleBefore, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		if err := tx.Commit(); err != nil {
			return nil, err
		}
		return nil, nil
	}

	placeholders := make([]string, 0, len(ids))
	args := make([]any, 0, 2+len(ids))
	args = append(args, now, now)
	for _, id := range ids {
		placeholders = append(placeholders, "?")
		args = append(args, id)
	}

	q := fmt.Sprintf(`
UPDATE sync_order_state
SET invoiced_to_dsco_status = 9,
    last_attempt_at = ?,
    retry_count = retry_count + 1,
    updated_at = ?
WHERE dsco_order_id IN (%s)
`, strings.Join(placeholders, ","))

	if _, err := tx.ExecContext(ctx, q, args...); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return ids, nil
}

func (s *OrderStateStore) MarkInvoiceSuccess(ctx context.Context, dscoOrderID, invoiceID string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	invoiceID = strings.TrimSpace(invoiceID)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if invoiceID == "" {
		return errors.New("invoiceID 不能为空")
	}
	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET invoiced_to_dsco_status = 1,
    invoiced_to_dsco_at = ?,
    dsco_invoice_id = ?,
    last_error = NULL,
    updated_at = ?
WHERE dsco_order_id = ?
`, now, invoiceID, now, dscoOrderID)
	return err
}

func (s *OrderStateStore) MarkInvoiceFailure(ctx context.Context, dscoOrderID, errMsg string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	errMsg = strings.TrimSpace(errMsg)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if errMsg == "" {
		errMsg = "unknown error"
	}
	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET invoiced_to_dsco_status = 2,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, errMsg, now, dscoOrderID)
	return err
}

func (s *OrderStateStore) MarkInvoiceManual(ctx context.Context, dscoOrderID, reason string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	reason = strings.TrimSpace(reason)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if reason == "" {
		reason = "manual"
	}
	now := time.Now().UTC()
	_, err := s.db.ExecContext(ctx, `
UPDATE sync_order_state
SET invoiced_to_dsco_status = 3,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, reason, now, dscoOrderID)
	return err
}
