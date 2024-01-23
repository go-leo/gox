package syncx

func WaitNotify(waiter interface{ Wait() }) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		defer close(c)
		waiter.Wait()
	}()
	return c
}

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
