package ops

import (
	"lingxingipass/infra/runtimecfg"
	"lingxingipass/integration"
)

func Register(reg *integration.Registry, d *Domain) {
	reg.Register(runtimecfg.JobCleanupExports, d.CleanupExports)
}
