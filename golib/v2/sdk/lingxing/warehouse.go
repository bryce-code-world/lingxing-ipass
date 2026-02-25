package lingxing

import (
	"context"
	"net/http"
)

// WarehouseService 提供仓库相关接口（销售出库单等）。
type WarehouseService struct {
	c *Client
}

// WmsOrderList 查询销售出库单列表。
//
// API Path: /erp/sc/routing/wms/order/wmsOrderList
func (s *WarehouseService) WmsOrderList(ctx context.Context, req WmsOrderListRequest) ([]WmsOrder, int, error) {
	var items []WmsOrder
	env, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/erp/sc/routing/wms/order/wmsOrderList", nil, req, &items)
	if err != nil {
		return nil, 0, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, nil
}

// WarehouseLists 查询仓库列表。
//
// API Path: /erp/sc/data/local_inventory/warehouse
func (s *WarehouseService) WarehouseLists(ctx context.Context, req WarehouseListsRequest) ([]WarehouseInfo, int, error) {
	var items []WarehouseInfo
	env, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/erp/sc/data/local_inventory/warehouse", nil, req, &items)
	if err != nil {
		return nil, 0, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, nil
}

// ListUsedLogisticsType 查询已启用的自发货物流方式列表。
//
// API Path: /erp/sc/routing/wms/WmsLogistics/listUsedLogisticsType
func (s *WarehouseService) ListUsedLogisticsType(ctx context.Context, req ListUsedLogisticsTypeRequest) ([]UsedLogisticsType, int, error) {
	var items []UsedLogisticsType
	env, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/erp/sc/routing/wms/WmsLogistics/listUsedLogisticsType", nil, req, &items)
	if err != nil {
		return nil, 0, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, nil
}
