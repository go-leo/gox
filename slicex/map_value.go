package slicex

// MapElem 两个切片s1，s2,，获取v所有s1同位置的s2的值
func MapElem[S1 ~[]E1, S2 ~[]E2, E1 comparable, E2 any](v E1, s1 S1, s2 S2) E2 {
	var r E2
	if len(s1) != len(s2) {
		panic("length of s1 and s2 not equal")
	}
	for i := range s1 {
		if s1[i] == v {
			return s2[i]
		}
	}
	return r
}
