package dsco

// 本文件按 DSCO OpenAPI 文档补齐 Invoice 相关模型：
// `lingxingipass/golib/v2/sdk/dsco/docs/dsco-api-spec.yaml` -> schema: Invoice（用于 /invoice/batch/small 等接口）。

// InvoiceGetQuery 表示按 key/value 查询发票的参数（GET /invoice）。
type InvoiceGetQuery struct {
	// Key 指定查询使用的发票标识类型（例如 invoiceId / dscoInvoiceId / poNumber 等）。
	Key string `url:"key"`
	// Value 指定 Key 对应的值。
	Value string `url:"value"`
}

// GetInvoicesByIDResponse 表示 GET /invoice 的响应体。
type GetInvoicesByIDResponse struct {
	// Invoices 发票列表。
	Invoices []Invoice `json:"invoices"`
}

// Invoice 表示发票对象（用于创建/查询）。
//
// 说明：
// - OpenAPI 要求：invoiceId、totalAmount 为必填。
// - dscoOrderId / poNumber / supplierOrderNumber 三选一用于关联订单。
// - DSCO 的 Invoice 模型不包含顶层 trackingNumber；运单号位于 ship.trackingNumber 或 lineItems[].trackingNumber。
type Invoice struct {
	// InvoiceID 合作方发票 ID（OpenAPI required）。
	InvoiceID string `json:"invoiceId,omitempty"`

	// DscoOrderID DSCO 订单 ID（与 PoNumber/SupplierOrderNumber 三选一）。
	DscoOrderID string `json:"dscoOrderId,omitempty"`
	// PoNumber 订单采购单号（PO）（与 DscoOrderID/SupplierOrderNumber 三选一）。
	PoNumber string `json:"poNumber,omitempty"`
	// SupplierOrderNumber 供货商侧订单号（与 DscoOrderID/PoNumber 三选一）。
	SupplierOrderNumber string `json:"supplierOrderNumber,omitempty"`

	// InvoiceDate 发票日期时间（RFC3339 / date-time）。
	InvoiceDate string `json:"invoiceDate,omitempty"`
	// CurrencyCode 币种（ISO 4217）。
	CurrencyCode string `json:"currencyCode,omitempty"`
	// TotalAmount 发票总金额（OpenAPI required）。
	TotalAmount float64 `json:"totalAmount,omitempty"`

	// LineItemsSubtotal 行项目小计（可选）。
	LineItemsSubtotal *float64 `json:"lineItemsSubtotal,omitempty"`
	// SubtotalExcludingLineItems 不包含行项目的小计（可选）。
	SubtotalExcludingLineItems *float64 `json:"subtotalExcludingLineItems,omitempty"`
	// NumberOfLineItems 行项目数量（可选）。
	NumberOfLineItems *int `json:"numberOfLineItems,omitempty"`
	// HandlingAmount 处理费（可选）。
	HandlingAmount *float64 `json:"handlingAmount,omitempty"`
	// FreightAmount 运费（可选）。
	FreightAmount *float64 `json:"freightAmount,omitempty"`
	// SalesTaxAmount 销售税总额（可选）。
	SalesTaxAmount *float64 `json:"salesTaxAmount,omitempty"`

	// BuyerID 买方 ID（可选）。
	BuyerID string `json:"buyerId,omitempty"`
	// SellerID 卖方 ID（可选）。
	SellerID string `json:"sellerId,omitempty"`
	// SellerInvoiceNumber 卖方发票号（可选）。
	SellerInvoiceNumber string `json:"sellerInvoiceNumber,omitempty"`
	// ConsumerOrderNumber 面向消费者的订单号（可选）。
	ConsumerOrderNumber string `json:"consumerOrderNumber,omitempty"`
	// ExternalBatchID 外部系统批次 ID（可选）。
	ExternalBatchID string `json:"externalBatchId,omitempty"`

	// ExpectedOrderTotalAmount 预期订单总金额（可选）。
	ExpectedOrderTotalAmount *float64 `json:"expectedOrderTotalAmount,omitempty"`
	// ExpectedOrderTotalDifference 预期订单总金额差异（可选）。
	ExpectedOrderTotalDifference *float64 `json:"expectedOrderTotalDifference,omitempty"`

	// Charges 发票费用明细（例如某项费用标题+金额）。
	Charges []InvoiceCharge `json:"charges,omitempty"`
	// Terms 账期条款信息。
	Terms *InvoiceTerms `json:"terms,omitempty"`

	// ShipFrom 发票发货方地址信息。
	ShipFrom *InvoiceShipFromTo `json:"shipFrom,omitempty"`
	// ShipTo 发票收货方地址信息。
	ShipTo *InvoiceShipFromTo `json:"shipTo,omitempty"`
	// Buyer 买方地址信息（注意：字段名 buyer 与 BuyerID 不同）。
	Buyer *InvoiceShipFromTo `json:"buyer,omitempty"`
	// RemitTo 汇款地址信息。
	RemitTo *InvoiceShipFromTo `json:"remitTo,omitempty"`
	// Seller 卖方地址信息（注意：字段名 seller 与 SellerID 不同）。
	Seller *InvoiceShipFromTo `json:"seller,omitempty"`
	// Ship 发货信息（含 trackingNumber/carrier/method 等）。
	Ship *InvoiceShipInfo `json:"ship,omitempty"`

	// Credits 发票抵扣/贷项信息。
	Credits []InvoiceCredit `json:"credits,omitempty"`
	// Taxes 发票层级税信息。
	Taxes []InvoiceTax `json:"taxes,omitempty"`

	// LineItems 发票行项目列表。
	LineItems []InvoiceLineItem `json:"lineItems,omitempty"`

	// ---- 以下为 DSCO 侧回填字段（readOnly）或扩展字段（用于查询时可能出现）----

	// DscoInvoiceID DSCO 侧发票唯一 ID（readOnly）。
	DscoInvoiceID string `json:"dscoInvoiceId,omitempty"`
	// DscoExpectedOrderTotalAmount DSCO 侧计算的预期订单总金额（readOnly）。
	DscoExpectedOrderTotalAmount *float64 `json:"dscoExpectedOrderTotalAmount,omitempty"`
	// DscoExpectedOrderTotalDifference DSCO 侧计算的预期订单总金额差异（readOnly）。
	DscoExpectedOrderTotalDifference *float64 `json:"dscoExpectedOrderTotalDifference,omitempty"`
	// DscoSupplierID DSCO 侧供应商 ID（readOnly）。
	DscoSupplierID string `json:"dscoSupplierId,omitempty"`
	// DscoRetailerID DSCO 侧零售商 ID（readOnly）。
	DscoRetailerID string `json:"dscoRetailerId,omitempty"`
	// DscoTradingPartnerID DSCO 侧交易伙伴 ID（readOnly）。
	DscoTradingPartnerID string `json:"dscoTradingPartnerId,omitempty"`
	// DscoTradingPartnerName 交易伙伴名称（可选）。
	DscoTradingPartnerName string `json:"dscoTradingPartnerName,omitempty"`
	// DscoTradingPartnerParentID 交易伙伴 ParentID（可选）。
	DscoTradingPartnerParentID string `json:"dscoTradingPartnerParentId,omitempty"`

	// RequestedWarehouseCode 请求的仓库编码（readOnly）。
	RequestedWarehouseCode string `json:"requestedWarehouseCode,omitempty"`
	// RequestedWarehouseRetailerCode 请求的仓库零售商编码（readOnly）。
	RequestedWarehouseRetailerCode string `json:"requestedWarehouseRetailerCode,omitempty"`
	// RequestedWarehouseDscoID 请求的仓库 DSCO ID（readOnly）。
	RequestedWarehouseDscoID string `json:"requestedWarehouseDscoId,omitempty"`

	// ExportDate 导出时间（date-time, readOnly）。
	ExportDate string `json:"exportDate,omitempty"`
	// CreateDate 创建时间（date-time, readOnly）。
	CreateDate string `json:"createDate,omitempty"`
	// LastUpdate 最后更新时间（date-time, readOnly）。
	LastUpdate string `json:"lastUpdate,omitempty"`

	// OriginID Origin ID（可选）。
	OriginID *int `json:"originId,omitempty"`
	// FlsaComplianceFlag 是否符合 FLSA（Fair Labor Standards Act）合规要求（可选）。
	FlsaComplianceFlag *bool `json:"flsaComplianceFlag,omitempty"`
	// OrderType 订单类型（Dropship / Marketplace / Wholesale / Mixed）（可选）。
	OrderType string `json:"orderType,omitempty"`

	// AmountOfSalesTaxCollected 已收取的销售税金额（可选；字段在文档中也存在于 InvoiceLineItem）。
	AmountOfSalesTaxCollected *float64 `json:"amountOfSalesTaxCollected,omitempty"`
	// TaxableAmountOfSale 可计税销售额（可选；字段在文档中也存在于 InvoiceLineItem）。
	TaxableAmountOfSale *float64 `json:"taxableAmountOfSale,omitempty"`
}

