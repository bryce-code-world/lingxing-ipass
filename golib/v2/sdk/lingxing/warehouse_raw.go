package lingxing

import (
	"context"
	"net/http"
)

func (s *WarehouseService) WmsOrderListWithRawBody(ctx context.Context, req WmsOrderListRequest) ([]WmsOrder, int, string, error) {
	var items []WmsOrder
	env, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/erp/sc/routing/wms/order/wmsOrderList", nil, req, &items)
	if err != nil {
		return nil, 0, raw, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, raw, nil
}

func (s *WarehouseService) WarehouseListsWithRawBody(ctx context.Context, req WarehouseListsRequest) ([]WarehouseInfo, int, string, error) {
	var items []WarehouseInfo
	env, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/erp/sc/data/local_inventory/warehouse", nil, req, &items)
	if err != nil {
		return nil, 0, raw, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, raw, nil
}

func (s *WarehouseService) ListUsedLogisticsTypeWithRawBody(ctx context.Context, req ListUsedLogisticsTypeRequest) ([]UsedLogisticsType, int, string, error) {
	var items []UsedLogisticsType
	env, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/erp/sc/routing/wms/WmsLogistics/listUsedLogisticsType", nil, req, &items)
	if err != nil {
		return nil, 0, raw, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, raw, nil
}
