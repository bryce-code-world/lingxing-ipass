package lingxing

import "encoding/json"

// PlatformCode 表示多平台平台 code（枚举值以领星文档为准）。
//
// 文档示例（节选）：10001 AMAZON、10002 Shopify、10009 自定义平台、10008 Walmart、10011 TikTok 等。
type PlatformCode int

const (
	PlatformCodeAmazon         PlatformCode = 10001 // AMAZON
	PlatformCodeShopify        PlatformCode = 10002 // Shopify
	PlatformCodeEbay           PlatformCode = 10003 // eBay
	PlatformCodeWish           PlatformCode = 10004 // Wish
	PlatformCodeAliExpress     PlatformCode = 10005 // AliExpress
	PlatformCodeShopee         PlatformCode = 10006 // Shopee
	PlatformCodeLazada         PlatformCode = 10007 // Lazada
	PlatformCodeWalmart        PlatformCode = 10008 // Walmart
	PlatformCodeCustom         PlatformCode = 10009 // 自定义平台
	PlatformCodeWayfair        PlatformCode = 10010 // Wayfair
	PlatformCodeTikTok         PlatformCode = 10011 // TikTok
	PlatformCodeMercado        PlatformCode = 10012 // MERCADO
	PlatformCodeCdiscount      PlatformCode = 10013 // CDISCOUNT
	PlatformCodeNewegg         PlatformCode = 10014 // NEWEGG
	PlatformCodeRakuten        PlatformCode = 10015 // RAKUTEN
	PlatformCodeShopline       PlatformCode = 10016 // SHOPLINE
	PlatformCodeTeapplix       PlatformCode = 10017 // TEAPPLIX
	PlatformCodeShoplazza      PlatformCode = 10018 // SHOPLAZZA
	PlatformCodeUeeshop        PlatformCode = 10019 // UEESHOP
	PlatformCodeCoupang        PlatformCode = 10020 // COUPANG
	PlatformCodeShein          PlatformCode = 10021 // SHEIN
	PlatformCodeTemuFull       PlatformCode = 10022 // Temu 全托管
	PlatformCodeTemuSemi       PlatformCode = 10024 // Temu 半托管
	PlatformCodeOtto           PlatformCode = 10025 // OTTO
	PlatformCodeOzon           PlatformCode = 10026 // OZON
	PlatformCodeSheinFull      PlatformCode = 10027 // SHEIN 全托管
	PlatformCodeSheinSemi      PlatformCode = 10028 // SHEIN 半托管
	PlatformCodeAliExpressSemi PlatformCode = 10029 // AliExpress 半托管
	PlatformCodeAliExpressFull PlatformCode = 10030 // AliExpress 全托管
	PlatformCodeQoo10          PlatformCode = 10033 // Qoo10
	PlatformCodeMirakl         PlatformCode = 10034 // Mirakl
	PlatformCodeAmazonVC       PlatformCode = 10035 // AMAZON VC
	PlatformCodeKaufland       PlatformCode = 10036 // Kaufland
	PlatformCodeAllegro        PlatformCode = 10037 // Allegro
	PlatformCodeLineShopping   PlatformCode = 10038 // Line Shopping
	PlatformCodeSPSCommerce    PlatformCode = 10039 // SPS Commerce
)

// MultiPlatformOrderDateType 表示“查询订单管理订单列表”的日期类型（枚举值以领星文档为准）。
type MultiPlatformOrderDateType string

const (
	MultiPlatformOrderDateTypeUpdateTime         MultiPlatformOrderDateType = "update_time"          // 更新时间
	MultiPlatformOrderDateTypeGlobalPurchaseTime MultiPlatformOrderDateType = "global_purchase_time" // 订购时间
	MultiPlatformOrderDateTypeGlobalDeliveryTime MultiPlatformOrderDateType = "global_delivery_time" // 发货时间
	MultiPlatformOrderDateTypeGlobalPaymentTime  MultiPlatformOrderDateType = "global_payment_time"  // 付款时间
	MultiPlatformOrderDateTypeDeliveryTime       MultiPlatformOrderDateType = "delivery_time"        // 平台发货时间
)

// MultiPlatformOrderStatus 表示系统订单状态（枚举值以领星文档为准）。
//
// 对接接口：
// - /pb/mp/order/v2/list
//   - 请求字段：order_status
//   - 返回字段：data>>list>>status
//
// 枚举：
// 1 同步中
// 2 已同步
// 3 待付款
// 4 待审核
// 5 待发货
// 6 已发货
// 7 已取消/不发货
// 8 不显示
// 9 平台发货
type MultiPlatformOrderStatus int

