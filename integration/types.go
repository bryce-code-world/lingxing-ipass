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

	Config runtimecfg.Config

	// Override is only used for some manual operations (e.g., manual pull range).
	Override any
}
