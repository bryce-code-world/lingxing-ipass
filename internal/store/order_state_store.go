package store

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// OrderStateStore 负责 sync_order_state 的读写（基于 GORM）。
type OrderStateStore struct {
	db *gorm.DB
}

func NewOrderStateStore(db *gorm.DB) (*OrderStateStore, error) {
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
		err := s.db.WithContext(ctx).Exec(`
INSERT INTO sync_order_state (dsco_order_id, created_at, updated_at)
VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at)
`, id, now, now).Error
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
	staleBefore := now.Add(-30 * time.Minute)

	var ids []string
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var rows []struct {
			DscoOrderID string `gorm:"column:dsco_order_id"`
		}
		if err := tx.Raw(`
SELECT dsco_order_id
FROM sync_order_state
WHERE pushed_to_lx_status IN (0, 2)
   OR (pushed_to_lx_status = 9 AND last_attempt_at < ?)
ORDER BY updated_at ASC
LIMIT ?
FOR UPDATE
`, staleBefore, limit).Scan(&rows).Error; err != nil {
			return err
		}
		for _, r := range rows {
			id := strings.TrimSpace(r.DscoOrderID)
			if id != "" {
				ids = append(ids, id)
			}
		}
		if len(ids) == 0 {
			return nil
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
SET pushed_to_lx_status = 9,
    last_attempt_at = ?,
    retry_count = retry_count + 1,
    updated_at = ?
WHERE dsco_order_id IN (%s)
`, strings.Join(placeholders, ","))
		return tx.Exec(q, args...).Error
	})
	if err != nil {
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
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET pushed_to_lx_status = 1,
    pushed_to_lx_at = ?,
    lingxing_global_order_no = ?,
    last_error = NULL,
    updated_at = ?
WHERE dsco_order_id = ?
`, now, globalOrderNo, now, dscoOrderID).Error
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
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET pushed_to_lx_status = 2,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, errMsg, now, dscoOrderID).Error
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
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET pushed_to_lx_status = 3,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, reason, now, dscoOrderID).Error
}

// TryClaimAck 抢占“待 ACK 回传”的订单，避免多实例重复处理。
func (s *OrderStateStore) TryClaimAck(ctx context.Context, dscoOrderID string) (bool, error) {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	if dscoOrderID == "" {
		return false, errors.New("dscoOrderID 不能为空")
	}
	now := time.Now().UTC()
	res := s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET acked_to_dsco_status = 9,
    last_attempt_at = ?,
    retry_count = retry_count + 1,
    updated_at = ?
WHERE dsco_order_id = ?
  AND acked_to_dsco_status IN (0, 2)
`, now, now, dscoOrderID)
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected > 0, nil
}

func (s *OrderStateStore) MarkAckSuccess(ctx context.Context, dscoOrderID string) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	now := time.Now().UTC()
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET acked_to_dsco_status = 1,
    acked_to_dsco_at = ?,
    last_error = NULL,
    updated_at = ?
WHERE dsco_order_id = ?
`, now, now, dscoOrderID).Error
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
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET acked_to_dsco_status = 2,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, errMsg, now, dscoOrderID).Error
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
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET acked_to_dsco_status = 3,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, reason, now, dscoOrderID).Error
}

// TryClaimShipment 抢占“待发货回传”的订单，避免多实例重复处理。
func (s *OrderStateStore) TryClaimShipment(ctx context.Context, dscoOrderID string) (bool, error) {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	if dscoOrderID == "" {
		return false, errors.New("dscoOrderID 不能为空")
	}
	now := time.Now().UTC()
	res := s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET shipped_to_dsco_status = 9,
    last_attempt_at = ?,
    retry_count = retry_count + 1,
    updated_at = ?
WHERE dsco_order_id = ?
  AND shipped_to_dsco_status IN (0, 2)
`, now, now, dscoOrderID)
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected > 0, nil
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
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET shipped_to_dsco_status = 1,
    shipped_to_dsco_at = ?,
    shipped_tracking_no = ?,
    last_error = NULL,
    updated_at = ?
WHERE dsco_order_id = ?
`, now, trackingNo, now, dscoOrderID).Error
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
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET shipped_to_dsco_status = 2,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, errMsg, now, dscoOrderID).Error
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
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET shipped_to_dsco_status = 3,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, reason, now, dscoOrderID).Error
}

