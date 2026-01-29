package integration

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"example.com/lingxing/golib/v2/tool/logger"
	"gorm.io/gorm"

	"lingxingipass/infra/lock"
	"lingxingipass/infra/runtimecfg"
)

var (
	ErrJobNotFound   = errors.New("job not found")
	ErrJobDisabled   = errors.New("job disabled")
	ErrJobRunning    = errors.New("job running")
	ErrConfigMissing = errors.New("runtime config missing")
)

type Runner struct {
	cfgMgr *runtimecfg.Manager
	reg    *Registry

	sqlDB *sql.DB
	locks *lock.Advisory
}

func NewRunner(cfgMgr *runtimecfg.Manager, reg *Registry, gdb *gorm.DB) *Runner {
	sqlDB, _ := gdb.DB()
	return &Runner{
		cfgMgr: cfgMgr,
		reg:    reg,
		sqlDB:  sqlDB,
		locks:  lock.NewAdvisory(sqlDB),
	}
}

func (r *Runner) Run(ctx context.Context, req RunRequest) error {
	task, ok := r.reg.Get(req.Job)
	if !ok {
		return ErrJobNotFound
	}
	rc, ok := r.cfgMgr.Snapshot(req.Domain)
	if !ok {
		return ErrConfigMissing
	}
	jobCfg, ok := rc.Config.Jobs[req.Job]
	if !ok {
		return fmt.Errorf("job config missing: %s", req.Job)
	}
	if !jobCfg.Enable && req.Trigger == TriggerScheduled {
		return ErrJobDisabled
	}
	// For manual trigger: keep simple and only allow enabled jobs by default.
	if !jobCfg.Enable && req.Trigger == TriggerManual {
		return ErrJobDisabled
	}
	size := req.Size
	if size <= 0 {
		size = jobCfg.Size
	}

	lockKey := lock.KeyFromString(req.Domain + ":" + string(req.Job))
	conn, err := r.locks.TryLock(ctx, lockKey)
	if err != nil {
		if errors.Is(err, lock.ErrLocked) {
			return ErrJobRunning
		}
		return err
	}
	defer func() { _ = r.locks.Unlock(context.Background(), conn, lockKey) }()

	logger.Info(ctx, "job start",
		"domain", req.Domain,
		"job", string(req.Job),
		"trigger", string(req.Trigger),
		"size", size,
		"cfg_updated_at", rc.UpdatedAt,
	)

	err = task(TaskContext{
		Ctx:      ctx,
		Domain:   req.Domain,
		Job:      req.Job,
		Trigger:  req.Trigger,
		Size:     size,
		Config:   rc.Config,
		Override: req.Override,
	})
	if err != nil {
		logger.Error(ctx, "job failed",
			"domain", req.Domain,
			"job", string(req.Job),
			"err", err,
		)
		return err
	}
	logger.Info(ctx, "job ok",
		"domain", req.Domain,
		"job", string(req.Job),
	)
	return nil
}
