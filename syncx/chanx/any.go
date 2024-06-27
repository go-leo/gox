package chanx

import (
	"context"
	"sync"
)

func Any[T any](channels ...<-chan T) T {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() 

	valueC := make(chan T, 1) 
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan T) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch:
				if ok {
					select {
					case valueC <- v:
						cancel()
					default: 
						return
					}
				}

			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(valueC)
	}()

	return <-valueC
}
