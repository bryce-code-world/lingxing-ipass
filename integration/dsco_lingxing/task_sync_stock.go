package dsco_lingxing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

func shortDSCOErr(err error) string {
	if err == nil {
		return ""
	}

	var apiErr *dsco.APIError
	if errors.As(err, &apiErr) {
		type apiMsg struct {
			Code        string `json:"code"`
			Severity    string `json:"severity"`
			Description any    `json:"description"`
		}
		type apiBody struct {
			Status    string   `json:"status"`
			RequestID string   `json:"requestId"`
			EventDate string   `json:"eventDate"`
			Messages  []apiMsg `json:"messages"`
		}
		var body apiBody
		if len(apiErr.Body) > 0 && json.Unmarshal(apiErr.Body, &body) == nil && len(body.Messages) > 0 {
			msg := body.Messages[0]
			desc := strings.TrimSpace(fmt.Sprintf("%v", msg.Description))
			if len(desc) > 240 {
				desc = desc[:240] + "..."
			}
			if apiErr.StatusCode == 404 && msg.Code == "notFound" {
				// 404 notFound 在库存对账里属于常见场景：领星 SKU 不存在于 DSCO inventory。
				if desc != "" {
					return "dsco notFound: " + desc
				}
				return "dsco notFound"
			}
			if msg.Code != "" && desc != "" {
				return fmt.Sprintf("dsco api %d %s: %s", apiErr.StatusCode, msg.Code, desc)
			}
			if msg.Code != "" {
				return fmt.Sprintf("dsco api %d %s", apiErr.StatusCode, msg.Code)
			}
		}
		return fmt.Sprintf("dsco api status=%d", apiErr.StatusCode)
	}

	s := strings.TrimSpace(err.Error())
	if len(s) > 300 {
		s = s[:300] + "..."
	}
	return s
}

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

	jobCfg, ok := ctx.Config.Jobs[ctx.Job]
	doSync := ok && jobCfg.Sync
	useStream := ok && jobCfg.UseStream
	if !doSync {
		logger.Info(taskCtx, "sync disabled: will only pull and record",
			append(base, "task", "sync_stock", "sync", false)...,
		)
	}

	// 初始化客户端：领星；DSCO（可选）
	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		retErr = err
		return retErr
	}

	var dscoCli *dsco.Client
	dscoCli, err = d.dscoClient()
	if err != nil {
		if doSync {
			retErr = err
			return retErr
		}
		logger.Warn(taskCtx, "dsco client init failed (sync disabled, will continue without dsco inventory compare)", "err", err)
		dscoCli = nil
	}

	type dscoInvCacheEntry struct {
		ByWarehouseCode map[string]int
		Err             string
	}
	dscoInvCache := make(map[string]dscoInvCacheEntry) // key=partnerSku
	if useStream && dscoCli != nil {
		wantedWarehouseCodes := make(map[string]struct{}, len(ctx.Config.Mapping.Warehouse))
		for code := range ctx.Config.Mapping.Warehouse {
			code = strings.TrimSpace(code)
			if code != "" {
				wantedWarehouseCodes[code] = struct{}{}
			}
		}

		streamStartedAt := time.Now().UTC()
		var pulled int
		logger.Info(taskCtx, "dsco stream pull begin",
			append(base, "task", "sync_stock", "use_stream", true)...,
		)
		err := dscoCli.Inventory.PullAll(taskCtx, dsco.InventoryFullPullOptions{
			Description:          "lingxing-ipass sync_stock dsco inventory pull",
			MaxEventsObjectCount: 1000,
			PollInterval:         2 * time.Second,
			IdleTimeout:          2 * time.Minute,
			MaxDuration:          30 * time.Minute,
		}, func(inv *dsco.ItemInventory) error {
			if inv == nil {
				return nil
			}
			sku := strings.TrimSpace(inv.SKU)
			if sku == "" {
				return nil
			}
			// 仅缓存本任务关心的仓库维度数量；如果该 SKU 在关心的仓库里没有记录（可能被 DSCO 省略为“0”），
			// 也要缓存一个空 entry，避免后续对每个 SKU 再发起 GetByKey 的逐条请求（会非常慢）。
			var m map[string]int
			for _, w := range inv.Warehouses {
				code := strings.TrimSpace(w.Code)
				if code == "" {
					continue
				}
				if _, ok := wantedWarehouseCodes[code]; !ok {
					continue
				}
				if m == nil {
					m = make(map[string]int, len(wantedWarehouseCodes))
				}
				q := 0
				if w.Quantity != nil {
					q = *w.Quantity
				}
				m[code] = q
			}
			dscoInvCache[sku] = dscoInvCacheEntry{ByWarehouseCode: m}
			if inv.PartnerSKU != nil {
				partner := strings.TrimSpace(*inv.PartnerSKU)
				if partner != "" {
					if _, ok := dscoInvCache[partner]; !ok {
						dscoInvCache[partner] = dscoInvCacheEntry{ByWarehouseCode: m}
					}
				}
			}
			pulled++
			if pulled%5000 == 0 {
				logger.Info(taskCtx, "dsco stream pull progress",
					append(base, "task", "sync_stock", "count", pulled)...,
				)
			}
			return nil
		})
		if err != nil {
			logger.Warn(taskCtx, "dsco stream pull failed (will fallback to get inventory by key)", "err", err)
			dscoInvCache = make(map[string]dscoInvCacheEntry)
		} else {
			logger.Info(taskCtx, "dsco stream pull end",
				append(base,
					"task", "sync_stock",
					"duration_ms", time.Since(streamStartedAt).Milliseconds(),
					"cached_skus", len(dscoInvCache),
				)...,
			)
		}
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
		var (
			dscoCacheHit        int
			dscoCacheMiss       int
			dscoFallbackFetch   int
			dscoFallbackFetchMs int64
		)

		// 拉取领星库存明细可以用更大的分页 size，显著减少接口请求次数。
		// 同步回写 DSCO 时仍然按 ctx.Size 分批，避免单次更新过大。
		pullLength := ctx.Size
		if pullLength < 200 {
			pullLength = 200
		}
		offset := 0
		for {
			// 分页拉取领星库存明细
			detailsStart := time.Now()
			items, total, err := lx.Inventory.InventoryDetails(taskCtx, lingxing.InventoryDetailsRequest{
				WID:    wid,
				Offset: offset,
				Length: pullLength,
			})
			logger.Info(taskCtx, "lingxing inventoryDetails done",
				append(base,
					"task", "sync_stock",
					"warehouse_code", dscoWarehouseCode,
					"wid", wid,
					"offset", offset,
					"length", pullLength,
					"got", len(items),
					"total", total,
					"duration_ms", time.Since(detailsStart).Milliseconds(),
				)...,
			)
			if err != nil {
				logger.Warn(taskCtx, "lingxing inventoryDetails failed", "wid", wid, "err", err)
				break
			}
			if len(items) == 0 {
				break
			}
			totalItems += len(items)

			var invs []dsco.ItemInventory
			var rows []store.DSCOWarehouseSyncRow
			for _, it := range items {
				// 领星 SKU -> DSCO partnerSku（优先反向映射；缺省同名直传）
				partner := strings.TrimSpace(reverseSKU[it.SKU])
				if partner == "" {
					partner = strings.TrimSpace(it.SKU)
				}
				if partner == "" {
					continue
				}

				lxQty := it.ProductValidNum

				dscoQty := 0
				rowStatus := int16(1)
				rowReason := ""
				if dscoCli == nil {
					rowStatus = 2
					rowReason = "dsco client unavailable"
				} else {
					if cached, ok := dscoInvCache[partner]; ok {
						dscoCacheHit++
						if cached.Err != "" {
							rowStatus = 2
							rowReason = cached.Err
						} else if cached.ByWarehouseCode != nil {
							dscoQty = cached.ByWarehouseCode[dscoWarehouseCode]
						}
					} else {
						dscoCacheMiss++

						// Stream 模式下 cache miss 很可能意味着 stream pull 不完整或 SKU mapping 配置不一致。
						// 逐条 fallback GET 会导致任务整体非常慢；这里只做少量采样请求用于排错，其余直接记录缺失。
						if useStream {
							if dscoCacheMiss > 5 {
								rowStatus = 2
								rowReason = "dsco inventory missing in stream (skip fallback get)"
								dscoInvCache[partner] = dscoInvCacheEntry{Err: rowReason}
								// 继续写入记录（dscoQty 按 0 处理）
								goto doneDSCOQty
							}
						}

						// 统计 fallback GET 的耗时（用于定位慢点）
						fetchOne := func(itemKey string) (*dsco.ItemInventory, error) {
							started := time.Now()
							inv, err := dscoCli.Inventory.GetByKey(taskCtx, dsco.InventoryGetQuery{
								ItemKey: itemKey,
								Value:   partner,
							})
							dscoFallbackFetch++
							dscoFallbackFetchMs += time.Since(started).Milliseconds()
							return inv, err
						}

						inv, invErr := fetchOne("partnerSku")
						if invErr != nil {
							inv, invErr = fetchOne("sku")
						}
						if invErr != nil {
							rowReason = shortDSCOErr(invErr)
							dscoInvCache[partner] = dscoInvCacheEntry{Err: rowReason}
							rowStatus = 2
						} else {
							m := make(map[string]int, len(inv.Warehouses))
							for _, w := range inv.Warehouses {
								code := strings.TrimSpace(w.Code)
								if code == "" {
									continue
								}
								q := 0
								if w.Quantity != nil {
									q = *w.Quantity
								}
								m[code] = q
							}
							dscoInvCache[partner] = dscoInvCacheEntry{ByWarehouseCode: m}
							dscoQty = m[dscoWarehouseCode]
						}
					}
				}
			doneDSCOQty:
				diff := lxQty - dscoQty

				if doSync {
					partnerCopy := partner
					invs = append(invs, dsco.ItemInventory{
						Item: dsco.Item{
							SKU:        partner,
							PartnerSKU: &partnerCopy,
						},
						Warehouses: []dsco.ItemWarehouse{
							{Code: dscoWarehouseCode, Quantity: &lxQty},
						},
						QuantityAvailable: &lxQty,
					})
				}

				rows = append(rows, store.DSCOWarehouseSyncRow{
					SyncTime:             time.Now().UTC().Unix(),
					DSCOWarehouseID:      dscoWarehouseCode,
					DSCOWarehouseSKU:     partner,
					DSCOWarehouseNum:     dscoQty,
					LingXingWarehouseID:  wid,
					LingXingWarehouseSKU: it.SKU,
					LingXingWarehouseNum: lxQty,
					Diff:                 diff,
					Status:               rowStatus,
					Reason:               rowReason,
				})
			}
			totalInvs += len(invs)

			var syncErr error
			if doSync && len(invs) > 0 {
				// 调用 DSCO 批量库存回写接口（按 ctx.Size 分批，减少失败概率与单次请求体大小）。
				batchSize := ctx.Size
				if batchSize <= 0 {
					batchSize = 50
				}
				for start := 0; start < len(invs); start += batchSize {
					end := start + batchSize
					if end > len(invs) {
						end = len(invs)
					}
					batch := invs[start:end]

					_, syncErr = dscoCli.Inventory.UpdateSmallBatch(taskCtx, batch, dsco.InventoryUpdateSmallBatchQuery{})
					if syncErr != nil {
						logger.Warn(taskCtx, "dsco inventory update failed", "warehouse_code", dscoWarehouseCode, "err", syncErr)
						break
					}
					logger.Info(taskCtx, "inventory batch sent",
						append(base,
							"task", "sync_stock",
							"warehouse_code", dscoWarehouseCode,
							"wid", wid,
							"offset", offset,
							"count", len(batch),
						)...,
					)
				}
			}

			// 写入同步记录（用于 Admin 查询/导出）。
			// 这里必须批量插入，否则会触发大量单条 INSERT，导致整体耗时明显变长，并出现 gorm slow sql。
			for i := range rows {
				if doSync && syncErr != nil && rows[i].Status == 1 {
					rows[i].Status = 2
					rows[i].Reason = shortDSCOErr(syncErr)
				}
			}
			insertStart := time.Now()
			if err := d.warehouseStore.InsertBatch(taskCtx, rows); err != nil {
				logger.Warn(taskCtx, "warehouse sync rows insert failed", "warehouse_code", dscoWarehouseCode, "wid", wid, "err", err)
			} else {
				logger.Info(taskCtx, "warehouse sync rows inserted",
					append(base,
						"task", "sync_stock",
						"warehouse_code", dscoWarehouseCode,
						"wid", wid,
						"count", len(rows),
						"duration_ms", time.Since(insertStart).Milliseconds(),
					)...,
				)
			}
			logger.Info(taskCtx, "dsco inventory compare stats",
				append(base,
					"task", "sync_stock",
					"warehouse_code", dscoWarehouseCode,
					"wid", wid,
					"cache_hit", dscoCacheHit,
					"cache_miss", dscoCacheMiss,
					"fallback_get", dscoFallbackFetch,
					"fallback_get_ms", dscoFallbackFetchMs,
				)...,
			)

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
