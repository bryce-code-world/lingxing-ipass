package lingxing

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

func (s *OrderService) NewPlatformOrderListWithRawBody(ctx context.Context, req NewPlatformOrderListRequest) (NewPlatformOrderListResponse, string, error) {
	var out NewPlatformOrderListResponse
	_, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/cepfPlatformOrder/open-api/newPlatformOrder/list", nil, req, &out)
	return out, raw, err
}

func (s *OrderService) CreateOrdersV2WithRawBody(ctx context.Context, req CreateOrdersV2Request) (CreateOrdersV2ResponseData, string, error) {
	var out CreateOrdersV2ResponseData
	_, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/pb/mp/order/v2/create", nil, req, &out)
	return out, raw, err
}

func (s *OrderService) ListOrdersV2WithRawBody(ctx context.Context, req OrderListV2Request) (OrderListV2ResponseData, string, error) {
	var out OrderListV2ResponseData
	_, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/pb/mp/order/v2/list", nil, req, &out)
	return out, raw, err
}

func (s *OrderService) GetOrderDetailV2WithRawBody(ctx context.Context, req OrderDetailV2Request) (OrderDetailV2, string, error) {
	platformOrderNo := strings.TrimSpace(req.PlatformOrderNo)
	platformOrderName := strings.TrimSpace(req.PlatformOrderName)
	if platformOrderNo == "" && platformOrderName == "" {
		return OrderDetailV2{}, "", errors.New("platform_order_no/platform_order_name 不能同时为空")
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
	_, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/pb/mp/order/v2/list", nil, listReq, &out)
	if err != nil {
		return OrderDetailV2{}, raw, err
	}
	if len(out.List) == 0 {
		return OrderDetailV2{}, raw, errors.New("未查询到订单")
	}
	return out.List[0], raw, nil
}

func (s *OrderService) EditOrderWithRawBody(ctx context.Context, req EditOrderRequest) (OrderBatchUpdateResult, string, error) {
	out, err := s.EditOrder(ctx, req)
	return out, out.RawBody, err
}

func (s *OrderService) UpdateOrderV2WithRawBody(ctx context.Context, req UpdateOrderV2Request) (OrderBatchUpdateResult, string, error) {
	out, err := s.UpdateOrderV2(ctx, req)
	return out, out.RawBody, err
}