// InvoiceCharge 表示发票费用明细（Invoice.charges[]）。
type InvoiceCharge struct {
	// Title 费用标题（OpenAPI required）。
	Title string `json:"title,omitempty"`
	// Amount 费用金额（OpenAPI required）。
	Amount float64 `json:"amount,omitempty"`
}

// InvoiceTerms 表示账期条款（Invoice.terms）。
type InvoiceTerms struct {
	// Type 条款类型（可选）。
	Type string `json:"type,omitempty"`
	// BasisDate 账期基准日期（date-time，可选）。
	BasisDate string `json:"basisDate,omitempty"`
	// DiscountPercent 折扣百分比（可选）。
	DiscountPercent *float64 `json:"discountPercent,omitempty"`
	// DiscountDueDate 折扣到期日（date-time，可选）。
	DiscountDueDate string `json:"discountDueDate,omitempty"`
	// DiscountDaysDue 折扣到期天数（可选）。
	DiscountDaysDue *int `json:"discountDaysDue,omitempty"`
	// NetDueDate 净到期日（date-time，可选）。
	NetDueDate string `json:"netDueDate,omitempty"`
	// NetDays 净到期天数（可选）。
	NetDays *int `json:"netDays,omitempty"`
	// DiscountAmount 折扣金额（可选）。
	DiscountAmount *float64 `json:"discountAmount,omitempty"`
	// DayOfMonth 月内日（可选）。
	DayOfMonth *int `json:"dayOfMonth,omitempty"`
	// TotalAmountSubjectToDiscount 可折扣的总金额（可选）。
	TotalAmountSubjectToDiscount *float64 `json:"totalAmountSubjectToDiscount,omitempty"`
}

