package store

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// ManualTaskStore 负责 manual_task 的读写（admin 独立实现）。
type ManualTaskStore struct {
	db *gorm.DB
}

func NewManualTaskStore(db *gorm.DB) (*ManualTaskStore, error) {
	if db == nil {
		return nil, errors.New("db 不能为空")
	}
	return &ManualTaskStore{db: db}, nil
}

type ManualTaskRow struct {
	ID          int64     `json:"id"`
	TaskType    string    `json:"task_type"`
	DscoOrderID *string   `json:"dsco_order_id,omitempty"`
	Payload     []byte    `json:"payload"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ListByStatus 按状态分页查询人工任务（admin 展示用）。
func (s *ManualTaskStore) ListByStatus(ctx context.Context, status, limit, offset int) ([]ManualTaskRow, error) {
	if limit <= 0 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	var rows []ManualTaskRow
	err := s.db.WithContext(ctx).Raw(`
SELECT id, task_type, dsco_order_id, payload, status, created_at, updated_at
FROM manual_task
WHERE status = ?
ORDER BY updated_at DESC, id DESC
LIMIT ? OFFSET ?
`, status, limit, offset).Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	for i := range rows {
		rows[i].TaskType = strings.TrimSpace(rows[i].TaskType)
		if rows[i].DscoOrderID != nil {
			v := strings.TrimSpace(*rows[i].DscoOrderID)
			rows[i].DscoOrderID = &v
		}
	}
	return rows, nil
}
