package contextx

import (
	"context"
	"os"
	"os/signal"
)

// SignalContext creates a new context that cancels on the given signals.
func SignalContext(signals ...os.Signal) (context.Context, context.CancelFunc) {
	return WithSignal(context.Background(), signals...)
}

// WithSignal creates a new context that cancels on the given signals.
func WithSignal(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, cancelFunc := context.WithCancel(ctx)
	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, signals...)
	go func() {
		select {
		case <-signalC:
			cancelFunc()
		case <-ctx.Done():
			signal.Stop(signalC)
		}
	}()
	return ctx, cancelFunc
}