// InvoiceShipFromTo 表示地址信息（Invoice.shipFrom/shipTo/buyer/remitTo/seller 等共用）。
//
// 注意：OpenAPI 描述中 City/Country/Postal 必填，Address1 或 Address 亦需提供；此处仅结构建模，不在 SDK 层强校验。
type InvoiceShipFromTo struct {
	// FirstName 名（deprecated，建议使用 Name）。
	FirstName string `json:"firstName,omitempty"`
	// LastName 姓（deprecated，建议使用 Name）。
	LastName string `json:"lastName,omitempty"`
	// Company 公司名。
	Company string `json:"company,omitempty"`

	// Address1 地址1（deprecated，建议使用 Address）。
	Address1 string `json:"address1,omitempty"`
	// Address2 地址2（deprecated，建议使用 Address）。
	Address2 string `json:"address2,omitempty"`
	// Address 地址行数组。
	Address []string `json:"address,omitempty"`

	// Country 国家（ISO 国家码或国家名，按 DSCO 约定）。
	Country string `json:"country,omitempty"`
	// Region 州/省（可选）。
	Region string `json:"region,omitempty"`
	// Postal 邮编。
	Postal string `json:"postal,omitempty"`
	// City 城市。
	City string `json:"city,omitempty"`
	// LocationCode 位置编码（可选）。
	LocationCode string `json:"locationCode,omitempty"`

	// Attention 收件注意/抬头（可选）。
	Attention string `json:"attention,omitempty"`
	// Email 邮箱（可选）。
	Email string `json:"email,omitempty"`
	// Phone 电话（可选）。
	Phone string `json:"phone,omitempty"`
	// StoreNumber 门店号（可选）。
	StoreNumber string `json:"storeNumber,omitempty"`
	// Name 姓名（推荐字段）。
	Name string `json:"name,omitempty"`
	// CustomerNumber 客户号（可选）。
	CustomerNumber string `json:"customerNumber,omitempty"`

	// AddressType 地址类型（Residential / Commercial）（可选）。
	AddressType string `json:"addressType,omitempty"`
	// TaxExemptNumber 免税号（可选）。
	TaxExemptNumber string `json:"taxExemptNumber,omitempty"`
	// TaxRegistrationNumber 税务登记号（可选）。
	TaxRegistrationNumber string `json:"taxRegistrationNumber,omitempty"`
}