const (
	MultiPlatformOrderStatusSyncing         MultiPlatformOrderStatus = 1 // 同步中
	MultiPlatformOrderStatusSynced          MultiPlatformOrderStatus = 2 // 已同步
	MultiPlatformOrderStatusPendingPayment  MultiPlatformOrderStatus = 3 // 待付款
	MultiPlatformOrderStatusPendingReview   MultiPlatformOrderStatus = 4 // 待审核
	MultiPlatformOrderStatusPendingShipment MultiPlatformOrderStatus = 5 // 待发货
	MultiPlatformOrderStatusShipped         MultiPlatformOrderStatus = 6 // 已发货
	MultiPlatformOrderStatusCancelledNoShip MultiPlatformOrderStatus = 7 // 已取消/不发货
	MultiPlatformOrderStatusHidden          MultiPlatformOrderStatus = 8 // 不显示
	MultiPlatformOrderStatusPlatformShipped MultiPlatformOrderStatus = 9 // 平台发货
)

// CreateOrderAddressType 表示“创建订单”的地址类型（枚举值以领星文档为准）。
type CreateOrderAddressType int

const (
	CreateOrderAddressTypeResidential CreateOrderAddressType = 1 // 住宅地址
	CreateOrderAddressTypeBusiness    CreateOrderAddressType = 2 // 商业地址
)

// StockDeductionType 表示“创建订单-商品行”的库存扣减类型（枚举值以领星文档为准）。
type StockDeductionType int

const (
	StockDeductionTypeEmpty            StockDeductionType = 1 // “空”
	StockDeductionTypeSKUAndOrderStore StockDeductionType = 2 // “SKU+订单店铺”
)

// NewPlatformOrderListRequest 表示“查询平台订单列表”的请求参数。
//
// API Path: /cepfPlatformOrder/open-api/newPlatformOrder/list
type NewPlatformOrderListRequest struct {
	// DateType 时间类型：
	// 0 平台数据变动时间
	// 1 订购时间
	// 2 订购时间-北京
	// 3 支付时间
	// 4 支付时间-北京
	// 5 发货时间
	// 6 发货时间-北京
	DateType int `json:"dateType"`

	// DeliveryTypeList 配送类型列表（可选，枚举口径以领星文档为准）。
	DeliveryTypeList []int `json:"deliveryTypeList,omitempty"`
	// PageNum 页码（可选）。
	PageNum int `json:"pageNum,omitempty"`
	// PageSize 每页大小（可选）。
	PageSize int `json:"pageSize,omitempty"`
	// PlatformCodeList 平台 code 列表（可选）。
	PlatformCodeList []string `json:"platformCodeList,omitempty"`
	// SearchMultiValue 多值搜索内容（可选）。
	SearchMultiValue []string `json:"searchMultiValue,omitempty"`
	// SearchSingleValue 单值搜索内容（可选）。
	SearchSingleValue string `json:"searchSingleValue,omitempty"`
	// SearchType 搜索类型（可选，枚举口径以领星文档为准）。
	SearchType int `json:"searchType,omitempty"`
	// SiteCodeList 站点/站点代码列表（可选）。
	SiteCodeList []string `json:"siteCodeList,omitempty"`
	// SortField 排序字段（可选）。
	SortField string `json:"sortField,omitempty"`
	// SortType 排序方式（可选）。
	SortType string `json:"sortType,omitempty"`
	// StartDate 开始时间（字符串格式以领星文档为准）。
	StartDate string `json:"startDate"`
	// EndDate 结束时间（字符串格式以领星文档为准）。
	EndDate string `json:"endDate"`
	// StatusList 订单状态列表（可选）。
	StatusList []string `json:"statusList,omitempty"`
	// StoreIDList 店铺 ID 列表（可选）。
	StoreIDList []string `json:"storeIdList,omitempty"`
}

// NewPlatformOrderListResponse 表示“查询平台订单列表”的 data 字段。
//
// 注意：订单列表字段非常多且变化频繁，这里将 list 用 json.RawMessage 保留原始结构，
// 业务侧如需强类型字段，可自行定义结构体再二次反序列化。
type NewPlatformOrderListResponse struct {
	// Current 当前页。
	Current int64 `json:"current"`
	// Total 总数。
	Total int64 `json:"total"`
	// List 订单列表元素（原始 JSON）。
	List []json.RawMessage `json:"list"`
}

