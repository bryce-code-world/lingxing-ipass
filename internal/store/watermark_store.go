package store

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// WatermarkStore 负责 job_watermark 的读写（基于 GORM）。
type WatermarkStore struct {
	db *gorm.DB
}

func NewWatermarkStore(db *gorm.DB) (*WatermarkStore, error) {
	if db == nil {
		return nil, errors.New("db 不能为空")
	}
	return &WatermarkStore{db: db}, nil
}

// Get 获取任务水位；若不存在返回 nil,false。
func (s *WatermarkStore) Get(ctx context.Context, jobName string) (json.RawMessage, bool, error) {
	jobName = strings.TrimSpace(jobName)
	if jobName == "" {
		return nil, false, errors.New("jobName 不能为空")
	}

	var row struct {
		Watermark []byte `gorm:"column:watermark"`
	}
	res := s.db.WithContext(ctx).Raw(`SELECT watermark FROM job_watermark WHERE job_name = ?`, jobName).Scan(&row)
	if res.Error != nil {
		return nil, false, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, false, nil
	}
	if len(row.Watermark) == 0 {
		return nil, false, nil
	}
	return json.RawMessage(row.Watermark), true, nil
}

// Set 写入任务水位（upsert）。
func (s *WatermarkStore) Set(ctx context.Context, jobName string, watermark json.RawMessage) error {
	jobName = strings.TrimSpace(jobName)
	if jobName == "" {
		return errors.New("jobName 不能为空")
	}
	if len(watermark) == 0 {
		return errors.New("watermark 不能为空")
	}

	return s.db.WithContext(ctx).Exec(`
INSERT INTO job_watermark (job_name, watermark)
VALUES (?, ?)
ON DUPLICATE KEY UPDATE watermark = VALUES(watermark)
`, jobName, []byte(watermark)).Error
}

type JobWatermarkRow struct {
	JobName   string          `json:"job_name"`
	Watermark json.RawMessage `json:"watermark"`
	UpdatedAt time.Time       `json:"updated_at"`
}

// ListAll 列出所有任务水位（一期用于管理后台展示）。
func (s *WatermarkStore) ListAll(ctx context.Context) ([]JobWatermarkRow, error) {
	var rows []JobWatermarkRow
	if err := s.db.WithContext(ctx).Raw(`SELECT job_name, watermark, updated_at FROM job_watermark ORDER BY job_name ASC`).Scan(&rows).Error; err != nil {
		return nil, err
	}
	for i := range rows {
		rows[i].JobName = strings.TrimSpace(rows[i].JobName)
	}
	return rows, nil
}
