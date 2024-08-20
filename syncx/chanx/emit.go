package chanx

import "context"

// Emit 创建一个通道，异步发送传入的值，并允许通过上下文取消发送过程。
func Emit[T any](ctx context.Context, values ...T) <-chan T {
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
