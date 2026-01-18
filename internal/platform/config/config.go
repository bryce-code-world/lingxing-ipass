package config

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Config 表示一期服务配置。
type Config struct {
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
	Stock struct {
		// LingXingWIDToDSCOWarehouseCode 用于库存同步：领星 WID -> DSCO warehouseCode。
		// 通过 env JSON 注入，例如：{"26":"WH1","27":"WH2"}
		LingXingWIDToDSCOWarehouseCode map[string]string

		// LingXingSKUToDSCOSKU 用于库存同步：领星 SKU -> DSCO SKU（可选，默认同名）。
		// 通过 env JSON 注入，例如：{"LXSKU-1":"DSCOSKU-1"}
		LingXingSKUToDSCOSKU map[string]string
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

// LoadFromEnv 从环境变量加载配置（一期最小化实现）。
func LoadFromEnv() (Config, error) {
	var cfg Config

	cfg.DB.DSN = strings.TrimSpace(os.Getenv("IPASS_DB_DSN"))
	if cfg.DB.DSN == "" {
		return Config{}, errors.New("缺少环境变量 IPASS_DB_DSN")
	}

	cfg.Log.Dir = strings.TrimSpace(os.Getenv("IPASS_LOG_DIR"))
	if cfg.Log.Dir == "" {
		cfg.Log.Dir = "logs"
	}

	cfg.HTTP.Enable = envBool("IPASS_HTTP_ENABLE", true)
	cfg.HTTP.Addr = strings.TrimSpace(os.Getenv("IPASS_HTTP_ADDR"))
	if cfg.HTTP.Addr == "" {
		cfg.HTTP.Addr = ":8080"
	}

	cfg.DSCO.BaseURL = strings.TrimSpace(os.Getenv("IPASS_DSCO_BASE_URL"))
	cfg.DSCO.Token = strings.TrimSpace(os.Getenv("IPASS_DSCO_TOKEN"))

	cfg.LingXing.BaseURL = strings.TrimSpace(os.Getenv("IPASS_LINGXING_BASE_URL"))
	cfg.LingXing.AppID = strings.TrimSpace(os.Getenv("IPASS_LINGXING_APP_ID"))
	cfg.LingXing.AccessToken = strings.TrimSpace(os.Getenv("IPASS_LINGXING_ACCESS_TOKEN"))
	cfg.LingXing.PlatformCode = envInt("IPASS_LINGXING_PLATFORM_CODE", 0)
	cfg.LingXing.StoreID = strings.TrimSpace(os.Getenv("IPASS_LINGXING_STORE_ID"))
	cfg.LingXing.SID = envInt("IPASS_LINGXING_SID", 0)

	cfg.Jobs.HeartbeatIntervalSec = envInt("IPASS_HEARTBEAT_INTERVAL_SEC", 10)
	if cfg.Jobs.HeartbeatIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_HEARTBEAT_INTERVAL_SEC 必须为正数")
	}

	cfg.Jobs.PullDSCOOrdersEnable = envBool("IPASS_JOB_PULL_DSCO_ORDERS_ENABLE", false)
	cfg.Jobs.PullDSCOOrdersIntervalSec = envInt("IPASS_JOB_PULL_DSCO_ORDERS_INTERVAL_SEC", 30)
	if cfg.Jobs.PullDSCOOrdersIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_PULL_DSCO_ORDERS_INTERVAL_SEC 必须为正数")
	}

	cfg.Jobs.PushOrdersToLingXingEnable = envBool("IPASS_JOB_PUSH_ORDERS_TO_LINGXING_ENABLE", false)
	cfg.Jobs.PushOrdersToLingXingIntervalSec = envInt("IPASS_JOB_PUSH_ORDERS_TO_LINGXING_INTERVAL_SEC", 30)
	if cfg.Jobs.PushOrdersToLingXingIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_PUSH_ORDERS_TO_LINGXING_INTERVAL_SEC 必须为正数")
	}
	cfg.Jobs.PushOrdersToLingXingBatchSize = envInt("IPASS_JOB_PUSH_ORDERS_TO_LINGXING_BATCH_SIZE", 50)
	if cfg.Jobs.PushOrdersToLingXingBatchSize <= 0 {
		return Config{}, errors.New("IPASS_JOB_PUSH_ORDERS_TO_LINGXING_BATCH_SIZE 必须为正数")
	}

	cfg.Jobs.AckToDSCOEnable = envBool("IPASS_JOB_ACK_TO_DSCO_ENABLE", false)
	cfg.Jobs.AckToDSCOIntervalSec = envInt("IPASS_JOB_ACK_TO_DSCO_INTERVAL_SEC", 30)
	if cfg.Jobs.AckToDSCOIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_ACK_TO_DSCO_INTERVAL_SEC 必须为正数")
	}

	cfg.Jobs.ShipToDSCOEnable = envBool("IPASS_JOB_SHIP_TO_DSCO_ENABLE", false)
	cfg.Jobs.ShipToDSCOIntervalSec = envInt("IPASS_JOB_SHIP_TO_DSCO_INTERVAL_SEC", 30)
	if cfg.Jobs.ShipToDSCOIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_SHIP_TO_DSCO_INTERVAL_SEC 必须为正数")
	}

	cfg.Jobs.InvoiceToDSCOEnable = envBool("IPASS_JOB_INVOICE_TO_DSCO_ENABLE", false)
	cfg.Jobs.InvoiceToDSCOIntervalSec = envInt("IPASS_JOB_INVOICE_TO_DSCO_INTERVAL_SEC", 60)
	if cfg.Jobs.InvoiceToDSCOIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_INVOICE_TO_DSCO_INTERVAL_SEC 必须为正数")
	}
	cfg.Jobs.InvoiceToDSCOBatchSize = envInt("IPASS_JOB_INVOICE_TO_DSCO_BATCH_SIZE", 50)
	if cfg.Jobs.InvoiceToDSCOBatchSize <= 0 {
		return Config{}, errors.New("IPASS_JOB_INVOICE_TO_DSCO_BATCH_SIZE 必须为正数")
	}

	cfg.Jobs.SyncStockEnable = envBool("IPASS_JOB_SYNC_STOCK_ENABLE", false)
	cfg.Jobs.SyncStockIntervalSec = envInt("IPASS_JOB_SYNC_STOCK_INTERVAL_SEC", 300)
	if cfg.Jobs.SyncStockIntervalSec <= 0 {
		return Config{}, errors.New("IPASS_JOB_SYNC_STOCK_INTERVAL_SEC 必须为正数")
	}
	cfg.Jobs.SyncStockBatchSize = envInt("IPASS_JOB_SYNC_STOCK_BATCH_SIZE", 200)
	if cfg.Jobs.SyncStockBatchSize <= 0 {
		return Config{}, errors.New("IPASS_JOB_SYNC_STOCK_BATCH_SIZE 必须为正数")
	}

	// 按需校验：避免只跑 heartbeat 时也必须配置全部外部系统参数。
	if cfg.Jobs.PullDSCOOrdersEnable || cfg.Jobs.PushOrdersToLingXingEnable || cfg.Jobs.AckToDSCOEnable || cfg.Jobs.ShipToDSCOEnable || cfg.Jobs.InvoiceToDSCOEnable || cfg.Jobs.SyncStockEnable {
		if cfg.DSCO.Token == "" {
			return Config{}, errors.New("开启 DSCO 相关任务时，必须配置 IPASS_DSCO_TOKEN")
		}
	}
	if cfg.Jobs.PushOrdersToLingXingEnable || cfg.Jobs.AckToDSCOEnable || cfg.Jobs.ShipToDSCOEnable || cfg.Jobs.SyncStockEnable {
		if cfg.LingXing.AppID == "" {
			return Config{}, errors.New("开启领星相关任务时，必须配置 IPASS_LINGXING_APP_ID")
		}
		if cfg.LingXing.AccessToken == "" {
			return Config{}, errors.New("开启领星相关任务时，必须配置 IPASS_LINGXING_ACCESS_TOKEN")
		}
		if cfg.LingXing.PlatformCode <= 0 {
			return Config{}, errors.New("开启领星相关任务时，必须配置 IPASS_LINGXING_PLATFORM_CODE（正整数）")
		}
		if cfg.LingXing.StoreID == "" {
			return Config{}, errors.New("开启领星相关任务时，必须配置 IPASS_LINGXING_STORE_ID")
		}
	}
	if cfg.Jobs.ShipToDSCOEnable && cfg.LingXing.SID <= 0 {
		return Config{}, errors.New("开启发货回传任务时，必须配置 IPASS_LINGXING_SID（正整数）")
	}
	if cfg.Jobs.SyncStockEnable {
		var err error
		cfg.Stock.LingXingWIDToDSCOWarehouseCode, err = envJSONMapStringString("IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON")
		if err != nil {
			return Config{}, err
		}
		if len(cfg.Stock.LingXingWIDToDSCOWarehouseCode) == 0 {
			return Config{}, errors.New("开启库存同步任务时，必须配置 IPASS_STOCK_WID_TO_DSCO_WAREHOUSE_CODE_JSON")
		}

		cfg.Stock.LingXingSKUToDSCOSKU, err = envJSONMapStringStringOptional("IPASS_STOCK_SKU_TO_DSCO_SKU_JSON")
		if err != nil {
			return Config{}, err
		}
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
