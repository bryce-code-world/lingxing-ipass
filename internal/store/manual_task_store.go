package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// ManualTaskStore 负责 manual_task 的写入。
type ManualTaskStore struct {
	db *sql.DB
}

func NewManualTaskStore(db *sql.DB) (*ManualTaskStore, error) {
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
	_, err := s.db.ExecContext(ctx, `
INSERT INTO manual_task (task_type, dsco_order_id, payload, status)
VALUES (?, ?, ?, 0)
`, task.TaskType, sql.NullString{String: task.DscoOrderID, Valid: task.DscoOrderID != ""}, []byte(task.Payload))
	return err
}

type ManualTaskRow struct {
	ID          int64
	TaskType    string
	DscoOrderID sql.NullString
	Payload     json.RawMessage
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
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

	rows, err := s.db.QueryContext(ctx, `
SELECT id, task_type, dsco_order_id, payload, status, created_at, updated_at
FROM manual_task
WHERE status = ?
ORDER BY id DESC
LIMIT ? OFFSET ?
`, status, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []ManualTaskRow
	for rows.Next() {
		var r ManualTaskRow
		var payload []byte
		if err := rows.Scan(&r.ID, &r.TaskType, &r.DscoOrderID, &payload, &r.Status, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, err
		}
		r.Payload = json.RawMessage(payload)
		r.TaskType = strings.TrimSpace(r.TaskType)
		out = append(out, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
