package config

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Config 表示一期服务配置。
//
// 约定：
// - System：系统配置（基础设施、外部系统连接、可靠性、调度等）。
// - Biz：业务配置（映射/口径等可随业务调整的内容）。
type Config struct {
	System SystemConfig
	Biz    BizConfig
}

// SystemConfig 表示系统配置（尽量稳定，不随业务频繁变更）。
type SystemConfig struct {
	DB struct {
		DSN string
	}
	Log struct {
		Dir string
	}
	HTTP struct {
		Enable bool
		Addr   string
	}
	Ops struct {
		// Password 用于业务进程 ops 接口鉴权（admin 将通过请求头 X-Ops-Password 调用）。
		Password string
	}
	Admin struct {
		// Password 用于管理后台访问控制（一期最小：单密码）。
		Password string
	}
	DSCO struct {
		BaseURL string
		Token   string
	}
	LingXing struct {
		BaseURL      string
		AppID        string
		AccessToken  string
		PlatformCode int
		StoreID      string
		// SID 用于 WMS/仓库相关接口（请求字段 sid_arr）。
		SID int
	}
	Reliability struct {
		// MaxRetryPerOrder 表示同一 dscoOrderId 在某个环节失败后，可被“再次处理”的最大次数。
		// 达到上限后会转人工，避免无限重试。
		MaxRetryPerOrder int
	}
	Jobs struct {
		HeartbeatIntervalSec int

		PullDSCOOrdersEnable      bool
		PullDSCOOrdersIntervalSec int

		PushOrdersToLingXingEnable      bool
		PushOrdersToLingXingIntervalSec int
		PushOrdersToLingXingBatchSize   int

		AckToDSCOEnable      bool
		AckToDSCOIntervalSec int

		ShipToDSCOEnable      bool
		ShipToDSCOIntervalSec int

		InvoiceToDSCOEnable      bool
		InvoiceToDSCOIntervalSec int
		InvoiceToDSCOBatchSize   int

		SyncStockEnable      bool
		SyncStockIntervalSec int
		SyncStockBatchSize   int
	}
}

// BizConfig 表示业务配置（口径/映射/策略等）。
type BizConfig struct {
	Stock struct {
		// LingXingWIDToDSCOWarehouseCode 用于库存同步：领星 WID -> DSCO warehouseCode。
		// 通过 env JSON 注入，例如：{"26":"WH1","27":"WH2"}
		LingXingWIDToDSCOWarehouseCode map[string]string

		// LingXingSKUToDSCOSKU 用于库存同步：领星 SKU -> DSCO SKU（可选，默认同名）。
		// 通过 env JSON 注入，例如：{"LXSKU-1":"DSCOSKU-1"}
		LingXingSKUToDSCOSKU map[string]string
	}

	Shipment struct {
		// ShipDateSource 指定发货回传的 shipDate 取值来源：
		// - delivered_at：使用领星 WMS 的 delivered_at
		// - stock_delivered_at：使用领星 WMS 的 stock_delivered_at
		// - none：不回传 shipDate
		ShipDateSource string
	}
}