// ClaimForInvoice 抢占一批“待回传发票”的 dscoOrderId，避免多实例重复处理。
func (s *OrderStateStore) ClaimForInvoice(ctx context.Context, limit int) ([]string, error) {
	if limit <= 0 {
		return nil, errors.New("limit 必须大于 0")
	}

	now := time.Now().UTC()
	staleBefore := now.Add(-30 * time.Minute)

	var ids []string
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var rows []struct {
			DscoOrderID string `gorm:"column:dsco_order_id"`
		}
		if err := tx.Raw(`
SELECT dsco_order_id
FROM sync_order_state
WHERE invoiced_to_dsco_status IN (0, 2)
   OR (invoiced_to_dsco_status = 9 AND last_attempt_at < ?)
ORDER BY updated_at ASC
LIMIT ?
FOR UPDATE
`, staleBefore, limit).Scan(&rows).Error; err != nil {
			return err
		}
		for _, r := range rows {
			id := strings.TrimSpace(r.DscoOrderID)
			if id != "" {
				ids = append(ids, id)
			}
		}
		if len(ids) == 0 {
			return nil
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
		return tx.Exec(q, args...).Error
	})
	if err != nil {
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
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET invoiced_to_dsco_status = 1,
    invoiced_to_dsco_at = ?,
    dsco_invoice_id = ?,
    last_error = NULL,
    updated_at = ?
WHERE dsco_order_id = ?
`, now, invoiceID, now, dscoOrderID).Error
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
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET invoiced_to_dsco_status = 2,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, errMsg, now, dscoOrderID).Error
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
	return s.db.WithContext(ctx).Exec(`
UPDATE sync_order_state
SET invoiced_to_dsco_status = 3,
    last_error = ?,
    updated_at = ?
WHERE dsco_order_id = ?
`, reason, now, dscoOrderID).Error
}

// GetRetryCount 返回某个 dscoOrderId 的 retry_count；若不存在返回 ok=false。
func (s *OrderStateStore) GetRetryCount(ctx context.Context, dscoOrderID string) (retryCount int, ok bool, err error) {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	if dscoOrderID == "" {
		return 0, false, errors.New("dscoOrderID 不能为空")
	}
	var row struct {
		DscoOrderID string `gorm:"column:dsco_order_id"`
		RetryCount  int    `gorm:"column:retry_count"`
	}
	if err := s.db.WithContext(ctx).Raw(`
SELECT dsco_order_id, retry_count
FROM sync_order_state
WHERE dsco_order_id = ?
LIMIT 1
`, dscoOrderID).Scan(&row).Error; err != nil {
		return 0, false, err
	}
	if strings.TrimSpace(row.DscoOrderID) == "" {
		return 0, false, nil
	}
	return row.RetryCount, true, nil
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

	var row OrderStateRow
	res := s.db.WithContext(ctx).Raw(`
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
`, dscoOrderID).Scan(&row)
	if res.Error != nil {
		return nil, false, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, false, nil
	}
	row.DscoOrderID = strings.TrimSpace(row.DscoOrderID)
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

	sql := `
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
`
	args := make([]any, 0, 8)
	if q.PushedToLXStatus != nil {
		sql += " AND pushed_to_lx_status = ?\n"
		args = append(args, *q.PushedToLXStatus)
	}
	if q.AckedToDSCOStatus != nil {
		sql += " AND acked_to_dsco_status = ?\n"
		args = append(args, *q.AckedToDSCOStatus)
	}
	if q.ShippedToDSCOStatus != nil {
		sql += " AND shipped_to_dsco_status = ?\n"
		args = append(args, *q.ShippedToDSCOStatus)
	}
	if q.InvoicedToDSCOStatus != nil {
		sql += " AND invoiced_to_dsco_status = ?\n"
		args = append(args, *q.InvoicedToDSCOStatus)
	}
	sql += " ORDER BY updated_at DESC\n LIMIT ? OFFSET ?\n"
	args = append(args, q.Limit, q.Offset)

	var rows []OrderStateRow
	if err := s.db.WithContext(ctx).Raw(sql, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}
	for i := range rows {
		rows[i].DscoOrderID = strings.TrimSpace(rows[i].DscoOrderID)
	}
	return rows, nil
}
