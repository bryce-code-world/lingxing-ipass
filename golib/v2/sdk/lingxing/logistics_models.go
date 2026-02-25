package lingxing

// ChannelListRequest 表示“查询头程物流渠道列表”的请求参数。
//
// API Path: /erp/sc/data/local_inventory/channelList
type ChannelListRequest struct {
	// Offset 分页偏移量。
	Offset int `json:"offset"`
	// Length 分页大小。
	Length int `json:"length"`
}

// LogisticsChannel 表示头程物流渠道信息。
type LogisticsChannel struct {
	// ID 渠道 ID。
	ID string `json:"id"`
	// ChannelName 物流渠道名称。
	ChannelName string `json:"channel_name"`
	// MethodID 运输方式/物流方式 ID。
	MethodID string `json:"method_id"`
	// MethodName 运输方式/物流方式名称。
	MethodName string `json:"method_name"`
	// BillingType 计费类型（枚举口径以领星文档为准，例如 0计费重/1体积）。
	BillingType int `json:"billing_type"`
	// ZipCode 邮编（口径以领星文档为准）。
	ZipCode string `json:"zip_code"`
	// ValidPeriod 有效期（口径以领星文档为准）。
	ValidPeriod int `json:"valid_period"`
	// Remark 备注。
	Remark string `json:"remark"`
	// Enabled 是否启用（枚举口径以领星文档为准）。
	Enabled int `json:"enabled"`
	// GmtModified 修改时间（字符串格式以领星文档为准）。
	GmtModified string `json:"gmt_modified"`
	// Provider 物流商信息（简要）。
	Provider LogisticsProviderBrief `json:"provider"`
	// Freight 运费规则列表。
	Freight []ChannelFreightRule `json:"freight"`

	// SendPlaceCode 提货地代码。
	SendPlaceCode string `json:"send_place_code"`
	// ReceiveCountryCode 收货国家二字码。
	ReceiveCountryCode string `json:"receive_country_code"`
	// IsIncludeTax 是否含税（枚举口径以领星文档为准）。
	IsIncludeTax int `json:"is_include_tax"`
	// IsPointsBehind 是否分抛（枚举口径以领星文档为准）。
	IsPointsBehind int `json:"is_points_behind"`
	// PointsBehindCoeffient 分抛系数（口径以领星文档为准）。
	PointsBehindCoeffient float64 `json:"points_behind_coeffient"`
}

// LogisticsProviderBrief 表示渠道返回中的物流商简要信息。
type LogisticsProviderBrief struct {
	// ID 物流商 ID。
	ID string `json:"id"`
	// LogisticsProviderName 物流商名称。
	LogisticsProviderName string `json:"logistics_provider_name"`
}

// ChannelFreightRule 表示渠道的运费规则。
type ChannelFreightRule struct {
	// CountryCode 国家二字码。
	CountryCode string `json:"country_code"`
	// RegionCode 州/省代码（口径以领星文档为准）。
	RegionCode string `json:"region_code"`
	// BillingWeightStart 计费重量起始区间（字符串口径以领星文档为准）。
	BillingWeightStart string `json:"billing_weight_start"`
	// BillingPrice 计费价格（字符串口径以领星文档为准）。
	BillingPrice string `json:"billing_price"`
}

// QueryHeadLogisticsProviderListRequest 表示“查询物流-头程物流商”的请求参数。
//
// API Path: /basicOpen/logistics/headLogisticsProvider/query/list
type QueryHeadLogisticsProviderListRequest struct {
	// Search 查询条件。
	Search HeadLogisticsProviderSearch `json:"search"`
}

// HeadLogisticsProviderSearch 表示查询条件。
type HeadLogisticsProviderSearch struct {
	// Page 页码。
	Page int `json:"page"`
	// Length 每页大小。
	Length int `json:"length"`
	// Enabled 是否启用（可选）。
	Enabled *int `json:"enabled,omitempty"`
	// IsAuth 是否授权（可选）。
	IsAuth *int `json:"isAuth,omitempty"`
	// PayMethod 结算方式（可选，枚举口径以领星文档为准）。
	PayMethod *int `json:"payMethod,omitempty"`
	// SearchField 搜索字段（可选）。
	SearchField string `json:"searchField,omitempty"`
	// SearchValue 搜索值（可选）。
	SearchValue string `json:"searchValue,omitempty"`
}

// HeadLogisticsProviderListResponse 表示“查询物流-头程物流商”的 data 字段。
type HeadLogisticsProviderListResponse struct {
	// Total 总数。
	Total int `json:"total"`
	// Providers 物流商列表。
	Providers []HeadLogisticsProvider `json:"providers"`
}

// HeadLogisticsProvider 表示物流商信息。
type HeadLogisticsProvider struct {
	// ProviderID 物流商 ID。
	ProviderID string `json:"providerId"`
	// Name 物流商名称。
	Name string `json:"name"`
	// Code 物流商编码。
	Code string `json:"code"`
	// Enabled 是否启用。
	Enabled int `json:"enabled"`
	// LogisticsType 物流类型（枚举口径以领星文档为准）。
	LogisticsType int `json:"logisticsType"`
	// IsAuth 是否授权。
	IsAuth int `json:"isAuth"`
	// SupplierCode 供应商编码（口径以领星文档为准）。
	SupplierCode int `json:"supplierCode"`
	// SupplierName 供应商名称（口径以领星文档为准）。
	SupplierName string `json:"supplierName"`
	// Status 状态（枚举口径以领星文档为准）。
	Status int `json:"status"`
	// Remark 备注。
	Remark string `json:"remark"`
	// PayMethod 结算方式（枚举口径以领星文档为准）。
	PayMethod int `json:"payMethod"`
	// ContactName 联系人姓名。
	ContactName string `json:"contactName"`
	// ContactPhone 联系人电话。
	ContactPhone string `json:"contactPhone"`
	// CreatorID 创建人 ID。
	CreatorID int64 `json:"creatorId"`
	// CreatorName 创建人名称。
	CreatorName string `json:"creatorName"`
	// CreatedAt 创建时间（时间戳，口径以领星文档为准）。
	CreatedAt int64 `json:"createdAt"`
}