// InvoiceShipInfo 表示发货信息（Invoice.ship）。
//
// 注意：trackingNumber 在此处是允许字段；顶层 Invoice 不接受 trackingNumber。
type InvoiceShipInfo struct {
	// NumberOfUnitsShipped 已发货件数（可选）。
	NumberOfUnitsShipped *int `json:"numberOfUnitsShipped,omitempty"`
	// UnitOfMeasure 单位（可选）。
	UnitOfMeasure string `json:"unitOfMeasure,omitempty"`
	// UnitOfMeasurementCode 单位码（可选；部分枚举在文档中已标记 deprecated）。
	UnitOfMeasurementCode string `json:"unitOfMeasurementCode,omitempty"`
	// Weight 重量（可选）。
	Weight *float64 `json:"weight,omitempty"`
	// WeightUnits 重量单位（lb/oz/g/kg）（可选）。
	WeightUnits string `json:"weightUnits,omitempty"`
	// Date 发货日期（date-time，可选）。
	Date string `json:"date,omitempty"`
	// TrackingNumber 运单号（可选）。
	TrackingNumber string `json:"trackingNumber,omitempty"`
	// Carrier 承运商（可选）。
	Carrier string `json:"carrier,omitempty"`
	// Method 物流方式（可选）。
	Method string `json:"method,omitempty"`
	// ServiceLevelCode 服务级别码（可选）。
	ServiceLevelCode string `json:"serviceLevelCode,omitempty"`

	// TransportationMethodCode 运输方式码（可选，枚举见 OpenAPI）。
	TransportationMethodCode string `json:"transportationMethodCode,omitempty"`
	// ReferenceNumberQualifier 引用号限定符（可选，枚举见 OpenAPI）。
	ReferenceNumberQualifier string `json:"referenceNumberQualifier,omitempty"`

	// DscoPackageID DSCO 包裹 ID（readOnly，可选）。
	DscoPackageID *int `json:"dscoPackageId,omitempty"`
	// WarehouseCode 仓库编码（可选）。
	WarehouseCode string `json:"warehouseCode,omitempty"`
	// WarehouseRetailerCode 仓库零售商编码（可选）。
	WarehouseRetailerCode string `json:"warehouseRetailerCode,omitempty"`
	// WarehouseDscoID 仓库 DSCO ID（可选）。
	WarehouseDscoID string `json:"warehouseDscoId,omitempty"`
	// SSCCBarcode SSCC 条码（可选）。
	SSCCBarcode string `json:"ssccBarcode,omitempty"`
	// CarrierManifestID 承运商 ManifestID（可选，LTL 用）。
	CarrierManifestID string `json:"carrierManifestId,omitempty"`
	// StoreNumber 门店号（可选）。
	StoreNumber string `json:"storeNumber,omitempty"`
}

// InvoiceCredit 表示发票抵扣/贷项（Invoice.credits[]）。
type InvoiceCredit struct {
	// Amount 抵扣金额（可选）。
	Amount *float64 `json:"amount,omitempty"`
	// Title 抵扣标题（可选）。
	Title string `json:"title,omitempty"`
}

