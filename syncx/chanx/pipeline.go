package chanx

import "context"

// Pipeline 创建一个数据处理管道。
// 它接收一个输入通道 in 和一个处理函数 f，返回一个输出通道 out。
// 在新的goroutine中，遍历输入通道中的每个元素，应用处理函数后将结果发送到输出通道。
// 如果上下文 ctx 被取消，则提前退出并关闭输出通道。
func Pipeline[T any, R any](ctx context.Context, in <-chan T, f func(T) R) <-chan R {
	out := make(chan R, len(in))
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case out <- f(v):
				}
			}
		}
	}()
	return out
}
