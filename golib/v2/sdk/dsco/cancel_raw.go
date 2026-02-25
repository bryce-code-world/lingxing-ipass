package dsco

import (
	"context"
	"errors"
	"net/http"
)

func (s *CancelService) OrderItemWithRawBody(ctx context.Context, req *OrderForCancel) (*SyncUpdateResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out SyncUpdateResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/order/item/cancel", nil, req, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *CancelService) OrderItemSmallBatchWithRawBody(ctx context.Context, req []OrderForCancel) (*AsyncUpdateResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out AsyncUpdateResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/order/item/cancel/batch/small", nil, req, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}
