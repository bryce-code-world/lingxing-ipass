package dsco

import "encoding/json"

// InventoryGetQuery 表示查询单个库存对象的参数。
// OpenAPI 中该组参数复用自 itemKey/value/dscoRetailerId/dscoTradingPartnerId/returnMultiple。
type InventoryGetQuery struct {
	// ItemKey 表示用于查询的商品标识类型（dscoItemId / sku / partnerSku / upc / ean / mpn / isbn / gtin）。
	ItemKey string `url:"itemKey"`
	// Value 表示 ItemKey 对应的值。
	Value string `url:"value"`
	// DscoRetailerID 当 itemKey=partnerSku 且调用方为 supplier 时，需要提供 dscoRetailerId 或 dscoTradingPartnerId 之一。
	DscoRetailerID string `url:"dscoRetailerId,omitempty"`
	// DscoTradingPartner 当 itemKey=partnerSku 且调用方为 supplier 时，需要提供 dscoRetailerId 或 dscoTradingPartnerId 之一。
	DscoTradingPartner string `url:"dscoTradingPartnerId,omitempty"`
	// ReturnMultiple 当存在多条匹配时，是否返回所有匹配项（默认只返回第一条）。
	ReturnMultiple *bool `url:"returnMultiple,omitempty"`
}

// Item 定义 ItemCatalog / ItemInventory 共享字段（对应 OpenAPI 锚点 a48）。
type Item struct {
	// SKU 供应商提供的唯一商品标识。
	SKU string `json:"sku"`

	// DscoItemID Dsco 的商品唯一标识。
	DscoItemID string `json:"dscoItemId,omitempty"`

	// UPC / EAN / MPN / ISBN / GTIN 为可选标识。
	UPC          *string `json:"upc,omitempty"`
	DepartmentID *string `json:"departmentId,omitempty"`
	EAN          *string `json:"ean,omitempty"`
	MPN          *string `json:"mpn,omitempty"`
	ISBN         *string `json:"isbn,omitempty"`
	GTIN         *string `json:"gtin,omitempty"`

	// PartnerSKU 为零售商侧 SKU（对 supplier 场景可能为空）。
	PartnerSKU *string `json:"partnerSku,omitempty"`

	// PartnerSKUMap 表示零售商 SKU 映射关系。
	PartnerSKUMap []PartnerSKUMapEntry `json:"partnerSkuMap,omitempty"`

	// Assortments 表示该商品关联的 assortment 列表。
	Assortments []string `json:"assortments,omitempty"`
}

// PartnerSKUMapEntry 对应 OpenAPI 的 PartnerSkuMap。
type PartnerSKUMapEntry struct {
	PartnerSKU string `json:"partnerSku"`

	// DscoRetailerID 与 DscoTradingPartnerID 二选一（由 DSCO 规则决定）。
	DscoRetailerID       string `json:"dscoRetailerId,omitempty"`
	DscoTradingPartnerID string `json:"dscoTradingPartnerId,omitempty"`
}

// ItemInventory 表示库存对象（对应 OpenAPI 锚点 a53）。
//
// 注意：该结构包含大量只读字段（例如 dscoCreateDate/retailerCode 等），
// 仍然保留是为了能完整解析 DSCO 返回值。
type ItemInventory struct {
	Item

	// Warehouses 表示各仓库库存信息列表（DSCO 要求至少一个仓库）。
	Warehouses []ItemWarehouse `json:"warehouses"`

	// QuantityAvailable 若指定，则应为各仓库 quantity 的准确总和（DSCO 文档说明未来可能转为只读）。
	QuantityAvailable *int `json:"quantityAvailable,omitempty"`

	// Cost 表示该商品成本/价格（具体口径以 DSCO 侧为准）。
	Cost *float64 `json:"cost,omitempty"`

	// ListingPrices 表示 marketplace retail model 下的售价（按币种）。
	ListingPrices *ListingPrices `json:"listingPrices,omitempty"`

	// Msrp 建议零售价。
	Msrp *float64 `json:"msrp,omitempty"`

	// RetailModelRules 表示允许的零售模型规则。
	RetailModelRules *RetailModelRules `json:"retailModelRules,omitempty"`

	// Status 库存状态（in-stock / out-of-stock / discontinued / hidden）。
	Status string `json:"status,omitempty"`

	// EstimatedAvailabilityDate 预计到货时间（RFC3339）。
	EstimatedAvailabilityDate *string `json:"estimatedAvailabilityDate,omitempty"`

	// QuantityOnOrder 在途/在单数量。
	QuantityOnOrder *int `json:"quantityOnOrder,omitempty"`

	// CurrencyCode 成本币种（ISO 4217）。
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Backstop 表示 backstop pool 相关属性。
	Backstop *ItemBackstop `json:"backstop,omitempty"`

	// Title / Brand 为商品展示信息（可用于 UI/审计）。
	Title     string            `json:"title,omitempty"`
	TitleI18n map[string]string `json:"titleI18n,omitempty"`
	Brand     *string           `json:"brand,omitempty"`
	BrandI18n map[string]string `json:"brandI18n,omitempty"`

	// DSCO 归属信息。
	DscoSupplierID             string  `json:"dscoSupplierId,omitempty"`
	DscoSupplierName           string  `json:"dscoSupplierName,omitempty"`
	DscoTradingPartnerID       string  `json:"dscoTradingPartnerId,omitempty"`
	DscoTradingPartnerName     string  `json:"dscoTradingPartnerName,omitempty"`
	DscoTradingPartnerParentID *string `json:"dscoTradingPartnerParentId,omitempty"`

	// DSCO 时间戳（RFC3339）。
	DscoCreateDate             string  `json:"dscoCreateDate,omitempty"`
	DscoLastQuantityUpdateDate string  `json:"dscoLastQuantityUpdateDate,omitempty"`
	DscoLastCostUpdateDate     string  `json:"dscoLastCostUpdateDate,omitempty"`
	DscoLastUpdateDate         *string `json:"dscoLastUpdateDate,omitempty"`

	// InventoryByContext 表示按 retailer context 的库存分配（可选）。
	InventoryByContext []InventoryContextMap `json:"inventoryByContext,omitempty"`

	// ActiveFulfillmentModel / RequestedFulfillmentModels 表示履约模型信息。
	ActiveFulfillmentModel     *string  `json:"activeFulfillmentModel,omitempty"`
	RequestedFulfillmentModels []string `json:"requestedFulfillmentModels,omitempty"`

	// RetailerOwnedAttributes 为“以 retailerId 为 key 的动态对象”，字段由零售商控制且可能扩展。
	// 为避免 SDK 在字段扩展时频繁变更，这里用 map[string]any 透传。
	RetailerOwnedAttributes map[string]any `json:"retailerOwnedAttributes,omitempty"`
}

