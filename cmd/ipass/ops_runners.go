package main

import (
	"context"
	"errors"

	"lingxingipass/internal/platform/adminhttp"
)

type opsRunnerDeps struct {
	PullDSCOOrders       adminhttp.JobRunner
	PushOrdersToLingXing adminhttp.JobRunner
	AckToDSCO            adminhttp.JobRunner
	ShipToDSCO           adminhttp.JobRunner
	InvoiceToDSCO        adminhttp.JobRunner
	SyncStock            adminhttp.JobRunner
}

func registerStandardOpsRunners(runners map[string]adminhttp.JobRunner, deps opsRunnerDeps) {
	if runners == nil {
		return
	}

	notReady := func(job string) adminhttp.JobRunner {
		return func(ctx context.Context) error {
			return errors.New("任务未就绪: " + job)
		}
	}

	if deps.PullDSCOOrders != nil {
		runners["pull_dsco_orders"] = deps.PullDSCOOrders
	} else {
		runners["pull_dsco_orders"] = notReady("pull_dsco_orders")
	}

	if deps.PushOrdersToLingXing != nil {
		runners["push_orders_to_lingxing"] = deps.PushOrdersToLingXing
	} else {
		runners["push_orders_to_lingxing"] = notReady("push_orders_to_lingxing")
	}

	if deps.AckToDSCO != nil {
		runners["ack_to_dsco"] = deps.AckToDSCO
	} else {
		runners["ack_to_dsco"] = notReady("ack_to_dsco")
	}

	if deps.ShipToDSCO != nil {
		runners["ship_to_dsco"] = deps.ShipToDSCO
	} else {
		runners["ship_to_dsco"] = notReady("ship_to_dsco")
	}

	if deps.InvoiceToDSCO != nil {
		runners["invoice_to_dsco"] = deps.InvoiceToDSCO
	} else {
		runners["invoice_to_dsco"] = notReady("invoice_to_dsco")
	}

	if deps.SyncStock != nil {
		runners["sync_stock"] = deps.SyncStock
	} else {
		runners["sync_stock"] = notReady("sync_stock")
	}
}

