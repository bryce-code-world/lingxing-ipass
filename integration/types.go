package integration

import (
	"context"

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

	// RunID 本次任务执行的唯一标识（由 Runner 生成），用于把同一轮任务的明细日志串起来。
	RunID string

	Config runtimecfg.Config

	// Override is only used for some manual operations (e.g., manual pull range).
	Override any
}

// BaseLogFields 返回本次任务的公共日志字段（用于所有明细日志）。
func (tc TaskContext) BaseLogFields() []any {
	return []any{
		"run_id", tc.RunID,
		"domain", tc.Domain,
		"job", string(tc.Job),
		"trigger", string(tc.Trigger),
		"size", tc.Size,
	}
}
