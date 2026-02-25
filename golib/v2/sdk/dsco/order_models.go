package dsco

import "encoding/json"

//	真实订单示例：
//
// {"billTo":{},"buyer":null,"cancelRequested":false,"cancelRequestedCode":null,"consumerOrderNumber":"5093358016","crossDockFlag":false,"crossDockLocation":null,"customer":null,"deliveryMethod":null,"expediteFlag":false,"giftFlag":false,"giftWrapFlag":false,"invoiceTo":null,"marketplaceType":null,"message":"Vendor agrees to comply with Compliance Manual: http://bit.ly/ChewyVManual","orderType":"dropship","poNumber":"5J7PMN8NH5DSBCXL2M1Y","requestedShipCarrier":"FedEx","requestedShipMethod":"Home Delivery","requestedShippingServiceLevelCode":"FEHD","requestedShippingServiceLevelCodeUnmapped":"FEHD","requestedWarehouseCode":"YQN-CA","requestedWarehouseDscoId":"w692005423c098655499533","requestedWarehouseRetailerCode":"HAU5","retailerCreateDate":"2026-01-28T00:00:14+00:00","shipByDate":"2026-01-29T04:59:00+00:00","shipCarrier":"FedEx","shipMethod":"Home Delivery","shipWarehouseCode":"YQN-CA","shipWarehouseDscoId":"w692005423c098655499533","shipWarehouseRetailerCode":"HAU5","shipping":{"address1":"1151 POWER AVE","address":["1151 POWER AVE"],"city":"PAYETTE","country":"US","email":"placeholder_email@chewy.com","firstName":"Deborah Bolman","lastName":";","name":"Deborah Bolman ;","phone":"5306400246","postal":"83661-3392","region":"ID","state":"ID"},"signatureRequiredFlag":false,"testFlag":false,"dscoCreateDate":"2026-01-28T01:15:40+00:00","lineItems":[{"rejectedQuantity":0,"readyForPickup":false,"activity":[{"quantity":1,"action":"add","uuid":"4bb14655bbe14513bf0862215baa9798","activityDate":"2026-01-28T01:15:40+00:00","updateDate":"2026-01-28T01:15:40+00:00"}],"acceptedQuantity":0,"expectedCost":54,"lineNumber":1,"cancelRequested":false,"productGroup":"FF003CV-GY05","sku":"FF003CV-GY05","quantity":1,"expectedCostAdjustmentAllowed":false,"retailerItemIds":["3795734"],"giftFlag":false,"upc":"760314276156","cancelledQuantity":0,"bogoFlag":false,"dscoItemId":"1310486994"}],"dscoLastUpdate":"2026-01-28T01:16:25+00:00","dscoLifecycle":"created","dscoRetailerId":"1000011436","dscoShipLateDate":"2026-01-29T04:59:00+00:00","shippingServiceLevelCode":"FEHD","dscoStatus":"created","dscoOrderId":"1075104801","dscoSupplierId":"1000061639","dscoSupplierName":"HAUTICE LL - V2C","dscoTradingPartnerId":null,"dscoTradingPartnerName":null}

