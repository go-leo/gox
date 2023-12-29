package mapx

func AppendValue[M ~map[K]V, R ~map[K]S, S ~[]V, K comparable, V any](r R, ms ...M) R {
	if r == nil {
		r = make(R)
	}
	for _, elem := range ms {
		for k, v := range elem {
			r[k] = append(r[k], v)
		}
	}
	return r
}
