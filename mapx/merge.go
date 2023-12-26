package mapx

func Append[M ~map[K]V, K comparable, V any](m M, ms ...M) M {
	if m == nil {
		return m
	}
	for _, elem := range ms {
		for k, v := range elem {
			m[k] = v
		}
	}
	return m
}
