package chanx

import "sync"


func All[T any](channels ...<-chan T) []T {
	valueCh := make(chan T, len(channels))

	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan T) {
			defer wg.Done()
			val, ok := <-ch
			if ok {
				valueCh <- val
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
