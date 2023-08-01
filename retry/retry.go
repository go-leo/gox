package retry

import (
	"context"
	"time"

	"github.com/go-leo/gox/backoff"
)

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

func Endpoint[Req any, Resp any](ctx context.Context, req Req, maxAttempts uint, backoffFunc backoff.BackoffFunc, endpoint func(ctx context.Context, req Req) (Resp, error)) (Resp, error) {
	for i := uint(0); i < maxAttempts; i++ {
		resp, err := endpoint(ctx, req)
		// return if err is nil.
		if err == nil {
			return resp, err
		}
		// sleep and wait retry
		time.Sleep(backoffFunc(ctx, i+1))
	}

	return endpoint(ctx, req)
}