// InvoiceTax 表示税信息（Invoice.taxes[] / InvoiceLineItem.taxes[]）。
type InvoiceTax struct {
	// Amount 税额（OpenAPI required）。
	Amount float64 `json:"amount,omitempty"`
	// Description 税描述（OpenAPI required）。
	Description string `json:"description,omitempty"`
	// JurisdictionQualifier 税辖区限定符（可选，枚举见 OpenAPI）。
	JurisdictionQualifier string `json:"jurisdictionQualifier,omitempty"`
	// Jurisdiction 税辖区（可选）。
	Jurisdiction string `json:"jurisdiction,omitempty"`
	// ExemptCode 免税码（可选，枚举见 OpenAPI）。
	ExemptCode string `json:"exemptCode,omitempty"`
	// Percentage 税率（可选）。
	Percentage *float64 `json:"percentage,omitempty"`
	// RegistrationNumber 税登记号（可选）。
	RegistrationNumber string `json:"registrationNumber,omitempty"`
	// TypeCode 税类型码（可选）。
	TypeCode string `json:"typeCode,omitempty"`
}

// InvoiceLineItem 表示发票行项目（OpenAPI required：quantity + unitPrice）。
//
// 注意：每个行项目需提供 dscoItemId/sku/partnerSku/upc/ean 中至少一个用于唯一标识商品。
type InvoiceLineItem struct {
	// DscoItemID DSCO 商品唯一 ID（用于匹配；可选，但推荐提供）。
	DscoItemID string `json:"dscoItemId,omitempty"`
	// SKU 供货商 SKU（可选）。
	SKU string `json:"sku,omitempty"`
	// UPC UPC（可选）。
	UPC string `json:"upc,omitempty"`
	// EAN EAN（可选）。
	EAN string `json:"ean,omitempty"`
	// MPN MPN（可选）。
	MPN string `json:"mpn,omitempty"`
	// ISBN ISBN（可选）。
	ISBN string `json:"isbn,omitempty"`
	// GTIN GTIN（可选）。
	GTIN string `json:"gtin,omitempty"`
	// PartnerSKU 合作方 SKU（可选，最大长度等约束见 OpenAPI）。
	PartnerSKU string `json:"partnerSku,omitempty"`

	// LineNumber 订单行号（可选，用于唯一定位行项目）。
	LineNumber int `json:"lineNumber,omitempty"`

	// Title 行项目标题（可选）。
	Title string `json:"title,omitempty"`
	// TitleI18n 标题多语言（LocalizedString：map[locale]text）（可选）。
	TitleI18n map[string]string `json:"titleI18n,omitempty"`

	// UnitPrice 单价（OpenAPI required）。
	UnitPrice float64 `json:"unitPrice"`
	// BasisOfUnitPrice 单价依据（可选）。
	BasisOfUnitPrice string `json:"basisOfUnitPrice,omitempty"`
	// Quantity 数量（OpenAPI required）。
	Quantity int `json:"quantity"`
	// UnitOfMeasure 单位（可选）。
	UnitOfMeasure string `json:"unitOfMeasure,omitempty"`

	// ExtendedAmount 扩展金额（可选）。
	ExtendedAmount *float64 `json:"extendedAmount,omitempty"`
	// HandlingAmount 行处理费（可选）。
	HandlingAmount *float64 `json:"handlingAmount,omitempty"`

	// ShipDate 发货日期（date-time，可选）。
	ShipDate string `json:"shipDate,omitempty"`
	// TrackingNumber 行级运单号（可选）。
	TrackingNumber string `json:"trackingNumber,omitempty"`
	// ShipAmount 行级运费（可选）。
	ShipAmount *float64 `json:"shipAmount,omitempty"`
	// ShipCarrier 行级承运商（可选）。
	ShipCarrier string `json:"shipCarrier,omitempty"`
	// ShipMethod 行级物流方式（可选）。
	ShipMethod string `json:"shipMethod,omitempty"`
	// ShippingServiceLevelCode 行级服务级别码（可选）。
	ShippingServiceLevelCode string `json:"shippingServiceLevelCode,omitempty"`

	// PromotionReference 促销引用（可选）。
	PromotionReference string `json:"promotionReference,omitempty"`
	// PromotionAmount 促销金额（可选）。
	PromotionAmount *float64 `json:"promotionAmount,omitempty"`

	// TaxAmount 税额（deprecated；文档提示可能会被替代）。
	TaxAmount *float64 `json:"taxAmount,omitempty"`
	// Subtotal 行小计（可选）。
	Subtotal *float64 `json:"subtotal,omitempty"`

	// ExpectedAmount 预期金额（deprecated；请使用 DscoExpectedAmount）。
	ExpectedAmount *float64 `json:"expectedAmount,omitempty"`
	// ExpectedDifference 预期差异（deprecated；请使用 DscoExpectedDifference）。
	ExpectedDifference *float64 `json:"expectedDifference,omitempty"`

	// DscoPackageID DSCO 包裹 ID（可选）。
	DscoPackageID *int `json:"dscoPackageId,omitempty"`
	// ShipWeight 发货重量（可选）。
	ShipWeight *float64 `json:"shipWeight,omitempty"`
	// ShipWeightUnits 发货重量单位（lb/oz/g/kg）（可选）。
	ShipWeightUnits string `json:"shipWeightUnits,omitempty"`

	// WarehouseCode 仓库编码（可选）。
	WarehouseCode string `json:"warehouseCode,omitempty"`
	// WarehouseRetailerCode 仓库零售商编码（可选）。
	WarehouseRetailerCode string `json:"warehouseRetailerCode,omitempty"`
	// WarehouseDscoID 仓库 DSCO ID（可选）。
	WarehouseDscoID string `json:"warehouseDscoId,omitempty"`
	// SSCCBarcode SSCC 条码（可选）。
	SSCCBarcode string `json:"ssccBarcode,omitempty"`

	// DscoExpectedAmount DSCO 侧预期金额（readOnly，可选）。
	DscoExpectedAmount *float64 `json:"dscoExpectedAmount,omitempty"`
	// DscoExpectedDifference DSCO 侧预期差异（readOnly，可选）。
	DscoExpectedDifference *float64 `json:"dscoExpectedDifference,omitempty"`

	// OriginalLineNumber 原始行号（readOnly，可选）。
	OriginalLineNumber *int `json:"originalLineNumber,omitempty"`
	// OriginalOrderQuantity 原始订单数量（readOnly，可选）。
	OriginalOrderQuantity *int `json:"originalOrderQuantity,omitempty"`

	// DepartmentID 部门 ID（readOnly，可选）。
	DepartmentID string `json:"departmentId,omitempty"`
	// DepartmentName 部门名称（readOnly，可选）。
	DepartmentName string `json:"departmentName,omitempty"`
	// MerchandisingAccountID 商品化账号 ID（readOnly，可选）。
	MerchandisingAccountID string `json:"merchandisingAccountId,omitempty"`
	// MerchandisingAccountName 商品化账号名称（readOnly，可选）。
	MerchandisingAccountName string `json:"merchandisingAccountName,omitempty"`

	// DscoOriginalOrderRetailerCreateDate 原始订单零售商创建时间（date-time, readOnly，可选）。
	DscoOriginalOrderRetailerCreateDate string `json:"dscoOriginalOrderRetailerCreateDate,omitempty"`

	// RequestedWarehouseCode 请求的仓库编码（readOnly，可选）。
	RequestedWarehouseCode string `json:"requestedWarehouseCode,omitempty"`
	// RequestedWarehouseRetailerCode 请求的仓库零售商编码（readOnly，可选）。
	RequestedWarehouseRetailerCode string `json:"requestedWarehouseRetailerCode,omitempty"`
	// RequestedWarehouseDscoID 请求的仓库 DSCO ID（readOnly，可选）。
	RequestedWarehouseDscoID string `json:"requestedWarehouseDscoId,omitempty"`

	// RetailerItemIDs 零售商 ItemId 列表（readOnly，可选）。
	RetailerItemIDs []string `json:"retailerItemIds,omitempty"`
	// RetailerLineID 零售商行 ID（readOnly，可选）。
	RetailerLineID string `json:"retailerLineId,omitempty"`

	// MSRP 建议零售价（可选）。
	MSRP *float64 `json:"msrp,omitempty"`

	// ShipReferenceNumberQualifier 运单引用号限定符（可选，枚举见 OpenAPI）。
	ShipReferenceNumberQualifier string `json:"shipReferenceNumberQualifier,omitempty"`
	// ShipTransportationMethodCode 行级运输方式码（可选，枚举见 OpenAPI）。
	ShipTransportationMethodCode string `json:"shipTransportationMethodCode,omitempty"`
	// ShipUnitOfMeasure 行级单位（可选）。
	ShipUnitOfMeasure string `json:"shipUnitOfMeasure,omitempty"`
	// ShipUnitOfMeasurementCode 行级单位码（可选）。
	ShipUnitOfMeasurementCode string `json:"shipUnitOfMeasurementCode,omitempty"`

	// CarrierManifestID 承运商 ManifestID（可选，LTL 用）。
	CarrierManifestID string `json:"carrierManifestId,omitempty"`
	// StoreNumber 门店号（可选）。
	StoreNumber string `json:"storeNumber,omitempty"`
	// SerialNumbers 序列号列表（可选）。
	SerialNumbers []string `json:"serialNumbers,omitempty"`

	// CarbCertifierID CARB 认证机构 ID（可选）。
	CarbCertifierID string `json:"carbCertifierId,omitempty"`
	// CarbFormaldehydeComplianceCode CARB 甲醛合规码（可选）。
	CarbFormaldehydeComplianceCode string `json:"carbFormaldehydeComplianceCode,omitempty"`

	// AmountOfSalesTaxCollected 行级已收取销售税金额（可选）。
	AmountOfSalesTaxCollected *float64 `json:"amountOfSalesTaxCollected,omitempty"`
	// TaxableAmountOfSale 行级可计税销售额（可选）。
	TaxableAmountOfSale *float64 `json:"taxableAmountOfSale,omitempty"`
	// Taxes 行级税信息（可选）。
	Taxes []InvoiceTax `json:"taxes,omitempty"`
}

