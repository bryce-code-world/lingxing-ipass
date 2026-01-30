package config

import (
	"example.com/lingxing/golib/v2/tool/db/gormx"
	"example.com/lingxing/golib/v2/tool/logger"
)

type EnvConfig struct {
	Base        BaseConfig        `yaml:"base"`
	Auth        AuthConfig        `yaml:"auth"`
	Integration IntegrationConfig `yaml:"integration"`
	DB          gormx.Config      `yaml:"db"`
	Admin       AdminConfig       `yaml:"admin"`
	Log         logger.Config     `yaml:"log"`
}

type BaseConfig struct {
	ListenAddr string `yaml:"listen_addr"`
	Env        string `yaml:"env"` // dev/test/release
}

type AuthConfig struct {
	DSCO     DSCOAuthConfig     `yaml:"dsco"`
	LingXing LingXingAuthConfig `yaml:"lingxing"`
}

type DSCOAuthConfig struct {
	Token string `yaml:"token"`
}

type LingXingAuthConfig struct {
	AppID     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
}

type IntegrationConfig struct {
	DSCO     DSCOIntegrationConfig     `yaml:"dsco"`
	LingXing LingXingIntegrationConfig `yaml:"lingxing"`
}

type DSCOIntegrationConfig struct {
	// BaseURL DSCO API 基础地址；为空则使用 SDK 默认 BaseURLProd。
	BaseURL string `yaml:"base_url"`
}

type LingXingIntegrationConfig struct {
	// BaseURL 领星 API 基础地址；为空则使用 SDK 默认 BaseURLProd。
	BaseURL string `yaml:"base_url"`
	// PlatformCode 领星平台编码（固定值；与业务相关，保留在 env.yaml）。
	PlatformCode int `yaml:"platform_code"`
}

type AdminConfig struct {
	Password        string       `yaml:"password"`
	DisplayTimezone string       `yaml:"display_timezone"`
	Export          ExportConfig `yaml:"export"`
}

type ExportConfig struct {
	Dir                   string `yaml:"dir"`
	MaxRangeDays          int    `yaml:"max_range_days"`
	CleanupThresholdBytes int64  `yaml:"cleanup_threshold_bytes"`
}