// Order 表示订单对象（按用例补齐“实用字段集”）。
//
// 说明：
//   - DSCO 的 Order 字段很多且会持续演进；这里优先补齐 boarding/同步闭环必需字段。
//   - 对于大块/多变的结构，仍使用 json.RawMessage 兜底，避免一次性全量建模带来维护成本。
type Order struct {
	// PoNumber 是零售商侧采购单号（PO），可作为部分接口的关联/查询口径。
	PoNumber string `json:"poNumber,omitempty"`

	// DscoOrderID 是 DSCO 侧订单唯一标识。
	DscoOrderID string `json:"dscoOrderId,omitempty"`

	// SupplierOrderNumber 供应商侧订单号（可为空）。
	SupplierOrderNumber *string `json:"supplierOrderNumber,omitempty"`

	// OrderType 用于标识订单类型（Dropship/Marketplace/Wholesale/Mixed 等，口径以 DSCO 返回为准）。
	OrderType *string `json:"orderType,omitempty"`

	// DeliveryMethod 配送方式/交付方式（in_store_pickup/ship_to_customer 等，口径以 DSCO 返回为准）。
	DeliveryMethod *string `json:"deliveryMethod,omitempty"`

	// MarketplaceType 市场/渠道类型（部分账号会返回；在 OpenAPI 中可能未强约束）。
	MarketplaceType *string `json:"marketplaceType,omitempty"`

	// CancelAfterDate：可取消截止时间（RFC3339）。
	CancelAfterDate *string `json:"cancelAfterDate,omitempty"`
	// CancelRequested 表示零售商是否请求取消订单（只读，可能为空）。
	CancelRequested *bool `json:"cancelRequested,omitempty"`
	// CancelRequestedCode 表示零售商请求取消时携带的原因码（只读，可能为空）。
	CancelRequestedCode *string `json:"cancelRequestedCode,omitempty"`

	// DscoLifecycle 订单生命周期（received/created/acknowledged/completed）。
	DscoLifecycle string `json:"dscoLifecycle,omitempty"`

	// DscoStatus 旧状态字段（created/shipment_pending/shipped/cancelled），文档标记 deprecated。
	DscoStatus string `json:"dscoStatus,omitempty"`

	// 关键时间字段（RFC3339，可为空）。
	ShipByDate        *string `json:"shipByDate,omitempty"`
	AcknowledgeByDate *string `json:"acknowledgeByDate,omitempty"`
	InvoiceByDate     *string `json:"invoiceByDate,omitempty"`

	// DSCO 延迟/超时阈值（RFC3339，可为空）。
	DscoShipLateDate        *string `json:"dscoShipLateDate,omitempty"`
	DscoAcknowledgeLateDate *string `json:"dscoAcknowledgeLateDate,omitempty"`
	DscoCancelLateDate      *string `json:"dscoCancelLateDate,omitempty"`
	DscoInvoiceLateDate     *string `json:"dscoInvoiceLateDate,omitempty"`

	// 期望/要求交付时间（RFC3339，可为空）。
	ExpectedDeliveryDate *string `json:"expectedDeliveryDate,omitempty"`
	RequiredDeliveryDate *string `json:"requiredDeliveryDate,omitempty"`

	// 期望的 CrossDock 送达时间（RFC3339，可为空）。
	ExpectedCrossDockDeliveryDate *string `json:"expectedCrossDockDeliveryDate,omitempty"`

	// SignatureRequiredFlag 是否要求签收。
	SignatureRequiredFlag *bool `json:"signatureRequiredFlag,omitempty"`

	// ExpediteFlag 是否加急。
	ExpediteFlag *bool `json:"expediteFlag,omitempty"`

	// ShipInstructions 发货备注/指示。
	ShipInstructions *string `json:"shipInstructions,omitempty"`

	// GiftWrapFlag 是否礼品包装。
	GiftWrapFlag *bool `json:"giftWrapFlag,omitempty"`
	// GiftWrapMessage 礼品包装说明。
	GiftWrapMessage *string `json:"giftWrapMessage,omitempty"`

	// LineItems 是订单行项目列表。
	LineItems []OrderLineItem `json:"lineItems,omitempty"`

	// Shipping 收件/收货信息（boarding/履约必用）。
	Shipping *OrderShipping `json:"shipping,omitempty"`

	// BillTo 账单/发票抬头信息（可为空）。
	BillTo *OrderBillTo `json:"billTo,omitempty"`

	// CrossDockFlag 是否 CrossDock 订单（可为空）。
	CrossDockFlag *bool `json:"crossDockFlag,omitempty"`

	// CrossDockLocation CrossDock 地址（可为空，口径以 DSCO 返回为准）。
	CrossDockLocation *OrderCrossDockLocation `json:"crossDockLocation,omitempty"`

	// 配送方式（部分场景 shipCarrier/shipMethod 与 shippingServiceLevelCode 二选一）。
	ShipCarrier              *string `json:"shipCarrier,omitempty"`
	ShipMethod               *string `json:"shipMethod,omitempty"`
	ShippingServiceLevelCode *string `json:"shippingServiceLevelCode,omitempty"`

	// 零售商请求的配送方式。
	RequestedShipCarrier                      *string `json:"requestedShipCarrier,omitempty"`
	RequestedShipMethod                       *string `json:"requestedShipMethod,omitempty"`
	RequestedShippingServiceLevelCode         *string `json:"requestedShippingServiceLevelCode,omitempty"`
	RequestedShippingServiceLevelCodeUnmapped *string `json:"requestedShippingServiceLevelCodeUnmapped,omitempty"`

	// DSCO 映射/标准化后的配送方式。
	DscoShipCarrier              *string `json:"dscoShipCarrier,omitempty"`
	DscoShipMethod               *string `json:"dscoShipMethod,omitempty"`
	DscoShippingServiceLevelCode *string `json:"dscoShippingServiceLevelCode,omitempty"`

	// Packages 包裹信息（发货后常见；boarding 后续步骤会用到 tracking）。
	Packages []OrderPackage `json:"packages,omitempty"`

	// ShippingAccountNumber 承运商账号/计费账号（可为空）。
	ShippingAccountNumber *string `json:"shippingAccountNumber,omitempty"`

	// PackingInstructions 打包/装箱指示（可为空）。
	PackingInstructions *string `json:"packingInstructions,omitempty"`

	// PackingSlip* 装箱单相关信息（可为空，口径以 DSCO 返回为准）。
	PackingSlipEmail    *string `json:"packingSlipEmail,omitempty"`
	PackingSlipPhone    *string `json:"packingSlipPhone,omitempty"`
	PackingSlipMessage  *string `json:"packingSlipMessage,omitempty"`
	PackingSlipTemplate *string `json:"packingSlipTemplate,omitempty"`

	// MarketingMessage 营销信息（可为空）。
	MarketingMessage *string `json:"marketingMessage,omitempty"`

	// BuyerMessage 订单备注/留言（可为空）。
	//
	// 兼容说明：历史代码中以 BuyerMessage 表达该字段，用于对接下游“买家留言/备注”口径；JSON 字段名仍为 message。
	BuyerMessage *string `json:"message,omitempty"`

	// NumberOfLineItems 行项目数量（可为空；通常等于 len(lineItems)）。
	NumberOfLineItems *int `json:"numberOfLineItems,omitempty"`

	// GiftFlag 是否礼品订单（可为空）。
	GiftFlag *bool `json:"giftFlag,omitempty"`

	// Gift* 礼品相关信息（可为空）。
	GiftMessage   *string `json:"giftMessage,omitempty"`
	GiftToName    *string `json:"giftToName,omitempty"`
	GiftFromName  *string `json:"giftFromName,omitempty"`
	GiftReceiptID *string `json:"giftReceiptId,omitempty"`

	// ShippingSurcharge 配送附加费（可为空）。
	ShippingSurcharge *float64 `json:"shippingSurcharge,omitempty"`

	// Taxes 订单级税费明细（可为空）。
	Taxes []OrderTax `json:"taxes,omitempty"`

	// AmountOfSalesTaxCollected 订单级“已收销售税”汇总（可为空）。
	AmountOfSalesTaxCollected *float64 `json:"amountOfSalesTaxCollected,omitempty"`

	// TaxableAmountOfSale 订单级“可税销售额”汇总（可为空）。
	TaxableAmountOfSale *float64 `json:"taxableAmountOfSale,omitempty"`

	// Coupons 优惠券信息（可为空）。
	Coupons []OrderCoupon `json:"coupons,omitempty"`

	// Payments 支付信息（可为空）。
	Payments []OrderPayment `json:"payments,omitempty"`

	// OrderTotalAmount 订单总金额（可为空）。
	OrderTotalAmount *float64 `json:"orderTotalAmount,omitempty"`

	// ExtendedExpectedCostTotal 扩展预估成本合计（可为空）。
	ExtendedExpectedCostTotal *float64 `json:"extendedExpectedCostTotal,omitempty"`

	// ReturnsMessage 退货相关备注（可为空）。
	ReturnsMessage *string `json:"returnsMessage,omitempty"`

	// ReturnsMessageI18n 多语言退货备注（可为空）。
	ReturnsMessageI18n map[string]string `json:"returnsMessageI18n,omitempty"`

	// RetailerCreateDate 零售商创建订单时间（RFC3339，可为空）。
	RetailerCreateDate *string `json:"retailerCreateDate,omitempty"`
	Channel            *string `json:"channel,omitempty"`

	// DscoRetailerID 零售商 ID（boarding 场景必用）。
	DscoRetailerID *string `json:"dscoRetailerId,omitempty"`

	// DscoSupplierId / DscoSupplierName 供应商信息（boarding 场景必用）。
	DscoSupplierId   *string `json:"dscoSupplierId,omitempty"`
	DscoSupplierName *string `json:"dscoSupplierName,omitempty"`

	// DscoTradingPartnerId / DscoTradingPartnerName / DscoTradingPartnerParentId
	// 表示零售商侧“交易伙伴”口径下的供应商标识（可为空，口径以 DSCO 返回为准）。
	DscoTradingPartnerId       *string `json:"dscoTradingPartnerId,omitempty"`
	DscoTradingPartnerName     *string `json:"dscoTradingPartnerName,omitempty"`
	DscoTradingPartnerParentId *string `json:"dscoTradingPartnerParentId,omitempty"`

	// DscoOnboardingRetailerId 表示 onboarding retailer 的 DSCO ID（可为空）。
	DscoOnboardingRetailerId *string `json:"dscoOnboardingRetailerId,omitempty"`

	// ConsumerOrderNumber 消费者订单号（部分零售商在后台展示的“Order #”口径）。
	ConsumerOrderNumber *string `json:"consumerOrderNumber,omitempty"`

	// ConsumerOrderDate 消费者订单时间（RFC3339，可为空）。
	ConsumerOrderDate *string `json:"consumerOrderDate,omitempty"`

	// ConsumerOrderCurrencyCode 消费者订单币种（可为空）。
	ConsumerOrderCurrencyCode *string `json:"consumerOrderCurrencyCode,omitempty"`

	// SecondaryConsumerOrderNumber 第二消费者订单号（可为空）。
	SecondaryConsumerOrderNumber *string `json:"secondaryConsumerOrderNumber,omitempty"`

	// ReceiptID 收据/凭证 ID（零售商侧标识，通常用于对账/追踪）。
	ReceiptID *string `json:"receiptId,omitempty"`

	// SecondaryReceiptID 第二收据/凭证 ID（可为空）。
	SecondaryReceiptID *string `json:"secondaryReceiptId,omitempty"`

	// InternalControlNumber 内部控制号（可为空，口径以 DSCO 返回为准）。
	InternalControlNumber *string `json:"internalControlNumber,omitempty"`

	// ReleaseNumber release number（可为空，口径以 DSCO 返回为准）。
	ReleaseNumber *string `json:"releaseNumber,omitempty"`

	// IssuerDivision 发行/所属事业部（可为空，口径以 DSCO 返回为准）。
	IssuerDivision *string `json:"issuerDivision,omitempty"`

	// AuthorizationForExpenseNumber 报销/费用授权号（可为空）。
	AuthorizationForExpenseNumber *string `json:"authorizationForExpenseNumber,omitempty"`

	// CustomerBalanceDue 客户应付余额（可为空）。
	CustomerBalanceDue *float64 `json:"customerBalanceDue,omitempty"`

	// ConsumerCreditAmountTotal 消费者侧“信用/抵扣金额合计”（可为空，口径以 DSCO 返回为准）。
	ConsumerCreditAmountTotal *float64 `json:"consumerCreditAmountTotal,omitempty"`

	// CustomerMembershipId 客户会员号（可为空）。
	CustomerMembershipId *string `json:"customerMembershipId,omitempty"`

	// CustomerBalanceDue/ConsumerBalanceDue 等字段存在多个口径；以 DSCO 返回为准。

	// DepartmentId/DepartmentName：订单级部门信息（可为空）。
	DepartmentId   *string `json:"departmentId,omitempty"`
	DepartmentName *string `json:"departmentName,omitempty"`

	// SalesRevenueCenter 销售收入中心（可为空）。
	SalesRevenueCenter *string `json:"salesRevenueCenter,omitempty"`
	// SalesAgent 销售人员/销售代表（可为空）。
	SalesAgent *string `json:"salesAgent,omitempty"`

	// PrimaryBatchNumber/SecondaryBatchNumber 批次号（可为空）。
	PrimaryBatchNumber   *string `json:"primaryBatchNumber,omitempty"`
	SecondaryBatchNumber *string `json:"secondaryBatchNumber,omitempty"`

	// BusinessRuleCode 业务规则代码（可为空）。
	BusinessRuleCode *string `json:"businessRuleCode,omitempty"`

	// RetailerAccountsPayableId 零售商应付账号 ID（可为空）。
	RetailerAccountsPayableId *string `json:"retailerAccountsPayableId,omitempty"`

	// ShipToStoreNumber 旧字段（deprecated）：Ship to Store 的门店号，推荐使用 shipping.storeNumber。
	ShipToStoreNumber *string `json:"shipToStoreNumber,omitempty"`

	// RequestedWarehouseCode 零售商“请求发货”的仓库代码（下单时指定/期望的仓库）。
	RequestedWarehouseCode *string `json:"requestedWarehouseCode,omitempty"`

	// RequestedWarehouseRetailerCode 零售商侧的仓库代码（可用于与零售商后台/对账口径对齐）。
	RequestedWarehouseRetailerCode *string `json:"requestedWarehouseRetailerCode,omitempty"`

	// RequestedWarehouseDscoID DSCO 侧仓库 ID（requestedWarehouse 的 DSCO 标识）。
	RequestedWarehouseDscoID *string `json:"requestedWarehouseDscoId,omitempty"`

	// DscoWarehouseCode DSCO 映射/标准化后的仓库代码（可能与 requested/ship 不同）。
	DscoWarehouseCode *string `json:"dscoWarehouseCode,omitempty"`

	// DscoWarehouseRetailerCode DSCO 映射后的零售商仓库代码。
	DscoWarehouseRetailerCode *string `json:"dscoWarehouseRetailerCode,omitempty"`

	// DscoWarehouseDscoID DSCO 映射后的仓库 DSCO ID。
	DscoWarehouseDscoID *string `json:"dscoWarehouseDscoId,omitempty"`

	// ShipWarehouseCode 实际发货仓库代码（供应商最终选择/使用的仓库）。
	ShipWarehouseCode *string `json:"shipWarehouseCode,omitempty"`

	// ShipWarehouseRetailerCode 实际发货仓库的零售商侧代码。
	ShipWarehouseRetailerCode *string `json:"shipWarehouseRetailerCode,omitempty"`

	// ShipWarehouseDscoID 实际发货仓库的 DSCO ID。
	ShipWarehouseDscoID *string `json:"shipWarehouseDscoId,omitempty"`

	// CurrencyCode 订单币种（ISO 4217），用于发票/金额口径。
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// TestFlag 是否测试订单。
	TestFlag *bool `json:"testFlag,omitempty"`

	// DSCO 时间戳（RFC3339，可为空）。
	DscoCreateDate *string `json:"dscoCreateDate,omitempty"`
	DscoLastUpdate *string `json:"dscoLastUpdate,omitempty"`

	// Buyer/Customer/InvoiceTo 结构较大且部分字段用 allOf/锚点复用：先保留原始 JSON 以便排查与后续扩展。
	Buyer     json.RawMessage `json:"buyer,omitempty"`
	Customer  json.RawMessage `json:"customer,omitempty"`
	InvoiceTo json.RawMessage `json:"invoiceTo,omitempty"`
}

