package dsco

import (
	"context"
	"errors"
	"net/http"
)

func (s *OrderService) CreateWithRawBody(ctx context.Context, order *Order) (*OrderCreatedResult, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out OrderCreatedResult
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/order/", nil, order, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *OrderService) GetByKeyWithRawBody(ctx context.Context, orderKey, value string, query interface{}) (*Order, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	q := map[string]string{
		"orderKey": orderKey,
		"value":    value,
	}
	merged := mergeQuery(q, query)

	var out Order
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodGet, "/order/", merged, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *OrderService) GetPageWithRawBody(ctx context.Context, q OrderPageQuery) (*PagedOrderResult, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out PagedOrderResult
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodGet, "/order/page", q, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *OrderService) GetPageRawWithRawBody(ctx context.Context, q OrderPageQuery) (*PagedOrderResultRaw, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out PagedOrderResultRaw
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodGet, "/order/page", q, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *OrderService) GetChangeLogWithRawBody(ctx context.Context, q OrderChangeLogQuery) (*OrderChangeLogResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out OrderChangeLogResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodGet, "/order/log", q, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *OrderService) AcknowledgeWithRawBody(ctx context.Context, items []OrderAcknowledgeRequest) (*AsyncUpdateResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out AsyncUpdateResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/order/acknowledge", nil, items, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}
