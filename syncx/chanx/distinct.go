package chanx

import "golang.org/x/exp/slices"

func Distinct[T any](in <-chan T, cmp func(value, stored T) bool) <-chan T {
	out := make(chan T, cap(in))
	go func() {
		defer close(out)
		values := make([]T, 0, cap(in))
		for value := range in {
			// 判断是否已经发送过
			if slices.ContainsFunc[[]T](values, func(stored T) bool { return cmp(value, stored) }) {
				continue
			}
			out <- value
			values = append(values, value)
		}
	}()
	return out
}
