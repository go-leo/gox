package chanx

import (
	"context"
	"github.com/go-leo/gox/slicex"
	"runtime"

	"github.com/go-leo/gox/mathx/randx"
)

// MovingOn is designed to randomly select from multiple given channels and receive the sent values until any of
// the channels is closed or the context is canceled.
//
// - **Behavior**:
// - It creates two local channels: `valueC` for receiving values and `errC` for potential errors.
// - In a loop, it selects a random channel from the input channels and attempts to read from it.
// - If the context is done, it sends the context's error into `errC` and returns.
// - If a value is received successfully, it's sent to `valueC`.
// - If a channel is closed (indicated by `!ok`), that channel is removed from the list and the process continues.
// - If no value can be read immediately, control is yielded back to the scheduler.
//
// - **Return Values**:
// - Returns the first received value from any of the channels and `nil` if successful.
// - Returns `nil` and an error if the context is canceled or if there are no more channels left to read from.
//
// See: [Go Concurrency Patterns: Timing out, moving on](https://go.dev/blog/concurrency-timeouts)
func MovingOn[T any](ctx context.Context, channels ...<-chan T) (T, error) {
	valueC := make(chan T, 1)
	errC := make(chan error, 1)
	go func(channels ...<-chan T) {
		defer close(valueC)
		defer close(errC)
		for len(channels) > 0 {
			index := randx.Intn(len(channels))
			ch := channels[index]
			select {
			case <-ctx.Done():
				errC <- ctx.Err()
				return
			case v, ok := <-ch:
				if !ok {
					channels = slicex.Delete(channels, index)
					runtime.Gosched()
					continue
				}
				select {
				case valueC <- v:
					return
				case <-ctx.Done():
					errC <- ctx.Err()
					return
				}
			default:
				runtime.Gosched()
			}
		}
	}(channels...)
	return <-valueC, <-errC
}

func Any[T any](ctx context.Context, channels ...<-chan T) (T, error) {
	return MovingOn[T](ctx, channels...)
}
