package gormx

import (
	"errors"
	"strings"
)

func buildSQLiteDSN(cfg SQLiteConfig) (string, error) {
	if strings.TrimSpace(cfg.DSN) != "" {
		return strings.TrimSpace(cfg.DSN), nil
	}
	if strings.TrimSpace(cfg.Path) != "" {
		return strings.TrimSpace(cfg.Path), nil
	}
	return "", errors.New("sqlite.dsn/path 不能为空")
}
