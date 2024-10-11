package backoff

import (
	"context"
	"time"
)

type key struct{}

// Inject 将一个回退函数注入到给定的上下文中，并返回一个新的上下文。这样新上下文就携带了回退信息。
func Inject(ctx context.Context, backoff BackoffFunc) context.Context {
	return context.WithValue(ctx, key{}, backoff)
}

// Context 从给定的上下文中获取一个回退函数，如果存在则调用它并返回回退时间；否则返回0。主要用于动态设置重试间的延迟。
func Context() BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		backoff, ok := ctx.Value(key{}).(BackoffFunc)
		if ok {
			return backoff(ctx, attempt)
		}
		return 0
	}
}
