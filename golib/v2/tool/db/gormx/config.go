package gormx

import "time"

// Config 表示 gorm 客户端总配置。
type Config struct {
	// DSN 为完整连接串，非空时优先使用。
	DSN string `json:"dsn" yaml:"dsn" mapstructure:"dsn"`

	MySQL        MySQLConfig        `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Postgres     PostgresConfig     `json:"postgres" yaml:"postgres" mapstructure:"postgres"`
	SQLite       SQLiteConfig       `json:"sqlite" yaml:"sqlite" mapstructure:"sqlite"`
	Pool         PoolConfig         `json:"pool" yaml:"pool" mapstructure:"pool"`
	Gorm         GormConfig         `json:"gorm" yaml:"gorm" mapstructure:"gorm"`
	StartupCheck StartupCheckConfig `json:"startupCheck" yaml:"startupCheck" mapstructure:"startupCheck"`
	Logger       LoggerConfig       `json:"logger" yaml:"logger" mapstructure:"logger"`
}

// MySQLConfig 表示 MySQL DSN 组装配置。
type MySQLConfig struct {
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`
	User     string `json:"user" yaml:"user" mapstructure:"user"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	DBName   string `json:"dbName" yaml:"dbName" mapstructure:"dbName"`

	// Net 连接协议，通常是 tcp。
	Net string `json:"net" yaml:"net" mapstructure:"net"`

	// 超时配置，会写入 DSN 参数。
	Timeout      time.Duration `json:"timeout" yaml:"timeout" mapstructure:"timeout"`
	ReadTimeout  time.Duration `json:"readTimeout" yaml:"readTimeout" mapstructure:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout" yaml:"writeTimeout" mapstructure:"writeTimeout"`

	// 编码相关配置。
	Charset   string `json:"charset" yaml:"charset" mapstructure:"charset"`
	Collation string `json:"collation" yaml:"collation" mapstructure:"collation"`

	// ParseTime 控制时间字段是否解析为 time.Time。
	ParseTime *bool `json:"parseTime" yaml:"parseTime" mapstructure:"parseTime"`

	// Loc 表示时区，例如 Local / Asia/Shanghai。
	Loc string `json:"loc" yaml:"loc" mapstructure:"loc"`

	// Params 是额外 DSN 参数。
	Params map[string]string `json:"params" yaml:"params" mapstructure:"params"`

	// SkipInitializeWithVersion 常用于 sqlmock 测试场景。
	SkipInitializeWithVersion bool `json:"skipInitializeWithVersion" yaml:"skipInitializeWithVersion" mapstructure:"skipInitializeWithVersion"`
}

// PostgresConfig 表示 PostgreSQL DSN 组装配置。
type PostgresConfig struct {
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`
	User     string `json:"user" yaml:"user" mapstructure:"user"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	DBName   string `json:"dbName" yaml:"dbName" mapstructure:"dbName"`

	// SSLMode 例如 disable / require / verify-full。
	SSLMode string `json:"sslMode" yaml:"sslMode" mapstructure:"sslMode"`

	// TimeZone 会写入 TimeZone query 参数。
	TimeZone string `json:"timeZone" yaml:"timeZone" mapstructure:"timeZone"`

	// ClientEncoding 会写入 client_encoding。
	ClientEncoding string `json:"clientEncoding" yaml:"clientEncoding" mapstructure:"clientEncoding"`

	// SearchPath 会写入 search_path。
	SearchPath string `json:"searchPath" yaml:"searchPath" mapstructure:"searchPath"`

	// ConnectTimeout 会写入 connect_timeout（秒）。
	ConnectTimeout time.Duration `json:"connectTimeout" yaml:"connectTimeout" mapstructure:"connectTimeout"`

	// StatementTimeout 会写入 statement_timeout（毫秒）。
	StatementTimeout time.Duration `json:"statementTimeout" yaml:"statementTimeout" mapstructure:"statementTimeout"`

	// ApplicationName 会写入 application_name。
	ApplicationName string `json:"applicationName" yaml:"applicationName" mapstructure:"applicationName"`

	// PreferSimpleProtocol 透传给 gorm postgres driver。
	PreferSimpleProtocol bool `json:"preferSimpleProtocol" yaml:"preferSimpleProtocol" mapstructure:"preferSimpleProtocol"`

	// Params 是额外 DSN query 参数。
	Params map[string]string `json:"params" yaml:"params" mapstructure:"params"`
}

// SQLiteConfig 表示 SQLite DSN 组装配置。
type SQLiteConfig struct {
	// Path 表示 sqlite 文件路径，例如 ./data.db。
	Path string `json:"path" yaml:"path" mapstructure:"path"`

	// DSN 非空时优先使用，例如 file::memory:?cache=shared。
	DSN string `json:"dsn" yaml:"dsn" mapstructure:"dsn"`
}

// PoolConfig 表示连接池配置。
type PoolConfig struct {
	MaxOpenConns    int           `json:"maxOpenConns" yaml:"maxOpenConns" mapstructure:"maxOpenConns"`
	MaxIdleConns    int           `json:"maxIdleConns" yaml:"maxIdleConns" mapstructure:"maxIdleConns"`
	ConnMaxLifetime time.Duration `json:"connMaxLifetime" yaml:"connMaxLifetime" mapstructure:"connMaxLifetime"`
	ConnMaxIdleTime time.Duration `json:"connMaxIdleTime" yaml:"connMaxIdleTime" mapstructure:"connMaxIdleTime"`
}

// GormConfig 表示 gorm 本身的常用开关。
type GormConfig struct {
	SkipDefaultTransaction bool `json:"skipDefaultTransaction" yaml:"skipDefaultTransaction" mapstructure:"skipDefaultTransaction"`
	PrepareStmt            bool `json:"prepareStmt" yaml:"prepareStmt" mapstructure:"prepareStmt"`
}

// StartupCheckConfig 控制启动连通性检查。
type StartupCheckConfig struct {
	// Enabled 为 nil 时默认开启。
	Enabled *bool         `json:"enabled" yaml:"enabled" mapstructure:"enabled"`
	Timeout time.Duration `json:"timeout" yaml:"timeout" mapstructure:"timeout"`
}

// LoggerConfig 控制 gorm 日志行为。
type LoggerConfig struct {
	Enabled bool   `json:"enabled" yaml:"enabled" mapstructure:"enabled"`
	Level   string `json:"level" yaml:"level" mapstructure:"level"`

	SlowThreshold             time.Duration `json:"slowThreshold" yaml:"slowThreshold" mapstructure:"slowThreshold"`
	IgnoreRecordNotFoundError bool          `json:"ignoreRecordNotFoundError" yaml:"ignoreRecordNotFoundError" mapstructure:"ignoreRecordNotFoundError"`
}
