package config

import (
	"errors"
	"os"
	"strings"
)

// Config 表示 admin 独立进程的配置（与业务进程完全分离）。
type Config struct {
	DB struct {
		DSN string
	}
	HTTP struct {
		Enable bool
		Addr   string
	}
	Auth struct {
		// Password 用于 Admin UI/API 登录与鉴权（一期最小：单密码）。
		Password string
	}
	Ops struct {
		// BaseURL 指向业务进程的 ops HTTP 服务，例如：http://127.0.0.1:8080
		BaseURL string
		// Password 用于调用业务 ops（通过请求头 X-Ops-Password）。
		Password string
	}
}

func LoadFromEnv() (Config, error) {
	var cfg Config

	cfg.DB.DSN = strings.TrimSpace(os.Getenv("ADMIN_DB_DSN"))
	if cfg.DB.DSN == "" {
		return Config{}, errors.New("缺少环境变量 ADMIN_DB_DSN")
	}

	cfg.HTTP.Enable = envBool("ADMIN_HTTP_ENABLE", true)
	cfg.HTTP.Addr = strings.TrimSpace(os.Getenv("ADMIN_HTTP_ADDR"))
	if cfg.HTTP.Addr == "" {
		cfg.HTTP.Addr = ":8081"
	}

	cfg.Auth.Password = strings.TrimSpace(os.Getenv("ADMIN_PASSWORD"))
	if cfg.Auth.Password == "" {
		return Config{}, errors.New("缺少环境变量 ADMIN_PASSWORD")
	}

	cfg.Ops.BaseURL = strings.TrimSpace(os.Getenv("ADMIN_OPS_BASE_URL"))
	if cfg.Ops.BaseURL == "" {
		return Config{}, errors.New("缺少环境变量 ADMIN_OPS_BASE_URL")
	}
	cfg.Ops.Password = strings.TrimSpace(os.Getenv("ADMIN_OPS_PASSWORD"))
	if cfg.Ops.Password == "" {
		return Config{}, errors.New("缺少环境变量 ADMIN_OPS_PASSWORD")
	}

	return cfg, nil
}

func envBool(key string, def bool) bool {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return def
	}
	switch strings.ToLower(raw) {
	case "1", "true", "yes", "y", "on":
		return true
	case "0", "false", "no", "n", "off":
		return false
	default:
		return def
	}
}
