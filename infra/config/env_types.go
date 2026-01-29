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
	Shipment ShipmentIntegrationConfig `yaml:"shipment"`
}

type DSCOIntegrationConfig struct {
	BaseURL string                 `yaml:"base_url"`
	Codes   map[string]any         `yaml:"codes"`
	Raw     map[string]interface{} `yaml:",inline"`
}

type LingXingIntegrationConfig struct {
	BaseURL      string `yaml:"base_url"`
	AccessToken  string `yaml:"access_token"`
	PlatformCode int    `yaml:"platform_code"`
	StoreID      string `yaml:"store_id"`
	SID          int    `yaml:"sid"`
}

type ShipmentIntegrationConfig struct {
	ShipDateSource string `yaml:"ship_date_source"` // delivered_at/stock_delivered_at/none
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
