package dsco

// SuccessFailResponse 对应 OpenAPI 中的 SuccessFailResponse（复用锚点 a2）。
type SuccessFailResponse struct {
	// Success 表示是否成功。
	Success bool `json:"success"`
}

// AsyncUpdateResponse 对应 OpenAPI 中的 AsyncUpdateResponse（锚点 a6），用于异步接口返回 requestId。
type AsyncUpdateResponse struct {
	// Status 表示处理状态（success/failure 等，取决于 DSCO 返回口径）。
	Status string `json:"status"`
	// RequestID 表示异步请求 ID，可用于后续查询 change log。
	RequestID string `json:"requestId"`
	// EventDate 表示事件时间（RFC3339）。
	EventDate string `json:"eventDate"`
	// Messages 表示返回的消息列表。
	Messages []APIResponseMessage `json:"messages,omitempty"`
}

// APIResponseMessage 对应 OpenAPI 中的 ApiResponseMessage（锚点 a8）。
type APIResponseMessage struct {
	// Code 表示返回码/错误码。
	Code string `json:"code"`
	// Severity 表示消息级别（可选）。
	Severity string `json:"severity,omitempty"`
	// Description 表示消息描述（可选，类型不固定）。
	Description any `json:"description,omitempty"`
}
