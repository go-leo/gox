package chanx

import "context"

// FanOut multiple functions can read from the same channel until that channel is closed; this is called fan-out.
// This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.
func FanOut[T any, R any](ctx context.Context, src <-chan T, f func(T) R) chan<- R {
	return Pipeline[T, R](ctx, src, f)
}
