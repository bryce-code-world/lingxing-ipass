package lingxing

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// OrderService 提供订单相关接口（多平台订单）。
type OrderService struct {
	c *Client
}

// NewPlatformOrderList 查询平台订单列表。
//
// API Path: /cepfPlatformOrder/open-api/newPlatformOrder/list
func (s *OrderService) NewPlatformOrderList(ctx context.Context, req NewPlatformOrderListRequest) (NewPlatformOrderListResponse, error) {
	var out NewPlatformOrderListResponse
	_, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/cepfPlatformOrder/open-api/newPlatformOrder/list", nil, req, &out)
	return out, err
}

// CreateOrdersV2 创建多平台店铺手工单（包含自定义平台订单）。
//
// API Path: /pb/mp/order/v2/create
func (s *OrderService) CreateOrdersV2(ctx context.Context, req CreateOrdersV2Request) (CreateOrdersV2ResponseData, error) {
	var out CreateOrdersV2ResponseData
	_, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/pb/mp/order/v2/create", nil, req, &out)
	return out, err
}

// ListOrdersV2 查询订单管理订单列表。
//
// API Path: /pb/mp/order/v2/list
func (s *OrderService) ListOrdersV2(ctx context.Context, req OrderListV2Request) (OrderListV2ResponseData, error) {
	var out OrderListV2ResponseData
	_, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/pb/mp/order/v2/list", nil, req, &out)
	return out, err
}

// GetOrderDetailV2 获取单个订单详情。
//
// 说明：
// - 当前实现通过 /pb/mp/order/v2/list 查询并取第一条作为详情（适用于你已知 platform_order_no/platform_order_name 的场景）。
func (s *OrderService) GetOrderDetailV2(ctx context.Context, req OrderDetailV2Request) (OrderDetailV2, error) {
	platformOrderNo := strings.TrimSpace(req.PlatformOrderNo)
	platformOrderName := strings.TrimSpace(req.PlatformOrderName)
	if platformOrderNo == "" && platformOrderName == "" {
		return OrderDetailV2{}, errors.New("platform_order_no/platform_order_name 不能同时为空")
	}

	listReq := OrderListV2Request{
		Offset: 0,
		// 领星该接口对 length 有最小值限制（线上返回 code=102: length 最小值为 20）。
		// 这里只查单个订单时仍用较小页数即可，但需满足服务端限制。
		Length: 20,
	}
	if platformOrderNo != "" {
		listReq.PlatformOrderNos = []string{platformOrderNo}
	}
	if platformOrderName != "" {
		listReq.PlatformOrderNames = []string{platformOrderName}
	}

	type respData struct {
		Total StringOrNumber  `json:"total"`
		List  []OrderDetailV2 `json:"list"`
	}

	var out respData
	_, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/pb/mp/order/v2/list", nil, listReq, &out)
	if err != nil {
		return OrderDetailV2{}, err
	}
	if len(out.List) == 0 {
		return OrderDetailV2{}, errors.New("未查询到订单")
	}
	return out.List[0], nil
}

// EditOrder 编辑订单。
//
// 说明：该接口的 code 可能返回 10001（部分成功），SDK 会将其视为可用结果并返回 error_details。
//
// API Path: /pb/mp/order/editOrder
func (s *OrderService) EditOrder(ctx context.Context, req EditOrderRequest) (OrderBatchUpdateResult, error) {
	return s.doOrderBatchUpdate(ctx, "/pb/mp/order/editOrder", req)
}

// UpdateOrderV2 批量编辑/更新自发货订单。
//
// 说明：该接口的 code 可能返回 10001（部分成功）/10002（成功），SDK 会将其视为可用结果并返回 error_details。
//
// API Path: /pb/mp/order/v2/updateOrder
func (s *OrderService) UpdateOrderV2(ctx context.Context, req UpdateOrderV2Request) (OrderBatchUpdateResult, error) {
	// 领星文档：order_item_list “可以传空列表”。这里需要确保即使调用方未设置 OrderItemList，
	// 也会以空数组 `[]` 的形式出现在请求体中，而不是被 omitempty 去掉或编码为 null。
	for i := range req.OrderList {
		if req.OrderList[i].OrderItemList == nil {
			req.OrderList[i].OrderItemList = []UpdateOrderV2OrderItem{}
		}
	}
	return s.doOrderBatchUpdate(ctx, "/pb/mp/order/v2/updateOrder", req)
}

func (s *OrderService) doOrderBatchUpdate(ctx context.Context, path string, req any) (OrderBatchUpdateResult, error) {
	var out OrderBatchUpdateResult
	env, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, path, nil, req, nil)
	out.Code = env.Code.Raw()
	out.Message = env.Message()
	out.RequestID = env.RequestID()
	out.RawBody = raw

	dataText := strings.TrimSpace(string(env.Data))
	if dataText != "" && dataText != "null" {
		_ = json.Unmarshal(env.Data, &out.Data)
	}

	if err != nil && isOrderBatchUpdateAcceptableCode(out.Code) {
		return out, nil
	}
	return out, err
}

func isOrderBatchUpdateAcceptableCode(code string) bool {
	switch strings.TrimSpace(code) {
	case "0", "10001", "10002":
		return true
	default:
		return false
	}
}
