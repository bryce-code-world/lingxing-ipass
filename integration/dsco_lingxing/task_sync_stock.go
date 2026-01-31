package dsco_lingxing

import (
	"context"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

// SyncStock 将领星库存回写到 DSCO（Inventory UpdateSmallBatch）。
//
// 一期口径：
//  1. 遍历 mapping.warehouse（DSCO warehouseCode -> 领星 WID）。
//  2. 对每个 WID 分页调用领星 InventoryDetails 拉取库存明细。
//  3. 将领星库存条目转换为 DSCO ItemInventory 批量回写：
//     - DSCO Item.SKU 与 Item.PartnerSKU 使用同一口径（partnerSku 口径；一期已确认先不区分）。
//     - SKU 映射：领星 SKU -> DSCO partnerSku，优先使用 mapping.sku 的反向映射；缺省同名直传。
//     - 数量使用 领星 ProductValidNum（可用库存）。
//  4. 写入 dsco_warehouse_sync 作为“同步记录”（一期 reason 字段默认不写）。
//
// 关键点：
// - 映射匹配不到一律跳过（warehouseCode/WID 为空直接跳过；SKU 为空直接跳过）。
// - 一期不并发；失败仅日志，等待下次定时任务重试。
func (d *Domain) SyncStock(ctx integration.TaskContext) (retErr error) {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	startedAt := time.Now().UTC()
	base := ctx.BaseLogFields()
	logger.Info(taskCtx, "task begin", append(base, "task", "sync_stock")...)
	defer func() {
		fields := append(base,
			"task", "sync_stock",
			"duration_ms", time.Since(startedAt).Milliseconds(),
		)
		if retErr != nil {
			logger.Error(taskCtx, "task end", append(fields, "result", "failed", "err", retErr)...)
			return
		}
		logger.Info(taskCtx, "task end", append(fields, "result", "ok")...)
	}()

	reverseSKU, err := buildReverseSKUMap(ctx.Config)
	if err != nil {
		retErr = err
		return retErr
	}

	// 初始化客户端：领星 + DSCO
	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		retErr = err
		return retErr
	}
	dscoCli, err := d.dscoClient()
	if err != nil {
		retErr = err
		return retErr
	}

	// mapping.warehouse: DSCO warehouseCode -> 领星 WID
	for dscoWarehouseCode, wid := range ctx.Config.Mapping.Warehouse {
		dscoWarehouseCode = strings.TrimSpace(dscoWarehouseCode)
		wid = strings.TrimSpace(wid)
		if dscoWarehouseCode == "" || wid == "" {
			continue
		}

		warehouseStart := time.Now().UTC()
		logger.Info(taskCtx, "warehouse sync begin",
			append(base,
				"task", "sync_stock",
				"warehouse_code", dscoWarehouseCode,
				"wid", wid,
			)...,
		)

		var (
			totalItems int
			totalInvs  int
		)
		offset := 0
		for {
			// 分页拉取领星库存明细
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
			totalItems += len(items)

			var invs []dsco.ItemInventory
			for _, it := range items {
				// 领星 SKU -> DSCO partnerSku（优先反向映射；缺省同名直传）
				partner := strings.TrimSpace(reverseSKU[it.SKU])
				if partner == "" {
					partner = strings.TrimSpace(it.SKU)
				}
				if partner == "" {
					continue
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

				// 写入同步记录（用于 Admin 查询/导出）
				_ = d.warehouseStore.Insert(taskCtx, store.DSCOWarehouseSyncRow{
					SyncTime:             time.Now().UTC().Unix(),
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
			totalInvs += len(invs)

			if len(invs) > 0 {
				// 调用 DSCO 批量库存回写接口
				_, err := dscoCli.Inventory.UpdateSmallBatch(taskCtx, invs, dsco.InventoryUpdateSmallBatchQuery{})
				if err != nil {
					logger.Warn(taskCtx, "dsco inventory update failed", "warehouse_code", dscoWarehouseCode, "err", err)
				}
				logger.Info(taskCtx, "inventory batch sent",
					append(base,
						"task", "sync_stock",
						"warehouse_code", dscoWarehouseCode,
						"wid", wid,
						"offset", offset,
						"count", len(invs),
						"req", integration.JSONForLog(invs),
					)...,
				)
			}

			offset += len(items)
			if total > 0 && offset >= total {
				break
			}
			if len(items) < ctx.Size {
				break
			}
		}

		logger.Info(taskCtx, "warehouse sync end",
			append(base,
				"task", "sync_stock",
				"warehouse_code", dscoWarehouseCode,
				"wid", wid,
				"duration_ms", time.Since(warehouseStart).Milliseconds(),
				"lingxing_items", totalItems,
				"dsco_invs", totalInvs,
			)...,
		)
	}

	return nil
}
