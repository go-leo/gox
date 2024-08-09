package chanx

import "context"

// Emit creates and returns a read-only channel that sequentially sends all the passed `values` through the channel.
// Once all `values` have been sent, it closes the channel.
// The function uses generics `[T any]`, making it applicable to values of any type.
// Internally, it starts a goroutine where it iterates over the `values`, sending each one through the channel.
// Finally, the function returns this channel.
//
// See: [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
func Emit[T any](values ...T) <-chan T {
	out := make(chan T, len(values))
	for _, value := range values {
		out <- value
	}
	close(out)
	return out
}

func AsStream[T any](ctx context.Context, values ...T) <-chan T {
	out := make(chan T) //创建一个unbuffered的channel
	go func() {         // 启动一个goroutine，往s中塞数据
		defer close(out)               // 退出时关闭chan
		for _, value := range values { // 遍历数组
			select {
			case <-ctx.Done():
				return
			case out <- value: // 将数组元素塞入到chan中
			}
		}
	}()
	return out
}
