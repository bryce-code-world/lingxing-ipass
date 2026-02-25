package lingxing

import (
	"context"
	"net/http"
)

func (s *InventoryService) InventoryDetailsWithRawBody(ctx context.Context, req InventoryDetailsRequest) ([]InventoryDetailsItem, int, string, error) {
	var items []InventoryDetailsItem
	env, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/erp/sc/routing/data/local_inventory/inventoryDetails", nil, req, &items)
	if err != nil {
		return nil, 0, raw, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, raw, nil
}

func (s *InventoryService) InventoryBinDetailsWithRawBody(ctx context.Context, req InventoryBinDetailsRequest) ([]InventoryBinDetailsItem, int, string, error) {
	var items []InventoryBinDetailsItem
	env, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/erp/sc/routing/data/local_inventory/inventoryBinDetails", nil, req, &items)
	if err != nil {
		return nil, 0, raw, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, raw, nil
}

func (s *InventoryService) WarehouseBinWithRawBody(ctx context.Context, req WarehouseBinRequest) ([]WarehouseBinItem, int, string, error) {
	var items []WarehouseBinItem
	env, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/erp/sc/routing/data/local_inventory/warehouseBin", nil, req, &items)
	if err != nil {
		return nil, 0, raw, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, raw, nil
}

func (s *InventoryService) WarehouseStatementWithRawBody(ctx context.Context, req WarehouseStatementRequest) ([]WarehouseStatementItem, int, string, error) {
	var items []WarehouseStatementItem
	env, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/erp/sc/routing/data/local_inventory/wareHouseStatement", nil, req, &items)
	if err != nil {
		return nil, 0, raw, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, raw, nil
}

func (s *InventoryService) WarehouseBinStatementWithRawBody(ctx context.Context, req WarehouseStatementRequest) ([]WarehouseStatementItem, int, string, error) {
	var items []WarehouseStatementItem
	env, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/erp/sc/routing/data/local_inventory/wareHouseBinStatement", nil, req, &items)
	if err != nil {
		return nil, 0, raw, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, raw, nil
}
