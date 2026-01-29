package dsco_lingxing

import (
	"context"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

func (d *Domain) SyncStock(ctx integration.TaskContext) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		return err
	}
	dscoCli, err := d.dscoClient()
	if err != nil {
		return err
	}

	skuRev, err := buildReverseSKUMap(ctx.Config)
	if err != nil {
		return err
	}

	// mapping.warehouse: DSCO warehouseCode -> LingXing WID
	for dscoWarehouseCode, wid := range ctx.Config.Mapping.Warehouse {
		if dscoWarehouseCode == "" || wid == "" {
			continue
		}

		offset := 0
		for {
			items, total, err := lx.Inventory.InventoryDetails(taskCtx, lingxing.InventoryDetailsRequest{
				WID:    wid,
				Offset: offset,
				Length: ctx.Size,
			})
			if err != nil {
				logger.Warn(taskCtx, "lingxing inventoryDetails failed", "wid", wid, "err", err)
				break
			}
			if len(items) == 0 {
				break
			}

			var invs []dsco.ItemInventory
			for _, it := range items {
				partner := skuRev[it.SKU]
				if partner == "" {
					partner = it.SKU
				}
				qty := it.ProductValidNum
				partnerCopy := partner
				invs = append(invs, dsco.ItemInventory{
					Item: dsco.Item{
						SKU:        partner,
						PartnerSKU: &partnerCopy,
					},
					Warehouses: []dsco.ItemWarehouse{
						{Code: dscoWarehouseCode, Quantity: &qty},
					},
					QuantityAvailable: &qty,
				})

				_ = d.warehouseStore.Insert(taskCtx, store.DSCOWarehouseSyncRow{
					DSCOWarehouseID:      dscoWarehouseCode,
					DSCOWarehouseSKU:     partner,
					DSCOWarehouseNum:     qty,
					LingXingWarehouseID:  wid,
					LingXingWarehouseSKU: it.SKU,
					LingXingWarehouseNum: qty,
					Status:               1,
					Reason:               "",
				})
			}

			if len(invs) > 0 {
				_, err := dscoCli.Inventory.UpdateSmallBatch(taskCtx, invs, dsco.InventoryUpdateSmallBatchQuery{})
				if err != nil {
					logger.Warn(taskCtx, "dsco inventory update failed", "warehouse_code", dscoWarehouseCode, "err", err)
				}
			}

			offset += len(items)
			if total > 0 && offset >= total {
				break
			}
			if len(items) < ctx.Size {
				break
			}
		}

		logger.Info(taskCtx, "sync stock done", "warehouse_code", dscoWarehouseCode, "wid", wid)
	}

	return nil
}
