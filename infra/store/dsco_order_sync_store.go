package store

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type DSCOOrderSyncStore struct {
	db *gorm.DB
}

func NewDSCOOrderSyncStore(db *gorm.DB) *DSCOOrderSyncStore {
	return &DSCOOrderSyncStore{db: db}
}

func toPGTextArrayLiteral(items []string) string {
	if len(items) == 0 {
		return "{}"
	}
	parts := make([]string, 0, len(items))
	for _, it := range items {
		s := strings.TrimSpace(it)
		s = strings.ReplaceAll(s, `\`, `\\`)
		s = strings.ReplaceAll(s, `"`, `\"`)
		parts = append(parts, `"`+s+`"`)
	}
	return "{" + strings.Join(parts, ",") + "}"
}

type DSCOOrderSyncListFilter struct {
	StartTime      *int64
	EndTime        *int64
	StatusIn       []int16
	DSCOStatus     string
	PONumberLike   string
	DSCOREtailerID string
	MSKU           string // filter: sku = ANY(mskus)

	Offset int
	Limit  int
}

func (s *DSCOOrderSyncStore) Upsert(ctx context.Context, row DSCOOrderSyncRow) error {
	if strings.TrimSpace(row.PONumber) == "" {
		return errors.New("po_number 不能为空")
	}
	now := time.Now().UTC().Unix()
	if row.CreatedAt == 0 {
		row.CreatedAt = now
	}
	row.UpdatedAt = now

	// Use SQL upsert to ensure overwrite payload/status is allowed.
	return s.db.WithContext(ctx).Exec(
		`INSERT INTO dsco_order_sync
		    (po_number, dsco_create_time, dsco_retailer_id, dsco_status, status, payload, mskus, warehouse_id, shipment, shipped_tracking_no, dsco_invoice_id, created_at, updated_at)
		 VALUES
		    (?, ?, ?, ?, ?, ?::jsonb, ?::text[], ?, ?, ?, ?, ?, ?)
		 ON CONFLICT (po_number) DO UPDATE SET
		    dsco_create_time=EXCLUDED.dsco_create_time,
		    dsco_retailer_id=EXCLUDED.dsco_retailer_id,
		    dsco_status=EXCLUDED.dsco_status,
		    status=EXCLUDED.status,
		    payload=EXCLUDED.payload,
		    mskus=EXCLUDED.mskus,
		    warehouse_id=EXCLUDED.warehouse_id,
		    shipment=EXCLUDED.shipment,
		    shipped_tracking_no=EXCLUDED.shipped_tracking_no,
		    dsco_invoice_id=EXCLUDED.dsco_invoice_id,
		    updated_at=EXCLUDED.updated_at`,
		row.PONumber, row.DSCOCreateTime, row.DSCOREtailerID, row.DSCOStatus, row.Status, string(row.Payload), toPGTextArrayLiteral([]string(row.MSKUs)),
		row.WarehouseID, row.Shipment, row.ShippedTrackingNo, row.DSCOInvoiceID,
		row.CreatedAt, row.UpdatedAt,
	).Error
}

func (s *DSCOOrderSyncStore) GetMaxDSCOCreateTime(ctx context.Context) (int64, bool, error) {
	var max int64
	err := s.db.WithContext(ctx).Raw(`SELECT COALESCE(MAX(dsco_create_time), 0) FROM dsco_order_sync`).Scan(&max).Error
	if err != nil {
		return 0, false, err
	}
	if max == 0 {
		return 0, false, nil
	}
	return max, true, nil
}

func (s *DSCOOrderSyncStore) List(ctx context.Context, f DSCOOrderSyncListFilter) ([]DSCOOrderSyncRow, int64, error) {
	if f.Limit <= 0 || f.Limit > 500 {
		f.Limit = 50
	}
	if f.Offset < 0 {
		f.Offset = 0
	}

	q := s.db.WithContext(ctx).Model(&DSCOOrderSyncRow{})
	if f.StartTime != nil {
		q = q.Where("dsco_create_time >= ?", *f.StartTime)
	}
	if f.EndTime != nil {
		q = q.Where("dsco_create_time < ?", *f.EndTime)
	}
	if len(f.StatusIn) > 0 {
		q = q.Where("status IN ?", f.StatusIn)
	}
	if strings.TrimSpace(f.DSCOStatus) != "" {
		q = q.Where("dsco_status = ?", strings.TrimSpace(f.DSCOStatus))
	}
	if strings.TrimSpace(f.PONumberLike) != "" {
		q = q.Where("po_number ILIKE ?", "%"+strings.TrimSpace(f.PONumberLike)+"%")
	}
	if strings.TrimSpace(f.DSCOREtailerID) != "" {
		q = q.Where("dsco_retailer_id = ?", strings.TrimSpace(f.DSCOREtailerID))
	}
	if strings.TrimSpace(f.MSKU) != "" {
		q = q.Where("? = ANY(mskus)", strings.TrimSpace(f.MSKU))
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []DSCOOrderSyncRow
	if err := q.Order("dsco_create_time DESC").Offset(f.Offset).Limit(f.Limit).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (s *DSCOOrderSyncStore) FindByStatus(ctx context.Context, status int16, limit int) ([]DSCOOrderSyncRow, error) {
	if limit <= 0 {
		limit = 50
	}
	var items []DSCOOrderSyncRow
	err := s.db.WithContext(ctx).
		Where("status = ?", status).
		Order("dsco_create_time ASC").
		Limit(limit).
		Find(&items).Error
	return items, err
}

func (s *DSCOOrderSyncStore) GetByID(ctx context.Context, id int64) (DSCOOrderSyncRow, bool, error) {
	var row DSCOOrderSyncRow
	err := s.db.WithContext(ctx).Where("id = ?", id).Take(&row).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return DSCOOrderSyncRow{}, false, nil
		}
		return DSCOOrderSyncRow{}, false, err
	}
	return row, true, nil
}

func (s *DSCOOrderSyncStore) UpdateStatus(ctx context.Context, poNumber string, status int16) error {
	if strings.TrimSpace(poNumber) == "" {
		return errors.New("po_number 不能为空")
	}
	now := time.Now().UTC().Unix()
	return s.db.WithContext(ctx).
		Model(&DSCOOrderSyncRow{}).
		Where("po_number = ?", strings.TrimSpace(poNumber)).
		Updates(map[string]any{
			"status":     status,
			"updated_at": now,
		}).Error
}

func (s *DSCOOrderSyncStore) UpdateStatusAndFields(ctx context.Context, poNumber string, status int16, trackingNo, invoiceID string) error {
	now := time.Now().UTC().Unix()
	return s.db.WithContext(ctx).
		Model(&DSCOOrderSyncRow{}).
		Where("po_number = ?", poNumber).
		Updates(map[string]any{
			"status":              status,
			"shipped_tracking_no": trackingNo,
			"dsco_invoice_id":     invoiceID,
			"updated_at":          now,
		}).Error
}
