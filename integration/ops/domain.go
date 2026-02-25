package ops

import (
	"context"
	"time"

	"lingxingipass/golib/v2/tool/logger"

	"lingxingipass/infra/config"
	"lingxingipass/infra/fileutil"
	"lingxingipass/integration"
)

type Domain struct {
	exportDir        string
	cleanupThreshold int64
}

func NewDomain(env config.EnvConfig) *Domain {
	return &Domain{
		exportDir:        env.Admin.Export.Dir,
		cleanupThreshold: env.Admin.Export.CleanupThresholdBytes,
	}
}

func (d *Domain) CleanupExports(ctx integration.TaskContext) (retErr error) {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}
	startedAt := time.Now().UTC()
	base := ctx.BaseLogFields()
	logger.Info(taskCtx, "task begin", append(base, "task", "cleanup_exports")...)
	defer func() {
		fields := append(base,
			"task", "cleanup_exports",
			"duration_ms", time.Since(startedAt).Milliseconds(),
		)
		if retErr != nil {
			logger.Error(taskCtx, "task end", append(fields, "result", "failed", "err", retErr)...)
			return
		}
		logger.Info(taskCtx, "task end", append(fields, "result", "ok")...)
	}()
	if d.exportDir == "" || d.cleanupThreshold <= 0 {
		return nil
	}

	before, _ := fileutil.DirSizeBytes(d.exportDir)
	deleted, freed, err := fileutil.CleanupDirByOldest(d.exportDir, d.cleanupThreshold)
	after, _ := fileutil.DirSizeBytes(d.exportDir)
	if err != nil {
		logger.Warn(taskCtx, "cleanup exports failed", "err", err, "dir", d.exportDir)
		retErr = err
		return retErr
	}
	if deleted > 0 {
		logger.Info(taskCtx, "cleanup exports done",
			"dir", d.exportDir,
			"deleted", deleted,
			"freed_bytes", freed,
			"size_before", before,
			"size_after", after,
		)
	}
	return nil
}
