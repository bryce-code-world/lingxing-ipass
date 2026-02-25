package dsco

// CancelReasonCode 表示 DSCO 的取消原因码（Cancel Code Mapping）。
//
// 注意：这些值是否可用取决于具体零售商/账号配置（Portal 的 Cancel Code Mapping 页面）。
type CancelReasonCode string

const (
	// CancelReasonBadAddress 地址错误。
	CancelReasonBadAddress CancelReasonCode = "CXSB"
	// CancelReasonBadSKU SKU 错误。
	CancelReasonBadSKU CancelReasonCode = "CXSBS"
	// CancelReasonCancelledAtRetailerRequest 零售商请求取消。
	CancelReasonCancelledAtRetailerRequest CancelReasonCode = "CXSC"
	// CancelReasonCannotShipUSPS 无法使用 USPS 承运。
	CancelReasonCannotShipUSPS CancelReasonCode = "CXSCNSUSPS"
	// CancelReasonCannotShipToCountry 无法配送到该国家/地区。
	CancelReasonCannotShipToCountry CancelReasonCode = "CXSCNSC"
	// CancelReasonCannotShipToPOBox 无法配送到 PO Box。
	CancelReasonCannotShipToPOBox CancelReasonCode = "CXSCNSPO"
	// CancelReasonCarrierDoesNotServiceDeliveryLocation 承运商不覆盖该投递地址/区域。
	CancelReasonCarrierDoesNotServiceDeliveryLocation CancelReasonCode = "CXSCDNSA"
	// CancelReasonCustomerRefusedDelivery 客户拒收。
	CancelReasonCustomerRefusedDelivery CancelReasonCode = "CXSCR"
	// CancelReasonDamagedInventory 库存损坏。
	CancelReasonDamagedInventory CancelReasonCode = "CXSDAI"
	// CancelReasonDefectiveInventory 库存有缺陷。
	CancelReasonDefectiveInventory CancelReasonCode = "CXSDEI"
	// CancelReasonDuplicateOrder 重复订单。
	CancelReasonDuplicateOrder CancelReasonCode = "CXSDO"
	// CancelReasonInvalidItemCost 商品成本无效。
	CancelReasonInvalidItemCost CancelReasonCode = "CXSIIC"
	// CancelReasonInvalidShipInstructions 发货/配送指示无效。
	CancelReasonInvalidShipInstructions CancelReasonCode = "CXSI"
	// CancelReasonInvalidUOM 计量单位（UOM）无效。
	CancelReasonInvalidUOM CancelReasonCode = "CXUOM"
	// CancelReasonInvalidMethodOfShipment 运输方式无效。
	CancelReasonInvalidMethodOfShipment CancelReasonCode = "CXSISM"
	// CancelReasonIssueWithSomeItemsInOrder 订单部分商品存在问题。
	CancelReasonIssueWithSomeItemsInOrder CancelReasonCode = "CXCI"
	// CancelReasonItemRecall 商品召回。
	CancelReasonItemRecall CancelReasonCode = "CXSIR"
	// CancelReasonMinimumOrderNotMet 未满足最低订单要求。
	CancelReasonMinimumOrderNotMet CancelReasonCode = "CXSMONM"
	// CancelReasonOrderEntryError 订单录入错误。
	CancelReasonOrderEntryError CancelReasonCode = "CXSOEE"
	// CancelReasonOrderInfoMissing 订单信息缺失。
	CancelReasonOrderInfoMissing CancelReasonCode = "CXSIM"
	// CancelReasonPreOrderCancellation 预售订单取消。
	CancelReasonPreOrderCancellation CancelReasonCode = "CXSPOC"
	// CancelReasonSupplierDetectedFraud 供应商检测到欺诈。
	CancelReasonSupplierDetectedFraud CancelReasonCode = "CXSDF"
	// CancelReasonToCloseOrderAndAllowReissue 关闭订单并允许重新下单/重发。
	CancelReasonToCloseOrderAndAllowReissue CancelReasonCode = "CXSCAR"
	// CancelReasonUnableToContactRecipient 无法联系收件人。
	CancelReasonUnableToContactRecipient CancelReasonCode = "CXSUC"
	// CancelReasonCantShipOnTime 无法按时发货。
	CancelReasonCantShipOnTime CancelReasonCode = "CXST"
	// CancelReasonCannotShipAsOrdered 无法按订单要求发货。
	CancelReasonCannotShipAsOrdered CancelReasonCode = "CXSCNSO"
	// CancelReasonCollateralImpact 连带影响（Collateral Impact）。
	CancelReasonCollateralImpact CancelReasonCode = "CXIOI"
	// CancelReasonDiscontinuedItem 商品停产/下架。
	CancelReasonDiscontinuedItem CancelReasonCode = "CXSD"
	// CancelReasonNotEnoughStock 库存不足。
	CancelReasonNotEnoughStock CancelReasonCode = "CXSN"
	// CancelReasonOther 其他。
	CancelReasonOther CancelReasonCode = "CXSO"
	// CancelReasonOutOfStock 缺货。
	CancelReasonOutOfStock CancelReasonCode = "CXSS"
)

// SyncUpdateResponse 对应 OpenAPI 中的 SyncUpdateResponse（锚点 a54）。
type SyncUpdateResponse struct {
	Status    string               `json:"status"`
	RequestID string               `json:"requestId"`
	EventDate string               `json:"eventDate"`
	Messages  []APIResponseMessage `json:"messages,omitempty"`
}

// OrderForCancel 对应 Cancel Order Item 接口的请求体（POST /order/item/cancel）。
type OrderForCancel struct {
	// ID 订单唯一标识（必填）。
	ID string `json:"id"`
	// Type 指定 ID 的口径类型（DSCO_ORDER_ID / PO_NUMBER / SUPPLIER_ORDER_NUMBER）。
	Type OrderAcknowledgeIDType `json:"type"`
	// LineItems 要取消的订单行项目（必填）。
	LineItems []OrderLineItemForCancel `json:"lineItems"`
}

// OrderLineItemForCancel 对应 OpenAPI 中的 OrderLineItemForCancel（锚点 a55）。
type OrderLineItemForCancel struct {
	DscoSupplierID             *string `json:"dscoSupplierId,omitempty"`
	DscoTradingPartnerID       *string `json:"dscoTradingPartnerId,omitempty"`
	DscoTradingPartnerParentID *string `json:"dscoTradingPartnerParentId,omitempty"`
	DscoItemID                 *string `json:"dscoItemId,omitempty"`
	SKU                        *string `json:"sku,omitempty"`
	UPC                        *string `json:"upc,omitempty"`
	EAN                        *string `json:"ean,omitempty"`
	PartnerSKU                 *string `json:"partnerSku,omitempty"`

	// CancelledQuantity 取消数量（必填）。
	CancelledQuantity int `json:"cancelledQuantity"`
	// CancelCode 取消原因码（必填；可参考 CancelReasonCode 常量）。
	CancelCode      CancelReasonCode `json:"cancelCode"`
	CancelledReason *string          `json:"cancelledReason,omitempty"`

	// LineNumber 行号（可选，用于唯一定位订单行项目）。
	LineNumber *int `json:"lineNumber,omitempty"`

	TransactionID   *string `json:"transactionId,omitempty"`
	TransactionDate *string `json:"transactionDate,omitempty"`
}
