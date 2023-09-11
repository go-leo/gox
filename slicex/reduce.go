package slicex

func Reduce[S ~[]E, E any, R any](s S, initValue R, f func(previousValue R, currentValue E, currentIndex int, s S) R) R {
	var r = initValue
	for i, e := range s {
		r = f(r, e, i, s)
	}
	return r
}
