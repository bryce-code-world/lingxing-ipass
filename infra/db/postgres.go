package db

import (
	"context"

	"gitee.com/lsy007/golibv2/v2/tool/db/gormx"
	"gorm.io/gorm"
)

func OpenPostgres(ctx context.Context, cfg gormx.Config) (*gorm.DB, error) {
	return gormx.OpenPostgres(ctx, cfg)
}

func Close(gdb *gorm.DB) error {
	return gormx.Close(gdb)
}
