package dsco_lingxing

import (
	"context"

	"example.com/lingxing/golib/v2/sdk/dsco"
	"example.com/lingxing/golib/v2/sdk/lingxing"

	"lingxingipass/integration"
)

func (d *Domain) AckToDSCO(ctx integration.TaskContext) error {
	taskCtx := ctx.Ctx
	if taskCtx == nil {
		taskCtx = context.Background()
	}

	items, err := d.orderStore.FindByStatus(taskCtx, 2, ctx.Size)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}

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
		detail, err := lx.Order.GetOrderDetailV2(taskCtx, lingxing.OrderDetailV2Request{
			PlatformOrderNo: po,
		})
		if err != nil {
			continue
		}
		if detail.Status == lingxing.MultiPlatformOrderStatusPendingShipment || detail.Status == lingxing.MultiPlatformOrderStatusShipped {
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
	_, err = dscoCli.Order.Acknowledge(taskCtx, reqs)
	if err != nil {
		return err
	}
	for _, po := range toUpdate {
		_ = d.orderStore.UpdateStatusAndFields(taskCtx, po, 3, "", "")
	}
	return nil
}
