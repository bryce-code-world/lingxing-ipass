package lingxing

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// StringOrNumber 用于兼容“字符串/数字/null”混用的字段（常见于各类 id、时间戳等）。
type StringOrNumber string

func (v StringOrNumber) String() string { return string(v) }

func (v *StringOrNumber) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) {
		*v = ""
		return nil
	}
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*v = StringOrNumber(s)
		return nil
	}
	*v = StringOrNumber(string(b))
	return nil
}

// Float64OrString 用于兼容“数字/字符串/null”混用的小数类字段（常见于重量、尺寸等）。
type Float64OrString float64

func (v *Float64OrString) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.Equal(b, []byte("null")) {
		*v = 0
		return nil
	}
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		s = strings.TrimSpace(s)
		if s == "" {
			*v = 0
			return nil
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		*v = Float64OrString(f)
		return nil
	}
	var f float64
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}
	*v = Float64OrString(f)
	return nil
}

// OrderDetailV2Request 表示“获取单个订单详情”的请求参数。
//
// 说明：
// - 当前 SDK 通过 `/pb/mp/order/v2/list` 查询并取第一条作为“订单详情”返回。
// - 若后续领星提供独立的“订单详情接口”，可再单独对接并替换实现。
type OrderDetailV2Request struct {
	// PlatformOrderNo 平台订单号（优先使用）。
	PlatformOrderNo string
	// PlatformOrderName 平台订单名称/编号（部分平台需要使用该字段查询）。
	PlatformOrderName string
}

// OrderDetailV2 表示订单详情对象（字段口径以领星文档与实际返回为准）。
//
// 对接接口：
// - 订单管理订单列表：/pb/mp/order/v2/list（取 data.list[0] 作为详情）
type OrderDetailV2 struct {
	// GlobalOrderNo 系统单号（全局订单号）。
	GlobalOrderNo StringOrNumber `json:"global_order_no"`
	// ReferenceNo 参考号。
	ReferenceNo string `json:"reference_no"`
	// StoreID 店铺 ID。
	StoreID StringOrNumber `json:"store_id"`

	// OrderFromName 订单来源名称（例如：手工订单）。
	OrderFromName string `json:"order_from_name"`

	// DeliveryType 发货类型（口径以文档为准）。
	DeliveryType int `json:"delivery_type"`
	// SplitType 拆分类型（口径以文档为准，常见为字符串数字）。
	SplitType string `json:"split_type"`
	// Status 订单状态（口径以文档为准）。
	//
	// 枚举值参考：`/pb/mp/order/v2/list` 返回字段 `data>>list>>status`（系统订单状态）。
	Status MultiPlatformOrderStatus `json:"status"`

	// GlobalPurchaseTime 购买时间（时间戳，秒）。
	GlobalPurchaseTime StringOrNumber `json:"global_purchase_time"`
	// GlobalPaymentTime 支付时间（时间戳，秒）。
	GlobalPaymentTime StringOrNumber `json:"global_payment_time"`
	// GlobalReviewTime 审核时间（时间戳，秒）。
	GlobalReviewTime StringOrNumber `json:"global_review_time"`
	// GlobalDistributionTime 配货/分配时间（时间戳，秒）。
	GlobalDistributionTime StringOrNumber `json:"global_distribution_time"`
	// GlobalPrintTime 打印时间（可能为 null）。
	GlobalPrintTime *StringOrNumber `json:"global_print_time"`
	// GlobalMarkTime 标记时间（时间戳，秒；可能为 0）。
	GlobalMarkTime StringOrNumber `json:"global_mark_time"`
	// GlobalDeliveryTime 发货/出库时间（时间戳，秒）。
	GlobalDeliveryTime StringOrNumber `json:"global_delivery_time"`

	// AmountCurrency 币种（例如：USD）。
	AmountCurrency string `json:"amount_currency"`
	// Remark 备注。
	Remark string `json:"remark"`

	// GlobalLatestShipTime 最晚发货时间（时间戳，秒；可能为 "0"）。
	GlobalLatestShipTime StringOrNumber `json:"global_latest_ship_time"`
	// GlobalCancelTime 取消时间（时间戳，秒；可能为 0）。
	GlobalCancelTime StringOrNumber `json:"global_cancel_time"`

	// UpdateTime 更新时间（时间戳，秒；文档中常为字符串时间戳）。
	UpdateTime StringOrNumber `json:"update_time"`

	// OrderTag 订单标签（结构以实际返回为准）。
	OrderTag json.RawMessage `json:"order_tag"`
	// PendingOrderTag 待处理标签（结构以实际返回为准）。
	PendingOrderTag json.RawMessage `json:"pending_order_tag"`
	// ExceptionOrderTag 异常标签（结构以实际返回为准）。
	ExceptionOrderTag json.RawMessage `json:"exception_order_tag"`

	// WID 仓库 ID。
	WID StringOrNumber `json:"wid"`
	// WarehouseName 仓库名称。
	WarehouseName string `json:"warehouse_name"`

	// BuyersInfo 买家信息。
	BuyersInfo OrderDetailV2BuyersInfo `json:"buyers_info"`
	// AddressInfo 地址信息。
	AddressInfo OrderDetailV2AddressInfo `json:"address_info"`

	// ItemInfo 商品信息（订单行）。
	ItemInfo []OrderDetailV2Item `json:"item_info"`
	// PlatformInfo 平台单信息。
	PlatformInfo []OrderDetailV2PlatformInfo `json:"platform_info"`
	// PaymentInfo 支付信息。
	PaymentInfo []OrderDetailV2PaymentInfo `json:"payment_info"`
	// LogisticsInfo 物流信息。
	LogisticsInfo OrderDetailV2LogisticsInfo `json:"logistics_info"`
	// TransactionInfo 交易汇总信息（金额展示字段通常自带币种符号/格式）。
	TransactionInfo []OrderDetailV2Transaction `json:"transaction_info"`

	// OriginalGlobalOrderNo 补发订单原系统单号（结构以实际返回为准，可能为 null）。
	OriginalGlobalOrderNo json.RawMessage `json:"original_global_order_no"`
	// CustomerShippingList 客付物流/客选物流列表（部分场景为空字符串数组）。
	CustomerShippingList []string `json:"customer_shipping_list"`
	// GlobalCreateTime 创建时间（字符串格式，例如：2006-01-02 15:04:05）。
	GlobalCreateTime string `json:"global_create_time"`
	// FlowNode 流程节点（口径以文档为准）。
	FlowNode int `json:"flow_node"`
	// SupplierID 供应商 ID（结构以实际返回为准，可能为 null）。
	SupplierID json.RawMessage `json:"supplier_id"`
	// IsDelete 是否删除标记（口径以文档为准，常见 0/1）。
	IsDelete int `json:"is_delete"`
	// OrderCustomFields 订单自定义字段（结构以实际返回为准，可能为 null）。
	OrderCustomFields json.RawMessage `json:"order_custom_fields"`
}

