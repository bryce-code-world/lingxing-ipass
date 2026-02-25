package dsco

import (
	"context"
	"errors"
	"net/http"
)

// CancelService 封装订单售后取消（Cancel）相关接口。
type CancelService struct {
	c *Client
}

// OrderItem 同步取消单个订单的一个或多个行项目（POST /order/item/cancel）。
func (s *CancelService) OrderItem(ctx context.Context, req *OrderForCancel) (*SyncUpdateResponse, error) {
	if s == nil || s.c == nil {
		return nil, errors.New("dsco client 不能为空")
	}
	var out SyncUpdateResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/order/item/cancel", nil, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// OrderItemSmallBatch 异步小批量取消多个订单的行项目（POST /order/item/cancel/batch/small）。
func (s *CancelService) OrderItemSmallBatch(ctx context.Context, req []OrderForCancel) (*AsyncUpdateResponse, error) {
	if s == nil || s.c == nil {
		return nil, errors.New("dsco client 不能为空")
	}
	var out AsyncUpdateResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/order/item/cancel/batch/small", nil, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
