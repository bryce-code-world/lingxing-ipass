package dsco

import (
	"context"
	"errors"
	"net/http"
)

func (s *ShipmentService) CreateSingleWithRawBody(ctx context.Context, req ShipmentsForUpdate) (*SuccessFailResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out SuccessFailResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/order/singleShipment", nil, req, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *ShipmentService) CreateSmallBatchWithRawBody(ctx context.Context, reqs []ShipmentsForUpdate) (*AsyncUpdateResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out AsyncUpdateResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/order/shipment/batch/small", nil, reqs, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}
