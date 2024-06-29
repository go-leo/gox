package chanx

import (
	"context"
	"sync"
)

// All the function collects first value from multiple given channels into a slice and returns it. 
func All[T any](ctx context.Context, channels ...<-chan T) []T {
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
