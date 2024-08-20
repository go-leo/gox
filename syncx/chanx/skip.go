package chanx

import "context"

// Skip 是一个泛型函数，用于从输入的通道in中跳过指定数量n的元素，并将剩余的元素传递到输出通道out中。
func Skip[T any](ctx context.Context, in <-chan T, n int) <-chan T {
	var out chan T
	if in == nil {
		return out
	}
	out = make(chan T, cap(in))
	go func() {
		defer close(out)
		for {
			var value T
			var ok bool
			select {
			case <-ctx.Done():
				return
			case value, ok = <-in:
				if !ok {
					return
				}
				n--
			}
			if n >= 0 {
				continue
			}
			select {
			case <-ctx.Done():
				return
			case out <- value:
			}
		}
	}()
	return out
}
