package chanx

import (
	"context"
	"reflect"
	"sync"
)

// FanIn 函数将多个输入通道的数据合并到一个输出通道中。
// See: [concurrency](https://go.dev/talks/2012/concurrency.slide#28)
func FanIn[T any](ctx context.Context, ins ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	for _, ch := range ins {
		wg.Add(1)
		go func(ch <-chan T) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-ch:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- v:
					}
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

// FanOut 函数将输入通道的数据分发到多个输出通道中。
func FanOut[T any](ctx context.Context, in <-chan T, length int) []chan<- T {
	outs := make([]chan<- T, length)
	for i := range outs {
		outs[i] = make(chan<- T)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				for _, out := range outs {
					select {
					case <-ctx.Done():
						return
					case out <- v:
					}
				}
			}
		}
	}()

	go func() {
		wg.Wait()
		for i := 0; i < len(outs); i++ {
			close(outs[i])
		}
	}()
	return outs
}

func _FanIn[T any](channels ...<-chan T) <-chan T {
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
