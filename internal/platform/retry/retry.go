package retry

import (
	"context"
	"errors"
	"time"
)

// Policy 表示重试策略（一期最小可用）。
type Policy struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
}

func DefaultPolicy() Policy {
	return Policy{
		MaxAttempts: 3,
		BaseDelay:   200 * time.Millisecond,
		MaxDelay:    2 * time.Second,
	}
}

// Do 执行 fn，按 shouldRetry 判断是否重试。
func Do(ctx context.Context, p Policy, shouldRetry func(error) bool, fn func() error) error {
	if fn == nil {
		return errors.New("fn 不能为空")
	}
	if p.MaxAttempts <= 0 {
		p.MaxAttempts = 1
	}
	if p.BaseDelay <= 0 {
		p.BaseDelay = 100 * time.Millisecond
	}
	if p.MaxDelay <= 0 {
		p.MaxDelay = 2 * time.Second
	}
	if shouldRetry == nil {
		shouldRetry = func(err error) bool { return false }
	}

	var lastErr error
	for attempt := 1; attempt <= p.MaxAttempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		err := fn()
		if err == nil {
			return nil
		}
		lastErr = err

		if attempt == p.MaxAttempts || !shouldRetry(err) {
			return err
		}

		delay := p.BaseDelay * time.Duration(1<<(attempt-1))
		if delay > p.MaxDelay {
			delay = p.MaxDelay
		}

		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}
	return lastErr
}
