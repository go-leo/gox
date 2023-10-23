package contextx

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

// Signal creates a new context that cancels on the given signals.
func Signal(signals ...os.Signal) (context.Context, context.CancelFunc) {
	return WithSignal(context.Background(), signals...)
}

// WithSignal like signal.NotifyContext.
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

type SignalError struct {
	Signal os.Signal
}

func (e SignalError) Error() string {
	return fmt.Sprintf("received [%s] signal", e.Signal)
}

// SignalCause creates a new context that cancels on the given signals.
func SignalCause(signals ...os.Signal) (context.Context, context.CancelCauseFunc) {
	return WithSignalCause(context.Background(), signals...)
}

func WithSignalCause(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelCauseFunc) {
	newCtx, cancelFunc := context.WithCancelCause(ctx)
	if Error(ctx) != nil {
		return newCtx, cancelFunc
	}
	go func() {
		signalC := make(chan os.Signal, 1)
		defer close(signalC)

		signal.Notify(signalC, signals...)
		defer signal.Stop(signalC)

		select {
		case sig := <-signalC:
			cancelFunc(SignalError{Signal: sig})
			return
		case <-newCtx.Done():
			return
		}
	}()
	return newCtx, cancelFunc
}
