package chanx

import (
	"context"
	"reflect"
	"sync"
)

// FanIn can read from multiple inputs and proceed until all are closed by multiplexing the input
// channels onto a single channel that’s closed when all the inputs are closed.
// See: [concurrency](https://go.dev/talks/2012/concurrency.slide#28)
func FanIn[T any](ctx context.Context, channels ...<-chan T) <-chan T {
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

func fanInReflect[T any](channels ...<-chan T) <-chan T {
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

func fanInRec[T any](channels ...<-chan T) <-chan T {
	switch len(channels) {
	case 0:
		out := make(chan T)
		close(out)
		return out
	case 1:
		return channels[0]
	case 2:
		return mergeTwo(channels[0], channels[1])
	default:
		m := len(channels) / 2
		return mergeTwo(fanInRec[T](channels[:m]...), fanInRec[T](channels[m:]...))
	}
}

func mergeTwo[T any](aChan, bChan <-chan T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for aChan != nil || bChan != nil { //只要还有可读的chan
			select {
			case v, ok := <-aChan:
				if !ok { // aChan 已关闭，设置为nil
					aChan = nil
					continue
				}
				out <- v
			case v, ok := <-bChan:
				if !ok { // bChan 已关闭，设置为nil
					bChan = nil
					continue
				}
				out <- v
			}
		}
	}()
	return out
}

// Merge consolidates multiple input channels into one output channel:
// Deprecated: Do not use. use FanIn instead.
func Merge[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	return FanIn[T](ctx, channels...)
}

// Combine multiple input channels into one output channel.
// Deprecated: Do not use. use FanIn instead.
func Combine[T any](channels ...<-chan T) <-chan T {
	return FanIn[T](context.Background(), channels...)
}
