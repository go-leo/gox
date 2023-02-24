package contextx

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
)

type signalReceivedError struct{ incomingSignal os.Signal }

func (s signalReceivedError) Error() string {
	return fmt.Sprintf("signal received, %s", s.incomingSignal)
}
func (s signalReceivedError) Timeout() bool   { return false }
func (s signalReceivedError) Temporary() bool { return false }

// Signal creates a new context that cancels on the given signals.
func Signal(signals ...os.Signal) (context.Context, context.CancelCauseFunc) {
	return WithSignal(context.Background(), signals...)
}

// WithSignal creates a new context that cancels on the given signals.
func WithSignal(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelCauseFunc) {
	ctx, cancelFunc := context.WithCancelCause(ctx)
	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, signals...)
	go func() {
		defer signal.Stop(signalC)
		defer close(signalC)
		select {
		case incomingSignal := <-signalC:
			cancelFunc(signalReceivedError{incomingSignal: incomingSignal})
		case <-ctx.Done():
			cancelFunc(errors.New("received context done"))
		}
	}()
	return ctx, cancelFunc
}
