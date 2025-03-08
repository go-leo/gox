package backoff

import (
	"context"
	"time"
)

// Func return the backoff duration between call retries.
// The context.Context can be used to extract context values.
type Func func(ctx context.Context, attempt uint) time.Duration

// Factory returns a Func.
// see: ConstantFactory, ExponentialFactory, Exponential2Factory, FibonacciFactory, LinearFactory, JitterUpFactory.
type Factory func(delta time.Duration) Func
