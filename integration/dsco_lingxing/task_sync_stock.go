package dsco_lingxing

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"gitee.com/lsy007/golibv2/v2/sdk/dsco"
	"gitee.com/lsy007/golibv2/v2/sdk/lingxing"
	"gitee.com/lsy007/golibv2/v2/tool/logger"

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
//     - DSCO Item.SKU 采用 DSCO sku 主口径（必要时兼容 partnerSku）。
//     - SKU 映射：领星 SKU -> DSCO sku，优先使用 mapping.sku 的反向映射；缺省同名直传。
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

	jobCfg, ok := ctx.Config.Jobs[ctx.Job]
	doSync := ok && jobCfg.Sync
	useStream := ok && jobCfg.UseStream

	var ov SyncStockOverride
	var hasOv bool
	switch v := ctx.Override.(type) {
	case SyncStockOverride:
		ov = v
		hasOv = true
	case *SyncStockOverride:
		if v != nil {
			ov = *v
			hasOv = true
		}
	}
	if hasOv && ov.ForceDoSync != nil {
		doSync = *ov.ForceDoSync
	}

	if hasOv && (ov.Source == SyncStockSourceDailyPulled || ov.Source == SyncStockSourceManualItems) {
		retErr = d.syncStockByOverride(ctx, base, doSync, ov)
		return retErr
	}

	if !doSync {
		logger.Info(taskCtx, "sync disabled: will only pull and record",
			append(base, "task", "sync_stock", "sync", false)...,
		)
	}

	reverseSKU, err := buildReverseSKUMap(ctx.Config)
	if err != nil {
		retErr = err
		return retErr
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
	dscoInvCache := make(map[string]dscoInvCacheEntry) // key=dsco sku
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
				// 领星 SKU -> DSCO sku（优先反向映射；缺省同名直传）
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
					status := dscoInventoryStatusForQty(lxQty)
					invs = append(invs, dsco.ItemInventory{
						Item: dsco.Item{
							SKU:        partner,
							PartnerSKU: &partnerCopy,
						},
						Status: status,
						Warehouses: []dsco.ItemWarehouse{
							{Code: dscoWarehouseCode, Quantity: &lxQty, Status: status},
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

func dscoInventoryStatusForQty(qty int) string {
	if qty <= 0 {
		return "out-of-stock"
	}
	return "in-stock"
}

func (d *Domain) syncStockByOverride(ctx integration.TaskContext, base []any, doSync bool, ov SyncStockOverride) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	source := string(ov.Source)
	if source == "" {
		source = "unknown"
	}

	maxKeys := ov.MaxKeys
	if maxKeys <= 0 {
		maxKeys = 5000
	}

	type updateKey struct {
		WID string
		SKU string
	}
	target := make(map[updateKey]int) // dsco (warehouse_id + sku) -> qty

	switch ov.Source {
	case SyncStockSourceManualItems:
		for i, it := range ov.ManualItems {
			if strings.TrimSpace(it.DSCOWarehouseID) == "" || strings.TrimSpace(it.DSCOSKU) == "" {
				return fmt.Errorf("manual_items[%d] missing dsco_wid/dsco_sku", i)
			}
			if it.Qty < 0 {
				return fmt.Errorf("manual_items[%d] qty must be >= 0", i)
			}
			k := updateKey{WID: strings.TrimSpace(it.DSCOWarehouseID), SKU: strings.TrimSpace(it.DSCOSKU)}
			target[k] = it.Qty
		}
	case SyncStockSourceDailyPulled:
		if ov.DailyPulled == nil {
			return errors.New("daily_pulled options is nil")
		}
		opt := *ov.DailyPulled
		if opt.StartTime <= 0 || opt.EndTime <= 0 || opt.EndTime <= opt.StartTime {
			return errors.New("daily_pulled invalid start/end time")
		}

		filter := store.DSCOWarehouseSyncListFilter{
			StartTime: &opt.StartTime,
			EndTime:   &opt.EndTime,

			DSCOWarehouseID:     strings.TrimSpace(opt.DSCOWarehouseID),
			LingXingWarehouseID: strings.TrimSpace(opt.LingXingWarehouseID),

			DSCOWarehouseSKUIn:     opt.DSCOSKUList,
			LingXingWarehouseSKUIn: opt.LingXingSKUList,

			DiffNotZero: opt.DiffOnly,

			Offset: 0,
			Limit:  500,
		}

		var offset int
		for {
			filter.Offset = offset
			items, total, err := d.warehouseStore.ListLatestByDSCOKey(taskCtx, filter)
			if err != nil {
				return err
			}
			if total > int64(maxKeys) {
				return fmt.Errorf("%w: %d (max=%d)", ErrSyncStockTooManyKeys, total, maxKeys)
			}
			for _, row := range items {
				k := updateKey{WID: row.DSCOWarehouseID, SKU: row.DSCOWarehouseSKU}
				target[k] = row.LingXingWarehouseNum
			}
			offset += len(items)
			if len(items) == 0 || offset >= int(total) {
				break
			}
		}
	default:
		return fmt.Errorf("unsupported override source: %s", ov.Source)
	}

	logger.Info(taskCtx, "sync_stock override selected",
		append(base,
			"task", "sync_stock",
			"source", source,
			"dry_run", !doSync,
			"keys", len(target),
		)...,
	)

	if !doSync {
		if ctx.Report != nil {
			ctx.Report(map[string]any{
				"run_id":   ctx.RunID,
				"job":      "sync_stock",
				"dry_run":  true,
				"source":   source,
				"keys":     len(target),
				"max_keys": maxKeys,
			})
		}
		return nil
	}

	dscoCli, err := d.dscoClient()
	if err != nil {
		return err
	}

	batchSize := ctx.Size
	if batchSize <= 0 {
		batchSize = 50
	}

	var dscoRequestIDs []string

	var invs []dsco.ItemInventory
	invs = make([]dsco.ItemInventory, 0, len(target))
	for k, qty := range target {
		partner := k.SKU
		qtyCopy := qty
		status := dscoInventoryStatusForQty(qtyCopy)
		invs = append(invs, dsco.ItemInventory{
			Item: dsco.Item{
				SKU: partner,
			},
			Status:            status,
			Warehouses:        []dsco.ItemWarehouse{{Code: k.WID, Quantity: &qtyCopy}},
			QuantityAvailable: &qtyCopy,
		})
		invs[len(invs)-1].Warehouses[0].Status = status
	}

	for start := 0; start < len(invs); start += batchSize {
		end := start + batchSize
		if end > len(invs) {
			end = len(invs)
		}
		batch := invs[start:end]
		resp, err := dscoCli.Inventory.UpdateSmallBatch(taskCtx, batch, dsco.InventoryUpdateSmallBatchQuery{})
		if err != nil {
			return err
		}
		if resp != nil {
			if strings.TrimSpace(resp.RequestID) != "" {
				dscoRequestIDs = append(dscoRequestIDs, strings.TrimSpace(resp.RequestID))
			}
			logger.Info(taskCtx, "dsco inventory update accepted",
				append(base,
					"task", "sync_stock",
					"source", source,
					"dsco_status", resp.Status,
					"dsco_request_id", resp.RequestID,
					"dsco_event_date", resp.EventDate,
				)...,
			)
		}
		logger.Info(taskCtx, "sync_stock override batch sent",
			append(base,
				"task", "sync_stock",
				"source", source,
				"offset", start,
				"count", len(batch),
			)...,
		)
	}

	if ctx.Report != nil {
		// 注意：UpdateSmallBatch 为异步接口；这里只能代表“请求被 DSCO 接受并返回 requestId”，
		// 最终处理结果需通过 DSCO stream/change log 跟踪（如有需要）。
		ctx.Report(map[string]any{
			"run_id":           ctx.RunID,
			"job":              "sync_stock",
			"dry_run":          false,
			"source":           source,
			"keys":             len(target),
			"max_keys":         maxKeys,
			"dsco_request_ids": dscoRequestIDs,
		})
	}

	return nil
}
