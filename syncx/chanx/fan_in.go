package chanx

import (
	"context"
)

// FanIn can read from multiple inputs and proceed until all are closed by multiplexing the input
// channels onto a single channel that’s closed when all the inputs are closed.
// See: [concurrency](https://go.dev/talks/2012/concurrency.slide#28)
func FanIn[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	return Merge[T](ctx, channels...)
}
