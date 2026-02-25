package lingxing

import "encoding/json"

// WmsOrderListRequest 表示“查询销售出库单列表”的请求参数。
//
// API Path: /erp/sc/routing/wms/order/wmsOrderList
type WmsOrderListRequest struct {
	// Page 页码（可选）。
	Page int `json:"page,omitempty"`
	// PageSize 每页大小（可选）。
	PageSize int `json:"page_size,omitempty"`

	// SIDArr 店铺/站点 SID 列表（可选，口径以领星文档为准）。
	SIDArr []int `json:"sid_arr,omitempty"`

	// StatusArr 状态：1物流下单 2发货中 3已发货 4已删除
	StatusArr []int `json:"status_arr,omitempty"`

	// LogisticsStatusArr 物流下单状态（可选）。
	LogisticsStatusArr []int `json:"logistics_status_arr,omitempty"`

	// PlatformOrderNoArr 平台单号列表（可选）。
	PlatformOrderNoArr []string `json:"platform_order_no_arr,omitempty"`
	// OrderNumberArr 订单号列表（可选）。
	OrderNumberArr []string `json:"order_number_arr,omitempty"`
	// WoNumberArr 出库单号列表（可选）。
	WoNumberArr []string `json:"wo_number_arr,omitempty"`

	// TimeType 时间类型：create_at/delivered_at/stock_delivered_at/update_at
	TimeType string `json:"time_type,omitempty"`
	// StartDate 起始时间（字符串格式以领星文档为准）。
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间（字符串格式以领星文档为准）。
	EndDate string `json:"end_date,omitempty"`
}

// ListUsedLogisticsTypeRequest 表示“查询已启用的自发货物流方式”的请求参数。
//
// API Path: /erp/sc/routing/wms/WmsLogistics/listUsedLogisticsType
type ListUsedLogisticsTypeRequest struct {
	Param ListUsedLogisticsTypeParam `json:"param"`
}

// ListUsedLogisticsTypeParam 表示查询条件。
type ListUsedLogisticsTypeParam struct {
	// ProviderType 物流商类型：0 API物流 / 1 自定义物流 / 2 海外仓物流 / 4 平台物流
	ProviderType int `json:"provider_type"`
	// Page 分页页码（可选）
	Page int `json:"page,omitempty"`
	// Length 分页长度（可选）
	Length int `json:"length,omitempty"`
}

// UsedLogisticsType 表示“已启用的自发货物流方式”条目（按文档字段 + 常用字段保留）。
type UsedLogisticsType struct {
	// Type 物流商类型（同 provider_type）
	Type int `json:"type"`

	// LogisticsProviderID 物流商 ID（文档为 int，实际常见为 string/number 混用）
	LogisticsProviderID StringOrNumber `json:"logistics_provider_id"`
	// LogisticsProviderName 物流商名称
	LogisticsProviderName string `json:"logistics_provider_name"`

	// TypeID 物流方式 ID（文档为 int，实际常见为 string/number 混用）
	TypeID StringOrNumber `json:"type_id"`
	// Name 物流方式名称
	Name string `json:"name"`
	// IsUsed 渠道是否启用：0 停用 / 1 启用
	IsUsed int `json:"is_used"`

	// Code 渠道 code（示例中存在）
	Code string `json:"code,omitempty"`
}

