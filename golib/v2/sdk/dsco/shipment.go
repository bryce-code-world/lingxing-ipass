package dsco

import (
	"context"
	"net/http"
)

// ShipmentService 封装 Shipment 相关接口（Create Shipment / Small Batch 等）。
type ShipmentService struct {
	c *Client
}

// CreateSingle 创建单次发货信息（POST /order/singleShipment）。
//
// 对应 OpenAPI：Create Shipment（operationId: singleShipment）。
func (s *ShipmentService) CreateSingle(ctx context.Context, req ShipmentsForUpdate) (*SuccessFailResponse, error) {
	var out SuccessFailResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/order/singleShipment", nil, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateSmallBatch 批量创建/追加发货信息（POST /order/shipment/batch/small）。
//
// 对应 OpenAPI：Create Shipment Small Batch（operationId: createShipmentSmallBatch）。
// 该接口为异步处理，返回 requestId，可通过 Order Change Log（GET /order/log）追踪每条结果。
func (s *ShipmentService) CreateSmallBatch(ctx context.Context, reqs []ShipmentsForUpdate) (*AsyncUpdateResponse, error) {
	var out AsyncUpdateResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/order/shipment/batch/small", nil, reqs, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
