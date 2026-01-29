package db

import (
	"context"

	"example.com/lingxing/golib/v2/tool/db/gormx"
	"gorm.io/gorm"
)

func OpenPostgres(ctx context.Context, cfg gormx.Config) (*gorm.DB, error) {
	return gormx.OpenPostgres(ctx, cfg)
}

func Close(gdb *gorm.DB) error {
	return gormx.Close(gdb)
}
