package chanx

// AsSlice 将一个只读通道中的所有元素收集并转换为切片返回。
func AsSlice[T any](c <-chan T) []T {
	var ts []T
	for t := range c {
		ts = append(ts, t)
	}
	return ts
}