// WmsOrder 表示销售出库单（对应领星“查询销售出库单列表”接口）。
//
// API Path: /erp/sc/routing/wms/order/wmsOrderList
//
// 时间字段说明：
// - 领星该接口的时间字段为字符串，常见格式为 "2006-01-02 15:04:05"；也可能返回空字符串。
// - 部分时间字段可能返回 "0000-00-00 00:00:00"（表示“未发生/未打印”等），调用方需要自行判空/兜底。
type WmsOrder struct {
	// WoID 出库单 ID。
	WoID int64 `json:"wo_id,omitempty"`
	// WoNumber 出库单号。
	WoNumber string `json:"wo_number"`

	// SID 店铺/站点 ID。
	SID int64 `json:"sid,omitempty"`
	// WID 仓库 ID。
	WID int `json:"wid,omitempty"`
	// WarehouseType 仓库类型：1 本地仓库 / 2 FBA仓 / 3 第三方海外仓。
	WarehouseType int `json:"warehouse_type,omitempty"`

	// OrderNumber 系统单号。
	OrderNumber string `json:"order_number"`

	// BatchNo 批次号。
	BatchNo string `json:"batch_no,omitempty"`
	// ReferenceNo 参考号。
	ReferenceNo string `json:"reference_no,omitempty"`

	// OrderFrom 订单来源（例如：手工订单/补发订单）。
	OrderFrom string `json:"order_from,omitempty"`

	// DeliverDeadline 发货时限。
	DeliverDeadline string `json:"deliver_deadline,omitempty"`

	// PurchaseTime 下单时间。
	PurchaseTime string `json:"purchase_time,omitempty"`
	// PaymentTime 付款时间。
	PaymentTime string `json:"payment_time,omitempty"`
	// PlatformPaymentTime 平台结算时间。
	PlatformPaymentTime string `json:"platform_payment_time,omitempty"`

	// SurfacePrintTime 面单打印时间（可能为 "0000-00-00 00:00:00"）。
	SurfacePrintTime string `json:"surface_print_time,omitempty"`
	// OrderPrintTime 订单打印时间（可能为 "0000-00-00 00:00:00"）。
	OrderPrintTime string `json:"order_print_time,omitempty"`

	// WaybillNo 运单号。
	WaybillNo string `json:"waybill_no,omitempty"`
	// TrackingNo 跟踪号/跟踪号（常用于承运商追踪）。
	TrackingNo string `json:"tracking_no"`
	// PackageNo 小包号（用于组包）。
	PackageNo string `json:"package_no,omitempty"`

	// SurfaceFileID 面单文件 ID（接口文档未描述，实际返回存在）。
	SurfaceFileID int64 `json:"surface_file_id,omitempty"`

	// TransferLogisticsCompanyCode 国内中转物流公司代码。
	TransferLogisticsCompanyCode string `json:"transfer_logistics_company_code,omitempty"`
	// TransferLogisticsCompanyID 国内中转物流公司 ID（文档为 string，且实际也可能为空）。
	TransferLogisticsCompanyID string `json:"transfer_logistics_company_id,omitempty"`
	// TransferTrackingNo 国内中转跟踪号。
	TransferTrackingNo string `json:"transfer_tracking_no,omitempty"`

	// Picker 拣货人。
	Picker string `json:"picker,omitempty"`
	// OrderType 订单类型：1 一单一件 / 2 多品多件 / 3 单品多件。
	OrderType int `json:"order_type,omitempty"`
	// Deliverer 发货人。
	Deliverer string `json:"deliverer,omitempty"`

	// Status 状态：1 物流下单 / 2 发货中 / 3 已发货 / 4 已删除。
	Status int `json:"status,omitempty"`
	// StatusName 状态名称。
	StatusName string `json:"status_name,omitempty"`

	// LogisticsStatus 物流下单状态：
	// 1 待导入 / 2 物流待下单 / 3 物流下单中 / 4 下单异常 / 5 下单完成 / 6 待海外仓下单 / 7 海外仓下单中
	// 11 待导入国内物流 / 41 物流取消中 / 42 物流取消异常 / 43 物流取消完成
	LogisticsStatus int `json:"logistics_status,omitempty"`
	// LogisticsStatusName 物流下单状态名称。
	LogisticsStatusName string `json:"logistics_status_name,omitempty"`
	// LogisticsMessage 物流下单消息。
	LogisticsMessage string `json:"logistics_message,omitempty"`

	// LogisticsProviderID 物流服务商 ID（文档为 int，实际可能存在 string/number 混用）。
	LogisticsProviderID StringOrNumber `json:"logistics_provider_id,omitempty"`
	// LogisticsProviderName 物流服务商名称。
	LogisticsProviderName string `json:"logistics_provider_name,omitempty"`
	// LogisticsTypeID 物流方式 ID（文档为 int，实际可能存在 string/number 混用）。
	LogisticsTypeID StringOrNumber `json:"logistics_type_id,omitempty"`
	// LogisticsTypeName 物流方式名称。
	LogisticsTypeName string `json:"logistics_type_name,omitempty"`

	// LogisticsEstimatedFreight 预估运费（字符串小数）。
	LogisticsEstimatedFreight string `json:"logistics_estimated_freight,omitempty"`
	// LogisticsEstimatedFreightCurrencyCode 预估运费币种。
	LogisticsEstimatedFreightCurrencyCode string `json:"logistics_estimated_freight_currency_code,omitempty"`
	// LogisticsFreight 物流运费（字符串小数）。
	LogisticsFreight string `json:"logistics_freight,omitempty"`
	// LogisticsFreightCurrencyCode 物流运费币种。
	LogisticsFreightCurrencyCode string `json:"logistics_freight_currency_code,omitempty"`

	// LogisticsSuccessTime 物流下单成功时间（接口文档未描述，实际返回存在）。
	LogisticsSuccessTime string `json:"logistics_success_time,omitempty"`
	// ActualCarrier 实际承运商编码（接口文档未描述，实际返回存在）。
	ActualCarrier string `json:"actual_carrier,omitempty"`

	// OrderCustomerServiceNotes 客服备注。
	OrderCustomerServiceNotes string `json:"order_customer_service_notes,omitempty"`
	// OrderBuyerNotes 买家留言。
	OrderBuyerNotes string `json:"order_buyer_notes,omitempty"`

	// IsCheck 是否验货：0 否 / 1 是。
	IsCheck int `json:"is_check,omitempty"`
	// IsWeigh 是否称重：0 否 / 1 是。
	IsWeigh int `json:"is_weigh,omitempty"`
	// IsSurfacePrint 面单是否打印：0 否 / 1 是。
	IsSurfacePrint int `json:"is_surface_print,omitempty"`
	// IsOrderPrint 订单是否打印：0 否 / 1 是。
	IsOrderPrint int `json:"is_order_print,omitempty"`

	// OrderOriginAmount 订单金额（字符串小数）。
	OrderOriginAmount string `json:"order_origin_amount,omitempty"`
	// OrderCurrencyCode 订单币种（例如 USD）。
	OrderCurrencyCode string `json:"order_currency_code,omitempty"`

	// PlatformOrderNo 平台单号（文档为 array）。
	PlatformOrderNo []string `json:"platform_order_no"`

	// DeliveredAt 出库时间（字符串格式以领星文档为准）。
	DeliveredAt string `json:"delivered_at"`
	// StockDeliveredAt 库存流水出库时间（一期用于回传 DSCO 的发货时间兜底口径）。
	StockDeliveredAt string `json:"stock_delivered_at"`

	// ProcessSN 加工单号。
	ProcessSN string `json:"process_sn,omitempty"`

	// CancelStatus 物流取消状态（接口文档未描述，实际返回存在）。
	CancelStatus int `json:"cancel_status,omitempty"`
	// CancelMessage 物流取消消息（接口文档未描述，实际返回存在）。
	CancelMessage string `json:"cancel_message,omitempty"`

	// DeliveryStatus 派送/妥投状态（接口文档未描述，实际返回存在）。
	DeliveryStatus int `json:"delivery_status,omitempty"`
	// DeliveryMessage 派送/妥投消息（接口文档未描述，实际返回存在）。
	DeliveryMessage string `json:"delivery_message,omitempty"`

	// ReportStatus 申报/报告状态（接口文档未描述，实际返回存在）。
	ReportStatus int `json:"report_status,omitempty"`
	// ReportMessage 申报/报告消息（接口文档未描述，实际返回存在）。
	ReportMessage string `json:"report_message,omitempty"`

	// MarkLabelStatus 标记面单状态（接口文档未描述，实际返回存在）。
	MarkLabelStatus int `json:"mark_label_status,omitempty"`
	// MarkLabelFileID 标记面单文件 ID（接口文档未描述，实际返回存在）。
	MarkLabelFileID int64 `json:"mark_label_file_id,omitempty"`

	// PickIndex 拣货序号（接口文档未描述，实际返回存在）。
	PickIndex int `json:"pick_index,omitempty"`
	// SurfaceFileType 面单文件类型。
	SurfaceFileType string `json:"surface_file_type,omitempty"`

	// IsLockStorage 是否已锁定库存：0 否 / 1 是。
	IsLockStorage int `json:"is_lock_storage,omitempty"`
	// IsAdvanceDelivery 是否预发货：0 否 / 1 是。
	IsAdvanceDelivery int `json:"is_advance_delivery,omitempty"`

	// ApportionStatus 费用分摊状态：1 未分摊 / 2 分摊失败 / 3 分摊成功（部分返回可能为 0，口径以领星为准）。
	ApportionStatus int `json:"apportion_status,omitempty"`
	// ApportionMessage 费用分摊消息。
	ApportionMessage string `json:"apportion_message,omitempty"`

	// RemarkAttachment 客服备注附件 JSON（字符串，示例为 "[]"）。
	RemarkAttachment string `json:"remark_attachment,omitempty"`

	// AutoAllocateStatus 自动分配状态（接口文档未描述，实际返回存在）。
	AutoAllocateStatus int `json:"auto_allocate_status,omitempty"`

	// SplitNum 拆分次数/拆分数量（接口文档未描述，实际返回存在）。
	SplitNum int `json:"split_num,omitempty"`

	// AutoComplete 自动完成标记（接口文档未描述，实际返回存在）。
	AutoComplete int `json:"auto_complete,omitempty"`

	// NeedInvoice 是否需要发票（接口文档未描述，实际返回存在）。
	NeedInvoice int `json:"need_invoice,omitempty"`
	// InvoiceStatus 发票状态（接口文档未描述，实际返回存在）。
	InvoiceStatus int `json:"invoice_status,omitempty"`

	// FirstMileStatus 头程状态（接口文档未描述，实际返回存在）。
	FirstMileStatus int `json:"first_mile_status,omitempty"`

	// DocumentsFileID 单证文件 ID（接口文档未描述，实际返回存在）。
	DocumentsFileID int64 `json:"documents_file_id,omitempty"`

	// TagNames 标签列表。
	TagNames []WmsOrderTagName `json:"tag_names,omitempty"`

	// OmsAttachments OMS 附加信息（接口文档未描述，实际返回存在）。
	OmsAttachments *WmsOrderOmsAttachments `json:"omsAttachments,omitempty"`

	// NoShippingProductList 未发货商品列表（接口文档未描述，结构不稳定，使用 RawMessage 保留原始数据）。
	NoShippingProductList []json.RawMessage `json:"noShippingProductList,omitempty"`

	// TrackRecord 物流轨迹（接口文档未描述，结构不稳定，使用 RawMessage 保留原始数据）。
	TrackRecord json.RawMessage `json:"track_record,omitempty"`

	// OrderTags 订单标签（接口文档未描述，结构不稳定，使用 RawMessage 保留原始数据）。
	OrderTags []json.RawMessage `json:"order_tags,omitempty"`

	// OrderSNs 订单号列表（接口文档未描述，结构不稳定，使用 RawMessage 保留原始数据）。
	OrderSNs []json.RawMessage `json:"order_sns,omitempty"`

	// SellerName 店铺名称。
	SellerName string `json:"seller_name,omitempty"`
	// SiteText 站点名称。
	SiteText string `json:"site_text,omitempty"`

	// ProductInfo 商品行列表。
	ProductInfo []WmsOrderProduct `json:"product_info"`

	// TargetCountry 收货国家。
	TargetCountry string `json:"target_country,omitempty"`

	// SurfaceFile 面单文件。
	SurfaceFile WmsOrderSurfaceFile `json:"surface_file,omitempty"`

	// ConsigneeFullAddress 收件地址（拼接后的完整地址）。
	ConsigneeFullAddress string `json:"consignee_full_address,omitempty"`

	// PlatformName 平台名称（例如 Amazon/Custom）。
	PlatformName string `json:"platform_name,omitempty"`

	// PackageDeliveredData 包裹出库信息（结构不稳定，使用 RawMessage 保留原始数据）。
	PackageDeliveredData []json.RawMessage `json:"package_delivered_data,omitempty"`

	// WarehouseName 仓库名。
	WarehouseName string `json:"warehouse_name,omitempty"`

	// 包裹尺寸与重量（字符串小数）。
	PkgVolume         string `json:"pkg_volume,omitempty"`
	PkgLength         string `json:"pkg_length,omitempty"`
	PkgWidth          string `json:"pkg_width,omitempty"`
	PkgHeight         string `json:"pkg_height,omitempty"`
	PkgWeight         string `json:"pkg_weight,omitempty"`
	PkgRealWeight     string `json:"pkg_real_weight,omitempty"`
	PkgFeeWeight      string `json:"pkg_fee_weight,omitempty"`
	PkgWeightUnit     string `json:"pkg_weight_unit,omitempty"`
	PkgRealWeightUnit string `json:"pkg_real_weight_unit,omitempty"`
	PkgFeeWeightUnit  string `json:"pkg_fee_weight_unit,omitempty"`
	PkgSizeUnit       string `json:"pkg_size_unit,omitempty"`

	// Consignee 收件人。
	Consignee string `json:"consignee,omitempty"`
	// ConsigneePhone 收件人电话。
	ConsigneePhone string `json:"consignee_phone,omitempty"`
	// ConsigneePostcode 收件人邮编。
	ConsigneePostcode string `json:"consignee_postcode,omitempty"`
	// ConsigneeAddress 收件人地址。
	ConsigneeAddress string `json:"consignee_address,omitempty"`

	// RecipientTaxNo 收件人税号。
	RecipientTaxNo string `json:"recipient_tax_no,omitempty"`
	// SenderTaxNo 发件人税号。
	SenderTaxNo string `json:"sender_tax_no,omitempty"`

	// CreateAt 创建时间。
	CreateAt string `json:"create_at,omitempty"`
	// UpdateAt 变更时间。
	UpdateAt string `json:"update_at,omitempty"`
}

