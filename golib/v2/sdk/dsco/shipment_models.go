package dsco

// ShipmentsForUpdate 表示 Create Shipment / Create Shipment Small Batch 的请求体。
//
// 对应 OpenAPI：ShipmentsForUpdate（锚点 a59）。
type ShipmentsForUpdate struct {
	// DscoOrderID / PoNumber / SupplierOrderNumber 三选一，用于定位要关联的订单。
	DscoOrderID string `json:"dscoOrderId,omitempty"`
	// PoNumber 采购单号（PO）。
	PoNumber string `json:"poNumber,omitempty"`
	// SupplierOrderNumber 供应商侧订单号。
	SupplierOrderNumber string `json:"supplierOrderNumber,omitempty"`

	// Shipments 发货信息列表。
	Shipments []ShipmentForUpdate `json:"shipments"`
}

// ShipmentForUpdate 表示单次发货信息（OpenAPI：ShipmentForUpdate，锚点 a58）。
type ShipmentForUpdate struct {
	// TrackingNumber 跟踪号/运单号。
	TrackingNumber string `json:"trackingNumber"`
	// ShipDate 发货时间（RFC3339，可选）。
	ShipDate string `json:"shipDate,omitempty"`

	// ShipCarrier/ShipMethod 与 ShippingServiceLevelCode 二选一（具体依账号策略，部分账号要求必须提供 carrier+method）。
	ShipCarrier              string `json:"shipCarrier,omitempty"`
	ShipMethod               string `json:"shipMethod,omitempty"`
	ShippingServiceLevelCode string `json:"shippingServiceLevelCode,omitempty"`

	// WarehouseCode 为供应商侧仓库 code（部分账号策略要求必须提供）。
	WarehouseCode string `json:"warehouseCode,omitempty"`

	// LineItems 发货行项目列表。
	LineItems []ShipmentLineItemForUpdate `json:"lineItems"`
}

// ShipmentLineItemForUpdate 表示发货行项目（OpenAPI：ShipmentLineItemForUpdate）。
//
// DSCO 文档：dscoItemId/sku/partnerSku/upc/ean 至少提供一个来唯一标识商品。
type ShipmentLineItemForUpdate struct {
	// Quantity 发货数量。
	Quantity int `json:"quantity"`

	LineNumber *int `json:"lineNumber,omitempty"`

	// DscoItemID DSCO 侧商品 ID（可选）。
	DscoItemID string `json:"dscoItemId,omitempty"`
	// SKU 商品 SKU（可选）。
	SKU string `json:"sku,omitempty"`
	// PartnerSKU 合作方 SKU（可选）。
	PartnerSKU string `json:"partnerSku,omitempty"`
	// UPC/EAN 可选。
	UPC string `json:"upc,omitempty"`
	EAN string `json:"ean,omitempty"`
	// GTIN/ISBN/MPN 为可选补充标识（按文档可出现）。
	GTIN string `json:"gtin,omitempty"`
	ISBN string `json:"isbn,omitempty"`
	MPN  string `json:"mpn,omitempty"`

	// PackageSpanFlag 标记该商品是否跨多个包裹（可为空）。
	// 对应 OpenAPI：packageSpanFlag，用于表示该行商品会出现在多个 trackingNumber/shipments 中。
	PackageSpanFlag *bool `json:"packageSpanFlag,omitempty"`
}
