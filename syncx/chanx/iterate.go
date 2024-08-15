package chanx

import "context"

// AsyncIterate 创建一个异步生成T类型值的channel，基于初始值和迭代函数不断产生新值，直至被上下文取消。
func AsyncIterate[T any](ctx context.Context, seed T, f func(context.Context, T) T) <-chan T {
	out := make(chan T)
	go func(seed T) {
		defer close(out)
		pre := seed
		for {
			value := f(ctx, pre)
			select {
			case out <- value:
				// 发送value
			case <-ctx.Done():
				// 被下游打断
				return
			}
			pre = value
		}
	}(seed)
	return out
}
