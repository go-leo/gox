package backoff

import (
	"context"
	"time"
)

type key struct{}

func NewContext(ctx context.Context, duration time.Duration) context.Context {
	return context.WithValue(ctx, key{}, duration)
}

func FromContext(ctx context.Context) (time.Duration, bool) {
	delta, ok := ctx.Value(key{}).(time.Duration)
	return delta, ok
}

func Context() BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		duration, ok := FromContext(ctx)
		if ok {
			return duration
		}
		return 0
	}
}