// LoadFromEnv 从环境变量加载配置（一期最小化实现）。
func LoadFromEnv() (Config, error) {
	var cfg Config

	cfg.System.DB.DSN = strings.TrimSpace(os.Getenv("IPASS_DB_DSN"))
	if cfg.System.DB.DSN == "" {
		return Config{}, errors.New("缺少环境变量 IPASS_DB_DSN")
	}

	cfg.System.Log.Dir = strings.TrimSpace(os.Getenv("IPASS_LOG_DIR"))
	if cfg.System.Log.Dir == "" {
		cfg.System.Log.Dir = "logs"
	}

	cfg.System.HTTP.Enable = envBool("IPASS_HTTP_ENABLE", true)
	cfg.System.HTTP.Addr = strings.TrimSpace(os.Getenv("IPASS_HTTP_ADDR"))
	if cfg.System.HTTP.Addr == "" {
		cfg.System.HTTP.Addr = ":8080"
	}

	cfg.System.Ops.Password = strings.TrimSpace(os.Getenv("IPASS_OPS_PASSWORD"))
	if cfg.System.HTTP.Enable && cfg.System.Ops.Password == "" {
		return Config{}, errors.New("开启业务 ops HTTP 时，必须配置 IPASS_OPS_PASSWORD")
	}

	cfg.System.Admin.Password = strings.TrimSpace(os.Getenv("IPASS_ADMIN_PASSWORD"))

	cfg.System.DSCO.BaseURL = strings.TrimSpace(os.Getenv("IPASS_DSCO_BASE_URL"))
	cfg.System.DSCO.Token = strings.TrimSpace(os.Getenv("IPASS_DSCO_TOKEN"))

	cfg.System.LingXing.BaseURL = strings.TrimSpace(os.Getenv("IPASS_LINGXING_BASE_URL"))
	cfg.System.LingXing.AppID = strings.TrimSpace(os.Getenv("IPASS_LINGXING_APP_ID"))
	cfg.System.LingXing.AccessToken = strings.TrimSpace(os.Getenv("IPASS_LINGXING_ACCESS_TOKEN"))
	cfg.System.LingXing.PlatformCode = envInt("IPASS_LINGXING_PLATFORM_CODE", 0)
	cfg.System.LingXing.StoreID = strings.TrimSpace(os.Getenv("IPASS_LINGXING_STORE_ID"))
	cfg.System.LingXing.SID = envInt("IPASS_LINGXING_SID", 0)

	cfg.System.Reliability.MaxRetryPerOrder = envInt("IPASS_MAX_RETRY_PER_ORDER", 5)
	if cfg.System.Reliability.MaxRetryPerOrder <= 0 {
		return Config{}, errors.New("IPASS_MAX_RETRY_PER_ORDER 必须为正数")
	}

	cfg.Biz.Shipment.ShipDateSource = strings.TrimSpace(os.Getenv("IPASS_SHIP_DATE_SOURCE"))
	if cfg.Biz.Shipment.ShipDateSource == "" {
		cfg.Biz.Shipment.ShipDateSource = "delivered_at"
	}
	switch cfg.Biz.Shipment.ShipDateSource {
	case "delivered_at", "stock_delivered_at", "none":
	default:
		return Config{}, errors.New("IPASS_SHIP_DATE_SOURCE 仅支持 delivered_at/stock_delivered_at/none")
	}

	cfg.System.Jobs.HeartbeatIntervalSec = envInt("IPASS_HEARTBEAT_INTERVAL_SEC", 10)
	if cfg.System.Jobs.HeartbeatIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_HEARTBEAT_INTERVAL_SEC 必须为正数")
	}

	cfg.System.Jobs.PullDSCOOrdersEnable = envBool("IPASS_JOB_PULL_DSCO_ORDERS_ENABLE", false)
	cfg.System.Jobs.PullDSCOOrdersIntervalSec = envInt("IPASS_JOB_PULL_DSCO_ORDERS_INTERVAL_SEC", 30)
	if cfg.System.Jobs.PullDSCOOrdersIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_PULL_DSCO_ORDERS_INTERVAL_SEC 必须为正数")
	}

	cfg.System.Jobs.PushOrdersToLingXingEnable = envBool("IPASS_JOB_PUSH_ORDERS_TO_LINGXING_ENABLE", false)
	cfg.System.Jobs.PushOrdersToLingXingIntervalSec = envInt("IPASS_JOB_PUSH_ORDERS_TO_LINGXING_INTERVAL_SEC", 30)
	if cfg.System.Jobs.PushOrdersToLingXingIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_PUSH_ORDERS_TO_LINGXING_INTERVAL_SEC 必须为正数")
	}
	cfg.System.Jobs.PushOrdersToLingXingBatchSize = envInt("IPASS_JOB_PUSH_ORDERS_TO_LINGXING_BATCH_SIZE", 50)
	if cfg.System.Jobs.PushOrdersToLingXingBatchSize <= 0 {
		return Config{}, errors.New("IPASS_JOB_PUSH_ORDERS_TO_LINGXING_BATCH_SIZE 必须为正数")
	}

	cfg.System.Jobs.AckToDSCOEnable = envBool("IPASS_JOB_ACK_TO_DSCO_ENABLE", false)
	cfg.System.Jobs.AckToDSCOIntervalSec = envInt("IPASS_JOB_ACK_TO_DSCO_INTERVAL_SEC", 30)
	if cfg.System.Jobs.AckToDSCOIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_ACK_TO_DSCO_INTERVAL_SEC 必须为正数")
	}

	cfg.System.Jobs.ShipToDSCOEnable = envBool("IPASS_JOB_SHIP_TO_DSCO_ENABLE", false)
	cfg.System.Jobs.ShipToDSCOIntervalSec = envInt("IPASS_JOB_SHIP_TO_DSCO_INTERVAL_SEC", 30)
	if cfg.System.Jobs.ShipToDSCOIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_SHIP_TO_DSCO_INTERVAL_SEC 必须为正数")
	}

	cfg.System.Jobs.InvoiceToDSCOEnable = envBool("IPASS_JOB_INVOICE_TO_DSCO_ENABLE", false)
	cfg.System.Jobs.InvoiceToDSCOIntervalSec = envInt("IPASS_JOB_INVOICE_TO_DSCO_INTERVAL_SEC", 60)
	if cfg.System.Jobs.InvoiceToDSCOIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_INVOICE_TO_DSCO_INTERVAL_SEC 必须为正数")
	}
	cfg.System.Jobs.InvoiceToDSCOBatchSize = envInt("IPASS_JOB_INVOICE_TO_DSCO_BATCH_SIZE", 50)
	if cfg.System.Jobs.InvoiceToDSCOBatchSize <= 0 {
		return Config{}, errors.New("IPASS_JOB_INVOICE_TO_DSCO_BATCH_SIZE 必须为正数")
	}

	cfg.System.Jobs.SyncStockEnable = envBool("IPASS_JOB_SYNC_STOCK_ENABLE", false)
	cfg.System.Jobs.SyncStockIntervalSec = envInt("IPASS_JOB_SYNC_STOCK_INTERVAL_SEC", 300)
	if cfg.System.Jobs.SyncStockIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_SYNC_STOCK_INTERVAL_SEC 必须为正数")
	}
	cfg.System.Jobs.SyncStockBatchSize = envInt("IPASS_JOB_SYNC_STOCK_BATCH_SIZE", 200)
	if cfg.System.Jobs.SyncStockBatchSize <= 0 {
		return Config{}, errors.New("IPASS_JOB_SYNC_STOCK_BATCH_SIZE 必须为正数")
	}

	// 按需校验：避免只跑 heartbeat 时也必须配置全部外部系统参数。
	if cfg.System.Jobs.PullDSCOOrdersEnable || cfg.System.Jobs.PushOrdersToLingXingEnable || cfg.System.Jobs.AckToDSCOEnable || cfg.System.Jobs.ShipToDSCOEnable || cfg.System.Jobs.InvoiceToDSCOEnable || cfg.System.Jobs.SyncStockEnable {
		if cfg.System.DSCO.Token == "" {
			return Config{}, errors.New("开启 DSCO 相关任务时，必须配置 IPASS_DSCO_TOKEN")
		}
	}
	if cfg.System.Jobs.PushOrdersToLingXingEnable || cfg.System.Jobs.AckToDSCOEnable || cfg.System.Jobs.ShipToDSCOEnable || cfg.System.Jobs.SyncStockEnable {
		if cfg.System.LingXing.AppID == "" {
			return Config{}, errors.New("开启领星相关任务时，必须配置 IPASS_LINGXING_APP_ID")
		}
		if cfg.System.LingXing.AccessToken == "" {
			return Config{}, errors.New("开启领星相关任务时，必须配置 IPASS_LINGXING_ACCESS_TOKEN")
		}
		if cfg.System.LingXing.PlatformCode <= 0 {
			return Config{}, errors.New("开启领星相关任务时，必须配置 IPASS_LINGXING_PLATFORM_CODE（正整数）")
		}
		if cfg.System.LingXing.StoreID == "" {
			return Config{}, errors.New("开启领星相关任务时，必须配置 IPASS_LINGXING_STORE_ID")
		}
	}
	if cfg.System.Jobs.ShipToDSCOEnable && cfg.System.LingXing.SID <= 0 {
		return Config{}, errors.New("开启发货回传任务时，必须配置 IPASS_LINGXING_SID（正整数）")
	}

	// 业务映射：允许为空（只在库存同步任务启用/执行时才要求必填）。
	var err error
	cfg.Biz.Stock.LingXingWIDToDSCOWarehouseCode, err = envJSONMapStringStringOptional("IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON")
	if err != nil {
		return Config{}, err
	}
	cfg.Biz.Stock.LingXingSKUToDSCOSKU, err = envJSONMapStringStringOptional("IPASS_STOCK_SKU_TO_DSCO_SKU_JSON")
	if err != nil {
		return Config{}, err
	}
	if cfg.System.Jobs.SyncStockEnable && len(cfg.Biz.Stock.LingXingWIDToDSCOWarehouseCode) == 0 {
		return Config{}, errors.New("开启库存同步任务时，必须配置 IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON")
	}

	return cfg, nil
}

func envInt(key string, def int) int {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return def
	}
	n, err := strconv.Atoi(raw)
	if err != nil {
		return def
	}
	return n
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

func envJSONMapStringString(key string) (map[string]string, error) {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return nil, nil
	}
	var m map[string]string
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		return nil, errors.New(key + " 不是合法 JSON")
	}
	return m, nil
}

func envJSONMapStringStringOptional(key string) (map[string]string, error) {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return nil, nil
	}
	var m map[string]string
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		return nil, errors.New(key + " 不是合法 JSON")
	}
	return m, nil
}
