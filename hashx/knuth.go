package hashx

import (
	"unsafe"
)

// 优化取模运算以减少偏差
// 使用 Knuth's multiplicative method 来计算索引
// 参考: https://en.wikipedia.org/wiki/Knuth_multiplicative_method
//
// 假设我们有一个哈希值 ( x )，需要将其映射到一个范围 ([0, m-1]) 内的索引值，其中 ( m ) 是目标范围的大小。使用 Knuth's multiplicative 方法，我们可以按照以下步骤计算索引值 ( y )：
// 选择一个乘数 ( a )。
// 计算 ( x * a )。
// 将结果向右移位 ( w - p ) 位，其中 ( w ) 是计算机字长的位数（例如 32 位或 64 位），( p ) 是目标范围 ( m ) 所需的位数。
// 对目标范围的大小 ( m ) 取模，即 ( y = ((x * a) >> (w - p)) % m )。
//
// 选择合适的乘数 ( a ) 对于 Knuth's multiplicative 方法来说非常重要，它直接影响到生成的伪随机数的质量。在 32 位和 64 位系统中，选择不同的乘数可以使方法更加有效。
// 32 位系统
// 在 32 位系统中，通常使用的乘数是 2654435761，这是一个常用的乘数值，它与 32 位计算机的字长相匹配。这个乘数的选择是基于它是一个大素数，并且与 32 位整数的最大值相接近，因此在 32 位系统中使用它是合理的。
// 64 位系统
// 在 64 位系统中，我们需要选择一个与 64 位整数范围相匹配的乘数。一个常用的 64 位乘数是 6364136223846793005，这也是一个大素数，并且接近 64 位整数的最大值。
const (
	knuthMultiplier32 uint64 = 2654435761
	knuthMultiplier64 uint64 = 6364136223846793005
)

func Knuth(hashSum uint) uint {
	var multiplier uint64
	switch unsafe.Sizeof(uintptr(0)) {
	case 4:
		multiplier = knuthMultiplier32
	case 8:
		multiplier = knuthMultiplier64
	default:
		panic("hashx: unsupported architecture")
	}
	return uint((multiplier * uint64(hashSum)) >> 32)
}
