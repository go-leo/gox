package chanx

import "context"

// AsSlice 将一个通道中的所有元素收集并转换为切片返回。
func AsSlice[T any](ctx context.Context, in <-chan T) <-chan []T {
	out := make(chan []T, 1)
	go func() {
		defer close(out)
		s := make([]T, 0)
		for {
			select {
			case <-ctx.Done():
				out <- s
				return
			case t, ok := <-in:
				if ok {
					s = append(s, t)
					continue
				}
				out <- s
				return
			}
		}
	}()
	return out
}
