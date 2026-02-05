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

	DiffMin *int
	DiffMax *int
	DiffEq  *int

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

func (s *DSCOWarehouseSyncStore) InsertBatch(ctx context.Context, rows []DSCOWarehouseSyncRow) error {
	if len(rows) == 0 {
		return nil
	}
	now := time.Now().UTC().Unix()
	for i := range rows {
		if rows[i].SyncTime == 0 {
			rows[i].SyncTime = now
		}
		if rows[i].CreatedAt == 0 {
			rows[i].CreatedAt = now
		}
		rows[i].UpdatedAt = now
	}
	// 批量插入用于显著降低 Postgres 往返次数，避免单条 INSERT 导致整体耗时过长。
	return s.db.WithContext(ctx).CreateInBatches(rows, 200).Error
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
	if f.DiffEq != nil {
		q = q.Where("diff = ?", *f.DiffEq)
	} else {
		if f.DiffMin != nil {
			q = q.Where("diff >= ?", *f.DiffMin)
		}
		if f.DiffMax != nil {
			q = q.Where("diff <= ?", *f.DiffMax)
		}
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

func (s *DSCOWarehouseSyncStore) ListDistinctDSCOWarehouseIDs(ctx context.Context, limit int) ([]string, error) {
	if limit <= 0 || limit > 500 {
		limit = 200
	}
	var out []string
	q := s.db.WithContext(ctx).Model(&DSCOWarehouseSyncRow{}).
		Select("dsco_warehouse_id").
		Distinct("dsco_warehouse_id").
		Order("dsco_warehouse_id ASC").
		Limit(limit)
	if err := q.Pluck("dsco_warehouse_id", &out).Error; err != nil {
		return nil, err
	}
	var cleaned []string
	for _, v := range out {
		if v != "" {
			cleaned = append(cleaned, v)
		}
	}
	return cleaned, nil
}

func (s *DSCOWarehouseSyncStore) ListDistinctLingXingWarehouseIDs(ctx context.Context, limit int) ([]string, error) {
	if limit <= 0 || limit > 500 {
		limit = 200
	}
	var out []string
	q := s.db.WithContext(ctx).Model(&DSCOWarehouseSyncRow{}).
		Select("lingxing_warehouse_id").
		Distinct("lingxing_warehouse_id").
		Order("lingxing_warehouse_id ASC").
		Limit(limit)
	if err := q.Pluck("lingxing_warehouse_id", &out).Error; err != nil {
		return nil, err
	}
	var cleaned []string
	for _, v := range out {
		if v != "" {
			cleaned = append(cleaned, v)
		}
	}
	return cleaned, nil
}
