package chanx

import "context"

// Limit 接收泛型输入通道in和上下文ctx，以及整数maxSize作为参数。
// 当从in读取数据时，向输出通道out发送数据项，直到maxSize达到0或in关闭。
// 如果in为nil，则直接返回nil。
// 支持通过ctx.Done()提前退出以取消操作。
func Limit[T any](ctx context.Context, in <-chan T, maxSize int) <-chan T {
	var out chan T
	if in == nil {
		return out
	}
	out = make(chan T, cap(in))
	go func() {
		defer close(out)
		for {
			if maxSize < 0 {
				return
			}
			var value T
			var ok bool
			select {
			case <-ctx.Done():
				return
			case value, ok = <-in:
				if !ok {
					return
				}
			}
			select {
			case <-ctx.Done():
				return
			case out <- value:
				maxSize--
			}
		}
	}()
	return out
}