// OrderDetailV2BuyersInfo 表示买家信息。
type OrderDetailV2BuyersInfo struct {
	// BuyerNo 买家编号。
	BuyerNo string `json:"buyer_no"`
	// BuyerEmail 买家邮箱。
	BuyerEmail string `json:"buyer_email"`
	// BuyerName 买家名称。
	BuyerName string `json:"buyer_name"`
	// BuyerNote 买家备注。
	BuyerNote string `json:"buyer_note"`
}

// OrderDetailV2AddressInfo 表示收件地址信息。
type OrderDetailV2AddressInfo struct {
	// ReceiverName 收件人姓名。
	ReceiverName string `json:"receiver_name"`
	// ReceiverMobile 收件人手机。
	ReceiverMobile string `json:"receiver_mobile"`
	// ReceiverTel 收件人电话。
	ReceiverTel string `json:"receiver_tel"`
	// ReceiverCountryCode 收件国家/地区二字码。
	ReceiverCountryCode string `json:"receiver_country_code"`
	// City 城市。
	City string `json:"city"`
	// StateOrRegion 州/省/地区。
	StateOrRegion string `json:"state_or_region"`
	// AddressLine1 地址 1。
	AddressLine1 string `json:"address_line1"`
	// AddressLine2 地址 2。
	AddressLine2 string `json:"address_line2"`
	// AddressLine3 地址 3。
	AddressLine3 string `json:"address_line3"`
	// District 区/县。
	District string `json:"district"`
	// PostalCode 邮编。
	PostalCode string `json:"postal_code"`
	// DoorplateNo 门牌号。
	DoorplateNo string `json:"doorplate_no"`
	// CompanyName 公司名（可能为 null）。
	CompanyName *string `json:"company_name"`
}

