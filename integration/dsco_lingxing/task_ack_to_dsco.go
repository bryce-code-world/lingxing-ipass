package dsco_lingxing

import (
	"context"
	"strings"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"
	"example.com/lingxing/golib/v2/tool/logger"

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
//  3. 将需要 ACK 的订单组装为 DSCO Acknowledge 批量请求，调用 /order/acknowledge。
func (d *Domain) AckToDSCO(ctx integration.TaskContext) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	// 1) 取待 ACK（status=2）
	items, err := d.orderStore.FindByStatus(taskCtx, 2, ctx.Size)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}

	// 2) 初始化客户端：DSCO + 领星
	dscoCli, err := d.dscoClient()
	if err != nil {
		return err
	}
	lx, err := d.lingxingClient(taskCtx)
	if err != nil {
		return err
	}

	// 3) 先查 DSCO 订单状态（避免重复 ACK）：
	poNumbers := make([]string, 0, len(items))
	for _, it := range items {
		poNumbers = append(poNumbers, it.PONumber)
	}
	dscoByPO := fetchDSCOOrdersByPONumbers(taskCtx, dscoCli, poNumbers, 5)

	needAck := make([]string, 0, len(items))
	for _, row := range items {
		po := strings.TrimSpace(row.PONumber)
		if po == "" {
			continue
		}

		if o, ok := dscoByPO[po]; ok {
			st := strings.TrimSpace(o.DscoStatus)
			// shipment_pending：已确认（等待发货）；shipped/completed：已进入更后阶段。
			if st == "shipment_pending" || st == "shipped" || st == "completed" {
				_ = d.orderStore.UpdateStatusAndFields(taskCtx, po, 3, "", "")
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
	detailsByPO := make(map[string]lingxing.OrderDetailV2, len(needAck))
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
					detailsByPO[po] = detail
				}
			}
			continue
		}
		for _, detail := range out.List {
			if po := poNumberFromLingXingOrderDetail(detail); po != "" {
				detailsByPO[po] = detail
			}
		}
	}

	// 5) 组装 DSCO ACK 请求：只对领星已到 5/6 的订单 ACK（已确认口径）
	var reqs []dsco.OrderAcknowledgeRequest
	var toUpdate []string
	for _, po := range needAck {
		detail, ok := detailsByPO[po]
		if !ok {
			continue
		}
		if detail.Status != lingxing.MultiPlatformOrderStatusPendingShipment && detail.Status != lingxing.MultiPlatformOrderStatusShipped {
			continue
		}
		reqs = append(reqs, dsco.OrderAcknowledgeRequest{ID: po, Type: dsco.OrderAcknowledgeIDTypePoNumber})
		toUpdate = append(toUpdate, po)
	}
	if len(reqs) == 0 {
		return nil
	}

	// 6) 调用 DSCO ACK 批量接口
	if _, err := dscoCli.Order.Acknowledge(taskCtx, reqs); err != nil {
		return err
	}

	// 7) 回传成功：推进状态到 3
	for _, po := range toUpdate {
		_ = d.orderStore.UpdateStatusAndFields(taskCtx, po, 3, "", "")
	}
	return nil
}