// CreateOrdersV2Request 表示“创建订单”的请求参数。
//
// API Path: /pb/mp/order/v2/create
type CreateOrdersV2Request struct {
	// PlatformCode 平台 code（自定义平台等）。
	PlatformCode PlatformCode `json:"platform_code"`
	// StoreID 店铺 id。
	StoreID string `json:"store_id"`
	// Orders 订单列表。
	Orders []CreateOrderV2 `json:"orders"`
}

// CreateOrderV2 表示创建订单的单个订单对象。
type CreateOrderV2 struct {
	// PlatformOrderNo 平台单号（同一店铺不支持重复）。
	PlatformOrderNo string `json:"platform_order_no"`
	// SiteCode 站点（可选）。
	SiteCode string `json:"site_code,omitempty"`
	// BuyerNote 买家备注（可选）。
	BuyerNote string `json:"buyer_note,omitempty"`
	// ReceiverCountryCode 国家/地区二字简码（必填）。
	ReceiverCountryCode string `json:"receiver_country_code"`
	// ReceiverName 收件人（必填）。
	ReceiverName string `json:"receiver_name"`
	// City 城市（必填）。
	City string `json:"city"`
	// AddressLine1 地址 1（必填）。
	AddressLine1 string `json:"address_line1"`
	// AddressType 地址类型（可选）：
	// 1 住宅地址
	// 2 商业地址
	AddressType CreateOrderAddressType `json:"address_type,omitempty"`
	// AmountCurrency 币种（可选）。
	AmountCurrency string `json:"amount_currency,omitempty"`

	// WID 仓库 ID（可选）。
	WID string `json:"wid,omitempty"`
	// LogisticsTypeID 物流渠道/方式 ID（可选）。
	LogisticsTypeID string `json:"logistics_type_id,omitempty"`

	// 自发货订单额外字段（这些字段需要向领星申请权限，权限申请下来后，通过 /pb/mp/order/v2/list 接口响应参数 data>>list>>address_info 会出现下列字段信息）
	// District 区/县（可选）。
	District string `json:"district,omitempty"`
	// StateOrRegion 州/省（可选）。
	StateOrRegion string `json:"state_or_region,omitempty"`
	// PostalCode 邮政编码（可选）。
	PostalCode string `json:"postal_code,omitempty"`
	// DoorplateNo 门牌号（可选）。
	DoorplateNo string `json:"doorplate_no,omitempty"`
	// AddressLine2 地址 2（可选）。
	AddressLine2 string `json:"address_line2,omitempty"`
	// AddressLine3 地址 3（可选）。
	AddressLine3 string `json:"address_line3,omitempty"`
	// ReceiverMobile 收件人手机号（可选）。
	ReceiverMobile string `json:"receiver_mobile,omitempty"`
	// BuyerEmail 买家邮箱（可选）。
	BuyerEmail string `json:"buyer_email,omitempty"`
	// BuyerName 买家姓名（可选）。
	BuyerName string `json:"buyer_name,omitempty"`

	// Items 商品行列表（必填）。
	Items []CreateOrderItemV2 `json:"items"`
}

// CreateOrderItemV2 表示创建订单的商品行。
type CreateOrderItemV2 struct {
	// SKU 本地 SKU，领星平台维度的SKU（SKU 和 MSKU 二选一）。
	SKU string `json:"sku,omitempty"`
	// MSKU Marketplace SKU，也叫“店铺 SKU / 渠道 SKU / 刊登 SKU”，通常是店铺维度的 SKU（SKU 和 MSKU 二选一）。
	MSKU string `json:"msku,omitempty"`
	// Quantity 数量（必填）。
	Quantity int `json:"quantity"`
	// UnitPrice 单价（必填）。
	UnitPrice float64 `json:"unit_price"`
	// StockDeductionType 库存扣减类型（可选，枚举口径以领星文档为准）。
	StockDeductionType StockDeductionType `json:"stock_deduction_type,omitempty"`
}

// CreateOrdersV2ResponseData 表示“创建订单”的 data 字段。
type CreateOrdersV2ResponseData struct {
	// ErrorDetails 失败明细列表。
	ErrorDetails []CreateOrdersV2ErrorDetail `json:"error_details"`
	// SuccessDetails 成功明细列表。
	SuccessDetails []CreateOrdersV2SuccessDetail `json:"success_details"`
}

// CreateOrdersV2ErrorDetail 表示创建订单失败详情。
type CreateOrdersV2ErrorDetail struct {
	// ErrorMessage 错误信息。
	ErrorMessage string `json:"error_message"`
	// PlatformOrderNo 平台单号。
	PlatformOrderNo string `json:"platform_order_no"`
}

