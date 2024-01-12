package slicex

func ToSet[S ~[]E, M ~map[E]struct{}, E comparable](s S) M {
	r := make(M)
	for _, e := range s {
		r[e] = struct{}{}
	}
	return r
}
