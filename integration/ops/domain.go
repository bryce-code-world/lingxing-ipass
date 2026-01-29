package ops

import (
	"context"

	"example.com/lingxing/golib/v2/tool/logger"

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

func (d *Domain) CleanupExports(ctx integration.TaskContext) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}
	if d.exportDir == "" || d.cleanupThreshold <= 0 {
		return nil
	}

	before, _ := fileutil.DirSizeBytes(d.exportDir)
	deleted, freed, err := fileutil.CleanupDirByOldest(d.exportDir, d.cleanupThreshold)
	after, _ := fileutil.DirSizeBytes(d.exportDir)
	if err != nil {
		logger.Warn(taskCtx, "cleanup exports failed", "err", err, "dir", d.exportDir)
		return err
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
