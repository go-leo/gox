package chanx

import (
	"context"
)

// Reduce 对输入channel中的元素执行归约操作，使用提供的累积函数和初始值，返回最终结果。
func Reduce[T any, R any](ctx context.Context, in <-chan T, identity R, accumulator func(R, T) R) <-chan R {
	return Pipeline[T, R](ctx, in, func(value T) R {
		identity = accumulator(identity, value)
		return identity
	})
}
