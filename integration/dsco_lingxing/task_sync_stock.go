package dsco_lingxing

import (
	"context"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

// SyncStock 将领星库存回写到 DSCO（Inventory UpdateSmallBatch）。
//
// 流程（一期口径）：
//  1. 遍历 mapping.warehouse：key=DSCO warehouseCode，value=领星 WID。
//  2. 对每个 WID，分页调用领星 InventoryDetails 拉取库存明细。
//  3. 将领星库存条目转换为 DSCO ItemInventory 批量回写：
//     - DSCO Item.SKU 与 Item.PartnerSKU 当前使用同一口径（partnerSku 口径），已确认先不区分。
//     - 数量使用领星 ProductValidNum（可用库存）。
//  4. 写入 dsco_warehouse_sync 作为“同步记录”（一期 reason 字段默认不写）。
//
// 关键点：
// - 映射匹配不到一律跳过（warehouseCode/WID 为空直接跳过）。
// - 一期不做并发；失败仅日志，等待下次定时任务重试。
func (d *Domain) SyncStock(ctx integration.TaskContext) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	// 1) 初始化客户端：领星 + DSCO
	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		return err
	}
	dscoCli, err := d.dscoClient()
	if err != nil {
		return err
	}

	// mapping.warehouse: DSCO warehouseCode -> LingXing WID
	for dscoWarehouseCode, wid := range ctx.Config.Mapping.Warehouse {
		// 2) 遍历仓库映射：缺少任意一边则跳过（大原则：映射缺失都跳过）
		if dscoWarehouseCode == "" || wid == "" {
			continue
		}

		offset := 0
		for {
			// 3) 分页拉取领星库存明细
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

			// 4) 组装 DSCO 批量库存回写
			var invs []dsco.ItemInventory
			for _, it := range items {
				partner := it.SKU
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

				// 5) 写入同步记录（用于 admin 查询/导出）
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
				// 6) 调用 DSCO 批量库存回写接口
				_, err := dscoCli.Inventory.UpdateSmallBatch(taskCtx, invs, dsco.InventoryUpdateSmallBatchQuery{})
				if err != nil {
					logger.Warn(taskCtx, "dsco inventory update failed", "warehouse_code", dscoWarehouseCode, "err", err)
				}
			}

			// 7) 翻页与退出条件
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
