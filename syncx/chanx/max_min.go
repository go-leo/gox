package chanx

import "context"

// Max 用于从输入通道 in 中找到最大值，并通过输出通道 out 返回。
func Max[T any](ctx context.Context, in <-chan T, cmp func(a, b T) int) <-chan T {
	out := make(chan T, 1)
	go func() {
		var maxValue T
		defer func() {
			out <- maxValue
			close(out)
		}()
		var ok bool
		select {
		case <-ctx.Done():
			return
		case maxValue, ok = <-in:
			if !ok {
				return
			}
		}
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				if cmp(value, maxValue) > 0 {
					maxValue = value
				}
			}
		}
	}()
	return out
}

// Min 用于从输入通道 in 中找到最小值，并通过输出通道 out 返回。
func Min[T any](ctx context.Context, in <-chan T, cmp func(a, b T) int) <-chan T {
	out := make(chan T, 1)
	go func() {
		var minValue T
		defer func() {
			out <- minValue
			close(out)
		}()
		var ok bool
		select {
		case <-ctx.Done():
			return
		case minValue, ok = <-in:
			if !ok {
				return
			}
		}
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				if cmp(value, minValue) < 0 {
					minValue = value
				}
			}
		}
	}()
	return out
}
