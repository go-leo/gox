package contextx

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

// Signal creates a new context that cancels on the given signals.
func Signal(signals ...os.Signal) (context.Context, context.CancelCauseFunc) {
	return WithSignal(context.Background(), signals...)
}

// WithSignal creates a new context that cancels on the given signals.
func WithSignal(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelCauseFunc) {
	ctx, cancelFunc := context.WithCancelCause(ctx)
	signalC := make(chan os.Signal, 1)
	defer close(signalC)
	signal.Notify(signalC, signals...)
	defer signal.Stop(signalC)
	go func() {
		select {
		case incomingSignal := <-signalC:
			cancelFunc(fmt.Errorf("receive signal, %s", incomingSignal))
		case <-ctx.Done():
			signal.Stop(signalC)
		}
	}()
	return ctx, cancelFunc
}