// ListingPrices 对应 OpenAPI 的 LocalizedCurrencyValue（示例显式列出 USD/GBP/EUR/CAD/AUD）。
type ListingPrices struct {
	USD *float64 `json:"USD,omitempty"`
	GBP *float64 `json:"GBP,omitempty"`
	EUR *float64 `json:"EUR,omitempty"`
	CAD *float64 `json:"CAD,omitempty"`
	AUD *float64 `json:"AUD,omitempty"`
}

// RetailModelRules 对应 OpenAPI 的 RetailModelRules。
type RetailModelRules struct {
	AllowedModels []string `json:"allowedModels"`
}

// ItemBackstop 对应 OpenAPI 的 ItemBackstop。
type ItemBackstop struct {
	Discoverable bool `json:"discoverable"`
}

// ItemWarehouse 表示单个仓库的库存信息（对应 OpenAPI 的 ItemWarehouse）。
type ItemWarehouse struct {
	// Code 供应商侧仓库唯一标识（DSCO 要求必填）。
	Code string `json:"code"`

	// DscoID Rithum/DSCO 侧仓库唯一标识。
	DscoID string `json:"dscoId,omitempty"`

	// RetailerCode 为 retailer 侧仓库标识（只读）。
	RetailerCode string `json:"retailerCode,omitempty"`

	// RetailerCodes 为“retailerId -> { retailerCode }”映射。
	RetailerCodes map[string]RetailerCode `json:"retailerCodes,omitempty"`

	// Quantity 库存数量（可为空）。
	Quantity *int `json:"quantity,omitempty"`

	// QuantityOnOrder 在途/在单数量（可为空）。
	QuantityOnOrder *int `json:"quantityOnOrder,omitempty"`

	// Cost 仓库维度成本（可为空）。
	Cost *float64 `json:"cost,omitempty"`

	// Status 仓库维度状态（口径以 DSCO 为准）。
	Status string `json:"status,omitempty"`

	// EstimatedAvailabilityDate 预计到货时间（RFC3339，可为空）。
	EstimatedAvailabilityDate *string `json:"estimatedAvailabilityDate,omitempty"`

	// InventoryByContext 仓库维度的 context 分配（可为空）。
	InventoryByContext []InventoryContextMap `json:"inventoryByContext,omitempty"`
}

// RetailerCode 表示 retailerCodes 的 value。
type RetailerCode struct {
	RetailerCode string `json:"retailerCode,omitempty"`
}

// InventoryContextMap 对应 OpenAPI 的 InventoryContextMap（锚点 a51）。
type InventoryContextMap struct {
	// Contexts 表示 retailer context -> quantity 的 map，且必须包含 default。
	Contexts InventoryRetailerContextMap `json:"contexts"`

	RetailerID       string `json:"retailerId,omitempty"`
	TradingPartnerID string `json:"tradingPartnerId,omitempty"`
}

// InventoryRetailerContextMap 表示 inventoryByContext.contexts。
//
// OpenAPI 定义：
// - 必须包含 default 键
// - 允许 additionalProperties（其他自定义 context）
type InventoryRetailerContextMap struct {
	Default    InventoryQuantity            `json:"default"`
	Additional map[string]InventoryQuantity `json:"-"`
}

func (m InventoryRetailerContextMap) MarshalJSON() ([]byte, error) {
	raw := map[string]InventoryQuantity{
		"default": m.Default,
	}
	for k, v := range m.Additional {
		if k == "" || k == "default" {
			continue
		}
		raw[k] = v
	}
	return json.Marshal(raw)
}

func (m *InventoryRetailerContextMap) UnmarshalJSON(b []byte) error {
	var raw map[string]InventoryQuantity
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["default"]; ok {
		m.Default = v
		delete(raw, "default")
	}
	if len(raw) > 0 {
		m.Additional = raw
	} else {
		m.Additional = nil
	}
	return nil
}

// InventoryQuantity 对应 OpenAPI 的 InventoryQuantity（包含 quantityAvailable/reservedQuantity/stockOnHand）。
type InventoryQuantity struct {
	QuantityAvailable float64  `json:"quantityAvailable"`
	ReservedQuantity  *float64 `json:"reservedQuantity,omitempty"`
	StockOnHand       *float64 `json:"stockOnHand,omitempty"`
}
