package dsco_lingxing

import (
	"lingxingipass/infra/runtimecfg"
	"lingxingipass/integration"
)

func Register(reg *integration.Registry, d *Domain) {
	reg.Register(runtimecfg.JobPullDSCOOrders, d.PullDSCOOrders)
	reg.Register(runtimecfg.JobPushToLingXing, d.PushToLingXing)
	reg.Register(runtimecfg.JobAckToDSCO, d.AckToDSCO)
	reg.Register(runtimecfg.JobShipToDSCO, d.ShipToDSCO)
	reg.Register(runtimecfg.JobInvoiceToDSCO, d.InvoiceToDSCO)
	reg.Register(runtimecfg.JobSyncStock, d.SyncStock)
	reg.Register(runtimecfg.JobPullSKUPair, d.PullSKUPair)
	reg.Register(runtimecfg.JobCheckOrders, d.CheckOrders)
}
