package chanx

import "context"

// AllMatch 检查通道中所有元素是否满足给定条件，若全部满足则返回true，否则返回false。
func AllMatch[T any](ctx context.Context, in <-chan T, predicate func(value T) bool) <-chan bool {
	var out chan bool
	if in == nil {
		return out
	}
	out = make(chan bool, 1)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-in:
				if !ok {
					out <- true
					return
				}
				if !predicate(value) {
					out <- false
					return
				}
			}
		}
	}()
	return out
}

// AnyMatch 检查通道中是否有元素满足给定条件，找到即返回 true，否则返回 false。
func AnyMatch[T any](ctx context.Context, in <-chan T, predicate func(value T) bool) <-chan bool {
	var out chan bool
	if in == nil {
		return out
	}
	out = make(chan bool, 1)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-in:
				if !ok {
					out <- false
					return
				}
				if predicate(value) {
					out <- true
					return
				}
			}
		}
	}()
	return out
}

// NoneMatch 检查通道中的所有元素是否都不满足给定条件，若所有元素都不满足则返回true，否则返回false。
func NoneMatch[T any](ctx context.Context, in <-chan T, predicate func(value T) bool) <-chan bool {
	var out chan bool
	if in == nil {
		return out
	}
	out = make(chan bool, 1)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-in:
				if !ok {
					out <- true
					return
				}
				if predicate(value) {
					out <- false
					return
				}
			}
		}
	}()
	return out
}
