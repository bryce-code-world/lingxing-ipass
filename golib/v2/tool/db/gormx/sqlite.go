package gormx

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	gormSQLite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

// OpenSQLite 初始化并打开 gorm 客户端（SQLite）。
func OpenSQLite(ctx context.Context, cfg Config) (*gorm.DB, error) {
	cfg, err := normalizeConfig(cfg)
	if err != nil {
		return nil, err
	}

	dsn := strings.TrimSpace(cfg.DSN)
	if dsn == "" {
		dsn, err = buildSQLiteDSN(cfg.SQLite)
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

	gdb, err := gorm.Open(gormSQLite.Dialector{
		DriverName: "sqlite",
		DSN:        dsn,
	}, gormCfg)
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
		if err := startupCheck(ctx, gdb, cfg.StartupCheck.Timeout, "sqlite"); err != nil {
			_ = sqlDB.Close()
			return nil, err
		}
	}
	return gdb, nil
}

// OpenSQLiteWithConn 使用外部传入的 *sql.DB 打开 gorm 客户端（SQLite）。
func OpenSQLiteWithConn(ctx context.Context, conn *sql.DB, cfg Config) (*gorm.DB, error) {
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

	dialector := gormSQLite.Dialector{Conn: conn}
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
		if err := startupCheck(ctx, gdb, cfg.StartupCheck.Timeout, "sqlite"); err != nil {
			return nil, err
		}
	}
	return gdb, nil
}
