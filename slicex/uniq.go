package slicex

func Uniq[S ~[]E, E comparable](s S) S {
	if s == nil {
		return nil
	}
	m := make(map[E]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	r := make(S, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}
