package chanx

import "context"

// Pipeline creates a data processing pipeline.
// It takes an input channel `src` and a processing function `f`, returning an output channel `dest`.
// In a background goroutine, it reads from `src`, applies `f` for processing,
// and writes the results to `dest`, while listening for cancellation signals from the `ctx` context.
func Pipeline[T any, R any](ctx context.Context, src <-chan T, f func(T) R) chan<- R {
	dest := make(chan R)
	go func() {
		defer close(dest)
		for v := range src {
			select {
			case <-ctx.Done():
				return
			case dest <- f(v):
			}
		}
	}()
	return dest
}
