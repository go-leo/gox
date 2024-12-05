package chanx

import (
	"fmt"
)

var ErrDefaultBranch = fmt.Errorf("chanx: default branch")

// TrySend 用于尝试将值发送到指定的通道。
func TrySend[T any](in chan<- T, v T) error {
	select {
	case in <- v:
		return nil
	default:
		return ErrDefaultBranch
	}
}

// TryReceive 用于尝试从通道接收数据。
func TryReceive[T any](in <-chan T) (T, error) {
	var v T
	select {
	case v = <-in:
		return v, nil
	default:
		return v, ErrDefaultBranch
	}
}
