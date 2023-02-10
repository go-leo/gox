package signalx

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-leo/syncx"
	"golang.org/x/sync/errgroup"
)

type SignalWaiter struct {
	Signals     []os.Signal
	Hooks       []func(os.Signal)
	WaitTimeout time.Duration
}

func (w *SignalWaiter) Wait() error {
	signalC := make(chan os.Signal)

	// 监听一个系统信号
	signal.Notify(signalC, w.Signals...)

	// 接收到第一个信号
	receivedSignals := append([]os.Signal{}, <-signalC)

	// 并发调用所有hook
	errGroup, ctx := errgroup.WithContext(context.Background())
	for _, hook := range w.Hooks {
		errGroup.Go(func() error {
			syncx.BraveDo(func() { hook(receivedSignals[0]) }, func(p any) {})
			return nil
		})
	}
	go func() { _ = errGroup.Wait() }()

	// 等待退出
	select {
	case sig := <-signalC:
		// 如果再接收到一个信号就直接退出。
		receivedSignals := append([]os.Signal{}, sig)
		return SignalError{Signals: receivedSignals}
	case <-ctx.Done():
		// 等待hook执行完毕后退出
		return SignalError{Signals: receivedSignals}
	case <-time.After(w.WaitTimeout):
		// 等待超时退出
		return SignalError{Signals: receivedSignals}
	}
}

type SignalError struct {
	Signals []os.Signal
}

func (e SignalError) Error() string {
	return fmt.Sprintf("received [%s] signal", e.Signals)
}

func (e SignalError) Contains(signals []os.Signal) bool {
	for _, sig := range e.Signals {
		for _, curr := range signals {
			if sig.String() == curr.String() {
				return true
			}
		}
	}
	return false
}
