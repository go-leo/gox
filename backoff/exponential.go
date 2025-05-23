package backoff

import (
	"context"
	"math"
	"time"
)

// Exponential it waits for "delta * e^attempts" time between calls.
func Exponential(delta time.Duration) Func {
	return func(ctx context.Context, attempt uint) time.Duration {
		return exponential(delta, attempt)
	}
}

// exponential return "delta * e^attempts" time.duration
func exponential(delta time.Duration, attempt uint) time.Duration {
	return delta * time.Duration(math.Exp(float64(attempt)))
}

func ExponentialFactory() Factory {
	return func(delta time.Duration) Func {
		return Exponential(delta)
	}
}
