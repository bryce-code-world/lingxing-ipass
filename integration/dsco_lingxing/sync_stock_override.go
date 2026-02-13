package dsco_lingxing

import "errors"

var ErrSyncStockTooManyKeys = errors.New("sync_stock too many keys")

// SyncStockSource 表示 sync_stock 手动触发时的“数据来源模式”。
type SyncStockSource string

const (
	SyncStockSourceDefault     SyncStockSource = ""
	SyncStockSourceDailyPulled SyncStockSource = "daily_pulled"
	SyncStockSourceManualItems SyncStockSource = "manual_items"
)

type SyncStockManualItem struct {
	DSCOWarehouseID string
	DSCOSKU         string
	Qty             int
}

type SyncStockDailyPulledOptions struct {
	// StartTime/EndTime 为 UTC 秒级区间：[StartTime, EndTime)。
	StartTime int64
	EndTime   int64

	DiffOnly bool

	DSCOWarehouseID     string
	LingXingWarehouseID string

	DSCOSKUList     []string
	LingXingSKUList []string
}

type SyncStockOverride struct {
	// ForceDoSync 用于覆盖 runtime_config.jobs.sync_stock.sync：
	// - nil：按 runtime_config
	// - true：强制回写 DSCO
	// - false：强制不回写（dry-run）
	ForceDoSync *bool

	Source SyncStockSource

	// DailyPulled 基于 dsco_warehouse_sync “当日最新记录”做回写。
	DailyPulled *SyncStockDailyPulledOptions

	// ManualItems 直接使用调用方传入的 DSCO（warehouse_id + sku + qty）做回写。
	ManualItems []SyncStockManualItem

	// MaxKeys 用于保护：当来源是 DailyPulled 时，最多允许处理多少个（dsco_warehouse_id + dsco_sku）维度。
	MaxKeys int
}
