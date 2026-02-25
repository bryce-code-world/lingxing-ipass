package dsco

import "encoding/json"

// ReturnReasonCode 表示 DSCO 的退货原因码（Return Code Mapping）。
//
// 注意：这些值是否可用取决于具体零售商/账号配置（Portal 的 Return Code Mapping 页面）。
type ReturnReasonCode string

const (
	// ReturnReasonArrivedTooLate 到货太晚。
	ReturnReasonArrivedTooLate ReturnReasonCode = "CRATL"
	// ReturnReasonCustomerOrderingError 客户下单错误。
	ReturnReasonCustomerOrderingError ReturnReasonCode = "CRCOE"
	// ReturnReasonExchange 换货。
	ReturnReasonExchange ReturnReasonCode = "CREX"
	// ReturnReasonLost 丢失。
	ReturnReasonLost ReturnReasonCode = "CRALIT"
	// ReturnReasonNotAsExpected 与预期不符。
	ReturnReasonNotAsExpected ReturnReasonCode = "CRNAE"
	// ReturnReasonDefectiveProduct 商品有缺陷。
	ReturnReasonDefectiveProduct ReturnReasonCode = "CRP"
	// ReturnReasonDamaged 商品损坏。
	ReturnReasonDamaged ReturnReasonCode = "CRD"
	// ReturnReasonCustomerRegret 客户后悔/不想要（Customer Regret）。
	ReturnReasonCustomerRegret ReturnReasonCode = "CRR"
	// ReturnReasonDifferentThanWebsite 与网站描述不符。
	ReturnReasonDifferentThanWebsite ReturnReasonCode = "CRDTW"
	// ReturnReasonUncollected 未被取件/未被收取。
	ReturnReasonUncollected ReturnReasonCode = "CRUNC"
	// ReturnReasonOther 其他。
	ReturnReasonOther ReturnReasonCode = "CRO"
	// ReturnReasonRefused 拒收。
	ReturnReasonRefused ReturnReasonCode = "CRRU"
	// ReturnReasonDifferentThanTV 与电视/广告展示不符。
	ReturnReasonDifferentThanTV ReturnReasonCode = "CRDTT"
	// ReturnReasonWrongAddress 地址错误。
	ReturnReasonWrongAddress ReturnReasonCode = "CRSTWA"
	// ReturnReasonShortShipment 少发/短装。
	ReturnReasonShortShipment ReturnReasonCode = "CRAGNF"
	// ReturnReasonCatalogMisprint 目录/画册印刷错误。
	ReturnReasonCatalogMisprint ReturnReasonCode = "CRCM"
	// ReturnReasonDislikeStyleFashion 不喜欢款式/风格。
	ReturnReasonDislikeStyleFashion ReturnReasonCode = "CRDSF"
	// ReturnReasonDissatisfiedWithColor 对颜色不满意。
	ReturnReasonDissatisfiedWithColor ReturnReasonCode = "CRDWC"
	// ReturnReasonDissatisfiedWithFit 对尺码/合身度不满意。
	ReturnReasonDissatisfiedWithFit ReturnReasonCode = "CRDWF"
	// ReturnReasonDissatisfiedWithQuality 对质量不满意。
	ReturnReasonDissatisfiedWithQuality ReturnReasonCode = "CRDWQ"
	// ReturnReasonGiftReturn 礼品退货。
	ReturnReasonGiftReturn ReturnReasonCode = "CRGR"
	// ReturnReasonFailedInspection 质检失败/验收未通过。
	ReturnReasonFailedInspection ReturnReasonCode = "CROFS"
	// ReturnReasonExcessQty 多余数量/超出数量。
	ReturnReasonExcessQty ReturnReasonCode = "CROS"
	// ReturnReasonItemNotSent 未发货/未寄出。
	ReturnReasonItemNotSent ReturnReasonCode = "CRINS"
	// ReturnReasonLostInbound 入库途中丢失（Lost Inbound）。
	ReturnReasonLostInbound ReturnReasonCode = "CRALITI"
	// ReturnReasonMispick 拣货错误。
	ReturnReasonMispick ReturnReasonCode = "CRMP"
	// ReturnReasonNotGoodValue 性价比不高/不值这个价。
	ReturnReasonNotGoodValue ReturnReasonCode = "CRNGV"
	// ReturnReasonUnauthorizedCOD 未授权的到付（COD）。
	ReturnReasonUnauthorizedCOD ReturnReasonCode = "CRRFCOD"
	// ReturnReasonUndeliverable 无法投递。
	ReturnReasonUndeliverable ReturnReasonCode = "CRU"
	// ReturnReasonProductNotAsDescribed 与描述不符（Product not as described）。
	ReturnReasonProductNotAsDescribed ReturnReasonCode = "CRN"
	// ReturnReasonTooLarge 太大。
	ReturnReasonTooLarge ReturnReasonCode = "CRTL"
	// ReturnReasonTooSmall 太小。
	ReturnReasonTooSmall ReturnReasonCode = "CRTS"
	// ReturnReasonUnwantedCancelledOrder 不想要：订单已取消。
	ReturnReasonUnwantedCancelledOrder ReturnReasonCode = "CRCO"
	// ReturnReasonUnwantedDuplicateOrder 不想要：重复订单。
	ReturnReasonUnwantedDuplicateOrder ReturnReasonCode = "CRDO"
	// ReturnReasonUnwantedQuoteOnly 不想要：仅询价。
	ReturnReasonUnwantedQuoteOnly ReturnReasonCode = "CRWQO"
	// ReturnReasonWrongItem 发错商品。
	ReturnReasonWrongItem ReturnReasonCode = "CRWI"
	// ReturnReasonWrongQuantity 数量错误。
	ReturnReasonWrongQuantity ReturnReasonCode = "CRWQ"
	// ReturnReasonWrongSize 尺码错误。
	ReturnReasonWrongSize ReturnReasonCode = "CRWS"
)

