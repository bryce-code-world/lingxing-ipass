package store

import (
	"context"
	"encoding/json"

	"gorm.io/gorm"
)

type RuntimeConfigStore struct {
	db *gorm.DB
}

func NewRuntimeConfigStore(db *gorm.DB) *RuntimeConfigStore {
	return &RuntimeConfigStore{db: db}
}

func (s *RuntimeConfigStore) Load(ctx context.Context, domain string) (json.RawMessage, int64, bool, error) {
	var row RuntimeConfigRow
	err := s.db.WithContext(ctx).Where("domain = ?", domain).Take(&row).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, false, nil
		}
		return nil, 0, false, err
	}
	return row.Config, row.UpdatedAt, true, nil
}

func (s *RuntimeConfigStore) Upsert(ctx context.Context, domain string, configJSON []byte, updatedAt int64) error {
	return s.db.WithContext(ctx).Exec(
		`INSERT INTO runtime_config(domain, config, updated_at)
		 VALUES (?, ?::jsonb, ?)
		 ON CONFLICT(domain) DO UPDATE SET config=EXCLUDED.config, updated_at=EXCLUDED.updated_at`,
		domain, string(configJSON), updatedAt,
	).Error
}
