package store

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type DSCOWarehouseSyncStore struct {
	db *gorm.DB
}

func NewDSCOWarehouseSyncStore(db *gorm.DB) *DSCOWarehouseSyncStore {
	return &DSCOWarehouseSyncStore{db: db}
}

type DSCOWarehouseSyncListFilter struct {
	StartTime *int64
	EndTime   *int64
	StatusIn  []int16

	DSCOWarehouseID      string
	DSCOWarehouseSKU     string
	LingXingWarehouseID  string
	LingXingWarehouseSKU string

	Offset int
	Limit  int
}

func (s *DSCOWarehouseSyncStore) Insert(ctx context.Context, row DSCOWarehouseSyncRow) error {
	now := time.Now().UTC().Unix()
	if row.SyncTime == 0 {
		row.SyncTime = now
	}
	if row.CreatedAt == 0 {
		row.CreatedAt = now
	}
	row.UpdatedAt = now
	return s.db.WithContext(ctx).Create(&row).Error
}

func (s *DSCOWarehouseSyncStore) List(ctx context.Context, f DSCOWarehouseSyncListFilter) ([]DSCOWarehouseSyncRow, int64, error) {
	if f.Limit <= 0 || f.Limit > 500 {
		f.Limit = 50
	}
	if f.Offset < 0 {
		f.Offset = 0
	}
	q := s.db.WithContext(ctx).Model(&DSCOWarehouseSyncRow{})
	if f.StartTime != nil {
		q = q.Where("sync_time >= ?", *f.StartTime)
	}
	if f.EndTime != nil {
		q = q.Where("sync_time < ?", *f.EndTime)
	}
	if len(f.StatusIn) > 0 {
		q = q.Where("status IN ?", f.StatusIn)
	}
	if f.DSCOWarehouseID != "" {
		q = q.Where("dsco_warehouse_id = ?", f.DSCOWarehouseID)
	}
	if f.DSCOWarehouseSKU != "" {
		q = q.Where("dsco_warehouse_sku = ?", f.DSCOWarehouseSKU)
	}
	if f.LingXingWarehouseID != "" {
		q = q.Where("lingxing_warehouse_id = ?", f.LingXingWarehouseID)
	}
	if f.LingXingWarehouseSKU != "" {
		q = q.Where("lingxing_warehouse_sku = ?", f.LingXingWarehouseSKU)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []DSCOWarehouseSyncRow
	if err := q.Order("sync_time DESC").Offset(f.Offset).Limit(f.Limit).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
