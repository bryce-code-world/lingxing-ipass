package store

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// OrderStateStore 负责 sync_order_state 的查询（admin 独立实现）。
type OrderStateStore struct {
	db *gorm.DB
}

func NewOrderStateStore(db *gorm.DB) (*OrderStateStore, error) {
	if db == nil {
		return nil, errors.New("db 不能为空")
	}
	return &OrderStateStore{db: db}, nil
}

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

type OrderStateListQuery struct {
	PushedToLXStatus     *int
	AckedToDSCOStatus    *int
	ShippedToDSCOStatus  *int
	InvoicedToDSCOStatus *int
	Limit                int
	Offset               int
}

// GetByDscoOrderID 查询单条订单状态；若不存在返回 ok=false。
func (s *OrderStateStore) GetByDscoOrderID(ctx context.Context, dscoOrderID string) (*OrderStateRow, bool, error) {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	if dscoOrderID == "" {
		return nil, false, errors.New("dscoOrderID 不能为空")
	}
	var row OrderStateRow
	err := s.db.WithContext(ctx).Raw(`
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
LIMIT 1
`, dscoOrderID).Scan(&row).Error
	if err != nil {
		return nil, false, err
	}
	if strings.TrimSpace(row.DscoOrderID) == "" {
		return nil, false, nil
	}
	row.DscoOrderID = strings.TrimSpace(row.DscoOrderID)
	return &row, true, nil
}

// List 按条件分页列出订单状态（admin 展示用）。
func (s *OrderStateStore) List(ctx context.Context, q OrderStateListQuery) ([]OrderStateRow, error) {
	if q.Limit <= 0 {
		q.Limit = 50
	}
	if q.Offset < 0 {
		q.Offset = 0
	}

	var args []any
	where := "WHERE 1=1"
	if q.PushedToLXStatus != nil {
		where += " AND pushed_to_lx_status = ?"
		args = append(args, *q.PushedToLXStatus)
	}
	if q.AckedToDSCOStatus != nil {
		where += " AND acked_to_dsco_status = ?"
		args = append(args, *q.AckedToDSCOStatus)
	}
	if q.ShippedToDSCOStatus != nil {
		where += " AND shipped_to_dsco_status = ?"
		args = append(args, *q.ShippedToDSCOStatus)
	}
	if q.InvoicedToDSCOStatus != nil {
		where += " AND invoiced_to_dsco_status = ?"
		args = append(args, *q.InvoicedToDSCOStatus)
	}
	args = append(args, q.Limit, q.Offset)

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
` + where + `
ORDER BY updated_at DESC
LIMIT ? OFFSET ?
`
	var rows []OrderStateRow
	if err := s.db.WithContext(ctx).Raw(sql, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}
	for i := range rows {
		rows[i].DscoOrderID = strings.TrimSpace(rows[i].DscoOrderID)
	}
	return rows, nil
}
