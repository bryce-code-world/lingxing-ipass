package gormx

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"lingxingipass/golib/v2/tool/logger"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type toolLoggerAdapterConfig struct {
	level                     gormLogger.LogLevel
	slowThreshold             time.Duration
	ignoreRecordNotFoundError bool
}

type toolLoggerAdapter struct {
	cfg toolLoggerAdapterConfig
}

func newToolLoggerAdapter(cfg toolLoggerAdapterConfig) gormLogger.Interface {
	return &toolLoggerAdapter{cfg: cfg}
}

func (l *toolLoggerAdapter) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	n := *l
	n.cfg.level = level
	return &n
}

func (l *toolLoggerAdapter) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.cfg.level < gormLogger.Info {
		return
	}
	logger.Info(ctx, fmt.Sprintf(msg, data...))
}

func (l *toolLoggerAdapter) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.cfg.level < gormLogger.Warn {
		return
	}
	logger.Warn(ctx, fmt.Sprintf(msg, data...))
}

func (l *toolLoggerAdapter) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.cfg.level < gormLogger.Error {
		return
	}
	logger.Error(ctx, fmt.Sprintf(msg, data...))
}

func (l *toolLoggerAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.cfg.level == gormLogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	// 错误优先，避免被慢查询或常规日志分支吞掉。
	if err != nil && (l.cfg.level >= gormLogger.Error) {
		// 可选忽略 gorm.ErrRecordNotFound，降低噪音。
		if l.cfg.ignoreRecordNotFoundError && errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		logger.Error(ctx, "gorm sql error",
			"duration", elapsed.String(),
			"rows", rows,
			"sql", sql,
			"err", err.Error(),
		)
		return
	}

	// 慢查询告警。
	if l.cfg.slowThreshold > 0 && elapsed > l.cfg.slowThreshold && l.cfg.level >= gormLogger.Warn {
		logger.Warn(ctx, "gorm slow sql",
			"duration", elapsed.String(),
			"rows", rows,
			"sql", sql,
		)
		return
	}

	// 常规 SQL 跟踪。
	if l.cfg.level >= gormLogger.Info {
		logger.Debug(ctx, "gorm sql",
			"duration", elapsed.String(),
			"rows", rows,
			"sql", sql,
		)
	}
}

func parseGormLogLevel(s string) (gormLogger.LogLevel, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "", "silent":
		return gormLogger.Silent, nil
	case "error":
		return gormLogger.Error, nil
	case "warn", "warning":
		return gormLogger.Warn, nil
	case "info":
		return gormLogger.Info, nil
	default:
		return gormLogger.Silent, fmt.Errorf("unknown gorm log level: %s", s)
	}
}
