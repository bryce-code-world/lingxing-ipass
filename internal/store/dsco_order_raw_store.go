package store

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// DscoOrderRawStore 负责 dsco_order_raw 的写入（基于 GORM）。
type DscoOrderRawStore struct {
	db *gorm.DB
}

func NewDscoOrderRawStore(db *gorm.DB) (*DscoOrderRawStore, error) {
	if db == nil {
		return nil, errors.New("db 不能为空")
	}
	return &DscoOrderRawStore{db: db}, nil
}

// UpsertLatest 写入/更新 DSCO 订单原始快照（只保留最新一份）。
func (s *DscoOrderRawStore) UpsertLatest(ctx context.Context, dscoOrderID string, payload []byte, fetchedAt time.Time) error {
	dscoOrderID = strings.TrimSpace(dscoOrderID)
	if dscoOrderID == "" {
		return errors.New("dscoOrderID 不能为空")
	}
	if len(payload) == 0 {
		return errors.New("payload 不能为空")
	}
	if fetchedAt.IsZero() {
		fetchedAt = time.Now().UTC()
	}

	sum := sha256.Sum256(payload)
	hash := hex.EncodeToString(sum[:])

	return s.db.WithContext(ctx).Exec(`
INSERT INTO dsco_order_raw (dsco_order_id, payload, payload_sha256, fetched_at)
VALUES (?, ?, ?, ?)
ON DUPLICATE KEY UPDATE payload = VALUES(payload), payload_sha256 = VALUES(payload_sha256), fetched_at = VALUES(fetched_at)
`, dscoOrderID, payload, hash, fetchedAt.UTC()).Error
}
