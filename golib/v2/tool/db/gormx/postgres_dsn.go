package gormx

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const defaultPostgresPort = 5432

func buildPostgresDSN(cfg PostgresConfig) (string, error) {
	if strings.TrimSpace(cfg.Host) == "" {
		return "", errors.New("postgres.host 不能为空")
	}
	if strings.TrimSpace(cfg.User) == "" {
		return "", errors.New("postgres.user 不能为空")
	}
	if strings.TrimSpace(cfg.DBName) == "" {
		return "", errors.New("postgres.dbName 不能为空")
	}

	port := cfg.Port
	if port <= 0 {
		port = defaultPostgresPort
	}

	q := url.Values{}
	for k, v := range cfg.Params {
		k = strings.TrimSpace(k)
		if k == "" {
			continue
		}
		q.Set(k, strings.TrimSpace(v))
	}

	if strings.TrimSpace(cfg.SSLMode) != "" {
		q.Set("sslmode", strings.TrimSpace(cfg.SSLMode))
	}
	if strings.TrimSpace(cfg.TimeZone) != "" {
		q.Set("TimeZone", strings.TrimSpace(cfg.TimeZone))
	}
	if strings.TrimSpace(cfg.ClientEncoding) != "" {
		q.Set("client_encoding", strings.TrimSpace(cfg.ClientEncoding))
	}
	if strings.TrimSpace(cfg.SearchPath) != "" {
		q.Set("search_path", strings.TrimSpace(cfg.SearchPath))
	}
	if strings.TrimSpace(cfg.ApplicationName) != "" {
		q.Set("application_name", strings.TrimSpace(cfg.ApplicationName))
	}

	if cfg.ConnectTimeout > 0 {
		sec := int((cfg.ConnectTimeout + time.Second - 1) / time.Second)
		if sec < 1 {
			sec = 1
		}
		q.Set("connect_timeout", fmt.Sprintf("%d", sec))
	}
	if cfg.StatementTimeout > 0 {
		ms := int((cfg.StatementTimeout + time.Millisecond - 1) / time.Millisecond)
		if ms < 1 {
			ms = 1
		}
		q.Set("statement_timeout", fmt.Sprintf("%d", ms))
	}

	u := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", cfg.Host, port),
		Path:   "/" + cfg.DBName,
	}
	if cfg.Password != "" {
		u.User = url.UserPassword(cfg.User, cfg.Password)
	} else {
		u.User = url.User(cfg.User)
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}
	return u.String(), nil
}