// ReturnCreateRequest 表示创建退货单请求体（POST /return/）。
type ReturnCreateRequest struct {
	// ReturnNumber 退货单号（调用方系统中的退货标识，创建时必填）。
	ReturnNumber string `json:"returnNumber"`
	// LineItems 退货行项目列表（创建时必填）。
	LineItems []ReturnLineItemCreate `json:"lineItems"`

	// PoNumber/DscoOrderID 二选一，用于定位关联的订单。
	PoNumber    string `json:"poNumber,omitempty"`
	DscoOrderID string `json:"dscoOrderId,omitempty"`
}

// ReturnLineItemCreate 表示创建退货单的行项目。
type ReturnLineItemCreate struct {
	// Quantity 退货数量（必填）。
	Quantity int `json:"quantity"`
	// ReasonCode 退货原因码（必填，枚举值由 DSCO 定义；可参考 ReturnReasonCode 常量）。
	ReasonCode ReturnReasonCode `json:"reasonCode"`

	// 以下为“商品标识”字段，OpenAPI 要求至少提供一个：
	// itemId / sku / upc / partnerSku / ean / mpn
	ItemID     string `json:"itemId,omitempty"`
	SKU        string `json:"sku,omitempty"`
	UPC        string `json:"upc,omitempty"`
	EAN        string `json:"ean,omitempty"`
	PartnerSKU string `json:"partnerSku,omitempty"`
	MPN        string `json:"mpn,omitempty"`

	// ISBN 可选（OpenAPI nullable）。
	ISBN *string `json:"isbn,omitempty"`

	// LineNumber 订单行号：当订单存在相同商品的多行，或项目要求必须提供 lineNumber 时使用。
	LineNumber *int `json:"lineNumber,omitempty"`

	// Adjustments：附加费用/调整项（key->amount）。
	Adjustments map[string]float64 `json:"adjustments,omitempty"`
}

