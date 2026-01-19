package db

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OpenMySQL 打开 MySQL 并返回 GORM DB（admin 独立使用）。
func OpenMySQL(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, errors.New("dsn 不能为空")
	}
	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := gdb.Raw("SELECT 1").Error; err != nil {
		return nil, err
	}
	return gdb, nil
}

