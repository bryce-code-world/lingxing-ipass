package runtimecfg

import (
	"context"
	"encoding/json"
	"errors"
	"sync/atomic"
	"time"

	"lingxingipass/infra/store"
)

type Manager struct {
	store *store.RuntimeConfigStore
	cur   atomic.Value // RuntimeConfig
}

func NewManager(s *store.RuntimeConfigStore) *Manager {
	return &Manager{store: s}
}

func (m *Manager) LoadOrInit(ctx context.Context, domain string) error {
	raw, updatedAt, ok, err := m.store.Load(ctx, domain)
	if err != nil {
		return err
	}
	if !ok {
		cfg := DefaultConfig(domain)
		raw, _ := json.Marshal(cfg)
		now := time.Now().UTC().Unix()
		if err := m.store.Upsert(ctx, domain, raw, now); err != nil {
			return err
		}
		rc := RuntimeConfig{Domain: domain, Config: cfg, UpdatedAt: now, LoadedAt: time.Now()}
		m.cur.Store(rc)
		return nil
	}
	var cfg Config
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return err
	}
	rc := RuntimeConfig{Domain: domain, Config: cfg, UpdatedAt: updatedAt, LoadedAt: time.Now()}
	m.cur.Store(rc)
	return nil
}

func (m *Manager) Snapshot(domain string) (RuntimeConfig, bool) {
	v := m.cur.Load()
	if v == nil {
		return RuntimeConfig{}, false
	}
	rc := v.(RuntimeConfig)
	if rc.Domain != domain {
		return RuntimeConfig{}, false
	}
	return rc, true
}

func (m *Manager) Update(ctx context.Context, domain string, cfg Config, supportedJobs map[JobName]struct{}) error {
	if cfg.Domain == "" {
		cfg.Domain = domain
	}
	if err := Validate(cfg, supportedJobs); err != nil {
		return err
	}
	raw, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	now := time.Now().UTC().Unix()
	if err := m.store.Upsert(ctx, domain, raw, now); err != nil {
		return err
	}
	m.cur.Store(RuntimeConfig{
		Domain:    domain,
		Config:    cfg,
		UpdatedAt: now,
		LoadedAt:  time.Now(),
	})
	return nil
}

var ErrNotLoaded = errors.New("runtime config not loaded")
