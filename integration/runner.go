package integration

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gitee.com/lsy007/golibv2/v2/tool/logger"
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
	_, err := r.runInternal(ctx, req, nil)
	return err
}

// RunWithResult 用于“手动检查类任务”同步获取结果（不落库）。任务需通过 TaskContext.Report 回传结果。
func (r *Runner) RunWithResult(ctx context.Context, req RunRequest) (any, error) {
	return r.runInternal(ctx, req, func(v any) {})
}

func (r *Runner) runInternal(ctx context.Context, req RunRequest, report func(v any)) (any, error) {
	startedAt := time.Now().UTC()
	runID := fmt.Sprintf("%d", startedAt.UnixNano())

	task, ok := r.reg.Get(req.Job)
	if !ok {
		return nil, ErrJobNotFound
	}
	rc, ok := r.cfgMgr.Snapshot(req.Domain)
	if !ok {
		return nil, ErrConfigMissing
	}
	jobCfg, ok := rc.Config.Jobs[req.Job]
	if !ok {
		// 兼容：runtime_config 可能是老数据，未包含新 job 的配置。此时回退到 DefaultConfig 的 job 配置。
		def := runtimecfg.DefaultConfig(req.Domain)
		if jc, ok2 := def.Jobs[req.Job]; ok2 {
			jobCfg = jc
		} else {
			return nil, fmt.Errorf("job config missing: %s", req.Job)
		}
	}
	if !jobCfg.Enable && req.Trigger == TriggerScheduled {
		return nil, ErrJobDisabled
	}
	// Manual trigger is allowed even when the job is disabled for scheduler.
	size := req.Size
	if size <= 0 {
		size = jobCfg.Size
	}

	lockKey := lock.KeyFromString(req.Domain + ":" + string(req.Job))
	conn, err := r.locks.TryLock(ctx, lockKey)
	if err != nil {
		if errors.Is(err, lock.ErrLocked) {
			return nil, ErrJobRunning
		}
		return nil, err
	}
	defer func() { _ = r.locks.Unlock(context.Background(), conn, lockKey) }()

	logger.Info(ctx, "job start",
		"run_id", runID,
		"domain", req.Domain,
		"job", string(req.Job),
		"trigger", string(req.Trigger),
		"size", size,
		"only_po_number", req.OnlyPONumber,
		"cfg_updated_at", rc.UpdatedAt,
	)

	supportedJobs := r.reg.SupportedJobsSet()

	var result any
	var reportFn func(v any)
	if report != nil {
		reportFn = func(v any) { result = v; report(v) }
	}

	err = task(TaskContext{
		Ctx:     ctx,
		Domain:  req.Domain,
		Job:     req.Job,
		Trigger: req.Trigger,
		Size:    size,
		Report:  reportFn,
		SnapshotRuntimeConfig: func(domain string) (runtimecfg.RuntimeConfig, bool) {
			return r.cfgMgr.Snapshot(domain)
		},
		UpdateRuntimeConfig: func(taskCtx context.Context, domain string, cfg runtimecfg.Config) error {
			return r.cfgMgr.Update(taskCtx, domain, cfg, supportedJobs)
		},
		OnlyPONumber: req.OnlyPONumber,
		RunID:        runID,
		Config:       rc.Config,
		Override:     req.Override,
	})
	if err != nil {
		logger.Error(ctx, "job failed",
			"run_id", runID,
			"domain", req.Domain,
			"job", string(req.Job),
			"only_po_number", req.OnlyPONumber,
			"duration_ms", time.Since(startedAt).Milliseconds(),
			"err", err,
		)
		return result, err
	}
	logger.Info(ctx, "job ok",
		"run_id", runID,
		"domain", req.Domain,
		"job", string(req.Job),
		"only_po_number", req.OnlyPONumber,
		"duration_ms", time.Since(startedAt).Milliseconds(),
	)
	return result, nil
}
