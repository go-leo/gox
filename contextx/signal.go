package contextx

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

type signalReceivedError struct{ incomingSignal os.Signal }

func (s signalReceivedError) Error() string {
	return fmt.Sprintf("the %s signal was received", s.incomingSignal.String())
}
func (s signalReceivedError) Timeout() bool   { return false }
func (s signalReceivedError) Temporary() bool { return false }

// Signal creates a new context that cancels on the given signals.
func Signal(signals ...os.Signal) (context.Context, context.CancelCauseFunc) {
	return WithSignal(context.Background(), signals...)
}

// WithSignal creates a new context that cancels on the given signals.
func WithSignal(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelCauseFunc) {
	newCtx, cancelFunc := context.WithCancelCause(ctx)
	if Error(ctx) != nil {
		return newCtx, cancelFunc
	}
	go func() {
		signalC := make(chan os.Signal)
		defer close(signalC)

		signal.Notify(signalC, signals...)
		defer signal.Stop(signalC)

		select {
		case incomingSignal := <-signalC:
			cancelFunc(signalReceivedError{incomingSignal: incomingSignal})
			return
		case <-newCtx.Done():
			return
		}
	}()
	return newCtx, cancelFunc
}
