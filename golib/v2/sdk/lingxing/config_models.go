package lingxing

// PairListV2Request 表示“查询多平台 SKU 配对列表”的请求参数。
//
// API Path: /pb/mp/listing/v2/getPairList
type PairListV2Request struct {
	// Length 分页条数（建议 >= 20；SDK 会在 Length<=0 时自动设置为 20）。
	Length int `json:"length"`
	// Offset 分页偏移量（默认 0）。
	Offset int `json:"offset,omitempty"`

	// MSKU MSKU 列表（可选）。
	MSKU []string `json:"msku,omitempty"`
	// SKU 本地 SKU 列表（可选）。
	SKU []string `json:"sku,omitempty"`

	// StartTime 操作开始时间（格式：2006-01-02 15:04:05；可选）。
	StartTime string `json:"start_time,omitempty"`
	// EndTime 操作结束时间（格式：2006-01-02 15:04:05；可选）。
	EndTime string `json:"end_time,omitempty"`

	// PlatformCodes 平台 code 列表（可选；示例："10003"）。
	PlatformCodes []string `json:"platform_codes,omitempty"`
	// StoreIDs 店铺 id 列表（可选）。
	StoreIDs []string `json:"store_ids,omitempty"`

	// UseCursor 分页游标（可选；为 true 时建议配合 CursorID 使用）。
	UseCursor *bool `json:"use_cursor,omitempty"`
	// CursorID 游标 id（当 UseCursor=true 时必填，第一次请求传 0）。
	CursorID *int64 `json:"cursor_id,omitempty"`
}

// PairListV2ResponseData 表示“查询多平台 SKU 配对列表”的 data 字段。
type PairListV2ResponseData struct {
	// Total 总数（未使用游标分页时返回；类型兼容 string/number）。
	Total StringOrNumber `json:"total,omitempty"`
	// HasNext 是否存在下一页（使用游标分页时返回）。
	HasNext bool `json:"has_next,omitempty"`
	// NextCursorID 下一页游标 id（使用游标分页时返回；类型兼容 string/number）。
	NextCursorID StringOrNumber `json:"next_cursor_id,omitempty"`
	// List 详细列表。
	List []PairListV2Item `json:"list"`
}

// PairListV2Item 表示“多平台 MSKU 与本地 SKU 的配对关系”条目。
type PairListV2Item struct {
	// StoreID 店铺 ID。
	StoreID StringOrNumber `json:"store_id"`
	// StoreName 店铺名称。
	StoreName string `json:"store_name"`
	// PlatformCode 平台 code。
	PlatformCode string `json:"platform_code"`
	// PlatformName 平台名称。
	PlatformName string `json:"platform_name"`
	// MSKU 多平台 MSKU。
	MSKU string `json:"msku"`
	// SKU 本地 SKU。
	SKU string `json:"sku"`
	// LocalName 本地 SKU 品名/名称。
	LocalName string `json:"local_name"`
	// ModifyTime 操作时间（格式：2006-01-02 15:04:05）。
	ModifyTime string `json:"modify_time"`
}
