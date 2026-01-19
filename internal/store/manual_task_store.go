package store

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// ManualTaskStore 负责 manual_task 的写入（基于 GORM）。
type ManualTaskStore struct {
	db *gorm.DB
}

func NewManualTaskStore(db *gorm.DB) (*ManualTaskStore, error) {
	if db == nil {
		return nil, errors.New("db 不能为空")
	}
	return &ManualTaskStore{db: db}, nil
}

type ManualTask struct {
	TaskType    string
	DscoOrderID string
	Payload     json.RawMessage
}

// Create 写入一条人工任务（payload 必须是脱敏后的最小上下文）。
func (s *ManualTaskStore) Create(ctx context.Context, task ManualTask) error {
	task.TaskType = strings.TrimSpace(task.TaskType)
	task.DscoOrderID = strings.TrimSpace(task.DscoOrderID)
	if task.TaskType == "" {
		return errors.New("task_type 不能为空")
	}
	if len(task.Payload) == 0 {
		task.Payload = []byte(`{}`)
	}

	var dscoOrderID any
	if task.DscoOrderID == "" {
		dscoOrderID = nil
	} else {
		dscoOrderID = task.DscoOrderID
	}

	return s.db.WithContext(ctx).Exec(`
INSERT INTO manual_task (task_type, dsco_order_id, payload, status)
VALUES (?, ?, ?, 0)
`, task.TaskType, dscoOrderID, []byte(task.Payload)).Error
}

type ManualTaskRow struct {
	ID          int64           `json:"id"`
	TaskType    string          `json:"task_type"`
	DscoOrderID *string         `json:"dsco_order_id,omitempty"`
	Payload     json.RawMessage `json:"payload"`
	Status      int             `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// ListByStatus 查询人工任务列表（一期：只做最小分页）。
func (s *ManualTaskStore) ListByStatus(ctx context.Context, status int, limit int, offset int) ([]ManualTaskRow, error) {
	if limit <= 0 {
		limit = 50
	}
	if limit > 200 {
		limit = 200
	}
	if offset < 0 {
		offset = 0
	}

	var out []ManualTaskRow
	if err := s.db.WithContext(ctx).Raw(`
SELECT id, task_type, dsco_order_id, payload, status, created_at, updated_at
FROM manual_task
WHERE status = ?
ORDER BY id DESC
LIMIT ? OFFSET ?
`, status, limit, offset).Scan(&out).Error; err != nil {
		return nil, err
	}
	for i := range out {
		out[i].TaskType = strings.TrimSpace(out[i].TaskType)
		if out[i].DscoOrderID != nil {
			v := strings.TrimSpace(*out[i].DscoOrderID)
			if v == "" {
				out[i].DscoOrderID = nil
			} else {
				out[i].DscoOrderID = &v
			}
		}
	}
	return out, nil
}

