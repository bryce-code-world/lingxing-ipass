package lingxing

// EditOrderRequest 表示“编辑订单”的请求参数。
//
// API Path: /pb/mp/order/editOrder
type EditOrderRequest struct {
	OrderList []EditOrderItem `json:"order_list"`
}

type EditOrderItem struct {
	// GlobalOrderNo 系统订单号
	GlobalOrderNo int64 `json:"global_order_no"`
	// Logistics 物流信息
	Logistics EditOrderLogistics `json:"logistics"`
}

type EditOrderLogistics struct {
	// LogisticsTypeID 物流方式 id（对应 listUsedLogisticsType 的 type_id）
	LogisticsTypeID int64 `json:"logistics_type_id"`
	// SysWID 仓库 id（对应 WarehouseLists 的 wid）
	SysWID int64 `json:"sys_wid"`
}

// UpdateOrderV2Request 表示“编辑/更新自发货订单”的请求参数。
//
// API Path: /pb/mp/order/v2/updateOrder
type UpdateOrderV2Request struct {
	OrderList []UpdateOrderV2Item `json:"order_list"`
}

type UpdateOrderV2Item struct {
	AddressInfo   *UpdateOrderV2AddressInfo `json:"address_info,omitempty"`
	GlobalOrderNo int64                     `json:"global_order_no,omitempty"`
	Logistics     *UpdateOrderV2Logistics   `json:"logistics,omitempty"`
	OrderItemList []UpdateOrderV2OrderItem  `json:"order_item_list"`
}

type UpdateOrderV2AddressInfo struct {
	AddressLine1        string `json:"address_line1,omitempty"`
	AddressLine2        string `json:"address_line2,omitempty"`
	AddressLine3        string `json:"address_line3,omitempty"`
	City                string `json:"city,omitempty"`
	District            string `json:"district,omitempty"`
	DoorplateNo         string `json:"doorplate_no,omitempty"`
	PostalCode          string `json:"postal_code,omitempty"`
	ReceiverCompanyName string `json:"receiver_company_name,omitempty"`
	ReceiverCountryCode string `json:"receiver_country_code,omitempty"`
	ReceiverMobile      string `json:"receiver_mobile,omitempty"`
	ReceiverName        string `json:"receiver_name,omitempty"`
	ReceiverTel         string `json:"receiver_tel,omitempty"`
	StateOrRegion       string `json:"state_or_region,omitempty"`
}

type UpdateOrderV2Logistics struct {
	CodType       string `json:"cod_type,omitempty"`
	SenderTaxNo   string `json:"sender_tax_no,omitempty"`
	SenderTaxType string `json:"sender_tax_type,omitempty"`
}

type UpdateOrderV2OrderItem struct {
	Mark            string `json:"mark,omitempty"`
	MSKU            string `json:"msku,omitempty"`
	Price           int64  `json:"price,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
	SKU             string `json:"sku,omitempty"`
	Type            int    `json:"type,omitempty"`
	ID              string `json:"id,omitempty"`
	PlatformOrderNo string `json:"platformOrderNo,omitempty"`
}

// OrderBatchUpdateResult 表示 editOrder / updateOrder 的结果（保留 code + error_details）。
type OrderBatchUpdateResult struct {
	// Code 接口返回的 code（文档可能为 0/10000/10001/10002 等）
	Code string
	// Message 接口返回的 message
	Message string
	// RequestID 接口返回的 request_id/requestId
	RequestID string
	// RawBody 原始响应体（便于排查）
	RawBody string

	Data OrderBatchUpdateData
}

type OrderBatchUpdateData struct {
	ErrorDetails []OrderBatchUpdateErrorDetail `json:"error_details"`
}

type OrderBatchUpdateErrorDetail struct {
	GlobalOrderNo string `json:"global_order_no"`
	ErrorMessage  string `json:"error_message"`
}
