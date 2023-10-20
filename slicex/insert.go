package slicex

// Insert 在切片的指定位置插入元素
func Insert[S ~[]E, E any](s S, i int, e E) S {
	// 创建一个新的切片，长度比原切片多1
	r := make(S, len(s)+1)

	// 将原切片的前index个元素复制到新切片
	copy(r, s[:i])

	// 插入要插入的元素
	r[i] = e

	// 将原切片的index及之后的元素复制到新切片
	copy(r[i+1:], s[i:])

	return r
}
