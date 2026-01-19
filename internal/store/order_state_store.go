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

// GetRetryCount 返回某个 dscoOrderId 的 retry_count；若不存在返回 ok=false。
func (s *OrderStateStore) GetRetryCount(ctx context.Context, dscoOrderID string) (retryCount int, ok bool, err error) {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	if dscoOrderID == "" {
		return 0, false, errors.New("dscoOrderID 不能为空")
	}
	var n int
	err = s.db.QueryRowContext(ctx, `SELECT retry_count FROM sync_order_state WHERE dsco_order_id = ?`, dscoOrderID).Scan(&n)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, false, nil
		}
		return 0, false, err
	}
	return n, true, nil
}

// OrderStateRow 表示 sync_order_state 的查询结果（用于管理端展示）。
type OrderStateRow struct {
	DscoOrderID           string     `json:"dsco_order_id"`
	LingxingGlobalOrderNo *string    `json:"lingxing_global_order_no,omitempty"`
	PushedToLXStatus      int        `json:"pushed_to_lx_status"`
	PushedToLXAt          *time.Time `json:"pushed_to_lx_at,omitempty"`
	AckedToDSCOStatus     int        `json:"acked_to_dsco_status"`
	AckedToDSCOAt         *time.Time `json:"acked_to_dsco_at,omitempty"`
	ShippedToDSCOStatus   int        `json:"shipped_to_dsco_status"`
	ShippedToDSCOAt       *time.Time `json:"shipped_to_dsco_at,omitempty"`
	ShippedTrackingNo     *string    `json:"shipped_tracking_no,omitempty"`
	InvoicedToDSCOStatus  int        `json:"invoiced_to_dsco_status"`
	InvoicedToDSCOAt      *time.Time `json:"invoiced_to_dsco_at,omitempty"`
	DscoInvoiceID         *string    `json:"dsco_invoice_id,omitempty"`
	RetryCount            int        `json:"retry_count"`
	LastError             *string    `json:"last_error,omitempty"`
	LastAttemptAt         *time.Time `json:"last_attempt_at,omitempty"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

// GetByDscoOrderID 查询单条订单状态；若不存在返回 ok=false。
func (s *OrderStateStore) GetByDscoOrderID(ctx context.Context, dscoOrderID string) (*OrderStateRow, bool, error) {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	if dscoOrderID == "" {
		return nil, false, errors.New("dscoOrderID 不能为空")
	}

	var (
		lingxingGlobalOrderNo sql.NullString
		pushedToLXAt          sql.NullTime
		ackedToDSCOAt         sql.NullTime
		shippedToDSCOAt       sql.NullTime
		shippedTrackingNo     sql.NullString
		invoicedToDSCOAt      sql.NullTime
		dscoInvoiceID         sql.NullString
		lastError             sql.NullString
		lastAttemptAt         sql.NullTime

		row OrderStateRow
	)
	err := s.db.QueryRowContext(ctx, `
SELECT dsco_order_id,
       lingxing_global_order_no,
       pushed_to_lx_status, pushed_to_lx_at,
       acked_to_dsco_status, acked_to_dsco_at,
       shipped_to_dsco_status, shipped_to_dsco_at, shipped_tracking_no,
       invoiced_to_dsco_status, invoiced_to_dsco_at, dsco_invoice_id,
       retry_count, last_error, last_attempt_at,
       created_at, updated_at
FROM sync_order_state
WHERE dsco_order_id = ?
`, dscoOrderID).Scan(
		&row.DscoOrderID,
		&lingxingGlobalOrderNo,
		&row.PushedToLXStatus, &pushedToLXAt,
		&row.AckedToDSCOStatus, &ackedToDSCOAt,
		&row.ShippedToDSCOStatus, &shippedToDSCOAt, &shippedTrackingNo,
		&row.InvoicedToDSCOStatus, &invoicedToDSCOAt, &dscoInvoiceID,
		&row.RetryCount, &lastError, &lastAttemptAt,
		&row.CreatedAt, &row.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}

	if lingxingGlobalOrderNo.Valid {
		v := strings.TrimSpace(lingxingGlobalOrderNo.String)
		if v != "" {
			row.LingxingGlobalOrderNo = &v
		}
	}
	if pushedToLXAt.Valid {
		v := pushedToLXAt.Time
		row.PushedToLXAt = &v
	}
	if ackedToDSCOAt.Valid {
		v := ackedToDSCOAt.Time
		row.AckedToDSCOAt = &v
	}
	if shippedToDSCOAt.Valid {
		v := shippedToDSCOAt.Time
		row.ShippedToDSCOAt = &v
	}
	if shippedTrackingNo.Valid {
		v := strings.TrimSpace(shippedTrackingNo.String)
		if v != "" {
			row.ShippedTrackingNo = &v
		}
	}
	if invoicedToDSCOAt.Valid {
		v := invoicedToDSCOAt.Time
		row.InvoicedToDSCOAt = &v
	}
	if dscoInvoiceID.Valid {
		v := strings.TrimSpace(dscoInvoiceID.String)
		if v != "" {
			row.DscoInvoiceID = &v
		}
	}
	if lastError.Valid {
		v := strings.TrimSpace(lastError.String)
		if v != "" {
			row.LastError = &v
		}
	}
	if lastAttemptAt.Valid {
		v := lastAttemptAt.Time
		row.LastAttemptAt = &v
	}

	return &row, true, nil
}

type OrderStateListQuery struct {
	PushedToLXStatus     *int
	AckedToDSCOStatus    *int
	ShippedToDSCOStatus  *int
	InvoicedToDSCOStatus *int
	Limit                int
	Offset               int
}

// List 查询订单状态列表（一期最小分页 + 可选按状态过滤）。
func (s *OrderStateStore) List(ctx context.Context, q OrderStateListQuery) ([]OrderStateRow, error) {
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Limit > 200 {
		q.Limit = 200
	}
	if q.Offset < 0 {
		q.Offset = 0
	}

	var sb strings.Builder
	sb.WriteString(`
SELECT dsco_order_id,
       lingxing_global_order_no,
       pushed_to_lx_status, pushed_to_lx_at,
       acked_to_dsco_status, acked_to_dsco_at,
       shipped_to_dsco_status, shipped_to_dsco_at, shipped_tracking_no,
       invoiced_to_dsco_status, invoiced_to_dsco_at, dsco_invoice_id,
       retry_count, last_error, last_attempt_at,
       created_at, updated_at
FROM sync_order_state
WHERE 1=1
`)
	args := make([]any, 0, 8)
	if q.PushedToLXStatus != nil {
		sb.WriteString(" AND pushed_to_lx_status = ?\n")
		args = append(args, *q.PushedToLXStatus)
	}
	if q.AckedToDSCOStatus != nil {
		sb.WriteString(" AND acked_to_dsco_status = ?\n")
		args = append(args, *q.AckedToDSCOStatus)
	}
	if q.ShippedToDSCOStatus != nil {
		sb.WriteString(" AND shipped_to_dsco_status = ?\n")
		args = append(args, *q.ShippedToDSCOStatus)
	}
	if q.InvoicedToDSCOStatus != nil {
		sb.WriteString(" AND invoiced_to_dsco_status = ?\n")
		args = append(args, *q.InvoicedToDSCOStatus)
	}
	sb.WriteString(" ORDER BY updated_at DESC\n LIMIT ? OFFSET ?\n")
	args = append(args, q.Limit, q.Offset)

	rows, err := s.db.QueryContext(ctx, sb.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []OrderStateRow
	for rows.Next() {
		var (
			lingxingGlobalOrderNo sql.NullString
			pushedToLXAt          sql.NullTime
			ackedToDSCOAt         sql.NullTime
			shippedToDSCOAt       sql.NullTime
			shippedTrackingNo     sql.NullString
			invoicedToDSCOAt      sql.NullTime
			dscoInvoiceID         sql.NullString
			lastError             sql.NullString
			lastAttemptAt         sql.NullTime

			r OrderStateRow
		)
		if err := rows.Scan(
			&r.DscoOrderID,
			&lingxingGlobalOrderNo,
			&r.PushedToLXStatus, &pushedToLXAt,
			&r.AckedToDSCOStatus, &ackedToDSCOAt,
			&r.ShippedToDSCOStatus, &shippedToDSCOAt, &shippedTrackingNo,
			&r.InvoicedToDSCOStatus, &invoicedToDSCOAt, &dscoInvoiceID,
			&r.RetryCount, &lastError, &lastAttemptAt,
			&r.CreatedAt, &r.UpdatedAt,
		); err != nil {
			return nil, err
		}

		if lingxingGlobalOrderNo.Valid {
			v := strings.TrimSpace(lingxingGlobalOrderNo.String)
			if v != "" {
				r.LingxingGlobalOrderNo = &v
			}
		}
		if pushedToLXAt.Valid {
			v := pushedToLXAt.Time
			r.PushedToLXAt = &v
		}
		if ackedToDSCOAt.Valid {
			v := ackedToDSCOAt.Time
			r.AckedToDSCOAt = &v
		}
		if shippedToDSCOAt.Valid {
			v := shippedToDSCOAt.Time
			r.ShippedToDSCOAt = &v
		}
		if shippedTrackingNo.Valid {
			v := strings.TrimSpace(shippedTrackingNo.String)
			if v != "" {
				r.ShippedTrackingNo = &v
			}
		}
		if invoicedToDSCOAt.Valid {
			v := invoicedToDSCOAt.Time
			r.InvoicedToDSCOAt = &v
		}
		if dscoInvoiceID.Valid {
			v := strings.TrimSpace(dscoInvoiceID.String)
			if v != "" {
				r.DscoInvoiceID = &v
			}
		}
		if lastError.Valid {
			v := strings.TrimSpace(lastError.String)
			if v != "" {
				r.LastError = &v
			}
		}
		if lastAttemptAt.Valid {
			v := lastAttemptAt.Time
			r.LastAttemptAt = &v
		}

		out = append(out, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
