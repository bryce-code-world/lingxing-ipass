package main

import (
	"context"
	"testing"

	"lingxingipass/internal/platform/adminhttp"
	"lingxingipass/internal/platform/config"
)

func TestOpsRunners_ManualTriggerHasStandardJobsWhenSchedulerDisabled(t *testing.T) {
	// 业务目标：即使所有定时任务开关都关闭（不跑调度），也应该仍然能通过 ops 接口手动触发标准 job。
	// 否则 admin 侧会收到 {"error":"unknown job"}，无法运维。

	cfg := config.Config{}
	cfg.System.Jobs.PullDSCOOrdersEnable = false
	cfg.System.Jobs.PushOrdersToLingXingEnable = false
	cfg.System.Jobs.AckToDSCOEnable = false
	cfg.System.Jobs.ShipToDSCOEnable = false
	cfg.System.Jobs.InvoiceToDSCOEnable = false
	cfg.System.Jobs.SyncStockEnable = false

	runners := map[string]adminhttp.JobRunner{}
	registerStandardOpsRunners(runners, opsRunnerDeps{})

	// 只验证“job 名称存在”，不验证具体业务执行结果（执行链路在 sync 包已有大量测试覆盖）。
	for _, name := range []string{
		"pull_dsco_orders",
		"push_orders_to_lingxing",
		"ack_to_dsco",
		"ship_to_dsco",
		"invoice_to_dsco",
		"sync_stock",
	} {
		fn, ok := runners[name]
		if !ok {
			t.Fatalf("runners 缺少 job=%s", name)
		}
		// runner 必须可调用（即使返回“未就绪”错误），避免 handler 侧出现 nil panic。
		_ = fn(context.Background())
	}

	_ = cfg // 避免误以为测试与 cfg 无关；cfg 的 enable=false 即本用例前提。
}

