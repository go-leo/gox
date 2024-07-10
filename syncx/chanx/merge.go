package chanx

import (
	"context"
	"sync"
)

// Merge consolidates multiple input channels into one output channel:
// - Creates a buffered output channel.
// - Spawns goroutines for each input channel to read and forward data to the output channel, controlled by `context.Context`.
// - Waits for all goroutines to finish and closes the output channel.
// - Returns the merged output channel, supporting concurrency and cancellation mechanisms.
func Merge[T any](ctx context.Context, channels ...<-chan T) <-chan T {
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

// Combine multiple input channels into one output channel.
// Deprecated: Do not use. use Merge instead.
func Combine[T any](channels ...<-chan T) <-chan T {
	return Merge[T](context.Background(), channels...)
}
