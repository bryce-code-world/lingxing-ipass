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

	DSCOWarehouseID        string
	DSCOWarehouseSKU       string
	DSCOWarehouseSKUIn     []string
	LingXingWarehouseID    string
	LingXingWarehouseSKU   string
	LingXingWarehouseSKUIn []string

	DiffMin     *int
	DiffMax     *int
	DiffEq      *int
	DiffNotZero bool

	Offset int
	Limit  int
}

func applyDSCOWarehouseSyncListFilter(q *gorm.DB, f DSCOWarehouseSyncListFilter) *gorm.DB {
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
	if len(f.DSCOWarehouseSKUIn) > 0 {
		q = q.Where("dsco_warehouse_sku IN ?", f.DSCOWarehouseSKUIn)
	}
	if f.LingXingWarehouseID != "" {
		q = q.Where("lingxing_warehouse_id = ?", f.LingXingWarehouseID)
	}
	if f.LingXingWarehouseSKU != "" {
		q = q.Where("lingxing_warehouse_sku = ?", f.LingXingWarehouseSKU)
	}
	if len(f.LingXingWarehouseSKUIn) > 0 {
		q = q.Where("lingxing_warehouse_sku IN ?", f.LingXingWarehouseSKUIn)
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
	if f.DiffNotZero {
		q = q.Where("diff <> 0")
	}
	return q
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
	q := applyDSCOWarehouseSyncListFilter(
		s.db.WithContext(ctx).Model(&DSCOWarehouseSyncRow{}),
		f,
	)

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

func (s *DSCOWarehouseSyncStore) ListLatestByFullKey(ctx context.Context, f DSCOWarehouseSyncListFilter) ([]DSCOWarehouseSyncRow, int64, error) {
	if f.Limit <= 0 || f.Limit > 500 {
		f.Limit = 50
	}
	if f.Offset < 0 {
		f.Offset = 0
	}

	base := applyDSCOWarehouseSyncListFilter(
		s.db.WithContext(ctx).Model(&DSCOWarehouseSyncRow{}),
		f,
	)

	var total int64
	countQ := s.db.WithContext(ctx).Table("(?) AS t",
		base.Select("1").Group("dsco_warehouse_id, dsco_warehouse_sku, lingxing_warehouse_id, lingxing_warehouse_sku"),
	)
	if err := countQ.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sub := applyDSCOWarehouseSyncListFilter(
		s.db.WithContext(ctx).Table("dsco_warehouse_sync"),
		f,
	).Select("DISTINCT ON (dsco_warehouse_id, dsco_warehouse_sku, lingxing_warehouse_id, lingxing_warehouse_sku) *").
		Order("dsco_warehouse_id, dsco_warehouse_sku, lingxing_warehouse_id, lingxing_warehouse_sku, sync_time DESC")

	var items []DSCOWarehouseSyncRow
	if err := s.db.WithContext(ctx).
		Table("(?) AS latest", sub).
		Order("sync_time DESC").
		Offset(f.Offset).
		Limit(f.Limit).
		Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (s *DSCOWarehouseSyncStore) ListLatestByDSCOKey(ctx context.Context, f DSCOWarehouseSyncListFilter) ([]DSCOWarehouseSyncRow, int64, error) {
	if f.Limit <= 0 || f.Limit > 500 {
		f.Limit = 50
	}
	if f.Offset < 0 {
		f.Offset = 0
	}

	base := applyDSCOWarehouseSyncListFilter(
		s.db.WithContext(ctx).Model(&DSCOWarehouseSyncRow{}),
		f,
	)

	var total int64
	countQ := s.db.WithContext(ctx).Table("(?) AS t",
		base.Select("1").Group("dsco_warehouse_id, dsco_warehouse_sku"),
	)
	if err := countQ.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sub := applyDSCOWarehouseSyncListFilter(
		s.db.WithContext(ctx).Table("dsco_warehouse_sync"),
		f,
	).Select("DISTINCT ON (dsco_warehouse_id, dsco_warehouse_sku) *").
		Order("dsco_warehouse_id, dsco_warehouse_sku, sync_time DESC")

	var items []DSCOWarehouseSyncRow
	if err := s.db.WithContext(ctx).
		Table("(?) AS latest", sub).
		Order("sync_time DESC").
		Offset(f.Offset).
		Limit(f.Limit).
		Find(&items).Error; err != nil {
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
