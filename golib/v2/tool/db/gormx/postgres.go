package gormx

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenPostgres 初始化并打开 gorm 客户端（PostgreSQL）。
//
// 约定：
// - cfg.DSN 非空时优先使用；否则根据 cfg.Postgres 组装 DSN。
// - 默认执行一次 SELECT 1 启动自检，可通过 cfg.StartupCheck.Enabled 关闭。
func OpenPostgres(ctx context.Context, cfg Config) (*gorm.DB, error) {
	cfg, err := normalizeConfig(cfg)
	if err != nil {
		return nil, err
	}

	dsn := strings.TrimSpace(cfg.DSN)
	if dsn == "" {
		dsn, err = buildPostgresDSN(cfg.Postgres)
		if err != nil {
			return nil, err
		}
	}
	if strings.TrimSpace(dsn) == "" {
		return nil, errors.New("dsn 不能为空")
	}

	gormCfg := &gorm.Config{
		SkipDefaultTransaction: cfg.Gorm.SkipDefaultTransaction,
		PrepareStmt:            cfg.Gorm.PrepareStmt,
		Logger:                 buildGormLogger(cfg),
	}

	gdb, err := gorm.Open(gormPostgres.New(gormPostgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: cfg.Postgres.PreferSimpleProtocol,
	}), gormCfg)
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
		if err := startupCheck(ctx, gdb, cfg.StartupCheck.Timeout, "postgres"); err != nil {
			_ = sqlDB.Close()
			return nil, err
		}
	}
	return gdb, nil
}

// OpenPostgresWithConn 使用外部传入的 *sql.DB 打开 gorm 客户端（PostgreSQL）。
func OpenPostgresWithConn(ctx context.Context, conn *sql.DB, cfg Config) (*gorm.DB, error) {
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

	gdb, err := gorm.Open(gormPostgres.New(gormPostgres.Config{
		Conn:                 conn,
		PreferSimpleProtocol: cfg.Postgres.PreferSimpleProtocol,
	}), gormCfg)
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
		if err := startupCheck(ctx, gdb, cfg.StartupCheck.Timeout, "postgres"); err != nil {
			return nil, err
		}
	}
	return gdb, nil
}
