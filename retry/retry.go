package retry

import (
	"context"
	"time"

	"github.com/go-leo/gox/backoff"
)

type Executor interface {
	Exec(ctx context.Context, cmd func(ctx context.Context, attempt uint) error) error
}

type Retry interface {
	Backoff(backoffFunc backoff.BackoffFunc) Executor
}

type retryStrategy struct {
	maxAttempts uint
	backoffFunc backoff.BackoffFunc
}


// 实现WithBackoff接口中的Backoff方法
func (r *retryStrategy) Backoff(backoffFunc backoff.BackoffFunc) Executor {
	r.backoffFunc = backoffFunc
    return r
}

func (r *retryStrategy) Exec(ctx context.Context, cmd func(ctx context.Context, attempt uint) error) error {
	var attempt uint
	for attempt < r.maxAttempts {
		// execte cmd
		err := cmd(ctx, attempt)
		// return if err is nil.
		if err == nil {
			return nil
		}
		// increase the number of attempts
		attempt++
		// sleep and wait retry
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(r.backoffFunc(ctx, attempt))
		}
	}
	// perform the last execution
	return cmd(ctx, attempt)
}

func MaxAttempts(maxAttempts uint) Retry {
	return &retryStrategy{
		maxAttempts: maxAttempts,
	}	
}

func Call(ctx context.Context, maxAttempts uint, backoffFunc backoff.BackoffFunc, method func(attemptTime int) error) error {
	var err error
	max := int(maxAttempts)
	for i := 0; i <= max; i++ {
		// call method
		err = method(i)

		// if method not return error, no need to retry
		if err == nil {
			break
		}

		// If the maximum number of attempts is exceeded, no need to retry
		if i >= max {
			break
		}

		// sleep and wait retry
		time.Sleep(backoffFunc(ctx, uint(i+1)))
	}
	return err
}
