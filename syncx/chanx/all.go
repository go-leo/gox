package chanx

import (
	"context"
	"sync"
)

// All 将一个只读通道中的所有元素收集并转换为切片返回。
func All[T any](c <-chan T) []T {
	out := make([]T, 0, len(c))
	for t := range c {
		out = append(out, t)
	}
	return out
}

// AsyncAll the function collects first value from multiple given channels into a slice and returns it.
func AsyncAll[T any](ctx context.Context, channels ...<-chan T) []T {
	if len(channels) <= 0 {
		return nil
	}

	valueCh := make(chan T, len(channels))

	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan T) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			case val, ok := <-ch:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case valueCh <- val:
					return
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(valueCh)
	}()

	values := make([]T, 0, len(channels))
	for v := range valueCh {
		values = append(values, v)
	}
	return values
}
