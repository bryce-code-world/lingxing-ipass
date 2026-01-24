package db

import (
	"errors"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OpenMySQL 打开 MySQL 连接（一期最小化，使用 GORM）。
//
// 约定：
// - 不直接使用 database/sql（由 GORM 承载）。
// - 使用最小可行的连接自检：SELECT 1。
func OpenMySQL(dsn string) (*gorm.DB, error) {
	if strings.TrimSpace(dsn) == "" {
		return nil, errors.New("dsn 不能为空")
	}

	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 最小自检：确保连接可用。
	var one int
	if err := gdb.Raw("SELECT 1").Scan(&one).Error; err != nil {
		return nil, err
	}
	return gdb, nil
}
