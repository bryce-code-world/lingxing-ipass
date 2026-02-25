package dsco

import (
	"context"
	"errors"
	"net/http"
)

func (s *InvoiceService) GetByIDWithRawBody(ctx context.Context, q InvoiceGetQuery) (*GetInvoicesByIDResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out GetInvoicesByIDResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodGet, "/invoice", q, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *InvoiceService) CreateSingleWithRawBody(ctx context.Context, inv *Invoice) (*SuccessFailResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out SuccessFailResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/invoice", nil, inv, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *InvoiceService) CreateSmallBatchWithRawBody(ctx context.Context, invs []Invoice) (*AsyncUpdateResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out AsyncUpdateResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/invoice/batch/small", nil, invs, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *InvoiceService) GetChangeLogWithRawBody(ctx context.Context, q InvoiceChangeLogQuery) (*InvoiceChangeLogResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out InvoiceChangeLogResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodGet, "/invoice/log", q, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}