// OrderDetailV2Item 表示订单商品行信息。
type OrderDetailV2Item struct {
	// ID 系统商品唯一 id。
	ID StringOrNumber `json:"id"`
	// PlatformOrderNo 平台订单号。
	PlatformOrderNo string `json:"platform_order_no"`
	// OrderItemNo 订单明细单号（平台侧）。
	OrderItemNo string `json:"order_item_no"`

	// ItemFromName 商品来源名称。
	ItemFromName string `json:"item_from_name"`
	// MSKU 店铺 MSKU。
	MSKU string `json:"msku"`
	// LocalSKU 本地 SKU。
	LocalSKU string `json:"local_sku"`
	// ProductNo 平台商品 id（部分平台为 ASIN）。
	ProductNo string `json:"product_no"`
	// LocalProductName 本地品名。
	LocalProductName string `json:"local_product_name"`

	// IsBundled 是否捆绑商品（0/1）。
	IsBundled int `json:"is_bundled"`
	// SubProducts 子产品列表（捆绑商品时使用）。
	SubProducts []OrderDetailV2SubItem `json:"sub_products"`

	// Title 商品标题。
	Title string `json:"title"`
	// VariantAttr 变体属性。
	VariantAttr string `json:"variant_attr"`

	// UnitPriceAmount 单价（字符串金额，币种取 amount_currency）。
	UnitPriceAmount string `json:"unit_price_amount"`
	// ItemPriceAmount 商品金额（字符串金额，币种取 amount_currency）。
	ItemPriceAmount string `json:"item_price_amount"`
	// Quantity 数量。
	Quantity int `json:"quantity"`
	// Remark 商品备注。
	Remark string `json:"remark"`
	// PlatformStatus 平台订单商品状态。
	PlatformStatus string `json:"platform_status"`

	// Type 商品类型（文档口径不稳定，保留原始结构）。
	Type json.RawMessage `json:"type"`
	// StockCost 商品出库成本（文档中常为 stockCost，口径不稳定，保留原始结构）。
	StockCost json.RawMessage `json:"stockCost"`

	// WmsOutboundCostAmount 实际出库成本（通常币种 CNY，来源销售出库单）。
	WmsOutboundCostAmount string `json:"wms_outbound_cost_amount"`
	// StockCostAmount 库存明细成本（通常币种 CNY）。
	StockCostAmount string `json:"stock_cost_amount"`
	// StockDeductID 备货店铺 ID。
	StockDeductID string `json:"stock_deduct_id"`
	// StockDeductName 备货店铺名称（可能为 null）。
	StockDeductName json.RawMessage `json:"stock_deduct_name"`
	// CGPriceAmount 采购成本（通常币种 CNY）。
	CGPriceAmount string `json:"cg_price_amount"`
	// ShippingAmount 预估运费（通常币种 CNY）。
	ShippingAmount string `json:"shipping_amount"`
	// WmsShippingPriceAmount 实际运费（币种取 logistics_info.cost_currency_code 或文档口径）。
	WmsShippingPriceAmount string `json:"wms_shipping_price_amount"`

	// CustomerShippingAmount 客付运费（币种取 amount_currency）。
	CustomerShippingAmount string `json:"customer_shipping_amount"`
	// DiscountAmount 折扣（币种取 amount_currency）。
	DiscountAmount string `json:"discount_amount"`
	// CustomerTipAmount 小费（Shopify 等，币种取 amount_currency）。
	CustomerTipAmount string `json:"customer_tip_amount"`
	// TaxAmount 税费（币种取 amount_currency）。
	TaxAmount string `json:"tax_amount"`
	// SalesRevenueAmount 销售收益（币种取 amount_currency）。
	SalesRevenueAmount string `json:"sales_revenue_amount"`
	// TransactionFeeAmount 交易费（币种取 amount_currency）。
	TransactionFeeAmount string `json:"transaction_fee_amount"`
	// OtherAmount 平台其他费用（币种取 amount_currency）。
	OtherAmount string `json:"other_amount"`

	// CustomizedURL 定制商品文件下载链接（结构以实际返回为准，可能为 null）。
	CustomizedURL json.RawMessage `json:"customized_url"`
	// PlatformSubsidyAmount 平台补贴（部分平台返回）。
	PlatformSubsidyAmount string `json:"platform_subsidy_amount"`
	// CodAmount COD 费用（币种取 amount_currency）。
	CodAmount string `json:"cod_amount"`
	// GiftWrapAmount 礼品包装费（币种取 amount_currency）。
	GiftWrapAmount string `json:"gift_wrap_amount"`
	// PlatformTaxAmount 平台销售税（币种取 amount_currency）。
	PlatformTaxAmount string `json:"platform_tax_amount"`
	// PointsGrantedAmount 积分成本（币种取 amount_currency）。
	PointsGrantedAmount string `json:"points_granted_amount"`
	// OtherFee 其他费用（用户手动导入，正负皆可能）。
	OtherFee string `json:"other_fee"`

	// DeliveryTime 平台发货时间（部分平台/条件下返回）。
	DeliveryTime json.RawMessage `json:"delivery_time"`
	// SourceName 来源名称（结构以实际返回为准，可能为 null）。
	SourceName json.RawMessage `json:"source_name"`
	// DataJSON 扩展字段（JSON 字符串）。
	DataJSON string `json:"data_json"`

	// ItemCustomFields 商品自定义字段（结构以实际返回为准）。
	ItemCustomFields json.RawMessage `json:"item_custom_fields"`
	// GlobalItemNo 系统商品行全局编号。
	GlobalItemNo StringOrNumber `json:"global_item_no"`
	// IsDelete 是否删除标记（0/1）。
	IsDelete int `json:"is_delete"`
}

