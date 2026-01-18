package sync

import (
	"context"
	"errors"
	"strings"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"
	"lingxingipass/internal/platform/retry"
	"lingxingipass/internal/store"
)

// StockPipeline 负责一期“库存同步（可用库存）”编排。
type StockPipeline struct {
	dscoCli     *dsco.Client
	lingxingCli *lingxing.Client
	manualTask  *store.ManualTaskStore

	widToWarehouseCode map[string]string
	skuToSKU           map[string]string
}

func NewStockPipeline(dscoCli *dsco.Client, lingxingCli *lingxing.Client, manualTask *store.ManualTaskStore, widToWarehouseCode map[string]string, skuToSKU map[string]string) (*StockPipeline, error) {
	if dscoCli == nil {
		return nil, errors.New("dscoCli 不能为空")
	}
	if lingxingCli == nil {
		return nil, errors.New("lingxingCli 不能为空")
	}
	if manualTask == nil {
		return nil, errors.New("manualTask 不能为空")
	}
	if len(widToWarehouseCode) == 0 {
		return nil, errors.New("widToWarehouseCode 不能为空")
	}
	return &StockPipeline{
		dscoCli:            dscoCli,
		lingxingCli:        lingxingCli,
		manualTask:         manualTask,
		widToWarehouseCode: widToWarehouseCode,
		skuToSKU:           skuToSKU,
	}, nil
}

// SyncStock 将领星“可用库存(product_valid_num)”同步到 DSCO inventory/singleItem。
//
// 一期策略：
// - 只同步配置里指定的仓库（WID->warehouseCode 映射）
// - SKU 映射缺失则默认同名；warehouse 映射缺失则转人工（missing_mapping）
func (p *StockPipeline) SyncStock(ctx context.Context, batchSize int) error {
	if batchSize <= 0 {
		return errors.New("batchSize 必须大于 0")
	}

	logger.Info(ctx, "stock_sync_start", "job", "sync_stock", "batch_size", batchSize)

	for wid, dscoWH := range p.widToWarehouseCode {
		wid = strings.TrimSpace(wid)
		dscoWH = strings.TrimSpace(dscoWH)
		if wid == "" || dscoWH == "" {
			continue
		}
		logger.Info(ctx, "stock_sync_wid_start", "wid", wid, "dsco_warehouse_code", dscoWH)

		offset := 0
		for {
			var items []lingxing.InventoryDetailsItem
			var total int
			callErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryLingXing, func() error {
				var err error
				items, total, err = p.lingxingCli.Inventory.InventoryDetails(ctx, lingxing.InventoryDetailsRequest{
					WID:    wid,
					Offset: offset,
					Length: batchSize,
				})
				return err
			})
			if callErr != nil {
				return callErr
			}
			if len(items) == 0 {
				break
			}
			logger.Info(ctx, "stock_sync_page", "wid", wid, "count", len(items), "offset", offset, "total", total)

			for _, it := range items {
				lxSKU := strings.TrimSpace(it.SKU)
				if lxSKU == "" {
					continue
				}
				dscoSKU := lxSKU
				if p.skuToSKU != nil {
					if mapped, ok := p.skuToSKU[lxSKU]; ok && strings.TrimSpace(mapped) != "" {
						dscoSKU = strings.TrimSpace(mapped)
					}
				}

				qty := it.ProductValidNum
				inv := &dsco.ItemInventory{
					SKU: dscoSKU,
					Warehouses: []dsco.ItemWarehouse{
						{Code: dscoWH, Quantity: &qty},
					},
				}
				upsertErr := retry.Do(ctx, retry.DefaultPolicy(), shouldRetryDSCO, func() error {
					_, err := p.dscoCli.Inventory.UpsertSingle(ctx, inv)
					return err
				})
				if upsertErr != nil {
					_ = p.manualTask.Create(ctx, store.ManualTask{
						TaskType:    "sync_failed",
						DscoOrderID: "",
						Payload:     []byte(`{"reason":"inventory/singleItem 失败"}`),
					})
					continue
				}
			}

			offset += len(items)
			if offset >= total {
				break
			}
		}
		logger.Info(ctx, "stock_sync_wid_end", "wid", wid)
	}

	logger.Info(ctx, "stock_sync_end", "job", "sync_stock")
	return nil
}
