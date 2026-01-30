package dsco_lingxing

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/tool/logger"

	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

// PullDSCOOrders 拉取 DSCO 订单并入库到 dsco_order_sync。
//
// 一期口径：
//  1. 计算拉取时间窗口：
//     - 手动触发：使用 Admin 传入的 [start,end)（UTC 秒级）范围；入库 status 由 DSCO 的 dsco_status 自动推导。
//     - 定时触发：使用 dsco_order_sync.dsco_create_time 的最大值作为游标（增量拉取）。
//     若表为空，则从 2025-01-01 00:00:00（UTC）开始往后拉取。
//  2. 调用 DSCO Order.GetPageRaw 分页拉取（scrollId）。
//  3. 对每条订单：
//     - dsco_create_time：仅使用 dscoCreateDate（已确认），解析为 UTC 秒级时间戳。
//     - status：根据 dsco_status 自动推导（created/shipment_pending/shipped/cancelled）。
//     - mskus：提取行项目 partnerSku/sku，用于列表筛选与 CSV 导出。
//     - warehouse_id：使用 requestedWarehouseCode（MVP 口径）。
//     - shipment：使用 shippingServiceLevelCode（已确认），用于列表筛选与 CSV 导出。
//     - dsco_retailer_id：写入 dscoRetailerId，用于 mapping.shop（店铺/渠道映射）。
//  4. Upsert 入库：允许覆盖 payload/status（用于初始化补数据、人工纠错）。
func (d *Domain) PullDSCOOrders(ctx integration.TaskContext) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	// 1) 初始化 DSCO 客户端
	cli, err := d.dscoClient()
	if err != nil {
		return err
	}

	var (
		since time.Time
		until time.Time
	)

	if ctx.Trigger == integration.TriggerManual && ctx.Override != nil {
		// 1.1) 手动触发：从 Override 中解析时间范围
		ov, ok := ctx.Override.(integration.PullDSCOOrdersOverride)
		if !ok {
			if p, ok2 := ctx.Override.(*integration.PullDSCOOrdersOverride); ok2 && p != nil {
				ov = *p
				ok = true
			}
		}
		if ok {
			since = time.Unix(ov.Start, 0).UTC()
			until = time.Unix(ov.End, 0).UTC()
		}
	}

	// Default: incremental cursor from dsco_order_sync.dsco_create_time.
	if since.IsZero() || until.IsZero() {
		// 1.2) 定时触发：用 dsco_create_time 最大值作为游标；表空则使用固定起点
		maxTime, ok, err := d.orderStore.GetMaxDSCOCreateTime(taskCtx)
		if err != nil {
			return err
		}
		if ok {
			since = time.Unix(maxTime, 0).UTC()
		} else {
			since = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		}
		until = time.Now().UTC().Add(-10 * time.Second)
	}

	// 2) 组装 DSCO 分页查询：按 created_at（dscoCreateDate）区间拉取
	q := dsco.OrderPageQuery{
		OrdersCreatedSince: since.Format(time.RFC3339),
		Until:              until.Format(time.RFC3339),
		OrdersPerPage:      ctx.Size,
	}

	var pulled int
	var scroll string
	for {
		// 2.1) scrollId 分页：DSCO 返回 scrollId 表示下一页
		if scroll != "" {
			q.ScrollID = scroll
		}
		resp, err := cli.Order.GetPageRaw(taskCtx, q)
		if err != nil {
			return err
		}
		scroll = resp.ScrollID
		if len(resp.Orders) == 0 {
			break
		}

		for _, raw := range resp.Orders {
			// 3) 解析订单 payload（保留原始 JSON，便于审计/排查）
			order, err := decodeDSCOOrder(raw)
			if err != nil {
				logger.Warn(taskCtx, "decode dsco order failed", "err", err)
				continue
			}

			// 3.1) dsco_create_time：仅使用 dscoCreateDate（RFC3339 -> UTC 秒）
			createStr := derefString(order.DscoCreateDate)
			createUnix, err := parseRFC3339ToUnixSec(createStr)
			if err != nil {
				logger.Warn(taskCtx, "parse dsco create time failed", "err", err)
				continue
			}

			// 手动拉单范围边界：[start,end)
			if ctx.Trigger == integration.TriggerManual && !since.IsZero() && !until.IsZero() {
				if createUnix < since.Unix() || createUnix >= until.Unix() {
					continue
				}
			}

			// 3.2) mskus：提取 partnerSku/sku（用于列表筛选、CSV 导出）
			mskus := make([]string, 0, len(order.LineItems))
			for _, li := range order.LineItems {
				if li.PartnerSKU != nil && *li.PartnerSKU != "" {
					mskus = append(mskus, *li.PartnerSKU)
				} else if li.SKU != nil && *li.SKU != "" {
					mskus = append(mskus, *li.SKU)
				}
			}

			// 3.3) 入库状态：优先根据 DSCO dsco_status 推导；未知状态则默认 1
			dscoStatus := strings.TrimSpace(order.DscoStatus)
			rowStatus := int16(1)
			if dscoStatus != "" {
				if mapped, ok := dscoStatusToSyncStatus(dscoStatus); ok {
					rowStatus = mapped
				}
			}

			// 3.4) 入库行：只写入一期需要的字段，其它信息存入 payload
			row := store.DSCOOrderSyncRow{
				PONumber:       order.PoNumber,
				DSCOCreateTime: createUnix,
				DSCOREtailerID: derefString(order.DscoRetailerID),
				DSCOStatus:     dscoStatus,
				Status:         rowStatus,
				Payload:        json.RawMessage(raw),
				MSKUs:          store.PGTextArray(mskus),
				WarehouseID:    getDSCOWarehouseCode(order),
				Shipment:       getDSCOShippingServiceLevelCode(order),
			}

			// 4) Upsert：允许覆盖更新（初始化补数据、人工纠错）
			if err := d.orderStore.Upsert(taskCtx, row); err != nil {
				logger.Warn(taskCtx, "upsert dsco_order_sync failed", "err", err)
				continue
			}
			pulled++
		}

		// Safety break: if scrollId empty, DSCO has no more pages.
		if scroll == "" {
			break
		}
	}

	logger.Info(taskCtx, "pull dsco orders done", "count", pulled)
	return nil
}

func derefString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
