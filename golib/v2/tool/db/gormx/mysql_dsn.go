package gormx

import (
	"errors"
	"fmt"
	"strings"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
)

const (
	defaultMySQLPort      = 3306
	defaultMySQLNet       = "tcp"
	defaultMySQLCharset   = "utf8mb4"
	defaultMySQLCollation = "utf8mb4_0900_ai_ci"
)

func normalizeConfig(cfg Config) (Config, error) {
	cfg.DSN = strings.TrimSpace(cfg.DSN)

	cfg.MySQL.Host = strings.TrimSpace(cfg.MySQL.Host)
	cfg.MySQL.User = strings.TrimSpace(cfg.MySQL.User)
	cfg.MySQL.DBName = strings.TrimSpace(cfg.MySQL.DBName)
	cfg.MySQL.Net = strings.TrimSpace(cfg.MySQL.Net)
	cfg.MySQL.Charset = strings.TrimSpace(cfg.MySQL.Charset)
	cfg.MySQL.Collation = strings.TrimSpace(cfg.MySQL.Collation)
	cfg.MySQL.Loc = strings.TrimSpace(cfg.MySQL.Loc)

	cfg.Postgres.Host = strings.TrimSpace(cfg.Postgres.Host)
	cfg.Postgres.User = strings.TrimSpace(cfg.Postgres.User)
	cfg.Postgres.DBName = strings.TrimSpace(cfg.Postgres.DBName)
	cfg.Postgres.SSLMode = strings.TrimSpace(cfg.Postgres.SSLMode)
	cfg.Postgres.TimeZone = strings.TrimSpace(cfg.Postgres.TimeZone)
	cfg.Postgres.ClientEncoding = strings.TrimSpace(cfg.Postgres.ClientEncoding)
	cfg.Postgres.SearchPath = strings.TrimSpace(cfg.Postgres.SearchPath)
	cfg.Postgres.ApplicationName = strings.TrimSpace(cfg.Postgres.ApplicationName)

	cfg.SQLite.Path = strings.TrimSpace(cfg.SQLite.Path)
	cfg.SQLite.DSN = strings.TrimSpace(cfg.SQLite.DSN)

	if cfg.MySQL.Port == 0 {
		cfg.MySQL.Port = defaultMySQLPort
	}
	if cfg.MySQL.Net == "" {
		cfg.MySQL.Net = defaultMySQLNet
	}
	if cfg.Postgres.Port == 0 {
		cfg.Postgres.Port = defaultPostgresPort
	}
	if strings.TrimSpace(cfg.Logger.Level) == "" {
		cfg.Logger.Level = "silent"
	}
	if cfg.StartupCheck.Timeout <= 0 {
		cfg.StartupCheck.Timeout = 5 * time.Second
	}

	return cfg, nil
}

func buildDSN(cfg Config) (string, error) {
	if cfg.DSN != "" {
		return cfg.DSN, nil
	}
	return buildMySQLDSN(cfg.MySQL)
}

func buildMySQLDSN(cfg MySQLConfig) (string, error) {
	if strings.TrimSpace(cfg.Host) == "" {
		return "", errors.New("mysql.host 不能为空")
	}
	if strings.TrimSpace(cfg.User) == "" {
		return "", errors.New("mysql.user 不能为空")
	}
	if strings.TrimSpace(cfg.DBName) == "" {
		return "", errors.New("mysql.dbName 不能为空")
	}

	port := cfg.Port
	if port <= 0 {
		port = defaultMySQLPort
	}

	netName := strings.TrimSpace(cfg.Net)
	if netName == "" {
		netName = defaultMySQLNet
	}

	charset := strings.TrimSpace(cfg.Charset)
	if charset == "" {
		charset = defaultMySQLCharset
	}

	collation := strings.TrimSpace(cfg.Collation)
	if collation == "" {
		collation = defaultMySQLCollation
	}

	locStr := strings.TrimSpace(cfg.Loc)
	if locStr == "" {
		locStr = "Local"
	}
	loc, err := time.LoadLocation(locStr)
	if err != nil {
		return "", fmt.Errorf("mysql.loc 非法: %w", err)
	}

	params := map[string]string{}
	for k, v := range cfg.Params {
		k = strings.TrimSpace(k)
		if k == "" {
			continue
		}
		params[k] = strings.TrimSpace(v)
	}
	params["charset"] = charset

	mysqlCfg := mysqlDriver.NewConfig()
	mysqlCfg.User = cfg.User
	mysqlCfg.Passwd = cfg.Password
	mysqlCfg.Net = netName
	mysqlCfg.Addr = fmt.Sprintf("%s:%d", cfg.Host, port)
	mysqlCfg.DBName = cfg.DBName
	mysqlCfg.Params = params
	mysqlCfg.Collation = collation
	mysqlCfg.Timeout = cfg.Timeout
	mysqlCfg.ReadTimeout = cfg.ReadTimeout
	mysqlCfg.WriteTimeout = cfg.WriteTimeout
	mysqlCfg.Loc = loc

	mysqlCfg.ParseTime = true
	if cfg.ParseTime != nil {
		mysqlCfg.ParseTime = *cfg.ParseTime
	}

	return mysqlCfg.FormatDSN(), nil
}
