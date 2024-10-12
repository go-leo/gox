package retry

import (
	"context"
	"time"

	"github.com/go-leo/gox/backoff"
)

// Executor 接口定义了一个 Exec 方法，用于执行一个命令，并允许传递上下文和当前尝试次数。
type Executor interface {
	Exec(ctx context.Context, cmd func(ctx context.Context, attempt uint) error) error
}

// Retry 接口定义了一个 Backoff 方法，用于设置重试间隔策略。
type Retry interface {
	Backoff(backoffFunc backoff.BackoffFunc) Executor
}

// retryStrategy 结构体实现了 Retry 和 Executor 接口，并持有最大尝试次数和重试间隔策略函数。
type retryStrategy struct {
	maxAttempts uint
	backoffFunc backoff.BackoffFunc
}

func (r *retryStrategy) Backoff(backoffFunc backoff.BackoffFunc) Executor {
	r.backoffFunc = backoffFunc
	return r
}

func (r *retryStrategy) Exec(ctx context.Context, cmd func(ctx context.Context, attempt uint) error) error {
	var attempt uint
	for attempt < r.maxAttempts {
		// execute cmd
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

// MaxAttempts 函数创建一个具有指定最大重试次数的 retryStrategy 实例。
func MaxAttempts(maxAttempts uint) Retry {
	return &retryStrategy{
		maxAttempts: maxAttempts,
	}
}

// Call 函数执行一个命令，并允许传递上下文、最大尝试次数和重试间隔策略函数。
func Call(ctx context.Context, maxAttempts uint, backoffFunc backoff.BackoffFunc, method func(attemptTime int) error) error {
	return MaxAttempts(maxAttempts).Backoff(backoffFunc).Exec(ctx, func(ctx context.Context, attempt uint) error {
		return method(int(attempt))
	})
}
