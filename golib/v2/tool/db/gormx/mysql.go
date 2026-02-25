package gormx

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	gormMySQL "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OpenMySQL 初始化并打开 gorm 客户端（MySQL）。
//
// 约定：
// - cfg.DSN 非空时优先使用；否则根据 cfg.MySQL 组装 DSN。
// - 默认执行一次 SELECT 1 启动自检，可通过 cfg.StartupCheck.Enabled 关闭。
func OpenMySQL(ctx context.Context, cfg Config) (*gorm.DB, error) {
	cfg, err := normalizeConfig(cfg)
	if err != nil {
		return nil, err
	}

	dsn, err := buildDSN(cfg)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(dsn) == "" {
		return nil, errors.New("dsn 不能为空")
	}

	gormCfg := &gorm.Config{
		SkipDefaultTransaction: cfg.Gorm.SkipDefaultTransaction,
		PrepareStmt:            cfg.Gorm.PrepareStmt,
		Logger:                 buildGormLogger(cfg),
	}

	gdb, err := gorm.Open(gormMySQL.Open(dsn), gormCfg)
	if err != nil {
		return nil, err
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, err
	}
	if err := applyPool(sqlDB, cfg.Pool); err != nil {
		_ = sqlDB.Close()
		return nil, err
	}

	if startupCheckEnabled(cfg.StartupCheck) {
		if err := startupCheck(ctx, gdb, cfg.StartupCheck.Timeout, "mysql"); err != nil {
			_ = sqlDB.Close()
			return nil, err
		}
	}
	return gdb, nil
}

// OpenMySQLWithConn 使用外部传入的 *sql.DB 打开 gorm 客户端（MySQL）。
func OpenMySQLWithConn(ctx context.Context, conn *sql.DB, cfg Config) (*gorm.DB, error) {
	if conn == nil {
		return nil, errors.New("conn 不能为空")
	}

	cfg, err := normalizeConfig(cfg)
	if err != nil {
		return nil, err
	}

	gormCfg := &gorm.Config{
		SkipDefaultTransaction: cfg.Gorm.SkipDefaultTransaction,
		PrepareStmt:            cfg.Gorm.PrepareStmt,
		Logger:                 buildGormLogger(cfg),
	}

	dialector := gormMySQL.New(gormMySQL.Config{
		Conn:                      conn,
		SkipInitializeWithVersion: cfg.MySQL.SkipInitializeWithVersion,
	})

	gdb, err := gorm.Open(dialector, gormCfg)
	if err != nil {
		return nil, err
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, err
	}
	if err := applyPool(sqlDB, cfg.Pool); err != nil {
		return nil, err
	}

	if startupCheckEnabled(cfg.StartupCheck) {
		if err := startupCheck(ctx, gdb, cfg.StartupCheck.Timeout, "mysql"); err != nil {
			return nil, err
		}
	}
	return gdb, nil
}
