package integration

import (
	"context"
	"strings"

	"lingxingipass/infra/runtimecfg"
)

type Trigger string

const (
	TriggerScheduled Trigger = "scheduled"
	TriggerManual    Trigger = "manual"
)

type RunRequest struct {
	Domain  string
	Job     runtimecfg.JobName
	Trigger Trigger
	Size    int

	// OnlyPONumber 非空时表示“只处理这一单”，任务取数必须按该 po_number 精确过滤。
	// 用于 Admin 手动触发单笔订单执行，不影响 scheduler 的批处理语义。
	OnlyPONumber string

	// Override is reserved for future manual parameters (e.g., pull range).
	Override any
}

type Task func(ctx TaskContext) error

type TaskContext struct {
	Ctx context.Context

	Domain  string
	Job     runtimecfg.JobName
	Trigger Trigger
	Size    int

	// SnapshotRuntimeConfig 返回当前内存中的 runtime_config 快照（用于做轻量并发保护等）。
	// 注意：Runner 在每次任务执行时会注入该函数。
	SnapshotRuntimeConfig func(domain string) (runtimecfg.RuntimeConfig, bool)

	// UpdateRuntimeConfig 将新的 runtime_config 写入 DB 并更新内存快照（等价于 Admin 保存后的效果）。
	// 注意：Runner 在每次任务执行时会注入该函数（内部会做 Validate）。
	UpdateRuntimeConfig func(ctx context.Context, domain string, cfg runtimecfg.Config) error

	// OnlyPONumber 非空时表示“只处理这一单”，任务取数必须按该 po_number 精确过滤。
	OnlyPONumber string

	// RunID 本次任务执行的唯一标识（由 Runner 生成），用于把同一轮任务的明细日志串起来。
	RunID string

	Config runtimecfg.Config

	// Override is only used for some manual operations (e.g., manual pull range).
	Override any
}

// BaseLogFields 返回本次任务的公共日志字段（用于所有明细日志）。
func (tc TaskContext) BaseLogFields() []any {
	fields := []any{
		"run_id", tc.RunID,
		"domain", tc.Domain,
		"job", string(tc.Job),
		"trigger", string(tc.Trigger),
		"size", tc.Size,
	}
	if strings.TrimSpace(tc.OnlyPONumber) != "" {
		fields = append(fields, "only_po_number", tc.OnlyPONumber)
	}
	return fields
}
