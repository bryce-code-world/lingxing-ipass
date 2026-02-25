package lingxing

// InventoryDetailsRequest 表示“查询仓库库存明细”的请求参数。
//
// API Path: /erp/sc/routing/data/local_inventory/inventoryDetails
type InventoryDetailsRequest struct {
	// WID 仓库 ID。
	WID string `json:"wid,omitempty"`
	// Offset 分页偏移量。
	Offset int `json:"offset,omitempty"`
	// Length 分页大小。
	Length int `json:"length,omitempty"`
	// SKU 按 领星本地 SKU 筛选（可选）。
	SKU string `json:"sku,omitempty"`
}

// InventoryDetailsItem 表示库存明细条目（仅保留一期需要的字段集合）。
type InventoryDetailsItem struct {
	// WID 仓库 ID。
	WID int `json:"wid"`
	// ProductID 商品 ID。
	ProductID int `json:"product_id"`
	// SKU 本地 SKU。
	SKU string `json:"sku"`
	// ProductTotal 实际库存总量（文档口径：可用量 + 次品量 + 待检/待上架量 + 锁定量等）。
	ProductTotal int `json:"product_total"`
	// ProductValidNum 可用量（接口文档字段：product_valid_num，一期同步该字段到 DSCO）。
	ProductValidNum int `json:"product_valid_num"`
}

// InventoryBinDetailsRequest 表示“查询仓位库存明细”的请求参数。
//
// API Path: /erp/sc/routing/data/local_inventory/inventoryBinDetails
type InventoryBinDetailsRequest struct {
	// WID 仓库 ID。
	WID string `json:"wid,omitempty"`
	// BinTypeList 仓位类型列表（字符串形式，具体口径以领星文档为准）。
	BinTypeList string `json:"bin_type_list,omitempty"`
	// Offset 分页偏移量。
	Offset int `json:"offset,omitempty"`
	// Length 分页大小。
	Length int `json:"length,omitempty"`
}

// InventoryBinDetailsItem 表示仓位库存明细条目（仅保留一期需要的字段集合）。
type InventoryBinDetailsItem struct {
	// WID 仓库 ID。
	WID int `json:"wid"`
	// WHBID 仓位 ID。
	WHBID int `json:"whb_id"`
	// SKU 本地 SKU。
	SKU string `json:"sku"`
	// Total 仓位库存数量。
	Total int `json:"total"`
}

// WarehouseBinRequest 表示“查询本地仓位列表”的请求参数。
//
// API Path: /erp/sc/routing/data/local_inventory/warehouseBin
type WarehouseBinRequest struct {
	// WID 仓库 ID（字符串形式，多个用英文逗号分隔）。
	WID string `json:"wid,omitempty"`
	// ID 仓位 ID（字符串形式，多个用英文逗号分隔）。
	ID string `json:"id,omitempty"`

	// Status 仓位状态：1 禁用；2 启用（字符串形式，口径以领星文档为准）。
	Status string `json:"status,omitempty"`
	// Type 仓位类型：5 可用；6 次品（字符串形式，口径以领星文档为准）。
	Type string `json:"type,omitempty"`

	// Offset 分页偏移量（默认 0）。
	Offset int `json:"offset,omitempty"`
	// Limit 分页大小（默认 20）。
	Limit int `json:"limit,omitempty"`
}

// WarehouseBinItem 表示本地仓位条目。
type WarehouseBinItem struct {
	// ID 仓位 ID。
	ID int `json:"id"`
	// WID 仓库 ID。
	WID int `json:"wid"`
	// WarehouseName 仓库名称（字段：Ware_house_name）。
	WarehouseName string `json:"Ware_house_name"`
	// StorageBin 仓位编码/名称（字段：storage_bin）。
	StorageBin string `json:"storage_bin"`
	// Status 仓位状态（字段：status）。
	Status int `json:"status"`
	// Type 仓位类型（字段：type）。
	Type int `json:"type"`
	// SKUFnSKU 仓位商品关系（字段：sku_fnsku）。
	SKUFnSKU []WarehouseBinSKUFnSKU `json:"sku_fnsku,omitempty"`
}

// WarehouseBinSKUFnSKU 表示仓位与商品的关系信息。
type WarehouseBinSKUFnSKU struct {
	// SKU SKU。
	SKU string `json:"sku"`
	// ProductID 商品 ID。
	ProductID int `json:"product_id"`
	// FNSKU FNSKU（可选）。
	FNSKU string `json:"fnsku"`
	// StoreID 店铺 ID（字符串，文档示例为 "0"）。
	StoreID string `json:"store_id"`
	// SellerName 店铺名称（可选）。
	SellerName string `json:"seller_name"`
	// ProductName 商品名称（可选）。
	ProductName string `json:"product_name"`
}

