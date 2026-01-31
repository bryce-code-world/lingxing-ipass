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

// AckToDSCO 将“已审核/已发货”的订单回传 ACK 给 DSCO。
//
// 状态机（一期口径）：
// - 本任务处理 dsco_order_sync.status = 2 的订单（待确认/待回传 ACK）。
// - 回传成功后将本地状态推进到 3（待回传发货信息）。
//
// 核心逻辑（你最新确认）：
//  1. 先查 DSCO 订单状态（dscoStatus）：
//     - dscoStatus == shipment_pending：表示 DSCO 已确认，直接推进本地状态到 3，跳过 ACK。
//     - dscoStatus == shipped/completed：表示已进入更后阶段，同样直接推进到 3。
//     - 其他状态：才考虑继续走 ACK。
//  2. 对“仍需 ACK”的订单，再查领星订单状态（5/6 才允许 ACK），避免过早 ACK。
//     - 注意：同一 poNumber 可能在领星侧拆成多条子订单；此时必须“全部子订单”满足 5/6 才允许 ACK。
//  3. 将需要 ACK 的订单组装为 DSCO Acknowledge 批量请求，调用 /order/acknowledge。
func (d *Domain) AckToDSCO(ctx integration.TaskContext) (retErr error) {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	startedAt := time.Now().UTC()
	base := ctx.BaseLogFields()
	logger.Info(taskCtx, "task begin", append(base, "task", "ack_to_dsco")...)

	var (
		total    int
		okCount  int
		skip     int
		fail     int
		advanced int
	)
	defer func() {
		fields := append(base,
			"task", "ack_to_dsco",
			"duration_ms", time.Since(startedAt).Milliseconds(),
			"total", total,
			"ok", okCount,
			"skip", skip,
			"fail", fail,
			"advanced", advanced,
		)
		if retErr != nil {
			logger.Error(taskCtx, "task end", append(fields, "result", "failed", "err", retErr)...)
			return
		}
		logger.Info(taskCtx, "task end", append(fields, "result", "ok")...)
	}()

	// 1) 取待 ACK（status=2）
	var items []store.DSCOOrderSyncRow
	var err error
	if strings.TrimSpace(ctx.OnlyPONumber) != "" {
		items, err = d.orderStore.FindByStatusAndPONumber(taskCtx, 2, ctx.OnlyPONumber)
	} else {
		items, err = d.orderStore.FindByStatus(taskCtx, 2, ctx.Size)
	}
	if err != nil {
		retErr = err
		return retErr
	}
	if len(items) == 0 {
		return nil
	}
	total = len(items)

	// 任务明细日志需要用到“本地保存的 DSCO 原始订单 payload”。
	payloadByPO := make(map[string][]byte, len(items))
	for _, it := range items {
		if po := strings.TrimSpace(it.PONumber); po != "" {
			payloadByPO[po] = it.Payload
		}
	}

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

	// 3) 先查 DSCO 订单状态（避免重复 ACK）：
	poNumbers := make([]string, 0, len(items))
	for _, it := range items {
		poNumbers = append(poNumbers, it.PONumber)
	}
	dscoByPO := fetchDSCOOrdersByPONumbers(taskCtx, dscoCli, poNumbers, 5)
	logger.Info(taskCtx, "dsco orders fetched",
		append(base,
			"task", "ack_to_dsco",
			"po_count", len(uniqueNonEmptyStrings(poNumbers)),
			"fetched", len(dscoByPO),
		)...,
	)

	needAck := make([]string, 0, len(items))
	for _, row := range items {
		po := strings.TrimSpace(row.PONumber)
		if po == "" {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ack_to_dsco",
					"po_number", "",
					"result", "skip",
					"reason", "po_number_empty",
				)...,
			)
			continue
		}

		if o, ok := dscoByPO[po]; ok {
			st := strings.TrimSpace(o.DscoStatus)
			// shipment_pending：已确认（等待发货）；shipped/completed：已进入更后阶段。
			if st == "shipment_pending" || st == "shipped" || st == "completed" {
				if uerr := d.orderStore.UpdateStatusAndFields(taskCtx, po, 3, "", ""); uerr != nil {
					fail++
					logger.Warn(taskCtx, "order done",
						append(base,
							"task", "ack_to_dsco",
							"po_number", po,
							"result", "fail",
							"reason", "update_local_status_failed",
							"dsco_status", st,
							"dsco_raw", integration.JSONForLog(o),
							"err", uerr,
						)...,
					)
					continue
				}
				skip++
				advanced++
				logger.Info(taskCtx, "order done",
					append(base,
						"task", "ack_to_dsco",
						"po_number", po,
						"result", "skip",
						"reason", "dsco_status_already_confirmed_or_later",
						"dsco_status", st,
						"dsco_raw", integration.JSONForLog(o),
						"new_status", 3,
					)...,
				)
				continue
			}
		}
		needAck = append(needAck, po)
	}
	if len(needAck) == 0 {
		return nil
	}

	// 4) 批量查询领星订单状态（避免逐单查询触发 API 限制）：
	// - 仅对“确实需要 ACK”的订单查询领星状态。
	detailsByPO := make(map[string][]lingxing.OrderDetailV2, len(needAck))
	includeDelete := true
	const maxBatch = 50
	for _, chunk := range chunkStrings(uniqueNonEmptyStrings(needAck), maxBatch) {
		out, err := lx.Order.ListOrdersV2(taskCtx, lingxing.OrderListV2Request{
			Offset:           0,
			Length:           len(chunk),
			PlatformOrderNos: chunk,
			IncludeDelete:    &includeDelete,
		})
		if err != nil {
			logger.Warn(taskCtx, "lingxing list orders failed, fallback to per-order check", "err", err)
			// 降级：逐单查，保证任务不被单次 list 失败卡死
			for _, po := range chunk {
				detail, derr := lx.Order.GetOrderDetailV2(taskCtx, lingxing.OrderDetailV2Request{PlatformOrderNo: po})
				if derr == nil {
					detailsByPO[po] = append(detailsByPO[po], detail)
				}
			}
			continue
		}
		for _, detail := range out.List {
			if po := poNumberFromLingXingOrderDetail(detail); po != "" {
				detailsByPO[po] = append(detailsByPO[po], detail)
			}
		}
	}

	// 5) 组装 DSCO ACK 请求：只对领星已到 5/6 的订单 ACK（已确认口径）
	var reqs []dsco.OrderAcknowledgeRequest
	var toUpdate []string
	for _, po := range needAck {
		details := detailsByPO[po]
		if len(details) == 0 {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ack_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "lingxing_order_not_found",
					"dsco_payload", integration.JSONForLog(payloadByPO[po]),
				)...,
			)
			continue
		}

		ready := true
		statuses := make([]int, 0, len(details))
		for _, d := range details {
			statuses = append(statuses, int(d.Status))
			if d.Status != lingxing.MultiPlatformOrderStatusPendingShipment && d.Status != lingxing.MultiPlatformOrderStatusShipped {
				ready = false
			}
		}
		if !ready {
			skip++
			logger.Info(taskCtx, "order done",
				append(base,
					"task", "ack_to_dsco",
					"po_number", po,
					"result", "skip",
					"reason", "lingxing_status_not_ready_for_ack",
					"lingxing_statuses", statuses,
					"lingxing_count", len(details),
					"lingxing_raw", integration.JSONForLog(details),
					"dsco_payload", integration.JSONForLog(payloadByPO[po]),
				)...,
			)
			continue
		}
		reqs = append(reqs, dsco.OrderAcknowledgeRequest{ID: po, Type: dsco.OrderAcknowledgeIDTypePoNumber})
		toUpdate = append(toUpdate, po)
	}
	if len(reqs) == 0 {
		return nil
	}

	// 6) 调用 DSCO ACK 批量接口
	logger.Info(taskCtx, "dsco acknowledge request",
		append(base,
			"task", "ack_to_dsco",
			"count", len(reqs),
			"reqs", integration.JSONForLog(reqs),
		)...,
	)
	if _, err := dscoCli.Order.Acknowledge(taskCtx, reqs); err != nil {
		retErr = err
		logger.Error(taskCtx, "dsco acknowledge failed",
			append(base,
				"task", "ack_to_dsco",
				"reqs", integration.JSONForLog(reqs),
				"err", err,
			)...,
		)
		return retErr
	}

	// 7) 回传成功：推进状态到 3
	for _, po := range toUpdate {
		if uerr := d.orderStore.UpdateStatusAndFields(taskCtx, po, 3, "", ""); uerr != nil {
			fail++
			logger.Warn(taskCtx, "order done",
				append(base,
					"task", "ack_to_dsco",
					"po_number", po,
					"result", "fail",
					"reason", "update_local_status_failed_after_ack",
					"err", uerr,
				)...,
			)
			continue
		}
		okCount++
		details := detailsByPO[po]
		statuses := make([]int, 0, len(details))
		for _, d := range details {
			statuses = append(statuses, int(d.Status))
		}
		dscoSt := ""
		dscoRaw := ""
		if o, ok := dscoByPO[po]; ok {
			dscoSt = strings.TrimSpace(o.DscoStatus)
			dscoRaw = integration.JSONForLog(o)
		}
		logger.Info(taskCtx, "order done",
			append(base,
				"task", "ack_to_dsco",
				"po_number", po,
				"result", "ok",
				"reason", "dsco_ack_sent",
				"dsco_status", dscoSt,
				"dsco_raw", dscoRaw,
				"dsco_payload", integration.JSONForLog(payloadByPO[po]),
				"lingxing_statuses", statuses,
				"lingxing_count", len(details),
				"lingxing_raw", integration.JSONForLog(details),
				"new_status", 3,
			)...,
		)
	}
	return nil
}
