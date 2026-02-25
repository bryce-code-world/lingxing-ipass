package dsco

import "encoding/json"

// StreamObjectType 表示 Stream.objectType。
type StreamObjectType string

const (
	StreamObjectTypeOrder     StreamObjectType = "order"
	StreamObjectTypeInvoice   StreamObjectType = "invoice"
	StreamObjectTypeInventory StreamObjectType = "inventory"
	StreamObjectTypeCatalog   StreamObjectType = "catalog"
)

// Stream 表示 DSCO Stream（用于持续/全量导出订单、库存等数据）。
//
// 说明：
// - DSCO 文档建议通过 Stream APIs 进行日常数据拉取（包括 inventory）。
// - 本结构仅覆盖 SDK 当前需要用到的字段；更多字段可按需补齐。
type Stream struct {
	ID          string         `json:"id,omitempty"`
	Description string         `json:"description,omitempty"`
	ObjectType  StreamObjectType `json:"objectType"`

	// Query 为 stream 的筛选条件；对于 inventory 全量拉取，通常为 InventoryQuery。
	Query any `json:"query"`

	// NumPartitions 可选；不传则由服务端默认值决定（通常为 1）。
	NumPartitions int `json:"numPartitions,omitempty"`

	// MaxEvents 可选；用于限制单次拉取的数量/字节数/持续时间。
	MaxEvents *StreamMaxEvents `json:"maxEvents,omitempty"`

	// Partitions 为服务端只读返回。
	Partitions []StreamPartition `json:"partitions,omitempty"`
}

type StreamMaxEvents struct {
	ObjectCount  *int `json:"objectCount,omitempty"`
	NumBytes     *int `json:"numBytes,omitempty"`
	DurationSec  *int `json:"durationSec,omitempty"`
}

type StreamPartition struct {
	PartitionID int    `json:"partitionId"`
	Status      string `json:"status,omitempty"`
	OwnerID     string `json:"ownerId,omitempty"`
	Position    string `json:"position,omitempty"`
	MaxPosition string `json:"maxPosition,omitempty"`
	LastUpdate  string `json:"lastUpdate,omitempty"`
}

// InventoryQuery 表示 Stream.query 的 inventory 查询（queryType=inventory）。
type InventoryQuery struct {
	QueryType string `json:"queryType"`
}

// SyncStreamOperation 表示 Create Stream Operation 的 sync 操作。
// 该操作会把对应类型（inventory/catalog）的所有对象 dump 到 stream 中，用于 true-up。
type SyncStreamOperation struct {
	OperationType string `json:"operationType"`

	DSCOTradingPartnerID       *string `json:"dscoTradingPartnerId,omitempty"`
	DSCOTradingPartnerParentID *string `json:"dscoTradingPartnerParentId,omitempty"`

	// SKU 用于只同步指定 SKU（供应商侧 SKU）。
	SKU []string `json:"sku,omitempty"`

	IncludeOnHoldSuppliers *bool `json:"includeOnHoldSuppliers,omitempty"`

	DoNotFilterOutTheseTradingPartnerIDs []string `json:"doNotFilterOutTheseTradingPartnerIds,omitempty"`
}

type StreamOperationResponse struct {
	OperationType string `json:"operationType"`
	OperationUUID string `json:"operationUuid"`
}

type StreamEventWrapper struct {
	PartitionID int          `json:"partitionId"`
	OwnerID     string       `json:"ownerId,omitempty"`
	Events      []StreamEvent `json:"events"`
}

type StreamEvent struct {
	ID      string          `json:"id"`
	Payload json.RawMessage `json:"payload"`
	Source  string          `json:"source,omitempty"`
}

