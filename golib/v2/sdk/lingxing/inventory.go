package lingxing

import (
	"context"
	"net/http"
)

// InventoryService 提供库存相关接口（仓库库存、仓位库存等）。
type InventoryService struct {
	c *Client
}

// InventoryDetails 查询仓库库存明细。
//
// API Path: /erp/sc/routing/data/local_inventory/inventoryDetails
func (s *InventoryService) InventoryDetails(ctx context.Context, req InventoryDetailsRequest) ([]InventoryDetailsItem, int, error) {
	var items []InventoryDetailsItem
	env, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/erp/sc/routing/data/local_inventory/inventoryDetails", nil, req, &items)
	if err != nil {
		return nil, 0, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, nil
}

// InventoryBinDetails 查询仓位库存明细。
//
// API Path: /erp/sc/routing/data/local_inventory/inventoryBinDetails
func (s *InventoryService) InventoryBinDetails(ctx context.Context, req InventoryBinDetailsRequest) ([]InventoryBinDetailsItem, int, error) {
	var items []InventoryBinDetailsItem
	env, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/erp/sc/routing/data/local_inventory/inventoryBinDetails", nil, req, &items)
	if err != nil {
		return nil, 0, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, nil
}

// WarehouseBin 查询本地仓位列表。
//
// API Path: /erp/sc/routing/data/local_inventory/warehouseBin
func (s *InventoryService) WarehouseBin(ctx context.Context, req WarehouseBinRequest) ([]WarehouseBinItem, int, error) {
	var items []WarehouseBinItem
	env, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/erp/sc/routing/data/local_inventory/warehouseBin", nil, req, &items)
	if err != nil {
		return nil, 0, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, nil
}

// WarehouseStatement 查询库存流水（旧）。
//
// API Path: /erp/sc/routing/data/local_inventory/wareHouseStatement
func (s *InventoryService) WarehouseStatement(ctx context.Context, req WarehouseStatementRequest) ([]WarehouseStatementItem, int, error) {
	var items []WarehouseStatementItem
	env, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/erp/sc/routing/data/local_inventory/wareHouseStatement", nil, req, &items)
	if err != nil {
		return nil, 0, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, nil
}

// WarehouseBinStatement 查询仓位流水。
//
// API Path: /erp/sc/routing/data/local_inventory/wareHouseBinStatement
func (s *InventoryService) WarehouseBinStatement(ctx context.Context, req WarehouseStatementRequest) ([]WarehouseStatementItem, int, error) {
	var items []WarehouseStatementItem
	env, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/erp/sc/routing/data/local_inventory/wareHouseBinStatement", nil, req, &items)
	if err != nil {
		return nil, 0, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, nil
}
