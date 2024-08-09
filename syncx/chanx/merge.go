package chanx

import (
	"context"
)

// Merge consolidates multiple input channels into one output channel:
// - Creates a buffered output channel.
// - Spawns goroutines for each input channel to read and forward data to the output channel, controlled by `context.Context`.
// - Waits for all goroutines to finish and closes the output channel.
// - Returns the merged output channel, supporting concurrency and cancellation mechanisms.
func Merge[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	return FanIn[T](ctx, channels...)
}

// Combine multiple input channels into one output channel.
// Deprecated: Do not use. use Merge instead.
func Combine[T any](channels ...<-chan T) <-chan T {
	return FanIn[T](context.Background(), channels...)
}
