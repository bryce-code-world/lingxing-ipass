package scheduler

import (
	"context"
	"sync"
	"time"

	"gitee.com/lsy007/golibv2/v2/tool/logger"

	"github.com/robfig/cron/v3"

	"lingxingipass/infra/runtimecfg"
	"lingxingipass/integration"
)

type Scheduler struct {
	cfgMgr *runtimecfg.Manager
	reg    *integration.Registry
	runner *integration.Runner

	mu        sync.Mutex
	cron      *cron.Cron
	lastStamp int64

	stopCh chan struct{}
	wg     sync.WaitGroup
}

func New(cfgMgr *runtimecfg.Manager, reg *integration.Registry, runner *integration.Runner) *Scheduler {
	return &Scheduler{
		cfgMgr: cfgMgr,
		reg:    reg,
		runner: runner,
		stopCh: make(chan struct{}),
	}
}

func (s *Scheduler) Start(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cron != nil {
		return nil
	}

	s.cron = cron.New(
		cron.WithSeconds(),
		cron.WithLocation(time.UTC),
	)
	s.cron.Start()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.pollLoop(ctx)
	}()

	return nil
}

func (s *Scheduler) Stop(ctx context.Context) error {
	s.mu.Lock()
	c := s.cron
	s.cron = nil
	s.mu.Unlock()

	close(s.stopCh)

	if c != nil {
		stopCtx := c.Stop()
		select {
		case <-stopCtx.Done():
		case <-ctx.Done():
		}
	}
	done := make(chan struct{})
	go func() { s.wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-ctx.Done():
	}
	return nil
}

func (s *Scheduler) pollLoop(ctx context.Context) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-s.stopCh:
			return
		case <-ticker.C:
			s.rebuildIfNeeded(ctx)
		}
	}
}

func (s *Scheduler) rebuildIfNeeded(ctx context.Context) {
	cfg, ok := s.cfgMgr.Snapshot(runtimecfg.DomainDSCOLingXing)
	if !ok {
		return
	}
	if cfg.UpdatedAt == s.lastStamp {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cron == nil {
		return
	}
	// Double check within lock.
	cfg, ok = s.cfgMgr.Snapshot(runtimecfg.DomainDSCOLingXing)
	if !ok {
		return
	}
	if cfg.UpdatedAt == s.lastStamp {
		return
	}
	s.lastStamp = cfg.UpdatedAt

	// Rebuild: stop current cron and create a new instance.
	old := s.cron
	s.cron = cron.New(
		cron.WithSeconds(),
		cron.WithLocation(time.UTC),
	)
	s.cron.Start()
	oldCtx := old.Stop()
	select {
	case <-oldCtx.Done():
	case <-ctx.Done():
	}

	for _, job := range s.reg.Jobs() {
		jobCfg, exists := cfg.Config.Jobs[job]
		if !exists || !jobCfg.Enable {
			continue
		}
		jobName := job
		size := jobCfg.Size
		cronSpec := jobCfg.Cron
		_, err := s.cron.AddFunc(cronSpec, func() {
			runCtx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
			defer cancel()
			_ = s.runner.Run(runCtx, integration.RunRequest{
				Domain:   runtimecfg.DomainDSCOLingXing,
				Job:      jobName,
				Trigger:  integration.TriggerScheduled,
				Size:     size,
				Override: nil,
			})
		})
		if err != nil {
			logger.Warn(ctx, "add cron failed",
				"job", string(jobName),
				"cron", cronSpec,
				"err", err,
			)
		}
	}

	logger.Info(ctx, "scheduler rebuilt", "updated_at", cfg.UpdatedAt)
}
