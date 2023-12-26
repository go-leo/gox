package mapx

func KeySet[M ~map[K]V, R map[K]struct{}, K comparable, V any](m M) R {
	r := make(R, len(m))
	for k := range m {
		r[k] = struct{}{}
	}
	return r
}