// OrderLineItem 表示订单行项目（按用例补齐“实用字段集”）。
type OrderLineItem struct {
	// Quantity 是该行商品数量。
	Quantity int `json:"quantity"`

	// LineNumber 唯一标识订单行（部分下游策略/接口会要求发货时回传该字段）。
	LineNumber *int `json:"lineNumber,omitempty"`

	// ConsumerLineNumber 消费者侧行号（可为空）。
	ConsumerLineNumber *int `json:"consumerLineNumber,omitempty"`

	// DscoItemID DSCO 侧商品 ID（可能为空）。
	DscoItemID *string `json:"dscoItemId,omitempty"`

	// DscoSupplierId / DscoTradingPartnerId：用于指定该行由哪个供应商/交易伙伴履约（可为空）。
	DscoSupplierId       *string `json:"dscoSupplierId,omitempty"`
	DscoTradingPartnerId *string `json:"dscoTradingPartnerId,omitempty"`

	// WarehouseCode/WarehouseRetailerCode/WarehouseDscoID：行项目级仓库信息（可为空）。
	WarehouseCode         *string `json:"warehouseCode,omitempty"`
	WarehouseRetailerCode *string `json:"warehouseRetailerCode,omitempty"`
	WarehouseDscoID       *string `json:"warehouseDscoId,omitempty"`

	// ProductGroup 商品组/品类分组（可为空）。
	ProductGroup *string `json:"productGroup,omitempty"`

	// RetailerProductGroup 零售商侧商品组（可为空）。
	RetailerProductGroup *string `json:"retailerProductGroup,omitempty"`

	// SupplierProductGroup 供应商侧商品组（可为空）。
	SupplierProductGroup *string `json:"supplierProductGroup,omitempty"`

	// SupplierQuoteNumber 供应商报价单号（可为空）。
	SupplierQuoteNumber *string `json:"supplierQuoteNumber,omitempty"`

	// Title 商品标题（可为空）。
	Title *string `json:"title,omitempty"`

	// SupplierTitle 供应商侧商品标题（可为空）。
	SupplierTitle *string `json:"supplierTitle,omitempty"`

	// PackingSlipTitle 装箱单标题（可为空）。
	PackingSlipTitle *string `json:"packingSlipTitle,omitempty"`

	// PackingSlipSKU 装箱单 SKU 展示（可为空）。
	PackingSlipSKU *string `json:"packingSlipSku,omitempty"`

	// TitleI18n 多语言标题（可为空）。
	TitleI18n map[string]string `json:"titleI18n,omitempty"`

	// SKU / PartnerSKU 为商品标识（可能为空/为 null，取决于 DSCO 侧订单口径）。
	SKU        *string `json:"sku,omitempty"`
	PartnerSKU *string `json:"partnerSku,omitempty"`

	// UPC/EAN/GTIN 等标识（可为空）。
	UPC  *string `json:"upc,omitempty"`
	EAN  *string `json:"ean,omitempty"`
	GTIN *string `json:"gtin,omitempty"`

	// MPN/ISBN 等可选标识（可为空）。
	MPN  *string `json:"mpn,omitempty"`
	ISBN *string `json:"isbn,omitempty"`

	// Color/Size：颜色/尺码（可为空）。
	Color *string `json:"color,omitempty"`
	Size  *string `json:"size,omitempty"`

	// SupplierColor/SupplierSize：供应商侧颜色/尺码（可为空）。
	SupplierColor *string `json:"supplierColor,omitempty"`
	SupplierSize  *string `json:"supplierSize,omitempty"`

	// ProductSpecifications 商品规格说明（可为空）。
	ProductSpecifications *string `json:"productSpecifications,omitempty"`

	// Personalization/PersonalizationMap 个性化定制信息（可为空）。
	Personalization    *string           `json:"personalization,omitempty"`
	PersonalizationMap map[string]string `json:"personalizationMap,omitempty"`

	// ConsumerPrice 终端消费者价格（部分接口/场景可能为空）。
	ConsumerPrice           *float64 `json:"consumerPrice,omitempty"`
	ConsumerPriceWithTax    *float64 `json:"consumerPriceWithTax,omitempty"`
	ConsumerPriceWithoutTax *float64 `json:"consumerPriceWithoutTax,omitempty"`

	// ConsumerBalanceDue 消费者应付余额（可为空）。
	ConsumerBalanceDue *float64 `json:"consumerBalanceDue,omitempty"`

	// ConsumerCreditAmountTotal 消费者侧“信用/抵扣金额合计”（可为空）。
	ConsumerCreditAmountTotal *float64 `json:"consumerCreditAmountTotal,omitempty"`

	// RetailPrice 零售价格（当 consumerPrice 缺失时的兜底候选，口径以 DSCO API spec 为准）。
	RetailPrice *float64 `json:"retailPrice,omitempty"`

	// ExpectedCost 预估成本（作为 unit_price 的兜底候选）。
	ExpectedCost *float64 `json:"expectedCost,omitempty"`

	// ExpectedCostAdjustmentAllowed 是否允许预估成本调整（可为空）。
	ExpectedCostAdjustmentAllowed *bool `json:"expectedCostAdjustmentAllowed,omitempty"`

	// ExtendedExpectedCostTotal 扩展预估成本合计（可为空）。
	ExtendedExpectedCostTotal *float64 `json:"extendedExpectedCostTotal,omitempty"`

	// MAP/MSRP 建议零售价（可为空）。
	Map  *float64 `json:"map,omitempty"`
	Msrp *float64 `json:"msrp,omitempty"`

	// Commission* 佣金信息（可为空）。
	CommissionPercentage *float64 `json:"commissionPercentage,omitempty"`
	CommissionAmount     *float64 `json:"commissionAmount,omitempty"`
	CommissionLabel      *string  `json:"commissionLabel,omitempty"`

	// HandlingAmount 处理费/手续费（可为空）。
	HandlingAmount *float64 `json:"handlingAmount,omitempty"`

	// ShippingSurcharge 配送附加费（行项目级，可为空）。
	ShippingSurcharge *float64 `json:"shippingSurcharge,omitempty"`

	// Taxes 税费明细（行项目级，可为空）。
	Taxes []OrderTax `json:"taxes,omitempty"`

	// AmountOfSalesTaxCollected 行项目级“已收销售税”（可为空）。
	AmountOfSalesTaxCollected *float64 `json:"amountOfSalesTaxCollected,omitempty"`

	// TaxableAmountOfSale 行项目级“可税销售额”（可为空）。
	TaxableAmountOfSale *float64 `json:"taxableAmountOfSale,omitempty"`

	// UnitOfMeasure 包装单位（EA/PR/CA/IG/PL 等，口径以 DSCO 返回为准）。
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`

	// Weight/WeightUnits 重量及单位（可为空）。
	Weight      *float64 `json:"weight,omitempty"`
	WeightUnits *string  `json:"weightUnits,omitempty"`

	// DscoTags DSCO 标签（可为空）。
	DscoTags []string `json:"dscoTags,omitempty"`

	// RetailerItemIDs 零售商侧商品 ID 列表（可为空）。
	RetailerItemIDs []string `json:"retailerItemIds,omitempty"`

	// RetailerLineID 零售商侧行号（可为空）。
	RetailerLineID *string `json:"retailerLineId,omitempty"`

	// DepartmentId/DepartmentName 部门信息（可为空）。
	DepartmentId   *string `json:"departmentId,omitempty"`
	DepartmentName *string `json:"departmentName,omitempty"`

	// MerchandisingAccountId/MerchandisingAccountName 货品/陈列账号信息（可为空）。
	MerchandisingAccountId   *string `json:"merchandisingAccountId,omitempty"`
	MerchandisingAccountName *string `json:"merchandisingAccountName,omitempty"`

	// Message 行项目级备注（可为空）。
	Message *string `json:"message,omitempty"`

	// PackingInstructions 行项目级打包指示（可为空）。
	PackingInstructions *string `json:"packingInstructions,omitempty"`

	// ShipInstructions 行项目级发货指示（可为空）。
	ShipInstructions *string `json:"shipInstructions,omitempty"`

	// Gift* 行项目级礼品信息（可为空）。
	GiftFlag      *bool   `json:"giftFlag,omitempty"`
	GiftMessage   *string `json:"giftMessage,omitempty"`
	GiftToName    *string `json:"giftToName,omitempty"`
	GiftFromName  *string `json:"giftFromName,omitempty"`
	GiftReceiptID *string `json:"giftReceiptId,omitempty"`

	// ReceiptID/SecondaryReceiptID 凭证/收据标识（可为空）。
	ReceiptID          *string `json:"receiptId,omitempty"`
	SecondaryReceiptID *string `json:"secondaryReceiptId,omitempty"`

	// ReturnsMessage/ReturnsMessageI18n 退货备注（可为空）。
	ReturnsMessage     *string           `json:"returnsMessage,omitempty"`
	ReturnsMessageI18n map[string]string `json:"returnsMessageI18n,omitempty"`

	// BogoFlag/BogoInstructions 买一赠一相关标记（可为空）。
	BogoFlag         *bool   `json:"bogoFlag,omitempty"`
	BogoInstructions *string `json:"bogoInstructions,omitempty"`

	// CancelAfterDate/CancelCode/Cancelled* 取消相关字段（可为空）。
	CancelAfterDate *string `json:"cancelAfterDate,omitempty"`
	CancelCode      *string `json:"cancelCode,omitempty"`

	CancelledQuantity *int    `json:"cancelledQuantity,omitempty"`
	CancelledReason   *string `json:"cancelledReason,omitempty"`

	// CancelRequested/CancelRequestedCode 零售商请求取消（只读，可能为空）。
	CancelRequested     *bool   `json:"cancelRequested,omitempty"`
	CancelRequestedCode *string `json:"cancelRequestedCode,omitempty"`

	// Accepted*/Rejected*/Shipped*/Remaining* 数量与状态相关字段（可为空）。
	AcceptedQuantity *int    `json:"acceptedQuantity,omitempty"`
	AcceptedReason   *string `json:"acceptedReason,omitempty"`

	RejectedQuantity *int    `json:"rejectedQuantity,omitempty"`
	RejectedReason   *string `json:"rejectedReason,omitempty"`

	ShippedQuantity   *float64 `json:"shippedQuantity,omitempty"`
	RemainingQuantity *float64 `json:"remainingQuantity,omitempty"`

	// ReadyForPickup 是否可取货（可为空）。
	ReadyForPickup *bool `json:"readyForPickup,omitempty"`

	// Status/StatusReason 行项目状态（可为空，口径以 DSCO 返回为准）。
	Status       *string `json:"status,omitempty"`
	StatusReason *string `json:"statusReason,omitempty"`

	// DscoLifecycle 行项目生命周期（可为空，口径以 DSCO 返回为准）。
	DscoLifecycle *string `json:"dscoLifecycle,omitempty"`

	// 关键时间字段（RFC3339，可为空）。
	ShipByDate        *string `json:"shipByDate,omitempty"`
	AcknowledgeByDate *string `json:"acknowledgeByDate,omitempty"`
	InvoiceByDate     *string `json:"invoiceByDate,omitempty"`

	// RequestedShipDate 零售商请求的发货时间（RFC3339，可为空）。
	RequestedShipDate *string `json:"requestedShipDate,omitempty"`

	// EstimatedShipDate/EstimatedShipReason 预计发货时间及原因（RFC3339，可为空）。
	EstimatedShipDate   *string `json:"estimatedShipDate,omitempty"`
	EstimatedShipReason *string `json:"estimatedShipReason,omitempty"`

	// ExpectedDeliveryDate/RequiredDeliveryDate 期望/要求交付时间（RFC3339，可为空）。
	ExpectedDeliveryDate *string `json:"expectedDeliveryDate,omitempty"`
	RequiredDeliveryDate *string `json:"requiredDeliveryDate,omitempty"`

	// ExpectedCrossDockDeliveryDate 期望 CrossDock 送达时间（RFC3339，可为空）。
	ExpectedCrossDockDeliveryDate *string `json:"expectedCrossDockDeliveryDate,omitempty"`

	// CrossDockLocation 行项目级 CrossDock 地址（可为空）。
	CrossDockLocation *OrderCrossDockLocation `json:"crossDockLocation,omitempty"`

	// Activity 行项目变更记录（可为空）。
	Activity []OrderLineItemActivity `json:"activity,omitempty"`

	// UpdateDate 行项目更新时间（RFC3339，可为空）。
	UpdateDate *string `json:"updateDate,omitempty"`
}

// OrderLineItemActivity 表示订单行项目的变更活动（OpenAPI: OrderLineItemActivity）。
type OrderLineItemActivity struct {
	// Action 动作（add/update/delete 等，口径以 DSCO 返回为准）。
	Action string `json:"action,omitempty"`

	// Quantity 变更数量（可为空）。
	Quantity *int `json:"quantity,omitempty"`

	// UUID 变更记录的唯一标识。
	UUID string `json:"uuid,omitempty"`

	// ActivityDate 活动时间（RFC3339）。
	ActivityDate *string `json:"activityDate,omitempty"`

	// UpdateDate 变更更新时间（RFC3339，可为空）。
	UpdateDate *string `json:"updateDate,omitempty"`

	// TransactionDate 交易时间（RFC3339，可为空）。
	TransactionDate *string `json:"transactionDate,omitempty"`

	// TransactionId 交易 ID（可为空）。
	TransactionId *string `json:"transactionId,omitempty"`

	// FormerStatus 变更前状态（可为空）。
	FormerStatus *string `json:"formerStatus,omitempty"`

	// Reason 变更原因（口径以 DSCO 返回为准；部分场景可能为空）。
	Reason string `json:"reason,omitempty"`
}

// OrderShipping 表示订单的收件/收货信息（OpenAPI: OrderShipping）。
type OrderShipping struct {
	// Attention 收件人/联系人备注（Attn）。
	Attention *string `json:"attention,omitempty"`

	// FirstName/LastName：旧字段（deprecated），DSCO 可能仍会返回。
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`

	// Company 收件公司/机构名称。
	Company *string `json:"company,omitempty"`

	// Address1/Address2：旧字段（deprecated），推荐使用 Address 数组。
	Address1 *string `json:"address1,omitempty"`
	Address2 *string `json:"address2,omitempty"`

	// Address 地址行数组（通常为 1~n 行）。
	Address []string `json:"address,omitempty"`

	// City 城市。
	City string `json:"city,omitempty"`
	// Region 地区/行政区（部分国家使用该字段）。
	Region *string `json:"region,omitempty"`
	// State 州/省（部分国家使用该字段）。
	State *string `json:"state,omitempty"`
	// Postal 邮编。
	Postal string `json:"postal,omitempty"`

	// Country 国家（通常为国家缩写或名称，口径以 DSCO 返回为准）。
	Country *string `json:"country,omitempty"`
	// Phone 电话。
	Phone *string `json:"phone,omitempty"`
	// Email 邮箱。
	Email *string `json:"email,omitempty"`

	// Name 收件人姓名（推荐使用该字段替代 firstName/lastName）。
	Name *string `json:"name,omitempty"`

	// AddressType 地址类型（例如 residential/commercial 等，口径以 DSCO 返回为准）。
	AddressType *string `json:"addressType,omitempty"`
	// StoreNumber 门店号/店铺号（部分 Ship to Store 订单会提供）。
	StoreNumber *string `json:"storeNumber,omitempty"`

	// ShipNightPhone 夜间联系电话（可选）。
	ShipNightPhone *string `json:"shipNightPhone,omitempty"`

	// ShipCustomerNumber 客户编号/收货方编号（可选）。
	ShipCustomerNumber *string `json:"shipCustomerNumber,omitempty"`

	// ShipSiteType 站点类型（customer/store/installer/access-point）。
	ShipSiteType *string `json:"shipSiteType,omitempty"`

	// Carrier/Method：旧字段（deprecated），不建议在新对接中写入。
	Carrier *string `json:"carrier,omitempty"`
	Method  *string `json:"method,omitempty"`
}

