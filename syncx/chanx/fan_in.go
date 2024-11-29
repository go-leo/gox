package chanx

import (
	"context"
	"reflect"
	"sync"
)

func FanIn[T any](channels ...<-chan T) <-chan T {
	return FanInContext[T](context.Background(), channels...)
}

// FanInContext can read from multiple inputs and proceed until all are closed by multiplexing the input
// channels onto a single channel that’s closed when all the inputs are closed.
// See: [concurrency](https://go.dev/talks/2012/concurrency.slide#28)
func FanInContext[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	out := make(chan T, len(channels))
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan T) {
			defer wg.Done()
			for v := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- v:
				}
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func fanIn[T any](channels ...<-chan T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		// 构造SelectCase slice
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 循环，从cases中选择一个可用的
		for len(cases) > 0 {
			chosen, recv, ok := reflect.Select(cases)
			if !ok { // 此channel已经close
				cases = append(cases[:chosen], cases[chosen+1:]...)
				continue
			}
			v, ok := recv.Interface().(T)
			if ok {
				out <- v
			}
		}
	}()
	return out
}
