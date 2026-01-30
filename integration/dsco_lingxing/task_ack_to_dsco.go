package dsco_lingxing

import (
	"context"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"

	"lingxingipass/integration"
)

// AckToDSCO 将“已审核/已发货”的订单回传 ACK 给 DSCO。
//
// 状态机（一期口径）：
// - 本任务处理 dsco_order_sync.status = 2 的订单（待确认/待回传 ACK）。
// - 当领星订单状态为 5/6（待发货/已发货）时，认为可以回传 ACK。
// - 回传成功后将本地状态推进到 3（待回传发货信息）。
//
// 关键点：
// - 幂等策略主要依赖本地状态机：只有 status=2 才会参与 ACK。
// - 本任务会先查询领星订单状态再决定是否发送 DSCO ACK，避免过早 ACK。
// - 一期不落库失败原因，仅日志；失败时不推进状态，等待下次重试。
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

	var reqs []dsco.OrderAcknowledgeRequest
	var toUpdate []string
	for _, row := range items {
		po := row.PONumber
		// 3) 先查领星订单状态：仅当 status=5/6 才回传 ACK（一期口径）
		detail, err := lx.Order.GetOrderDetailV2(taskCtx, lingxing.OrderDetailV2Request{
			PlatformOrderNo: po,
		})
		if err != nil {
			continue
		}
		if detail.Status == lingxing.MultiPlatformOrderStatusPendingShipment || detail.Status == lingxing.MultiPlatformOrderStatusShipped {
			// 4) 组装 DSCO ACK 批量请求：用 poNumber 作为唯一键
			reqs = append(reqs, dsco.OrderAcknowledgeRequest{
				ID:   po,
				Type: dsco.OrderAcknowledgeIDTypePoNumber,
			})
			toUpdate = append(toUpdate, po)
		}
	}
	if len(reqs) == 0 {
		return nil
	}
	// 5) 调用 DSCO ACK 接口：批量回传
	_, err = dscoCli.Order.Acknowledge(taskCtx, reqs)
	if err != nil {
		return err
	}
	// 6) 回传成功：推进状态到 3
	for _, po := range toUpdate {
		_ = d.orderStore.UpdateStatusAndFields(taskCtx, po, 3, "", "")
	}
	return nil
}
