package db

import (
	"context"

	"gorm.io/gorm"
	"lingxingipass/golib/v2/tool/db/gormx"
)

func OpenPostgres(ctx context.Context, cfg gormx.Config) (*gorm.DB, error) {
	return gormx.OpenPostgres(ctx, cfg)
}

func Close(gdb *gorm.DB) error {
	return gormx.Close(gdb)
}