// ReturnCompleteRequest 表示完成退货单请求体（PUT /return/）。
//
// 注意：完成退货单要求每个行项目提供 response=accepted/rejected。
type ReturnCompleteRequest struct {
	// DscoReturnID DSCO 侧退货单 ID（dscoReturnId）。
	DscoReturnID string `json:"dscoReturnId,omitempty"`
	// PartnerReturnNumber 创建方设置的退货标识（partnerReturnNumber），可用于定位退货单。
	PartnerReturnNumber string `json:"partnerReturnNumber,omitempty"`
	// DscoAccountID 账号 ID（部分场景需要，用于指定零售商）。
	DscoAccountID string `json:"dscoAccountId,omitempty"`
	// DscoTradingPartnerID 交易伙伴 ID（部分场景需要，用于指定零售商）。
	DscoTradingPartnerID string `json:"dscoTradingPartnerId,omitempty"`

	// LineItems 退货行项目列表（必填）。
	LineItems []ReturnLineItemComplete `json:"lineItems"`
}

// ReturnLineItemComplete 表示完成退货单的行项目。
type ReturnLineItemComplete struct {
	// Response 处理结果（必填）：accepted 或 rejected。
	Response string `json:"response"`

	// Quantity/ReasonCode 已在 DSCO 文档中标记为 deprecated；保留用于兼容。
	Quantity   int    `json:"quantity,omitempty"`
	ReasonCode string `json:"reasonCode,omitempty"`
}

// ReturnResponse 表示退货接口的响应体。
type ReturnResponse struct {
	// Status 操作状态（success/failure 等，取决于 DSCO 返回口径）。
	Status string `json:"status"`
	// Return 退货单对象（可选）。
	Return *Return `json:"return,omitempty"`
}

// Return 表示退货单对象（仅定义一期用到的字段）。
type Return struct {
	// ReturnNumber 退货单号。
	ReturnNumber string `json:"returnNumber,omitempty"`
}

// -------- Return Change Log（GET /return/log）-------

// ReturnChangeLogQuery 表示 GET /return/log 的查询参数。
//
// 注意：
// - 若提供 scrollId，则其它参数会被 DSCO 忽略。
// - 若未提供 scrollId，则 requestId 与 startDate/endDate 二选一（DSCO 文档说明）。
type ReturnChangeLogQuery struct {
	ScrollID string `url:"scrollId,omitempty"`

	StartDate string `url:"startDate,omitempty"`
	EndDate   string `url:"endDate,omitempty"`

	RequestID string `url:"requestId,omitempty"`

	// Status 过滤条件：pending / success / failure / success_or_failure
	Status string `url:"status,omitempty"`
}

// ReturnChangeLogResponse 对应 OpenAPI 的 ReturnChangeLogResponse。
type ReturnChangeLogResponse struct {
	ScrollID string `json:"scrollId,omitempty"`

	// Status 仅在传 requestId 时返回：PROCESSING / COMPLETED
	Status string `json:"status,omitempty"`

	Logs []ReturnChangeLog `json:"logs"`
}

// ReturnChangeLog 对应 OpenAPI 的 ReturnChangeLog。
//
// payload 的结构为 Return 对象（字段较多且持续演进）；这里保留原始 JSON 便于排查与审计。
type ReturnChangeLog struct {
	Payload json.RawMessage `json:"payload"`

	DateProcessed string `json:"dateProcessed"`

	// Status：pending / success / failure
	Status string `json:"status"`

	// RequestMethod：api / portal / file
	RequestMethod string `json:"requestMethod"`

	// RequestMethodDetail：string 或 object
	RequestMethodDetail any `json:"requestMethodDetail,omitempty"`

	RequestID string `json:"requestId,omitempty"`
	ProcessID string `json:"processId,omitempty"`

	Results []APIResponseMessage `json:"results,omitempty"`
}
