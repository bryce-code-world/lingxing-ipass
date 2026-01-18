package scheduler

import (
	"context"
	"errors"
	"sync"
	"time"

	"example.com/lingxing/golib/v2/tool/logger"
)

// JobFunc 表示一个可被定时触发的任务。
type JobFunc func(ctx context.Context) error

type job struct {
	name     string
	interval time.Duration
	fn       JobFunc

	mu      sync.Mutex
	running bool
}

// Scheduler 是一期最小化定时调度器（ticker 驱动）。
type Scheduler struct {
	mu   sync.Mutex
	jobs []*job

	wg     sync.WaitGroup
	stopCh chan struct{}
}

func New() *Scheduler {
	return &Scheduler{
		stopCh: make(chan struct{}),
	}
}

func (s *Scheduler) Add(name string, interval time.Duration, fn JobFunc) error {
	if name == "" {
		return errors.New("name 不能为空")
	}
	if interval <= 0 {
		return errors.New("interval 必须为正数")
	}
	if fn == nil {
		return errors.New("fn 不能为空")
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs = append(s.jobs, &job{name: name, interval: interval, fn: fn})
	return nil
}

func (s *Scheduler) Start(ctx context.Context) {
	s.mu.Lock()
	jobs := append([]*job(nil), s.jobs...)
	s.mu.Unlock()

	for _, j := range jobs {
		j := j
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			ticker := time.NewTicker(j.interval)
			defer ticker.Stop()

			// 立即执行一次，避免首次等待。
			_ = s.runOnce(ctx, j)

			for {
				select {
				case <-ctx.Done():
					return
				case <-s.stopCh:
					return
				case <-ticker.C:
					_ = s.runOnce(ctx, j)
				}
			}
		}()
	}
}

func (s *Scheduler) Stop() {
	s.mu.Lock()
	select {
	case <-s.stopCh:
		// 已关闭
	default:
		close(s.stopCh)
	}
	s.mu.Unlock()
	s.wg.Wait()
}

func (s *Scheduler) runOnce(ctx context.Context, j *job) error {
	j.mu.Lock()
	if j.running {
		j.mu.Unlock()
		return nil
	}
	j.running = true
	j.mu.Unlock()

	defer func() {
		j.mu.Lock()
		j.running = false
		j.mu.Unlock()
	}()

	ctx2, _ := logger.EnsureTraceID(ctx)
	start := time.Now()
	logger.Info(ctx2, "job_start", "job", j.name, "interval_sec", int(j.interval.Seconds()))

	err := j.fn(ctx2)
	cost := time.Since(start)
	if err != nil {
		logger.Error(ctx2, "job_end", "job", j.name, "cost_ms", cost.Milliseconds(), "err", err.Error())
		return err
	}
	logger.Info(ctx2, "job_end", "job", j.name, "cost_ms", cost.Milliseconds())
	return nil
}
