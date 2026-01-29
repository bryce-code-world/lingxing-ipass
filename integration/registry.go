package integration

import (
	"sort"
	"sync"

	"lingxingipass/infra/runtimecfg"
)

type Registry struct {
	mu    sync.RWMutex
	tasks map[runtimecfg.JobName]Task
	jobs  []runtimecfg.JobName
}

func NewRegistry() *Registry {
	return &Registry{
		tasks: map[runtimecfg.JobName]Task{},
	}
}

func (r *Registry) Register(job runtimecfg.JobName, task Task) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tasks[job]; ok {
		panic("duplicate job: " + string(job))
	}
	r.tasks[job] = task
	r.jobs = append(r.jobs, job)
	sort.Slice(r.jobs, func(i, j int) bool { return r.jobs[i] < r.jobs[j] })
}

func (r *Registry) Jobs() []runtimecfg.JobName {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]runtimecfg.JobName, 0, len(r.jobs))
	out = append(out, r.jobs...)
	return out
}

func (r *Registry) SupportedJobsSet() map[runtimecfg.JobName]struct{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make(map[runtimecfg.JobName]struct{}, len(r.tasks))
	for k := range r.tasks {
		out[k] = struct{}{}
	}
	return out
}

func (r *Registry) Get(job runtimecfg.JobName) (Task, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.tasks[job]
	return t, ok
}
