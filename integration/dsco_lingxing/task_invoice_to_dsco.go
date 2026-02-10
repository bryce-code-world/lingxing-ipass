package dsco_lingxing

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"sort"
	"strings"
	"time"

	"gitee.com/lsy007/golibv2/v2/sdk/dsco"
	"gitee.com/lsy007/golibv2/v2/sdk/lingxing"
	"gitee.com/lsy007/golibv2/v2/tool/logger"

	"lingxingipass/infra/store"
	"lingxingipass/integration"
)

// InvoiceToDSCO 将领星侧“实际出库”的发票信息回传给 DSCO（Invoice CreateSmallBatch）。
//
// 状态机（一期口径）：
// - 本任务处理 dsco_order_sync.status = 4 的订单（已回传发货，待回传发票）。
// - 回传 invoice 成功后将状态推进到 5（完成态）。
//
// 关键点：
// - 幂等：若本地 dsco_invoice_id 已有值，则认为该 poNumber 已回传过发票，直接跳过。
// - 发票数据来源：
//   - 数量：优先使用 WmsOrderList 的 ProductInfo（以实际出库为准）。
//   - 单价：优先从 DSCO 原始订单行项目中挑选（见 pickUnitPrice），以 DSCO 订单字段为准。
//   - 订单总金额：按“数量 * 单价”汇总，并做 2 位小数 round。
//
// - invoiceDate：优先 DeliveredAt；为空则用 StockDeliveredAt 兜底（已确认）。
// - 发票回传口径（一期最新确认）：
//   - 以 poNumber 为维度：同一 poNumber 即使对应领星拆单/多行，也只回传“1 张汇总发票”给 DSCO。
//   - 前置条件：poNumber 下所有 SKU 均已发货（以 WmsOrderList 聚合数量 >= DSCO 原始订单数量 为准）。
//   - invoiceId：固定 invoiceId = poNumber。
//   - tracking：DSCO 发票接口不接受顶层 trackingNumber 字段；本任务仅在日志中记录运单号（若多个用逗号拼接），tracking 由 `ship_to_dsco` 回传。
func (d *Domain) InvoiceToDSCO(ctx integration.TaskContext) (retErr error) {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	startedAt := time.Now().UTC()
	base := ctx.BaseLogFields()
	logger.Info(taskCtx, "task begin", append(base, "task", "invoice_to_dsco")...)

	var (
		total   int
		okCount int
		skip    int
		fail    int
	)
	defer func() {
		fields := append(base,
			"task", "invoice_to_dsco",
			"duration_ms", time.Since(startedAt).Milliseconds(),
			"total", total,
			"ok", okCount,
			"skip", skip,
			"fail", fail,
		)
		if retErr != nil {
			logger.Error(taskCtx, "task end", append(fields, "result", "failed", "err", retErr)...)
			return
		}
		logger.Info(taskCtx, "task end", append(fields, "result", "ok")...)
	}()

	// 1) 取待回传发票（status=4）
	var items []store.DSCOOrderSyncRow
	var err error
	if strings.TrimSpace(ctx.OnlyPONumber) != "" {
		items, err = d.orderStore.FindByStatusAndPONumber(taskCtx, 4, ctx.OnlyPONumber)
	} else {
		items, err = d.orderStore.FindByStatus(taskCtx, 4, ctx.Size)
	}
	if err != nil {
		retErr = err
		return retErr
	}
	if len(items) == 0 {
		return nil
	}
	total = len(items)

	// 2) 初始化客户端：DSCO + 领星
	dscoCli, err := d.dscoClient()
	if err != nil {
		retErr = err
		return retErr
	}
	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		retErr = err
		return retErr
	}

	// SKU 反向映射：领星 SKU -> DSCO partnerSku（mapping.sku 的反向映射；缺省同名直传）。
	reverseSKU, err := buildReverseSKUMap(ctx.Config)
	if err != nil {
		retErr = err
		return retErr
	}

	multiBan := false
	if jc, ok := ctx.Config.Jobs[ctx.Job]; ok {
		multiBan = jc.MultiBan
	}
	// 手动单号测试（OnlyPONumber 非空）允许通过：此时是人工验证单个订单，不做 multi_ban 限制。
	if strings.TrimSpace(ctx.OnlyPONumber) != "" {
		multiBan = false
	}

	var invs []dsco.Invoice
	var toUpdate []struct {
		po        string
		invoiceID string
	}

	for _, row := range items {
		po := strings.TrimSpace(row.PONumber)
		if po == "" {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", "",
					"result", "skip",
					"reason", "po_number_empty",
				)...,
			)
			continue
		}

		// 0) 幂等：
		// - 默认：已写入 dsco_invoice_id 的订单视为已回传过发票，直接跳过。
		// - 手动指定 po_number（OnlyPONumber 非空）时属于人工重试/验证：即使本地已有 invoice_id 也继续执行。
		// - 兜底：若本地已有 invoice_id，但 DSCO 侧实际查不到该 invoice，则不跳过（避免“本地标记成功但 DSCO 无数据”导致永久卡死）。
		if shouldSkipInvoiceBecauseAlreadyHasInvoiceID(ctx.OnlyPONumber, row.DSCOInvoiceID) {
			exists, exErr := dscoInvoiceExists(taskCtx, dscoCli, row.DSCOInvoiceID)
			if exErr != nil {
				logger.Warn(taskCtx, "order note",
					append(base,
						"task", "invoice_to_dsco",
						"po_number", po,
						"reason", "check_dsco_invoice_exists_failed_before_skip",
						"invoice_id", row.DSCOInvoiceID,
						"err", exErr,
					)...,
				)
				// 查询失败时继续执行回传逻辑：避免本地标记与 DSCO 真实状态不一致。
			} else if !exists {
				logger.Warn(taskCtx, "order note",
					append(base,
						"task", "invoice_to_dsco",
						"po_number", po,
						"reason", "local_has_invoice_id_but_dsco_missing",
						"invoice_id", row.DSCOInvoiceID,
					)...,
				)
				// DSCO 侧查不到，继续执行回传逻辑。
			} else {
				skip++
				logger.Info(taskCtx, "order done",
					append(base,
						"task", "invoice_to_dsco",
						"po_number", po,
						"result", "skip",
						"reason", "already_has_invoice_id",
						"invoice_id", row.DSCOInvoiceID,
					)...,
				)
				continue
			}
		}
		if strings.TrimSpace(ctx.OnlyPONumber) != "" && strings.TrimSpace(row.DSCOInvoiceID) != "" {
			logger.Info(taskCtx, "order note",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", po,
					"reason", "manual_force_with_existing_invoice_id",
					"invoice_id", row.DSCOInvoiceID,
				)...,
			)
		}

		if multiBan {
			isMulti, mi := detectMultiOrderByMSKUs([]string(row.MSKUs))
			if isMulti {
				skip++
				logger.Info(taskCtx, "order done",
					append(base,
						"task", "invoice_to_dsco",
						"po_number", po,
						"result", "skip",
						"reason", "multi_banned",
						"multi_ban", true,
						"mskus", []string(row.MSKUs),
						"multi_info", integration.JSONForLog(mi),
						"dsco_raw", integration.JSONForLog(row.Payload),
					)...,
				)
				continue
			}
		}

		// 3) 解析 DSCO 原始订单（用于 dscoItemId/币种/参考字段等）
		var dscoOrder dsco.Order
		if err := json.Unmarshal(row.Payload, &dscoOrder); err != nil {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "fail",
					"reason", "decode_dsco_payload_failed",
					"dsco_raw", integration.JSONForLog(row.Payload),
					"err", err,
				)...,
			)
			continue
		}

		// 3.1) sid：用于 WmsOrderList 精准查询（从 mapping.shop 解析）
		sid, ok := lingxingSIDFromMapping(ctx.Config, dscoOrder)
		if !ok {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "missing_mapping_shop_sid",
					"dsco_raw", integration.JSONForLog(row.Payload),
				)...,
			)
			continue
		}

		wmsOrders, _, err := lx.Warehouse.WmsOrderList(taskCtx, lingxing.WmsOrderListRequest{
			Page:               1,
			PageSize:           200,
			SIDArr:             []int{sid},
			PlatformOrderNoArr: []string{po},
		})
		if err != nil || len(wmsOrders) == 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "lingxing_wms_order_not_found",
					"err", err,
				)...,
			)
			continue
		}

		// 4) DSCO 行项目口径（以 partnerSku 为主键）：
		// - expectedQty：用于判断“是否全部已发货”
		// - dscoItemId/lineNumber：用于发票行项目（尽量补齐以提升 DSCO 侧匹配成功率）
		expectedQtyByPartner := map[string]int{}
		dscoItemIDByPartner := map[string]string{}
		dscoItemIDConflict := map[string]bool{}
		lineNumberByPartner := map[string]int{}
		lineNumberConflict := map[string]bool{}
		for _, li := range dscoOrder.LineItems {
			partner := ""
			if li.PartnerSKU != nil && strings.TrimSpace(*li.PartnerSKU) != "" {
				partner = strings.TrimSpace(*li.PartnerSKU)
			} else if li.SKU != nil && strings.TrimSpace(*li.SKU) != "" {
				partner = strings.TrimSpace(*li.SKU)
			}
			if partner == "" {
				continue
			}
			if li.Quantity > 0 {
				expectedQtyByPartner[partner] += li.Quantity
			}
			if li.DscoItemID != nil && strings.TrimSpace(*li.DscoItemID) != "" {
				id := strings.TrimSpace(*li.DscoItemID)
				if prev, ok := dscoItemIDByPartner[partner]; ok && prev != id {
					dscoItemIDConflict[partner] = true
				} else {
					dscoItemIDByPartner[partner] = id
				}
			}
			if li.LineNumber != nil && *li.LineNumber > 0 {
				if prev, ok := lineNumberByPartner[partner]; ok && prev != *li.LineNumber {
					lineNumberConflict[partner] = true
				}
				lineNumberByPartner[partner] = *li.LineNumber
			}
		}

		priceByPartner := map[string]float64{}
		for _, li := range dscoOrder.LineItems {
			p := ""
			if li.PartnerSKU != nil {
				p = *li.PartnerSKU
			} else if li.SKU != nil {
				p = *li.SKU
			}
			if p == "" {
				continue
			}
			if price, ok := pickUnitPrice(li); ok {
				priceByPartner[p] = price
			}
		}

		// 5) 币种：默认 USD；若 DSCO 原始订单带 currencyCode 则使用它
		currency := "USD"
		if dscoOrder.CurrencyCode != nil && *dscoOrder.CurrencyCode != "" {
			currency = *dscoOrder.CurrencyCode
		}

		// 6) 计算“实际已发货数量”（WMS 聚合）与运单号集合：
		// - shippedQty：用于判断“是否全部已发货”
		// - tracking：仅用于日志排查；DSCO 发票接口不接受顶层 trackingNumber 字段
		shippedQtyByPartner := map[string]int{}
		trackingList := make([]string, 0, len(wmsOrders))
		var latestInvoiceDate string
		for _, w := range wmsOrders {
			if t := strings.TrimSpace(w.TrackingNo); t != "" {
				trackingList = append(trackingList, t)
			}
			// invoiceDate：取“最晚一次发货时间”（优先 DeliveredAt，否则 StockDeliveredAt）
			var chosen string
			if rawTime := strings.TrimSpace(w.DeliveredAt); rawTime != "" {
				if t, err := parseLingXingDateTimeToRFC3339UTC(rawTime); err == nil {
					chosen = t
				}
			} else if rawTime := strings.TrimSpace(w.StockDeliveredAt); rawTime != "" {
				if t, err := parseLingXingDateTimeToRFC3339UTC(rawTime); err == nil {
					chosen = t
				}
			}
			if chosen != "" {
				// RFC3339 字符串同一时区下可按字典序比较，但这里仍以字符串比较做最小实现。
				if latestInvoiceDate == "" || chosen > latestInvoiceDate {
					latestInvoiceDate = chosen
				}
			}
			for _, p := range w.ProductInfo {
				lxSKU := strings.TrimSpace(p.SKU)
				if lxSKU == "" || p.Count <= 0 {
					continue
				}
				dscoPartner := strings.TrimSpace(reverseSKU[lxSKU])
				if dscoPartner == "" {
					dscoPartner = lxSKU
				}
				if dscoPartner == "" {
					continue
				}
				shippedQtyByPartner[dscoPartner] += p.Count
			}
		}
		if latestInvoiceDate == "" {
			latestInvoiceDate = time.Now().UTC().Format(time.RFC3339)
		}
		trackingJoined := strings.Join(uniqueNonEmptyStrings(trackingList), ",")

		// 7) 全量发货判断：poNumber 下所有 SKU 均已发货才允许回传发票
		notReady := false
		for partner, expected := range expectedQtyByPartner {
			if expected <= 0 {
				continue
			}
			if shippedQtyByPartner[partner] < expected {
				notReady = true
				break
			}
		}
		if notReady {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "not_fully_shipped",
					"expected_qty", integration.JSONForLog(expectedQtyByPartner),
					"shipped_qty", integration.JSONForLog(shippedQtyByPartner),
					"wms_orders_raw", integration.JSONForLog(wmsOrders),
				)...,
			)
			continue
		}

		// 8) 组装“汇总发票”：
		// - 发票行数量以 DSCO 订单数量为准（已确认“必须全量发货”后再回传发票）
		// - 单价以 DSCO 原始订单为准（pickUnitPrice）
		// - 如缺少单价导致无法计算总额，则跳过（避免回传不完整发票）
		var missingPrice []string
		var lineItems []dsco.InvoiceLineItem
		var totalAmount float64
		for partner, expected := range expectedQtyByPartner {
			if expected <= 0 {
				continue
			}
			unit := priceByPartner[partner]
			if unit <= 0 {
				missingPrice = append(missingPrice, partner)
				continue
			}
			line := dsco.InvoiceLineItem{
				PartnerSKU: partner,
				Quantity:   expected,
				UnitPrice:  unit,
			}
			if !dscoItemIDConflict[partner] {
				if id := dscoItemIDByPartner[partner]; id != "" {
					line.DscoItemID = id
				}
			}
			if !lineNumberConflict[partner] {
				if n := lineNumberByPartner[partner]; n > 0 {
					line.LineNumber = n
				}
			}
			lineItems = append(lineItems, line)
			totalAmount += float64(expected) * unit
		}
		if len(missingPrice) > 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "missing_unit_price",
					"missing_partner_sku", strings.Join(uniqueNonEmptyStrings(missingPrice), ","),
					"dsco_raw", integration.JSONForLog(row.Payload),
				)...,
			)
			continue
		}
		if len(lineItems) == 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "no_invoice_line_items",
					"dsco_raw", integration.JSONForLog(row.Payload),
					"wms_orders_raw", integration.JSONForLog(wmsOrders),
				)...,
			)
			continue
		}

		totalAmount = math.Round(totalAmount*100) / 100
		invoiceID := po
		inv := dsco.Invoice{
			InvoiceID:           invoiceID,
			PoNumber:            po,
			ConsumerOrderNumber: derefString(dscoOrder.ConsumerOrderNumber),
			InvoiceDate:         latestInvoiceDate,
			CurrencyCode:        currency,
			TotalAmount:         totalAmount,
			LineItems:           lineItems,
		}
		invs = append(invs, inv)
		toUpdate = append(toUpdate, struct {
			po        string
			invoiceID string
		}{po: po, invoiceID: invoiceID})
		logger.Info(taskCtx, "order prepared",
			append(base,
				"task", "invoice_to_dsco",
				"po_number", po,
				"invoice_id", invoiceID,
				"tracking", trackingJoined,
				"invoice_request", integration.JSONForLog(inv),
				"wms_orders_raw", integration.JSONForLog(wmsOrders),
			)...,
		)
	}
	if len(invs) == 0 {
		return nil
	}
	// 11) 调用 DSCO invoice 批量接口
	resp, raw, err := dscoCli.Invoice.CreateSmallBatchWithRawBody(taskCtx, invs)
	if err != nil {
		retErr = err
		logger.Error(taskCtx, "dsco invoice createSmallBatch failed",
			append(base,
				"task", "invoice_to_dsco",
				"batch", integration.JSONForLog(invs),
				"resp_raw", raw,
				"err", err,
			)...,
		)
		return retErr
	}

	logger.Info(taskCtx, "dsco invoice createSmallBatch ok",
		append(base,
			"task", "invoice_to_dsco",
			"resp", integration.JSONForLog(resp),
			"resp_raw", raw,
		)...,
	)

	// 12) 校验 DSCO 侧是否真正落库（CreateSmallBatch 为异步接口；仅 HTTP 成功不代表每条都成功）。
	// - 优先：按 requestId 查询 change log，读取每条 invoice 的 success/failure 结果。
	// - 兜底：按 invoiceId 查询 /invoice，确认可见后再更新本地状态。
	verifyByInvoiceID := make(map[string]bool, len(toUpdate))
	verifyMsgByInvoiceID := make(map[string]any, len(toUpdate))

	remaining := make(map[string]struct{}, len(toUpdate))
	for _, it := range toUpdate {
		remaining[it.invoiceID] = struct{}{}
	}

	if resp != nil && strings.TrimSpace(resp.RequestID) != "" {
		m, msg, completed := pollDSCOInvoiceChangeLog(taskCtx, dscoCli, resp.RequestID, remaining, 60*time.Second, 3*time.Second)
		for k, v := range m {
			verifyByInvoiceID[k] = v
			if mm, ok := msg[k]; ok {
				verifyMsgByInvoiceID[k] = mm
			}
			delete(remaining, k)
		}

		var confirmedOK int
		var confirmedFail int
		for _, ok := range m {
			if ok {
				confirmedOK++
			} else {
				confirmedFail++
			}
		}
		logger.Info(taskCtx, "dsco invoice verify by requestId",
			append(base,
				"task", "invoice_to_dsco",
				"request_id", resp.RequestID,
				"completed", completed,
				"confirmed_ok", confirmedOK,
				"confirmed_fail", confirmedFail,
				"remaining_count", len(remaining),
				"remaining_invoice_ids", strings.Join(mapKeysToSortedSlice(remaining), ","),
			)...,
		)

		if !completed {
			logger.Warn(taskCtx, "dsco invoice verify note",
				append(base,
					"task", "invoice_to_dsco",
					"reason", "dsco_invoice_change_log_not_completed_in_time",
					"request_id", resp.RequestID,
				)...,
			)
		}
	}

	// 兜底：对未拿到明确结果的 invoiceId，再做一次可见性查询（避免 change log 延迟）。
	for invID := range remaining {
		ok, exErr := waitForDSCOInvoiceVisible(taskCtx, dscoCli, invID, 12, 2500*time.Millisecond)
		if exErr != nil {
			verifyMsgByInvoiceID[invID] = exErr.Error()
		}
		verifyByInvoiceID[invID] = ok
	}

	// 13) 仅在 DSCO 侧确认成功后，才写回 invoiceId 并推进到 status=5；否则保留 status=4 以便下次重试。
	for _, it := range toUpdate {
		ok := verifyByInvoiceID[it.invoiceID]
		if !ok {
			fail++
			// 注意：Warn 级别可能写入独立文件，这里同时打一条 Info，确保 info.log 可见失败原因。
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", it.po,
					"result", "fail",
					"reason", "dsco_invoice_not_confirmed",
					"invoice_id", it.invoiceID,
					"verify_detail", verifyMsgByInvoiceID[it.invoiceID],
				)...,
			)
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", it.po,
					"result", "fail",
					"reason", "dsco_invoice_not_confirmed",
					"invoice_id", it.invoiceID,
					"verify_detail", verifyMsgByInvoiceID[it.invoiceID],
				)...,
			)
			continue
		}

		if err := d.orderStore.UpdateStatusAndInvoiceID(taskCtx, it.po, 5, it.invoiceID); err != nil {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "invoice_to_dsco",
					"po_number", it.po,
					"result", "fail",
					"reason", "update_local_status_failed_after_invoice",
					"invoice_id", it.invoiceID,
					"err", err,
				)...,
			)
			continue
		}
		okCount++
		logger.Info(taskCtx, "order done",
			append(base,
				"task", "invoice_to_dsco",
				"po_number", it.po,
				"result", "ok",
				"reason", "dsco_invoice_sent",
				"invoice_id", it.invoiceID,
				"new_status", 5,
			)...,
		)
	}

	// 手动单号回传：若未确认成功，返回错误让 Admin 端明确感知失败。
	if strings.TrimSpace(ctx.OnlyPONumber) != "" && okCount == 0 && fail > 0 {
		return errors.New("dsco invoice not confirmed")
	}
	return nil
}

