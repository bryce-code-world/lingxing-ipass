package runtimecfg

import "time"

const DomainDSCOLingXing = "dsco_lingxing"

type Domain string

type JobName string

const (
	JobPullDSCOOrders JobName = "pull_dsco_orders"
	JobPushToLingXing JobName = "push_to_lingxing"
	JobAckToDSCO      JobName = "ack_to_dsco"
	JobShipToDSCO     JobName = "ship_to_dsco"
	JobInvoiceToDSCO  JobName = "invoice_to_dsco"
	JobSyncStock      JobName = "sync_stock"
	JobCleanupExports JobName = "cleanup_exports"
)

type RuntimeConfig struct {
	Domain    string    `json:"domain"`
	Config    Config    `json:"config"`
	UpdatedAt int64     `json:"updated_at"`
	LoadedAt  time.Time `json:"-"`
}

type Config struct {
	Domain  string                `json:"domain"`
	Jobs    map[JobName]JobConfig `json:"jobs"`
	Mapping Mapping               `json:"mapping"`
}

type JobConfig struct {
	Enable bool   `json:"enable"`
	Cron   string `json:"cron"`
	Size   int    `json:"size"`

	// MultiBan 为 true 时，禁止处理“多行/多数量”的订单（多 SKU / 单 SKU 多数量）。
	// 说明：
	// - MultiBan=false：不禁止（默认），允许执行该类订单。
	// - MultiBan=true：禁止，任务会跳过该类订单（不推进状态）。
	MultiBan bool `json:"multi_ban"`
}

// Mapping is DSCO -> LingXing (key=DSCO, value=LingXing).
type Mapping struct {
	Shop      map[string]string `json:"shop"`
	Warehouse map[string]string `json:"warehouse"`
	SKU       map[string]string `json:"sku"`
	Shipment  map[string]string `json:"shipment"`
}
