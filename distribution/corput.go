package distribution

// VanDerCorputSequence 生成范德科鲁普序列。
// 该序列是一种低差异序列，适用于在均匀分布中生成数值。
//
// 参数:
//
//	n: 序列中的位置，从0开始。
//	base: 序列的基数，用于确定序列的分形性质。
//
// 返回值:
//
//	生成的范德科鲁普序列的浮点数值。
func VanDerCorputSequence(n int, base int) float64 {
	// 初始化q为0，用于累加每个位的倒数权重。
	q := 0.0
	// 初始化bk为基数的倒数，用于计算每一位的权重。
	bk := 1.0 / float64(base)
	// 循环处理每一位直到n为0。
	for n > 0 {
		// 将当前位的值乘以权重bk加到q上。
		q += float64(n%base) * bk
		// 移动到下一位，更新n和bk。
		n /= base
		bk /= float64(base)
	}
	// 返回计算得到的范德科鲁普序列值。
	return q
}

// BinaryVanDerCorputSequence 生成二进制范德科鲁普序列。
// 该函数是范德科鲁普序列的一个特例，针对基数为2的情况。
// 范德科鲁普序列是一种低差异序列，常用于随机数生成和蒙特卡洛模拟。
// 在二进制情况下，这种序列特别适用于需要均匀分布的二进制位场景。
//
// 参数:
//
//	n - 序列中的第n个数的索引。
//
// 返回值:
//
//	生成的范德科鲁普序列中第n个数的值。
func BinaryVanDerCorputSequence(n int) float64 {
	return VanDerCorputSequence(n, 2)
}
