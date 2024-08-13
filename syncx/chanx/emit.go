package chanx

import "context"

// Emit 函数接收任意数量和类型的参数，创建一个同类型的通道，将这些值依次放入通道后返回，并最终关闭此通道。
func Emit[T any](values ...T) <-chan T {
	out := make(chan T, len(values))
	for _, value := range values {
		out <- value
	}
	close(out)
	return out
}

// AsyncEmit 创建一个通道，异步发送传入的值，并允许通过上下文取消发送过程。
func AsyncEmit[T any](ctx context.Context, values ...T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, value := range values {
			select {
			case <-ctx.Done():
				return
			case out <- value:
			}
		}
	}()
	return out
}
