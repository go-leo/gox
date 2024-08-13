package chanx

import (
	"context"
	"sync"
)

// Reduce 对输入channel中的元素执行归约操作，使用提供的累积函数和初始值，返回最终结果。
// 若channel为空，则直接返回初始值。
func Reduce[T any, R any](in <-chan T, identity R, accumulator func(R, T) R) R {
	if in == nil {
		return identity
	}
	for value := range in {
		identity = accumulator(identity, value)
	}
	return identity
}

func AsyncReduce[T any, R any](ctx context.Context, in <-chan T, identity R, accumulator func(R, T) R, combiner func(R, R) R) R {
	if in == nil {
		return identity
	}
	var wg sync.WaitGroup
	out := make(chan R, cap(in))
	for value := range in {
		wg.Add(1)
		go func(value T) {
			defer wg.Done()
			r := accumulator(identity, value)
			select {
			case out <- r:
			case <-ctx.Done():
				return
			}
		}(value)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for r := range out {
		identity = combiner(identity, r)
	}

	return identity
}
