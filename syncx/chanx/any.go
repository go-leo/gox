package chanx

import (
	"context"
	"runtime"

	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/gox/slicex"
)

// Any the function's purpose is to randomly select from multiple provided channels and return
// a value once received from the chosen channel.
func Any[T any](ctx context.Context, channels ...<-chan T) T {
	if len(channels) <= 0 {
		var zero T
		return zero
	}

	valueC := make(chan T, 1)

	go func(channels []<-chan T) {
		defer close(valueC)
		for {
			i := randx.Intn(len(channels))
			ch := channels[i]
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch:
				if ok {
					select {
					case valueC <- v:
						return
					case <-ctx.Done():
						return
					}
				}
				channels = slicex.Delete(channels, i)
				runtime.Gosched()
			default:
				runtime.Gosched()
			}
		}
	}(channels)

	return <-valueC
}
