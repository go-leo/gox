package stringx

import "strings"

func Max[S ~string](a, b S) S {
	if strings.Compare(string(a), string(b)) >= 0 {
		return a
	}
	return b
}

func Min[S ~string](a, b S) S {
	if strings.Compare(string(a), string(b)) <= 0 {
		return a
	}
	return b
}
