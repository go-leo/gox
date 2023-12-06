package slicex

func SafeSlice[S ~[]E, E comparable](s S, start, length int) S {
	var r S
	if start < 0 || length < 0 {
		return r
	}
	low := start
	if len(s) < low {
		return r
	}
	high := start + length
	if len(s) < high {
		high = len(s)
	}
	return s[low:high]
}