// WmsOrderTagName 表示标签信息。
type WmsOrderTagName struct {
	// TagName 标签名称。
	TagName string `json:"tag_name,omitempty"`
	// TagTypeID 标签类型 ID（示例：4-3）。
	TagTypeID string `json:"tag_type_id,omitempty"`
	// TagColor 标签颜色（HEX，不含 #）。
	TagColor string `json:"tag_color,omitempty"`
	// TagType 标签类型（示例：4）。
	TagType int `json:"tag_type,omitempty"`
}

// WmsOrderSurfaceFile 表示面单文件信息。
type WmsOrderSurfaceFile struct {
	// URI 面单文件链接。
	URI string `json:"uri,omitempty"`
	// Ext 文件后缀。
	Ext string `json:"ext,omitempty"`
	// Size 文件尺寸（字符串，口径以领星返回为准）。
	Size string `json:"size,omitempty"`
}

// WmsOrderOmsAttachments 表示 OMS 附加信息（字段名大小写以实际返回为准）。
type WmsOrderOmsAttachments struct {
	OrderNumber string `json:"order_number,omitempty"`

	OrderRemark string `json:"order_remark,omitempty"`

	// RemarkAttachment 备注附件（数组结构不稳定，使用 RawMessage 保留原始数据）。
	RemarkAttachment []json.RawMessage `json:"remark_attachment,omitempty"`

	// OmsOrderItems OMS 订单行（数组结构不稳定，使用 RawMessage 保留原始数据）。
	OmsOrderItems []json.RawMessage `json:"oms_order_items,omitempty"`
}

