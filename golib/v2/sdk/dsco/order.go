package dsco

import (
	"context"
	"net/http"
)

// OrderService 封装 Order 相关接口。
type OrderService struct {
	c *Client
}

// Create 创建单个订单（POST /order/）。
func (s *OrderService) Create(ctx context.Context, order *Order) (*OrderCreatedResult, error) {
	var out OrderCreatedResult
	if err := s.c.doJSON(ctx, http.MethodPost, "/order/", nil, order, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetByKey 获取单个订单对象（GET /order/）。
//
// orderKey/value 为必填，额外参数可通过 query 传入（例如 dscoAccountId 等）。
func (s *OrderService) GetByKey(ctx context.Context, orderKey, value string, query interface{}) (*Order, error) {
	q := map[string]string{
		"orderKey": orderKey,
		"value":    value,
	}
	// 允许调用方额外追加参数（简单优先：只支持 map[string]string / url.Values / struct url tag）。
	merged := mergeQuery(q, query)

	var out Order
	if err := s.c.doJSON(ctx, http.MethodGet, "/order/", merged, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPage 拉取订单分页（GET /order/page）。
func (s *OrderService) GetPage(ctx context.Context, q OrderPageQuery) (*PagedOrderResult, error) {
	var out PagedOrderResult
	if err := s.c.doJSON(ctx, http.MethodGet, "/order/page", q, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPageRaw 拉取订单分页（GET /order/page），并保留每笔订单的原始 JSON。
func (s *OrderService) GetPageRaw(ctx context.Context, q OrderPageQuery) (*PagedOrderResultRaw, error) {
	var out PagedOrderResultRaw
	if err := s.c.doJSON(ctx, http.MethodGet, "/order/page", q, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetChangeLog 获取订单变更日志（GET /order/log）。
func (s *OrderService) GetChangeLog(ctx context.Context, q OrderChangeLogQuery) (*OrderChangeLogResponse, error) {
	var out OrderChangeLogResponse
	if err := s.c.doJSON(ctx, http.MethodGet, "/order/log", q, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Acknowledge 批量回传订单 ACK（POST /order/acknowledge）。
func (s *OrderService) Acknowledge(ctx context.Context, items []OrderAcknowledgeRequest) (*AsyncUpdateResponse, error) {
	var out AsyncUpdateResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/order/acknowledge", nil, items, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func mergeQuery(base map[string]string, extra interface{}) interface{} {
	if extra == nil {
		return base
	}
	switch v := extra.(type) {
	case map[string]string:
		for k, val := range v {
			base[k] = val
		}
		return base
	default:
		// 交给 encodeQuery 去处理 struct/url.Values 等。
		// 这里用一个轻量 wrapper：先把 base 写入 url.Values，再把 extra 的 encodeQuery 合并进去。
		// 为了避免引入接口/抽象层，直接在 request.go 的 encodeQuery 上复用。
		values := encodeQuery(base)
		for k, arr := range encodeQuery(v) {
			for _, val := range arr {
				values.Add(k, val)
			}
		}
		return values
	}
}
