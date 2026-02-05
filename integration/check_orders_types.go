package integration

// CheckOrdersOverride 表示“订单一致性检查”任务的手动入参。
// - Start/End：按 dsco_create_time（Unix 秒，UTC）检查 DSCO 订单创建时间范围：[start,end)。
// - 当 OnlyPONumber 非空时，任务优先按单号检查，此时 Start/End 可为空。
type CheckOrdersOverride struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type OrderCheckResult struct {
	PONumber string `json:"po_number"`
	OK       bool   `json:"ok"`
	Code     string `json:"code"`
	Message  string `json:"message"`

	LocalStatus int16  `json:"local_status"`
	DSCOStatus  string `json:"dsco_status,omitempty"`

	Detail any `json:"detail,omitempty"`
}

type CheckOrdersResponse struct {
	Results      []OrderCheckResult `json:"results"`
	MissingLocal []MissingLocalOrder `json:"missing_local"`
	Meta         map[string]any     `json:"meta,omitempty"`
}

type MissingLocalOrder struct {
	PONumber    string `json:"po_number"`
	DSCOStatus  string `json:"dsco_status,omitempty"`
	DSCOOrderID string `json:"dsco_order_id,omitempty"`
}