// WarehouseStatementRequest 表示“查询库存流水（旧）/查询仓位流水”的请求参数。
//
// API Path:
// - /erp/sc/routing/data/local_inventory/wareHouseStatement
// - /erp/sc/routing/data/local_inventory/wareHouseBinStatement
type WarehouseStatementRequest struct {
	// WID 仓库 ID（多个用英文逗号分隔；不填默认所有仓库）。
	WID string `json:"wid,omitempty"`
	// Type 流水类型（多个用英文逗号分隔；不填默认全部类型）。
	Type string `json:"type,omitempty"`
	// StartDate 操作开始时间（格式：Y-m-d，闭区间，联合 EndDate 使用）。
	StartDate string `json:"start_date,omitempty"`
	// EndDate 操作结束时间（格式：Y-m-d，开区间，联合 StartDate 使用）。
	EndDate string `json:"end_date,omitempty"`
	// Offset 分页偏移量（默认 0）。
	Offset int `json:"offset,omitempty"`
	// Length 分页长度（默认 20）。
	Length int `json:"length,omitempty"`
}

// WarehouseStatementItem 表示库存流水条目（旧）。
type WarehouseStatementItem struct {
	// StatementID 流水 ID。
	StatementID string `json:"statement_id"`
	// WID 仓库 ID。
	WID int `json:"wid"`
	// WarehouseName 仓库名称（字段：ware_house_name）。
	WarehouseName string `json:"ware_house_name"`
	// BID 商品品牌 ID。
	BID int `json:"bid"`
	// BrandName 品牌名称。
	BrandName string `json:"brand_name"`

	// OrderSN 单据号（字段：order_sn）。
	OrderSN string `json:"order_sn"`
	// RefOrderSN 关联单据号（可选，字段：ref_order_sn）。
	RefOrderSN string `json:"ref_order_sn,omitempty"`

	// ProductID 商品 ID。
	ProductID int `json:"product_id"`
	// ProductName 品名。
	ProductName string `json:"product_name"`
	// SKU SKU。
	SKU string `json:"sku"`
	// SellerID 店铺 ID（字符串，文档示例为 "0"）。
	SellerID string `json:"seller_id,omitempty"`
	// FNSKU FNSKU（可选）。
	FNSKU string `json:"fnsku"`

	// ProductTotal 商品总量。
	ProductTotal float64 `json:"product_total"`
	// ProductGoodNum 良品量。
	ProductGoodNum float64 `json:"product_good_num"`
	// ProductBadNum 次品量。
	ProductBadNum float64 `json:"product_bad_num"`
	// ProductQCNum 质检量。
	ProductQCNum float64 `json:"product_qc_num"`
	// ProductLockNum 锁定量。
	ProductLockNum float64 `json:"product_lock_num"`

	// Price 单价（文档示例为字符串/数字混用，使用 StringOrNumber 兼容）。
	Price StringOrNumber `json:"price"`
	// Amount 金额（文档示例为字符串/数字混用，使用 StringOrNumber 兼容）。
	Amount StringOrNumber `json:"amount"`
	// ProductAmounts 货值（字段：product_amounts）。
	ProductAmounts StringOrNumber `json:"product_amounts"`

	// Type 流水类型（数值）。
	Type int `json:"type"`
	// TypeText 流水类型文本。
	TypeText string `json:"type_text"`
	// Remark 备注。
	Remark string `json:"remark"`

	// FeeCost 费用成本。
	FeeCost StringOrNumber `json:"fee_cost"`
	// SingleFeeCost 单位费用成本。
	SingleFeeCost StringOrNumber `json:"single_fee_cost"`
	// SingleCGPrice 采购单价。
	SingleCGPrice StringOrNumber `json:"single_cg_price"`

	// OptUID 操作人员 ID。
	OptUID int `json:"opt_uid"`
	// OptTime 操作时间（字符串格式以领星文档为准）。
	OptTime string `json:"opt_time"`
	// OptRealName 操作人员姓名。
	OptRealName string `json:"opt_realname"`
	// CancelTime 撤销时间（字符串格式以领星文档为准）。
	CancelTime string `json:"cancel_time"`
}
