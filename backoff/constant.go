package backoff

import (
	"context"
	"time"
)

// Constant it waits for a fixed period of time between calls.
func Constant(delta time.Duration) Func {
	return func(ctx context.Context, attempt uint) time.Duration {
		return delta
	}
}

// Zero it waits for zero time between calls.
func Zero() Func {
	return Constant(0)
}

func ConstantFactory() Factory {
	return func(delta time.Duration) Func {
		return Constant(delta)
	}
}
