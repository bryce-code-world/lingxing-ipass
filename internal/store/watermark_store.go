package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strings"
)

// WatermarkStore 负责 job_watermark 的读写。
type WatermarkStore struct {
	db *sql.DB
}

func NewWatermarkStore(db *sql.DB) (*WatermarkStore, error) {
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

	var raw []byte
	err := s.db.QueryRowContext(ctx, `SELECT watermark FROM job_watermark WHERE job_name = ?`, jobName).Scan(&raw)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return json.RawMessage(raw), true, nil
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

	_, err := s.db.ExecContext(ctx, `
INSERT INTO job_watermark (job_name, watermark)
VALUES (?, ?)
ON DUPLICATE KEY UPDATE watermark = VALUES(watermark)
`, jobName, []byte(watermark))
	return err
}
