package scheduler

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestScheduler_Add_Validation(t *testing.T) {
	t.Parallel()

	s := New()
	if err := s.Add("", time.Second, func(ctx context.Context) error { return nil }); err == nil {
		t.Fatalf("want err for empty name")
	}
	if err := s.Add("a", 0, func(ctx context.Context) error { return nil }); err == nil {
		t.Fatalf("want err for non-positive interval")
	}
	if err := s.Add("a", time.Second, nil); err == nil {
		t.Fatalf("want err for nil fn")
	}
}

func TestScheduler_NoReentry_Behavior(t *testing.T) {
	t.Parallel()

	var called int64
	block := make(chan struct{})

	s := New()
	if err := s.Add("job", 10*time.Millisecond, func(ctx context.Context) error {
		atomic.AddInt64(&called, 1)
		<-block
		return nil
	}); err != nil {
		t.Fatalf("Add err=%v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s.Start(ctx)

	time.Sleep(50 * time.Millisecond)
	// 由于 job 一直阻塞，后续 tick 不应重入，called 应该仍为 1（允许极小抖动，但这里应稳定）。
	if got := atomic.LoadInt64(&called); got != 1 {
		close(block)
		s.Stop()
		t.Fatalf("called=%d want=1", got)
	}

	close(block)
	time.Sleep(30 * time.Millisecond)
	cancel()
	s.Stop()
}
