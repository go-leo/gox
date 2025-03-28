package backoff

import (
	"context"
	"time"
)

// Linear it waits for "delta * attempt" time between calls.
func Linear(delta time.Duration) Func {
	return func(ctx context.Context, attempt uint) time.Duration {
		return linear(delta, attempt)
	}
}

func linear(delta time.Duration, attempt uint) time.Duration {
	return delta * time.Duration(attempt)
}

func LinearFactory() Factory {
	return func(delta time.Duration) Func {
		return Linear(delta)
	}
}
