package backoff

import (
	"context"
	"time"
)

// Constant it waits for a fixed period of time between calls.
func Constant(delta time.Duration) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return delta
	}
}

// Zero it waits for zero time between calls.
func Zero() BackoffFunc {
	return func(_ context.Context, _ uint) time.Duration {
		return 0
	}
}
