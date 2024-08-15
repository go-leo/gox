package chanx

import (
	"context"
)

// Reduce 对输入channel中的元素执行归约操作，使用提供的累积函数和初始值，返回最终结果。
// 若channel为空，则直接返回初始值。
func Reduce[T any, R any](in <-chan T, identity R, accumulator func(R, T) R) R {
	if in == nil {
		return identity
	}
	for value := range in {
		identity = accumulator(identity, value)
	}
	return identity
}

// AsyncReduce 异步处理通道数据，使用累积器函数累加元素，并能响应上下文取消操作，最终返回累积结果。
func AsyncReduce[T any, R any](ctx context.Context, in <-chan T, identity R, accumulator func(context.Context, R, T) R) R {
	if in == nil {
		return identity
	}
	out := make(chan R)
	go func(identity R) {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-in:
				if !ok {
					out <- identity
					return
				}
				identity = accumulator(ctx, identity, value)
			}
		}
	}(identity)
	return <-out
}
