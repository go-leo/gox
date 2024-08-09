package chanx

import (
	"reflect"
)

// OrRecursion 接收多个输入通道，并返回一个输出通道。一旦任一输入通道有值，就关闭输出通道。
// 适用于等待任意一个异步操作完成的场景。
func OrRecursion[T any](channels ...<-chan T) <-chan T {
	//特殊情况，只有0个或者1个
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan T)
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2: // 2个也是一种特殊情况
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: //超过两个，二分法递归处理
			m := len(channels) / 2
			select {
			case <-OrRecursion(channels[:m]...):
			case <-OrRecursion(channels[m:]...):
			}
		}
	}()
	return orDone
}

// OrReflect 接收多个输入通道，并返回一个输出通道。一旦任一输入通道有值，就关闭输出通道。
// 适用于等待任意一个异步操作完成的场景。
func OrReflect[T any](channels ...<-chan T) <-chan T {
	//特殊情况，只有0个或者1个
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan T)
	go func() {
		defer close(orDone)
		// 利用反射构建SelectCase
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		// 随机选择一个可用的case
		_, _, _ = reflect.Select(cases)
	}()
	return orDone
}