// -------- Invoice Change Log（GET /invoice/log）--------

// InvoiceChangeLogQuery 表示 GET /invoice/log 的查询参数。
//
// 注意：
// - 若提供 scrollId，则其它参数会被 DSCO 忽略。
// - 若未提供 scrollId，则 requestId 与 startDate/endDate 二选一。
type InvoiceChangeLogQuery struct {
	// ScrollID 翻页游标（可选）。
	ScrollID string `url:"scrollId,omitempty"`

	// StartDate 起始时间（date-time，可选；未传 scrollId 时可用）。
	StartDate string `url:"startDate,omitempty"`
	// EndDate 结束时间（date-time，可选；未传 scrollId 时可用）。
	EndDate string `url:"endDate,omitempty"`

	// RequestID 异步批量接口返回的 requestId（可选；未传 scrollId 时可用）。
	RequestID string `url:"requestId,omitempty"`

	// Status 过滤条件：pending / success / failure / success_or_failure（可选）。
	Status string `url:"status,omitempty"`
}

// InvoiceChangeLogResponse 对应 OpenAPI 的 InvoiceChangeLogResponse。
type InvoiceChangeLogResponse struct {
	// ScrollID 翻页游标（可选）。
	ScrollID string `json:"scrollId,omitempty"`

	// Status 仅在按 requestId 查询时返回：PROCESSING / COMPLETED。
	Status string `json:"status,omitempty"`

	// Logs 变更日志列表。
	Logs []InvoiceChangeLog `json:"logs"`
}

// InvoiceChangeLog 对应 OpenAPI 的 InvoiceChangeLog。
type InvoiceChangeLog struct {
	// Payload 该条变更日志对应的发票对象快照。
	Payload Invoice `json:"payload"`

	// DateProcessed DSCO 处理时间（date-time）。
	DateProcessed string `json:"dateProcessed"`

	// Status 处理状态：pending / success / failure。
	Status string `json:"status"`

	// RequestMethod 请求方式：api / portal / file。
	RequestMethod string `json:"requestMethod"`

	// RequestMethodDetail 请求方式明细（string 或 object）。
	RequestMethodDetail any `json:"requestMethodDetail,omitempty"`

	// RequestID 异步批量调用的 requestId。
	RequestID string `json:"requestId,omitempty"`
	// ProcessID 处理流程 ID。
	ProcessID string `json:"processId,omitempty"`

	// Results 处理结果明细（包含错误码/描述等）。
	Results []APIResponseMessage `json:"results,omitempty"`
}
