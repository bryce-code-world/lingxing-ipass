package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadEnvYAML(path string) (EnvConfig, error) {
	var cfg EnvConfig
	b, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return cfg, err
	}
	applyEnvOverride(&cfg)
	applyDefaults(&cfg)
	return cfg, nil
}

func applyDefaults(cfg *EnvConfig) {
	if cfg.Admin.DisplayTimezone == "" {
		cfg.Admin.DisplayTimezone = "UTC"
	}
}

func applyEnvOverride(cfg *EnvConfig) {
	// Minimal, explicit override set (predictable + documented).
	// Priority: env var > env.yaml.
	if v := os.Getenv("IPASS_BASE_LISTEN_ADDR"); v != "" {
		cfg.Base.ListenAddr = v
	}
	if v := os.Getenv("IPASS_BASE_ENV"); v != "" {
		cfg.Base.Env = v
	}

	if v := os.Getenv("IPASS_ADMIN_PASSWORD"); v != "" {
		cfg.Admin.Password = v
	}
	if v := os.Getenv("IPASS_ADMIN_DISPLAY_TIMEZONE"); v != "" {
		cfg.Admin.DisplayTimezone = v
	}
	if v := os.Getenv("IPASS_ADMIN_EXPORT_DIR"); v != "" {
		cfg.Admin.Export.Dir = v
	}

	if v := os.Getenv("IPASS_AUTH_DSCO_TOKEN"); v != "" {
		cfg.Auth.DSCO.Token = v
	}
	if v := os.Getenv("IPASS_AUTH_LINGXING_APP_ID"); v != "" {
		cfg.Auth.LingXing.AppID = v
	}
	if v := os.Getenv("IPASS_AUTH_LINGXING_APP_SECRET"); v != "" {
		cfg.Auth.LingXing.AppSecret = v
	}

	if v := os.Getenv("IPASS_INTEGRATION_DSCO_BASE_URL"); v != "" {
		cfg.Integration.DSCO.BaseURL = v
	}
	if v := os.Getenv("IPASS_INTEGRATION_LINGXING_BASE_URL"); v != "" {
		cfg.Integration.LingXing.BaseURL = v
	}
}

func ValidateEnv(cfg EnvConfig) error {
	if cfg.Base.ListenAddr == "" {
		return errors.New("base.listen_addr 不能为空")
	}
	switch cfg.Base.Env {
	case "dev", "test", "release":
	default:
		return errors.New("base.env 必须为 dev/test/release")
	}
	if cfg.Admin.Password == "" {
		return errors.New("admin.password 不能为空")
	}
	// display_timezone default is applied in LoadEnvYAML.
	if cfg.Admin.Export.Dir == "" {
		return errors.New("admin.export.dir 不能为空")
	}
	if cfg.Admin.Export.MaxRangeDays <= 0 {
		return errors.New("admin.export.max_range_days 必须为正整数")
	}
	if cfg.Admin.Export.CleanupThresholdBytes <= 0 {
		return errors.New("admin.export.cleanup_threshold_bytes 必须为正整数")
	}
	if cfg.Auth.DSCO.Token == "" {
		return errors.New("auth.dsco.token 不能为空")
	}
	if cfg.Auth.LingXing.AppID == "" || cfg.Auth.LingXing.AppSecret == "" {
		return errors.New("auth.lingxing.app_id/app_secret 不能为空")
	}
	if cfg.Integration.DSCO.BaseURL == "" {
		// optional: SDK has defaults
	}
	if cfg.Integration.LingXing.BaseURL == "" {
		// optional: SDK has defaults
	}
	if cfg.Integration.LingXing.PlatformCode == 0 {
		return errors.New("integration.lingxing.platform_code 不能为空/必须为正整数")
	}

	return nil
}