// OrderDetailV2SubItem 表示捆绑商品的子产品信息。
type OrderDetailV2SubItem struct {
	// SKU 子产品 SKU。
	SKU string `json:"sku"`
	// Qty 子产品发货数量。
	Qty int `json:"qty"`
}

// OrderDetailV2PlatformInfo 表示平台单信息。
type OrderDetailV2PlatformInfo struct {
	// OrderFrom 平台来源/渠道名称。
	OrderFrom string `json:"order_from"`
	// PlatformOrderNo 平台订单号。
	PlatformOrderNo string `json:"platform_order_no"`
	// PlatformOrderName 平台订单名称/编号。
	PlatformOrderName string `json:"platform_order_name"`
	// PlatformCode 平台 code（如 10001 Amazon / 10009 自定义等）。
	PlatformCode string `json:"platform_code"`
	// StoreCountryCode 店铺国家/地区代码（示例字段名：store_Country_code）。
	StoreCountryCode string `json:"store_Country_code"`

	// Status 平台订单状态（可能为空字符串）。
	Status string `json:"status"`
	// PaymentStatus 平台支付状态（可能为空字符串）。
	PaymentStatus string `json:"payment_status"`
	// ShippingStatus 平台物流/发货状态（可能为空字符串）。
	ShippingStatus string `json:"shipping_status"`

	// PurchaseTime 平台下单时间（时间戳，秒）。
	PurchaseTime StringOrNumber `json:"purchase_time"`
	// PaymentTime 平台支付时间（时间戳，秒）。
	PaymentTime StringOrNumber `json:"payment_time"`
	// LatestShipTime 平台最晚发货时间（时间戳，秒）。
	LatestShipTime StringOrNumber `json:"latest_ship_time"`
	// CancelTime 平台取消时间（时间戳，秒）。
	CancelTime StringOrNumber `json:"cancel_time"`
	// DeliveryTime 平台发货时间（时间戳，秒）。
	DeliveryTime StringOrNumber `json:"delivery_time"`
}

// OrderDetailV2PaymentInfo 表示支付信息。
type OrderDetailV2PaymentInfo struct {
	// PlatformOrderNo 平台订单号。
	PlatformOrderNo string `json:"platform_order_no"`
	// PaymentMethod 支付方式。
	PaymentMethod string `json:"payment_method"`
	// TransactionNo 交易号。
	TransactionNo string `json:"transaction_no"`
	// Currency 币种。
	Currency string `json:"currency"`
	// PaymentAmount 支付金额（字符串金额）。
	PaymentAmount string `json:"payment_amount"`
	// PaymentTime 支付时间（时间戳，秒）。
	PaymentTime StringOrNumber `json:"payment_time"`
}

