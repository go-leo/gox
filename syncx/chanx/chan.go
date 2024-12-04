package chanx

import (
	"context"
	"fmt"
)

var ErrDefaultBranch = fmt.Errorf("chanx: default branch")

// TrySend 用于尝试将值发送到指定的通道。
func TrySend[T any](ctx context.Context, in chan<- T, v T) error {
	select {
	case in <- v:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	default:
		return ErrDefaultBranch
	}
}

// TryReceive 用于尝试从通道接收数据。
func TryReceive[T any](ctx context.Context, in <-chan T) (T, error) {
	var v T
	select {
	case v = <-in:
		return v, nil
	case <-ctx.Done():
		return v, ctx.Err()
	default:
		return v, ErrDefaultBranch
	}
}
