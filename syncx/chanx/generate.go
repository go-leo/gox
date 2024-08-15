package chanx

import "context"

// AsyncGenerate 异步生成类型为T的值，并通过通道输出。它接收上下文和一个供应商函数，当上下文被取消时停止生成。
func AsyncGenerate[T any](ctx context.Context, supplier func(context.Context) T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for {
			value := supplier(ctx)
			select {
			case out <- value:
				// 发送value
			case <-ctx.Done():
				// 被下游打断
				return
			}
		}
	}()
	return out
}