// WmsOrderProduct 表示出库单商品行。
type WmsOrderProduct struct {
	// WodID 出库单明细 ID。
	WodID int64 `json:"wod_id,omitempty"`

	// GlobalItemNo 全局商品编号（接口文档未描述，实际返回存在）。
	GlobalItemNo string `json:"global_item_no,omitempty"`
	// PlatformOrderNo 平台单号（接口文档未描述，实际返回存在）。
	PlatformOrderNo string `json:"platform_order_no,omitempty"`

	// ProductID 商品 ID。
	ProductID int64 `json:"product_id,omitempty"`
	// SKU SKU。
	SKU string `json:"sku"`

	// SKUIdentifier SKU 标识（接口文档未描述，实际返回存在）。
	SKUIdentifier string `json:"sku_identifier,omitempty"`

	// SID 店铺/站点 ID（明细级，实际返回为字符串）。
	SID string `json:"sid,omitempty"`
	// StockSID 库存颗粒度店铺 ID（字符串）。
	StockSID string `json:"stock_sid,omitempty"`
	// StockSellerName 库存颗粒度店铺名称。
	StockSellerName string `json:"stock_seller_name,omitempty"`

	// BundleType 捆绑类型：0 普通商品 / 1 捆绑产品 / 2 捆绑产品拆分子产品。
	BundleType int `json:"bundle_type,omitempty"`
	// BundleWodID 捆绑产品 wod_id（只有拆分子产品才有）。
	BundleWodID int64 `json:"bundle_wod_id,omitempty"`

	// ProductName 商品名。
	ProductName string `json:"product_name,omitempty"`
	// BrandName 品牌名（接口文档未描述，实际返回存在）。
	BrandName string `json:"brand_name,omitempty"`
	// Model 型号（接口文档未描述，实际返回存在）。
	Model string `json:"model,omitempty"`

	// IsCombo 是否组合（接口文档未描述，实际返回存在）。
	IsCombo int `json:"is_combo,omitempty"`

	// Count 数量。
	Count int `json:"count"`

	// SellerSKU MSKU。
	SellerSKU string `json:"seller_sku,omitempty"`
	// Customization 商品备注。
	Customization string `json:"customization,omitempty"`

	// ItemUnitPrice 销售单价（字符串小数）。
	ItemUnitPrice string `json:"item_unit_price,omitempty"`
	// ItemTotalPrice 销售总价（字符串小数）。
	ItemTotalPrice string `json:"item_total_price,omitempty"`

	// RealWeightTotal 费用分摊-总实重（字符串小数）。
	RealWeightTotal string `json:"real_weight_total,omitempty"`
	// FeeWeightTotal 费用分摊-总计费重（字符串小数）。
	FeeWeightTotal string `json:"fee_weight_total,omitempty"`
	// VolumeWeightTotal 费用分摊-总体积重（字符串小数）。
	VolumeWeightTotal string `json:"volume_weight_total,omitempty"`

	// ApportionFreight 分摊运费（总计，字符串小数）。
	ApportionFreight string `json:"apportion_freight,omitempty"`
	// ApportionFreightSingle 分摊运费（单个，字符串小数）。
	ApportionFreightSingle string `json:"apportion_freight_single,omitempty"`

	// CNName 中文申报名。
	CNName string `json:"cn_name,omitempty"`
	// ENName 英文申报名。
	ENName string `json:"en_name,omitempty"`

	// UnitPrice 商品单价（字符串）。
	UnitPrice string `json:"unit_price,omitempty"`
	// CurrencyCode 币种（示例：$）。
	CurrencyCode string `json:"currency_code,omitempty"`
	// DeclaredCurrencyIcon 申报币种图标（接口文档字段名为 declared_currency_icon）。
	DeclaredCurrencyIcon string `json:"declared_currency_icon,omitempty"`

	// LogisticsFreightCurrencyCode 物流运费币种。
	LogisticsFreightCurrencyCode string `json:"logistics_freight_currency_code,omitempty"`

	// StockCost 库存成本（总计，字符串小数）。
	StockCost string `json:"stock_cost,omitempty"`

	// PurchaseFeeUnit 采购费单价（接口文档未描述，实际返回存在）。
	PurchaseFeeUnit string `json:"purchase_fee_unit,omitempty"`
	// HeadFeeUnit 头程费单价（接口文档未描述，实际返回存在）。
	HeadFeeUnit string `json:"head_fee_unit,omitempty"`
	// OtherFeeUnit 其他费用单价（接口文档未描述，实际返回存在）。
	OtherFeeUnit string `json:"other_fee_unit,omitempty"`

	// Material 材质（接口文档未描述，实际返回存在）。
	Material string `json:"material,omitempty"`

	// ThirdProductName 三方仓品名。
	ThirdProductName string `json:"third_product_name,omitempty"`
	// ThirdProductCode 三方仓 SKU。
	ThirdProductCode string `json:"third_product_code,omitempty"`

	// ProductTags 商品标签（接口文档未描述，结构不稳定，使用 RawMessage 保留原始数据）。
	ProductTags []json.RawMessage `json:"product_tags,omitempty"`

	// EarliestTimeBuyShippingLabel 最早可购买面单时间（接口文档未描述，实际返回存在）。
	EarliestTimeBuyShippingLabel string `json:"earliest_time_buy_shipping_label,omitempty"`
}

