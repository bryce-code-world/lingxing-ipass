package dsco

import (
	"context"
	"net/http"
)

// ReturnService 封装 Return（退货）相关接口。
type ReturnService struct {
	c *Client
}

// Create 创建单个退货单（POST /return/）。
func (s *ReturnService) Create(ctx context.Context, req *ReturnCreateRequest) (*ReturnResponse, error) {
	var out ReturnResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/return/", nil, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Complete 完成单个退货单（PUT /return/）。
func (s *ReturnService) Complete(ctx context.Context, req *ReturnCompleteRequest) (*ReturnResponse, error) {
	var out ReturnResponse
	if err := s.c.doJSON(ctx, http.MethodPut, "/return/", nil, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetChangeLog 获取退货变更日志（GET /return/log）。
func (s *ReturnService) GetChangeLog(ctx context.Context, q ReturnChangeLogQuery) (*ReturnChangeLogResponse, error) {
	var out ReturnChangeLogResponse
	if err := s.c.doJSON(ctx, http.MethodGet, "/return/log", q, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
