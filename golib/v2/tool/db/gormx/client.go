package gormx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Close 关闭 gorm 底层连接。
func Close(gdb *gorm.DB) error {
	if gdb == nil {
		return nil
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func startupCheck(ctx context.Context, gdb *gorm.DB, timeout time.Duration, dbType string) error {
	checkCtx := ctx
	cancel := func() {}
	if timeout > 0 {
		checkCtx, cancel = context.WithTimeout(ctx, timeout)
	}
	defer cancel()

	var one int
	if err := gdb.WithContext(checkCtx).Raw("SELECT 1").Scan(&one).Error; err != nil {
		return fmt.Errorf("%s 启动自检失败: %w", dbType, err)
	}
	if one != 1 {
		return fmt.Errorf("%s 启动自检失败: SELECT 1 返回结果异常", dbType)
	}
	return nil
}

func applyPool(sqlDB *sql.DB, pool PoolConfig) error {
	if sqlDB == nil {
		return errors.New("sqlDB 不能为空")
	}

	if pool.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(pool.MaxOpenConns)
	}
	if pool.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(pool.MaxIdleConns)
	}
	if pool.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(pool.ConnMaxLifetime)
	}
	if pool.ConnMaxIdleTime > 0 {
		sqlDB.SetConnMaxIdleTime(pool.ConnMaxIdleTime)
	}
	return nil
}

func buildGormLogger(cfg Config) gormLogger.Interface {
	if !cfg.Logger.Enabled {
		return gormLogger.Default.LogMode(gormLogger.Silent)
	}
	lv, err := parseGormLogLevel(cfg.Logger.Level)
	if err != nil {
		return gormLogger.Default.LogMode(gormLogger.Silent)
	}
	return newToolLoggerAdapter(toolLoggerAdapterConfig{
		level:                     lv,
		slowThreshold:             cfg.Logger.SlowThreshold,
		ignoreRecordNotFoundError: cfg.Logger.IgnoreRecordNotFoundError,
	})
}

func startupCheckEnabled(cfg StartupCheckConfig) bool {
	if cfg.Enabled == nil {
		return true
	}
	return *cfg.Enabled
}