// WarehouseListsRequest 表示“查询仓库列表”的请求参数。
//
// API Path: /erp/sc/data/local_inventory/warehouse
type WarehouseListsRequest struct {
	// Type 仓库类型：1 本地仓（默认）；3 海外仓；4 亚马逊平台仓；6 AWD 仓。
	Type int `json:"type,omitempty"`
	// SubType 海外仓子类型：1 无 API 海外仓；2 有 API 海外仓（仅在 Type=3 时生效）。
	SubType int `json:"sub_type,omitempty"`
	// IsDelete 是否删除，多个使用英文逗号分隔：0 未删除（默认）；1 已删除。
	//
	// 说明：文档字段类型为 string，且支持多个值；因此这里使用字符串承载（例如 "0" 或 "0,1"）。
	IsDelete string `json:"is_delete,omitempty"`
	// Offset 分页偏移量（默认 0）。
	Offset int `json:"offset,omitempty"`
	// Length 分页长度（默认 1000）。
	Length int `json:"length,omitempty"`
}

// WarehouseInfo 表示仓库信息。
type WarehouseInfo struct {
	// WID 系统仓库 ID。
	WID int `json:"wid"`
	// Name 仓库名。
	Name string `json:"name"`
	// Type 仓库类型：1 本地仓；3 海外仓；4 平台仓；6 AWD 仓。
	Type int `json:"type"`
	// IsDelete 是否删除：0 未删除；1 已删除。
	IsDelete StringOrNumber `json:"is_delete"`

	// TCountryAreaName 第三方仓库国家/地区（字段：t_country_area_name）。
	TCountryAreaName string `json:"t_country_area_name"`
	// TStatus 状态：0 未启用；1 启用（字段：t_status）。
	//
	// 说明：该字段在真实返回中可能是 number / string / 空字符串 混用，使用 StringOrNumber 兼容。
	TStatus StringOrNumber `json:"t_status"`
	// TWarehouseCode 第三方仓库代码（字段：t_warehouse_code）。
	TWarehouseCode string `json:"t_warehouse_code"`
	// TWarehouseName 第三方仓库名（字段：t_warehouse_name）。
	TWarehouseName string `json:"t_warehouse_name"`
	// CountryCode 国家代码（字段：country_code）。
	CountryCode string `json:"country_code"`

	// WPID 服务商 ID（仅 type=3 且仓库为第三方海外仓时有值，字段：wp_id）。
	WPID int `json:"wp_id"`
	// WPName 系统服务商名称（字段：wp_name）。
	WPName string `json:"wp_name"`
}