// CreateOrdersV2SuccessDetail 表示创建订单成功详情。
type CreateOrdersV2SuccessDetail struct {
	// GlobalOrderNo 系统单号（领星全局单号）。
	GlobalOrderNo string `json:"global_order_no"`
	// PlatformOrderNo 平台单号。
	PlatformOrderNo string `json:"platform_order_no"`
}

// OrderListV2Request 表示“查询订单管理订单列表”的请求参数。
//
// API Path: /pb/mp/order/v2/list
type OrderListV2Request struct {
	// Offset 分页偏移量（必填）。
	Offset int `json:"offset"`
	// Length 分页大小（必填）。
	Length int `json:"length"`

	// DateType 时间类型：update_time/global_purchase_time/global_delivery_time/global_payment_time/delivery_time 等。
	DateType MultiPlatformOrderDateType `json:"date_type,omitempty"`
	// StartTime/EndTime 时间戳（秒）。
	StartTime int64 `json:"start_time,omitempty"`
	EndTime   int64 `json:"end_time,omitempty"`

	// StoreID 店铺 ID 列表（可选）。
	StoreID []string `json:"store_id,omitempty"`
	// PlatformCode 平台 code 列表（可选）。
	PlatformCode []PlatformCode `json:"platform_code,omitempty"`

	// PlatformOrderNos 平台单号列表（可选）。
	PlatformOrderNos []string `json:"platform_order_nos,omitempty"`
	// PlatformOrderNames 平台订单名称列表（可选）。
	PlatformOrderNames []string `json:"platform_order_names,omitempty"`

	// OrderStatus 订单状态（领星：3待付款/4待审核/5待发货/6已发货...）。
	OrderStatus MultiPlatformOrderStatus `json:"order_status,omitempty"`

	// IncludeDelete 是否包含已删除订单（可选）。
	IncludeDelete *bool `json:"include_delete,omitempty"`
}

// OrderListV2ResponseData 表示“查询订单管理订单列表”的 data 字段。
type OrderListV2ResponseData struct {
	// Total 总数。
	Total StringOrNumber `json:"total"`
	// List 订单列表元素（原始 JSON）。
	List []OrderDetailV2 `json:"list"`
}

// OrderListV2Item 表示订单列表单条数据（一期最小字段集合）。
//
// 说明：
// - 领星订单列表字段很多且口径可能调整；一期只固化当前系统会用到的字段。
// - 如后续需要更多字段，再按真实用例增补即可，避免一次性把全量字段都“翻译”进系统。
type OrderListV2Item struct {
	// GlobalOrderNo 系统单号（领星全局单号）。
	GlobalOrderNo string `json:"global_order_no"`
	// UpdateTime 订单更新时间（秒级时间戳字符串，口径以领星文档为准）。
	UpdateTime string `json:"update_time,omitempty"`
	// OrderStatus 系统订单状态。
	//
	// 说明：领星文档在响应体中使用 `status` 字段；历史上也见过 `order_status`。
	// 为兼容两种返回，SDK 通过 UnmarshalJSON 做了兜底。
	OrderStatus MultiPlatformOrderStatus `json:"status,omitempty"`
	// PlatformInfo 平台单信息列表。
	PlatformInfo []OrderListV2PlatformInfo `json:"platform_info,omitempty"`
}

func (i *OrderListV2Item) UnmarshalJSON(b []byte) error {
	var aux struct {
		GlobalOrderNo string                    `json:"global_order_no"`
		UpdateTime    string                    `json:"update_time,omitempty"`
		Status        *MultiPlatformOrderStatus `json:"status"`
		OrderStatus   *MultiPlatformOrderStatus `json:"order_status"`
		PlatformInfo  []OrderListV2PlatformInfo `json:"platform_info,omitempty"`
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	i.GlobalOrderNo = aux.GlobalOrderNo
	i.UpdateTime = aux.UpdateTime
	i.PlatformInfo = aux.PlatformInfo

	if aux.Status != nil {
		i.OrderStatus = *aux.Status
		return nil
	}
	if aux.OrderStatus != nil {
		i.OrderStatus = *aux.OrderStatus
		return nil
	}
	return nil
}

// OrderListV2PlatformInfo 表示订单列表中的平台单信息（一期最小字段集合）。
type OrderListV2PlatformInfo struct {
	// PlatformOrderNo 平台单号。
	PlatformOrderNo string `json:"platform_order_no"`
	// PlatformOrderName 平台订单名称/编号（可选）。
	PlatformOrderName string `json:"platform_order_name,omitempty"`
}
