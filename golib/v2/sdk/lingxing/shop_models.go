package lingxing

import "encoding/json"

// AmazonSeller 表示“查询亚马逊店铺列表”返回的店铺信息。
type AmazonSeller struct {
	// SID 店铺ID（领星ERP对已授权店铺的唯一标识）。
	SID int64 `json:"sid"`
	// MID 站点ID。
	MID int64 `json:"mid"`
	// Name 店铺名称。
	Name string `json:"name"`
	// SellerID 亚马逊店铺 seller_id（卖家记号）。
	SellerID string `json:"seller_id"`
	// AccountName 店铺账号名称。
	AccountName string `json:"account_name"`
	// SellerAccountID 店铺账号ID。
	SellerAccountID int64 `json:"seller_account_id"`
	// Region 站点简称（例如 NA/EU）。
	Region string `json:"region"`
	// Country 店铺所在国家名称。
	Country string `json:"country"`
	// HasAdsSetting 是否授权广告：0 否，1 是。
	HasAdsSetting int `json:"has_ads_setting"`
	// MarketplaceID 市场ID。
	MarketplaceID string `json:"marketplace_id"`
	// Status 店铺状态：0 停止同步，1 正常，2 授权异常，3 欠费停服。
	Status int `json:"status"`
}

// AmazonConceptSellerListResponse 表示“查询亚马逊概念店铺列表”的返回结构（包含顶层 total）。
type AmazonConceptSellerListResponse struct {
	// Total 总数。
	Total int
	// List 概念店铺列表。
	List []AmazonConceptSeller
}

// AmazonConceptSeller 表示“查询亚马逊概念店铺列表”返回的概念店铺信息。
type AmazonConceptSeller struct {
	// ID 概念店铺ID（唯一标识）。
	ID int64 `json:"id"`
	// MID 概念市场ID。
	MID int64 `json:"mid"`
	// Name 概念店铺名称。
	Name string `json:"name"`
	// SellerID 亚马逊卖家记号。
	SellerID string `json:"seller_id"`
	// SellerAccountName 店铺账号名称。
	SellerAccountName string `json:"seller_account_name"`
	// SellerAccountID 店铺账号ID。
	SellerAccountID int64 `json:"seller_account_id"`
	// Region 站点简称（例如 NA/EU）。
	Region string `json:"region"`
	// Country 店铺所在国家名称（例如 北美共享/欧洲共享）。
	Country string `json:"country"`
	// Status 概念店铺状态：1 启用，2 禁用。
	Status int `json:"status"`
}

// AmazonSellerBatchRenameRequest 表示“批量修改店铺名称”的请求参数。
type AmazonSellerBatchRenameRequest struct {
	// SIDNameList 批量修改店铺数组（最多 10 个）。
	SIDNameList []AmazonSellerSIDName `json:"sid_name_list"`
}

// AmazonSellerSIDName 表示待修改的店铺信息。
type AmazonSellerSIDName struct {
	// SID 店铺ID（对应 SellerLists 返回字段 sid）。
	SID int64 `json:"sid"`
	// Name 店铺名称。
	Name string `json:"name"`
}

// AmazonSellerBatchRenameResponseData 表示“批量修改店铺名称”返回的 data 字段。
//
// 兼容两类返回：
// 1) 按文档：success_num / failure_num / failure_detail
// 2) 按示例：sid / name
type AmazonSellerBatchRenameResponseData struct {
	// SuccessNum 成功个数。
	SuccessNum int `json:"success_num"`
	// FailureNum 失败个数。
	FailureNum int `json:"failure_num"`
	// FailureDetail 失败详情。
	FailureDetail []AmazonSellerBatchRenameFailure `json:"failure_detail"`

	// SID 示例返回的店铺ID（部分示例会直接返回 sid/name）。
	SID int64 `json:"sid"`
	// Name 示例返回的店铺名称（部分示例会直接返回 sid/name）。
	Name string `json:"name"`

	// Extra 用于保留未建模字段，避免后续新增字段导致反序列化失败（不会参与序列化）。
	Extra map[string]json.RawMessage `json:"-"`
}

// AmazonSellerBatchRenameFailure 表示批量改名失败的明细项。
type AmazonSellerBatchRenameFailure struct {
	// SID 店铺ID。
	SID string `json:"sid"`
	// Name 店铺名称。
	Name string `json:"name"`
	// Error 失败原因。
	Error string `json:"error"`
}

// MultiPlatformStoreListV2Request 表示“查询多平台店铺信息”的请求参数。
type MultiPlatformStoreListV2Request struct {
	// Offset 分页偏移量。
	Offset int `json:"offset,omitempty"`
	// Length 分页长度，上限 200。
	Length int `json:"length,omitempty"`
	// PlatformCode 平台 code 列表（例如 10001 AMAZON、10002 Shopify、10009 自定义平台等）。
	PlatformCode []int `json:"platform_code,omitempty"`
	// IsSync 店铺同步状态：1 启用，0 停用。
	IsSync *int `json:"is_sync,omitempty"`
	// Status 店铺授权状态：1 正常授权，0 授权失败。
	Status *int `json:"status,omitempty"`
}

// MultiPlatformStoreListV2ResponseData 表示“查询多平台店铺信息”返回的 data 字段。
//
// 注意：total 可能为 string 或 number，这里用 RawMessage 保留原始值，便于兼容。
type MultiPlatformStoreListV2ResponseData struct {
	// Total 总数（可能为 string 或 number）。
	Total json.RawMessage `json:"total"`
	// List 店铺数据。
	List []MultiPlatformStoreItem `json:"list"`
}

// TotalInt 尝试把 total 解析为 int。
func (d MultiPlatformStoreListV2ResponseData) TotalInt() (int, bool) {
	return parseIntFromRaw(d.Total)
}

// MultiPlatformStoreItem 表示“查询多平台店铺信息”返回的单个店铺条目。
type MultiPlatformStoreItem struct {
	// StoreID 多平台店铺ID（多平台店铺唯一标识）。
	StoreID string `json:"store_id"`
	// SID 店铺ID（亚马逊店铺时对应 SellerLists 返回字段 sid；其他平台可能为空）。
	SID string `json:"sid"`
	// StoreName 店铺名称。
	StoreName string `json:"store_name"`
	// PlatformCode 平台 code。
	PlatformCode string `json:"platform_code"`
	// PlatformName 平台名称。
	PlatformName string `json:"platform_name"`
	// Currency 店铺币种。
	Currency string `json:"currency"`
	// IsSync 店铺同步状态：1 启用，0 停用。
	IsSync int `json:"is_sync"`
	// Status 店铺授权状态：1 正常授权，0 授权失败。
	Status int `json:"status"`
}
