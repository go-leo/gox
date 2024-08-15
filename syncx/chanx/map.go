package chanx

import "context"

// Map 函数接收一个输入通道和一个转换函数，将输入通道的每个元素通过转换函数处理后放入新通道，并返回此新通道。
func Map[T any, R any](in <-chan T, mapper func(T) R) <-chan R {
	var out chan R
	if in == nil {
		return out
	}
	out = make(chan R, cap(in))
	for value := range in {
		out <- mapper(value)
	}
	close(out)
	return out
}

// AsyncMap 函数AsyncMap接收数据通道和映射函数，异步处理通道中的每个元素，并通过另一个通道输出处理结果。
// 当输入通道关闭或上下文取消时，处理停止。若输入通道为空，则直接返回空输出通道。
func AsyncMap[T any, R any](ctx context.Context, in <-chan T, mapper func(T) R) <-chan R {
	var out chan R
	if in == nil {
		return out
	}
	out = make(chan R, cap(in))
	go func() {
		defer close(out)
		for {
			select {
			case value, ok := <-in:
				if !ok {
					return
				}
				out <- mapper(value)
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func FlatMap() {

}