const dscoInvoiceGetKeyInvoiceID = "invoiceId"

func dscoInvoiceExists(ctx context.Context, cli *dsco.Client, invoiceID string) (bool, error) {
	invoiceID = strings.TrimSpace(invoiceID)
	if cli == nil || invoiceID == "" {
		return false, nil
	}
	out, err := cli.Invoice.GetByID(ctx, dsco.InvoiceGetQuery{Key: dscoInvoiceGetKeyInvoiceID, Value: invoiceID})
	if err != nil {
		return false, err
	}
	return out != nil && len(out.Invoices) > 0, nil
}

func waitForDSCOInvoiceVisible(ctx context.Context, cli *dsco.Client, invoiceID string, attempts int, delay time.Duration) (bool, error) {
	if attempts <= 0 {
		attempts = 1
	}
	if delay <= 0 {
		delay = 500 * time.Millisecond
	}
	var lastErr error
	for i := 0; i < attempts; i++ {
		exists, err := dscoInvoiceExists(ctx, cli, invoiceID)
		if err == nil && exists {
			return true, nil
		}
		if err != nil {
			lastErr = err
		}
		if i < attempts-1 {
			time.Sleep(delay)
		}
	}
	return false, lastErr
}

func pollDSCOInvoiceChangeLog(
	ctx context.Context,
	cli *dsco.Client,
	requestID string,
	want map[string]struct{},
	maxWait time.Duration,
	interval time.Duration,
) (map[string]bool, map[string]any, bool) {
	requestID = strings.TrimSpace(requestID)
	if cli == nil || requestID == "" || len(want) == 0 {
		return map[string]bool{}, map[string]any{}, true
	}
	if maxWait <= 0 {
		maxWait = 10 * time.Second
	}
	if interval <= 0 {
		interval = 1 * time.Second
	}

	deadline := time.Now().Add(maxWait)
	result := make(map[string]bool, len(want))
	detail := make(map[string]any, len(want))

	for {
		out, err := cli.Invoice.GetChangeLog(ctx, dsco.InvoiceChangeLogQuery{
			RequestID: requestID,
			Status:    "success_or_failure",
		})
		if err == nil && out != nil {
			for _, l := range out.Logs {
				invID := strings.TrimSpace(l.Payload.InvoiceID)
				if invID == "" {
					continue
				}
				if _, ok := want[invID]; !ok {
					continue
				}
				st := strings.ToLower(strings.TrimSpace(l.Status))
				switch st {
				case "success":
					result[invID] = true
				case "failure":
					result[invID] = false
					// 失败时记录完整变更日志，便于排查 DSCO 返回的细节信息（例如 processId / requestMethodDetail 等）。
					detail[invID] = l
				default:
					// pending/unknown：不写入结果，等待下一轮
				}
			}
			// 若 change log 已完成，且本轮已经看到了需要的结果，则结束。
			if strings.EqualFold(strings.TrimSpace(out.Status), "COMPLETED") {
				return result, detail, true
			}
			// 若已拿到所有需要的结果（success/failure），也可提前结束。
			allDone := true
			for invID := range want {
				if _, ok := result[invID]; !ok {
					allDone = false
					break
				}
			}
			if allDone {
				return result, detail, true
			}
		}

		if time.Now().After(deadline) {
			return result, detail, false
		}
		time.Sleep(interval)
	}
}

func mapKeysToSortedSlice(m map[string]struct{}) []string {
	if len(m) == 0 {
		return nil
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		s := strings.TrimSpace(k)
		if s == "" {
			continue
		}
		keys = append(keys, s)
	}
	sort.Strings(keys)
	const maxKeys = 20
	if len(keys) > maxKeys {
		keys = append(keys[:maxKeys], "...(truncated)")
	}
	return keys
}
