package groupx

import "context"

// WaitNotify waiter.Wait()
func WaitNotify(waiter interface{ Wait() }) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		defer close(c)
		waiter.Wait()
	}()
	return c
}

// WaitNotifyE waiter.Wait()
func WaitNotifyE(waiter interface{ Wait() error }) <-chan error {
	c := make(chan error, 1)
	go func() {
		defer close(c)
		if err := waiter.Wait(); err != nil {
			c <- err
		}
	}()
	return c
}

// WaitContentNotifyE waiter.Wait()
func WaitContentNotifyE(ctx context.Context, waiter interface{ Wait(context.Context) error }) <-chan error {
	c := make(chan error, 1)
	go func() {
		defer close(c)
		if err := waiter.Wait(ctx); err != nil {
			c <- err
		}
	}()
	return c
}
