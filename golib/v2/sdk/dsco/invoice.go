package dsco

import (
	"context"
	"net/http"
)

// InvoiceService 封装 Invoice 相关接口。
type InvoiceService struct {
	c *Client
}

// GetByID 获取发票列表（GET /invoice）。
func (s *InvoiceService) GetByID(ctx context.Context, q InvoiceGetQuery) (*GetInvoicesByIDResponse, error) {
	var out GetInvoicesByIDResponse
	if err := s.c.doJSON(ctx, http.MethodGet, "/invoice", q, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateSingle 创建单个发票（POST /invoice）。
func (s *InvoiceService) CreateSingle(ctx context.Context, inv *Invoice) (*SuccessFailResponse, error) {
	var out SuccessFailResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/invoice", nil, inv, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateSmallBatch 异步小批量创建发票（POST /invoice/batch/small）。
func (s *InvoiceService) CreateSmallBatch(ctx context.Context, invs []Invoice) (*AsyncUpdateResponse, error) {
	var out AsyncUpdateResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/invoice/batch/small", nil, invs, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetChangeLog 获取发票变更日志（GET /invoice/log）。
func (s *InvoiceService) GetChangeLog(ctx context.Context, q InvoiceChangeLogQuery) (*InvoiceChangeLogResponse, error) {
	var out InvoiceChangeLogResponse
	if err := s.c.doJSON(ctx, http.MethodGet, "/invoice/log", q, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