// OrderDetailV2LogisticsInfo 表示物流信息。
type OrderDetailV2LogisticsInfo struct {
	// Status 物流状态（口径以文档为准）。
	Status int `json:"status"`
	// LogisticsTypeID 物流方式 id。
	LogisticsTypeID string `json:"logistics_type_id"`
	// LogisticsTypeName 物流方式名称。
	LogisticsTypeName string `json:"logistics_type_name"`
	// LogisticsProviderID 物流商 id。
	LogisticsProviderID StringOrNumber `json:"logistics_provider_id"`
	// LogisticsProviderName 物流商名称。
	LogisticsProviderName string `json:"logistics_provider_name"`
	// ActualCarrier 实际承运人。
	ActualCarrier string `json:"actual_carrier"`
	// WaybillNo 运单号。
	WaybillNo string `json:"waybill_no"`

	// PreWeight 预估重量。
	PreWeight Float64OrString `json:"pre_weight"`
	// PreFeeWeight 预估计费重。
	PreFeeWeight Float64OrString `json:"pre_fee_weight"`
	// PreFeeWeightUnit 预估计费重单位。
	PreFeeWeightUnit string `json:"pre_fee_weight_unit"`
	// PrePkgLength 预估包裹长。
	PrePkgLength Float64OrString `json:"pre_pkg_length"`
	// PrePkgHeight 预估包裹高。
	PrePkgHeight Float64OrString `json:"pre_pkg_height"`
	// PrePkgWidth 预估包裹宽。
	PrePkgWidth Float64OrString `json:"pre_pkg_width"`
	// Weight 实际重量。
	Weight Float64OrString `json:"weight"`
	// PkgFeeWeight 实际计费重。
	PkgFeeWeight Float64OrString `json:"pkg_fee_weight"`
	// PkgFeeWeightUnit 实际计费重单位。
	PkgFeeWeightUnit string `json:"pkg_fee_weight_unit"`
	// PkgLength 实际包裹长。
	PkgLength Float64OrString `json:"pkg_length"`
	// PkgWidth 实际包裹宽。
	PkgWidth Float64OrString `json:"pkg_width"`
	// PkgHeight 实际包裹高。
	PkgHeight Float64OrString `json:"pkg_height"`
	// WeightUnit 重量单位。
	WeightUnit string `json:"weight_unit"`
	// PkgSizeUnit 尺寸单位。
	PkgSizeUnit string `json:"pkg_size_unit"`
	// CostCurrencyCode 运费币种 code。
	CostCurrencyCode string `json:"cost_currency_code"`
	// PreCostAmount 预估运费（字符串金额，可能包含币种符号）。
	PreCostAmount string `json:"pre_cost_amount"`
	// CostAmount 实际运费（字符串金额）。
	CostAmount string `json:"cost_amount"`
	// LogisticsTime 物流下单时间（结构以实际返回为准，可能为 null）。
	LogisticsTime json.RawMessage `json:"logistics_time"`
	// TrackingNo 跟踪号。
	TrackingNo string `json:"tracking_no"`
	// MarkNo 标记号。
	MarkNo string `json:"mark_no"`
}

// OrderDetailV2Transaction 表示交易汇总信息（展示字段通常带币种符号/格式）。
type OrderDetailV2Transaction struct {
	// OrderItemAmount 商品金额展示。
	OrderItemAmount string `json:"order_item_amount"`
	// CustomerTaxAmountShow 客付税费展示。
	CustomerTaxAmountShow string `json:"customer_tax_amount_show"`
	// DiscountAmount 折扣展示。
	DiscountAmount string `json:"discount_amount"`
	// CustomerShippingAmount 客付运费展示。
	CustomerShippingAmount string `json:"customer_shipping_amount"`
	// CustomerTipAmount 小费展示。
	CustomerTipAmount string `json:"customer_tip_amount"`
	// OrderTotalAmount 订单总金额展示。
	OrderTotalAmount string `json:"order_total_amount"`
	// CGPriceAmount 采购成本展示。
	CGPriceAmount string `json:"cg_price_amount"`
	// StockCostAmount 库存成本展示。
	StockCostAmount string `json:"stock_cost_amount"`
	// OutboundCostAmount 出库成本展示。
	OutboundCostAmount string `json:"outbound_cost_amount"`
	// WmsOutboundCostAmount 实际出库成本展示。
	WmsOutboundCostAmount string `json:"wms_outbound_cost_amount"`
	// PreCostAmount 预估运费展示。
	PreCostAmount string `json:"pre_cost_amount"`
	// WmsShippingPriceAmount 实际运费展示。
	WmsShippingPriceAmount string `json:"wms_shipping_price_amount"`
	// TransactionFeeAmount 交易费展示。
	TransactionFeeAmount string `json:"transaction_fee_amount"`
	// ProfitAmount 利润展示。
	ProfitAmount string `json:"profit_amount"`
	// OtherAmount 其他费用展示。
	OtherAmount string `json:"other_amount"`
	// PlatformSubsidyAmount 平台补贴展示。
	PlatformSubsidyAmount string `json:"platform_subsidy_amount"`
	// CodAmount COD 费用展示。
	CodAmount string `json:"cod_amount"`
	// GiftWrapAmount 礼品包装费展示。
	GiftWrapAmount string `json:"gift_wrap_amount"`
	// PlatformTaxAmount 平台税费展示。
	PlatformTaxAmount string `json:"platform_tax_amount"`
	// PointsGrantedAmount 积分成本展示。
	PointsGrantedAmount string `json:"points_granted_amount"`
	// OtherFee 其他费用展示。
	OtherFee string `json:"other_fee"`
}
