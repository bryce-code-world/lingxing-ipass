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
	JobPullSKUPair    JobName = "pull_sku_pair"
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

	// Sync 为 true 时，表示执行“写回/同步”动作；为 false 时仅拉取并记录数据，不回写外部系统。
	// 当前用于：sync_stock（领星库存 -> DSCO inventory）是否实际回写 DSCO。
	Sync bool `json:"sync,omitempty"`

	// UseStream 为 true 时，优先使用 DSCO Stream + sync operation 拉取 inventory（全量/批量），
	// 用于替代循环调用 GET /inventory（该接口不适合批量，且有严格限流）。
	// 当前用于：sync_stock 对账读取 DSCO 当前库存数量。
	UseStream bool `json:"use_stream,omitempty"`

	// MultiBan 为 true 时，禁止处理“多行/多数量”的订单（多 SKU / 单 SKU 多数量）。
	// 说明：
	// - MultiBan=false：不禁止（默认），允许执行该类订单。
	// - MultiBan=true：禁止，任务会跳过该类订单（不推进状态）。
	MultiBan bool `json:"multi_ban,omitempty"`
}

// Mapping is DSCO -> LingXing (key=DSCO, value=LingXing).
type Mapping struct {
	Shop      map[string]string `json:"shop"`
	Warehouse map[string]string `json:"warehouse"`
	SKU       map[string]string `json:"sku"`
	Shipment  map[string]string `json:"shipment"`
}
