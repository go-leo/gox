//go:build !go1.20

package contextx

import (
	"context"
	"os"
	"os/signal"
)

// Signal creates a new context that cancels on the given signals.
func Signal(signals ...os.Signal) (context.Context, context.CancelFunc) {
	return WithSignal(context.Background(), signals...)
}

// WithSignal creates a new context that cancels on the given signals.
func WithSignal(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	newCtx, cancelFunc := context.WithCancel(ctx)
	if Error(ctx) != nil {
		return newCtx, cancelFunc
	}
	go func() {
		signalC := make(chan os.Signal, 1)
		defer close(signalC)

		signal.Notify(signalC, signals...)
		defer signal.Stop(signalC)

		select {
		case <-signalC:
			cancelFunc()
			return
		case <-newCtx.Done():
			return
		}
	}()
	return newCtx, cancelFunc
}