// OrderBillTo 表示订单的账单信息（OpenAPI: OrderBillTo）。
type OrderBillTo struct {
	Attention *string `json:"attention,omitempty"`

	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Name      *string `json:"name,omitempty"`

	Company *string `json:"company,omitempty"`

	Address1 *string  `json:"address1,omitempty"`
	Address2 *string  `json:"address2,omitempty"`
	Address  []string `json:"address,omitempty"`

	City    *string `json:"city,omitempty"`
	Region  *string `json:"region,omitempty"`
	Postal  *string `json:"postal,omitempty"`
	Country *string `json:"country,omitempty"`

	Phone *string `json:"phone,omitempty"`
	Email *string `json:"email,omitempty"`

	AddressType *string `json:"addressType,omitempty"`
}

// OrderCrossDockLocation 表示 CrossDock 地址（OpenAPI: crossDockLocation，锚点 a36）。
//
// 注意：该结构在 OpenAPI 中未命名为独立 schema，但在多个字段复用；这里以“最小可用建模”补齐字段。
type OrderCrossDockLocation struct {
	Address []string `json:"address,omitempty"`

	City   *string `json:"city,omitempty"`
	Region *string `json:"region,omitempty"`
	Postal *string `json:"postal,omitempty"`
	County *string `json:"county,omitempty"`

	Country *string `json:"country,omitempty"`

	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`

	// NightPhone 对应 OpenAPI 的 nightPhone。
	NightPhone *string `json:"nightPhone,omitempty"`

	Company *string `json:"company,omitempty"`

	// CustomerNumber 对应 OpenAPI 的 customerNumber。
	CustomerNumber *string `json:"customerNumber,omitempty"`

	// AddressType 通常为 commercial/residential（口径以 DSCO 返回为准）。
	AddressType *string `json:"addressType,omitempty"`

	Attention *string `json:"attention,omitempty"`

	// LocationCode 对应 OpenAPI 的 locationCode。
	LocationCode *string `json:"locationCode,omitempty"`

	StoreNumber *string `json:"storeNumber,omitempty"`
}

// OrderTax 表示税费项（OpenAPI: OrderTax）。
type OrderTax struct {
	Percentage *float64 `json:"percentage,omitempty"`
	TypeCode   *string  `json:"typeCode,omitempty"`
	Amount     *float64 `json:"amount,omitempty"`

	JurisdictionQualifier *string `json:"jurisdictionQualifier,omitempty"`
	Jurisdiction          *string `json:"jurisdiction,omitempty"`
	ExemptCode            *string `json:"exemptCode,omitempty"`
	RegistrationNumber    *string `json:"registrationNumber,omitempty"`
	Description           *string `json:"description,omitempty"`
}

// OrderCoupon 表示订单优惠券信息（OpenAPI: OrderCoupon）。
type OrderCoupon struct {
	Amount     *float64 `json:"amount,omitempty"`
	Percentage *float64 `json:"percentage,omitempty"`
}

// OrderPayment 表示订单支付信息（OpenAPI: OrderPayment）。
type OrderPayment struct {
	CardType        string  `json:"cardType,omitempty"`
	CardLastFour    string  `json:"cardLastFour,omitempty"`
	Description     *string `json:"description,omitempty"`
	CardTypeDetails *string `json:"cardTypeDetails,omitempty"`
}

// OrderPackage 对应 OpenAPI 的 Package（orders.packages[]）。
type OrderPackage struct {
	// TrackingNumber 包裹级运单号。
	TrackingNumber string `json:"trackingNumber,omitempty"`
	// TrackingURL 运单查询链接。
	TrackingURL *string `json:"trackingUrl,omitempty"`

	// Untracked 是否为“无追踪/不提供运单号”的包裹。
	Untracked *bool `json:"untracked,omitempty"`

	// 交易信息（可为空）。
	TransactionDate *string `json:"transactionDate,omitempty"`
	TransactionID   *string `json:"transactionId,omitempty"`

	// CarrierManifestID 载货清单/交接单 ID（可为空）。
	CarrierManifestID *string `json:"carrierManifestId,omitempty"`

	// TransportationMethodCode/UnitOfMeasurementCode/SSCCBarcode 等字段用于物流对接（可为空）。
	TransportationMethodCode *string `json:"transportationMethodCode,omitempty"`
	UnitOfMeasurementCode    *string `json:"unitOfMeasurementCode,omitempty"`
	SSCCBarcode              *string `json:"ssccBarcode,omitempty"`

	// ReferenceNumberQualifier 参考号限定符（可为空）。
	ReferenceNumberQualifier *string `json:"referenceNumberQualifier,omitempty"`

	// ShipCost 运费（金额口径以 DSCO 返回为准）。
	ShipCost *float64 `json:"shipCost,omitempty"`
	// ShipDate 发货日期/时间（RFC3339）。
	ShipDate *string `json:"shipDate,omitempty"`

	// ShipCarrier 承运商名称（例如 FedEx/USPS）。
	ShipCarrier *string `json:"shipCarrier,omitempty"`
	// ShipMethod 运输方式名称（例如 Home Delivery/2 Day）。
	ShipMethod *string `json:"shipMethod,omitempty"`

	// ShippingServiceLevelCode 配送服务级别码（可为空）。
	ShippingServiceLevelCode *string `json:"shippingServiceLevelCode,omitempty"`

	// CurrencyCode 币种（可为空）。
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// WarehouseCode/WarehouseRetailerCode/WarehouseDscoID 包裹级仓库信息（可为空）。
	WarehouseCode         *string `json:"warehouseCode,omitempty"`
	WarehouseRetailerCode *string `json:"warehouseRetailerCode,omitempty"`
	WarehouseDscoID       *string `json:"warehouseDscoId,omitempty"`

	// DscoPackageID DSCO 侧包裹 ID（可为空）。
	DscoPackageID *int `json:"dscoPackageId,omitempty"`

	// DscoTradingPartnerParentId 交易伙伴父级 ID（可为空）。
	DscoTradingPartnerParentId *string `json:"dscoTradingPartnerParentId,omitempty"`

	// BalanceDue 应付余额（可为空）。
	BalanceDue *float64 `json:"balanceDue,omitempty"`

	// NumberOfLineItems 行项目数量（可为空）。
	NumberOfLineItems *int `json:"numberOfLineItems,omitempty"`

	// ShipWeight/ShipWeightUnits 包裹重量与单位（可为空）。
	ShipWeight      *float64 `json:"shipWeight,omitempty"`
	ShipWeightUnits *string  `json:"shipWeightUnits,omitempty"`

	// PackageShipFrom/PackageShipTo 包裹级发货/收货地址（可为空）。
	PackageShipFrom *PackageShipFrom `json:"packageShipFrom,omitempty"`
	PackageShipTo   *PackageShipTo   `json:"packageShipTo,omitempty"`

	// DSCO 实际履约信息（可为空）。
	DscoActualShipMethod               *string  `json:"dscoActualShipMethod,omitempty"`
	DscoActualShipCarrier              *string  `json:"dscoActualShipCarrier,omitempty"`
	DscoActualShippingServiceLevelCode *string  `json:"dscoActualShippingServiceLevelCode,omitempty"`
	DscoActualShipCost                 *float64 `json:"dscoActualShipCost,omitempty"`
	DscoActualDeliveryDate             *string  `json:"dscoActualDeliveryDate,omitempty"`
	DscoActualPickupDate               *string  `json:"dscoActualPickupDate,omitempty"`

	// 退货信息（可为空）。
	ReturnedFlag *bool   `json:"returnedFlag,omitempty"`
	ReturnDate   *string `json:"returnDate,omitempty"`
	ReturnNumber *string `json:"returnNumber,omitempty"`
	ReturnReason *string `json:"returnReason,omitempty"`

	// VendorInvoiceNumber 供应商发票号（可为空）。
	VendorInvoiceNumber *string `json:"vendorInvoiceNumber,omitempty"`

	// Items 包裹内的行项目明细。
	Items []OrderPackageLineItem `json:"items,omitempty"`
}

// PackageShipFrom 表示包裹发货方地址（OpenAPI: PackageShipFrom）。
type PackageShipFrom struct {
	Attention *string `json:"attention,omitempty"`

	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Name      *string `json:"name,omitempty"`

	Company *string `json:"company,omitempty"`

	Address1 *string  `json:"address1,omitempty"`
	Address2 *string  `json:"address2,omitempty"`
	Address  []string `json:"address,omitempty"`

	City    *string `json:"city,omitempty"`
	Region  *string `json:"region,omitempty"`
	Postal  *string `json:"postal,omitempty"`
	Country *string `json:"country,omitempty"`

	Phone *string `json:"phone,omitempty"`
	Email *string `json:"email,omitempty"`

	LocationCode *string `json:"locationCode,omitempty"`

	AddressType *string `json:"addressType,omitempty"`

	TaxExemptNumber       *string `json:"taxExemptNumber,omitempty"`
	TaxRegistrationNumber *string `json:"taxRegistrationNumber,omitempty"`
}

// PackageShipTo 表示包裹收货方地址（OpenAPI: PackageShipTo）。
type PackageShipTo struct {
	Attention *string `json:"attention,omitempty"`

	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Name      *string `json:"name,omitempty"`

	Company *string `json:"company,omitempty"`

	Address1 *string  `json:"address1,omitempty"`
	Address2 *string  `json:"address2,omitempty"`
	Address  []string `json:"address,omitempty"`

	City    *string `json:"city,omitempty"`
	Region  *string `json:"region,omitempty"`
	Postal  *string `json:"postal,omitempty"`
	Country *string `json:"country,omitempty"`

	Phone *string `json:"phone,omitempty"`
	Email *string `json:"email,omitempty"`

	StoreNumber *string `json:"storeNumber,omitempty"`

	AddressType *string `json:"addressType,omitempty"`

	TaxExemptNumber       *string `json:"taxExemptNumber,omitempty"`
	TaxRegistrationNumber *string `json:"taxRegistrationNumber,omitempty"`
}

// OrderPackageLineItem 对应 OpenAPI 的 PackageLineItem（orders.packages[].items[]）。
type OrderPackageLineItem struct {
	// Quantity 本包裹内该商品的数量。
	Quantity int `json:"quantity"`

	// TitleI18n 多语言标题（可为空）。
	TitleI18n map[string]string `json:"titleI18n,omitempty"`

	// 商品标识（DSCO 会返回其一或多个，口径以 DSCO 为准）。
	DscoItemID *string `json:"dscoItemId,omitempty"`
	SKU        *string `json:"sku,omitempty"`
	PartnerSKU *string `json:"partnerSku,omitempty"`
	UPC        *string `json:"upc,omitempty"`
	EAN        *string `json:"ean,omitempty"`

	// LineNumber 对应订单的行号。
	LineNumber *int `json:"lineNumber,omitempty"`
	// OriginalLineNumber 原始行号（部分场景用于替换/拆分等）。
	OriginalLineNumber *int `json:"originalLineNumber,omitempty"`
	// OriginalOrderQuantity 原订单行的原始数量（部分场景用于对账/追溯）。
	OriginalOrderQuantity *int `json:"originalOrderQuantity,omitempty"`

	// RetailerItemIDs 零售商侧商品 ID 列表（可为空）。
	RetailerItemIDs []string `json:"retailerItemIds,omitempty"`

	// RetailerLineID 零售商侧行号（可为空）。
	RetailerLineID *string `json:"retailerLineId,omitempty"`

	// DepartmentId/DepartmentName 部门信息（可为空）。
	DepartmentId   *string `json:"departmentId,omitempty"`
	DepartmentName *string `json:"departmentName,omitempty"`

	// MerchandisingAccountId/MerchandisingAccountName 货品/陈列账号信息（可为空）。
	MerchandisingAccountId   *string `json:"merchandisingAccountId,omitempty"`
	MerchandisingAccountName *string `json:"merchandisingAccountName,omitempty"`

	// SerialNumbers 序列号列表（可为空）。
	SerialNumbers []string `json:"serialNumbers,omitempty"`

	// PackageSpanFlag 标记该商品是否跨多个包裹（可为空）。
	PackageSpanFlag *bool `json:"packageSpanFlag,omitempty"`
}

// OrderCreatedResult 表示创建订单的响应（锚点 a39）。
type OrderCreatedResult struct {
	// Status 表示处理状态（success/failure 等，取决于 DSCO 返回口径）。
	Status string `json:"status"`
	// DscoOrderIDs 表示 DSCO 侧订单 ID 列表。
	DscoOrderIDs []string `json:"dscoOrderIds,omitempty"`
	// DscoOrders 表示订单对象列表（部分接口会返回）。
	DscoOrders []Order `json:"dscoOrders,omitempty"`
	// EventDate 表示事件时间（RFC3339）。
	EventDate string `json:"eventDate,omitempty"`
}

// OrderPageQuery 表示“拉取订单分页”的查询参数（GET /order/page）。
type OrderPageQuery struct {
	// ScrollID 用于翻页；一旦传入 scrollId，其它查询参数会被 DSCO 忽略。
	ScrollID string `url:"scrollId,omitempty"`

	// ConsumerOrderNumber 消费者订单号；用于“按单号检索”口径。
	ConsumerOrderNumber string `url:"consumerOrderNumber,omitempty"`

	// OrdersCreatedSince 按创建时间过滤的起始时间（ISO 8601/RFC3339）。
	OrdersCreatedSince string `url:"ordersCreatedSince,omitempty"`
	// OrdersUpdatedSince 按更新时间过滤的起始时间（ISO 8601/RFC3339）。
	OrdersUpdatedSince string `url:"ordersUpdatedSince,omitempty"`
	// Until 过滤窗口的结束时间（ISO 8601/RFC3339），要求必须在过去至少 5 秒。
	Until string `url:"until,omitempty"`

	// Status 支持多值，示例：?status=created&status=shipment_pending
	Status []string `url:"status,omitempty"`

	// IncludeTestOrders 是否包含测试订单。
	IncludeTestOrders *bool `url:"includeTestOrders,omitempty"`
	// ReturnedOnly 是否只返回已退货订单。
	ReturnedOnly *bool `url:"returnedOnly,omitempty"`

	// OrdersPerPage 每页返回订单数量（10~1000）。
	OrdersPerPage int `url:"ordersPerPage,omitempty"`
}

// PagedOrderResult 表示 GET /order/page 返回。
type PagedOrderResult struct {
	// ScrollID 用于翻页。
	ScrollID string `json:"scrollId,omitempty"`
	// Orders 是订单对象列表。
	Orders []Order `json:"orders"`
}

// PagedOrderResultRaw 表示 GET /order/page 返回（保留订单原始 JSON）。
type PagedOrderResultRaw struct {
	// ScrollID 用于翻页。
	ScrollID string `json:"scrollId,omitempty"`
	// Orders 是订单原始 JSON 列表（用于“少翻译”策略）。
	Orders []json.RawMessage `json:"orders"`
}

// -------- Order Change Log（GET /order/log）--------

// OrderChangeLogQuery 表示 GET /order/log 的查询参数。
//
// 注意：
// - 若提供 scrollId，则其它参数会被 DSCO 忽略。
// - 若未提供 scrollId，则 requestId 与 startDate/endDate 二选一（DSCO 文档说明）。
type OrderChangeLogQuery struct {
	ScrollID string `url:"scrollId,omitempty"`

	StartDate string `url:"startDate,omitempty"`
	EndDate   string `url:"endDate,omitempty"`

	RequestID string `url:"requestId,omitempty"`

	// Status 过滤条件：pending / success / failure / success_or_failure
	Status string `url:"status,omitempty"`
}

// OrderChangeLogResponse 对应 OpenAPI 的 OrderChangeLogResponse。
type OrderChangeLogResponse struct {
	ScrollID string `json:"scrollId,omitempty"`

	// Status 仅在传 requestId 时返回：PROCESSING / COMPLETED
	Status string `json:"status,omitempty"`

	Logs []OrderChangeLog `json:"logs"`
}

// OrderChangeLog 对应 OpenAPI 的 OrderChangeLog。
//
// payload 为 oneOf：Order / ACK payload / cancel payload / shipment payload 等。
// 为避免“为了测试而一次性全量建模”，这里用 RawMessage 保留原始 JSON 供排查与审计。
type OrderChangeLog struct {
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

// OrderAcknowledgeIDType 表示 DSCO acknowledge 的 type 枚举。
type OrderAcknowledgeIDType string

const (
	OrderAcknowledgeIDTypeDscoOrderID         OrderAcknowledgeIDType = "DSCO_ORDER_ID"
	OrderAcknowledgeIDTypePoNumber            OrderAcknowledgeIDType = "PO_NUMBER"
	OrderAcknowledgeIDTypeSupplierOrderNumber OrderAcknowledgeIDType = "SUPPLIER_ORDER_NUMBER"
)

// OrderAcknowledgeRequest 表示 POST /order/acknowledge 的单项入参。
type OrderAcknowledgeRequest struct {
	// ID 是订单标识值，配合 Type 使用（string 或 integer，取决于 Type）。
	ID any `json:"id"`
	// Type 指定 ID 的口径类型（DSCO_ORDER_ID / PO_NUMBER / SUPPLIER_ORDER_NUMBER）。
	Type OrderAcknowledgeIDType `json:"type"`
	// SupplierOrderNumber 可选：供应商侧订单号，ACK 时可写入 DSCO 订单。
	SupplierOrderNumber string `json:"supplierOrderNumber,omitempty"`
}
