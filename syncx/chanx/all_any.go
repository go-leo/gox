package chanx

import (
	"context"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/gox/slicex"
	"runtime"
)

// All 将多个输入通道合并为一个输出通道，然后返回一个包含所有输入通道当前值的切片的通道
// See: [Go Concurrency Patterns: Timing out, moving on](https://go.dev/blog/concurrency-timeouts)
func All[T any](ctx context.Context, ins ...<-chan T) <-chan []T {
	return AsSlice(ctx, FanIn(ctx, ins...))
}

// Any 从多个输入通道中任意选择一个值，返回一个包含此值的一个通道，
func Any[T any](ctx context.Context, ins ...<-chan T) <-chan T {
	out := make(chan T, 1)
	go _any(ctx, out, ins...)
	return out
}

// _any 用于从多个输入通道中选择一个值并发送到 valueC 通道，或者在上下文取消时将错误发送到 errC 通道。
func _any[T any](ctx context.Context, out chan T, ins ...<-chan T) {
	// 确保在函数退出时关闭 out 和 errC 通道。
	defer close(out)
	for len(ins) > 0 {
		// 当 ins 列表不为空时，随机选择一个通道 ch。
		sel := randx.Intn(len(ins))
		ch := ins[sel]
		select {
		case <-ctx.Done():
			return
		case v, ok := <-ch:
			if !ok {
				// ch被关闭，则从 ins 列表中删除 ch，并继续循环。
				ins = slicex.RemoveAt(ins, sel)
				continue
			}
			select {
			case out <- v:
				return
			case <-ctx.Done():
				return
			}
		default:
			// 防止死锁
			runtime.Gosched()
		}
	}
}
