package backoff

import (
	"context"
	"time"
)

type key struct{}

func NewContext(ctx context.Context, delta time.Duration) context.Context {
	return context.WithValue(ctx, key{}, delta)
}

func FromContext(ctx context.Context) (time.Duration, bool) {
	delta, ok := ctx.Value(key{}).(time.Duration)
	return delta, ok
}

// BackoffFactory denotes a function that creates a BackoffFunc.
type BackoffFactory func(delta time.Duration) BackoffFunc

//
//func Context(backoff) BackoffFactory {
//	return func(delta time.Duration) BackoffFunc {
//		return func(ctx context.Context, attempt uint) time.Duration {
//
//		}
//	}
//}
