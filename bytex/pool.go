package bytex

import (
	"bytes"
	"math/bits"
	"sync"
)

// sizedPool 代表一个特定大小的缓冲池
type sizedPool struct {
	size int       // 该池中缓冲区的大小
	pool sync.Pool // 实际的同步池
}

// Get 从池中获取一个重置后的 bytes.Buffer
func (sp *sizedPool) Get() *bytes.Buffer {
	buf := sp.pool.Get().(*bytes.Buffer)
	buf.Reset() // 重置缓冲区以便复用
	return buf
}

// Put 将缓冲区重置后放回池中
func (sp *sizedPool) Put(b *bytes.Buffer) {
	b.Reset() // 清空缓冲区内容
	sp.pool.Put(b)
}

// newSizedPool 创建一个指定大小的缓冲池
func newSizedPool(size int) *sizedPool {
	return &sizedPool{
		size: size,
		pool: sync.Pool{
			New: func() any { return bytes.NewBuffer(make([]byte, size)) }, // 创建指定容量的新缓冲区
		},
	}
}

// Pool 是桶式内存池的主结构体
type Pool struct {
	minSize int          // 最小缓冲区大小
	maxSize int          // 最大缓冲区大小
	pools   []*sizedPool // 不同大小的缓冲池数组
}

// New 创建一个新的桶式内存池
// minSize: 最小缓冲区大小
// maxSize: 最大缓冲区大小
func New(minSize, maxSize int) *Pool {
	if maxSize < minSize {
		panic("maxSize can't be less than minSize") // 确保最大值不小于最小值
	}
	// 每个桶大小的倍数因子
	multiplier := 2
	var pools []*sizedPool
	curSize := minSize
	// 创建从 minSize 到 maxSize 的一系列缓冲池，每个池的大小是前一个的 multiplier 倍
	for curSize < maxSize {
		pools = append(pools, newSizedPool(curSize))
		curSize *= multiplier
	}
	pools = append(pools, newSizedPool(maxSize)) // 添加最大尺寸的池
	return &Pool{
		minSize: minSize,
		maxSize: maxSize,
		pools:   pools,
	}
}

// Get 根据请求的大小获取一个合适的 bytes.Buffer
func (p *Pool) Get(size int) *bytes.Buffer {
	sp := p.findPool(size) // 查找合适的池
	if sp == nil {
		// 如果找不到合适的池(请求大小超过 maxSize)，则直接创建新的缓冲区
		return bytes.NewBuffer(make([]byte, size))
	}
	return sp.Get() // 从找到的池中获取缓冲区
}

// Put 将缓冲区放回合适的池中
func (p *Pool) Put(b *bytes.Buffer) {
	sp := p.findPool(b.Cap()) // 根据缓冲区容量查找对应的池
	if sp == nil {
		return // 如果容量超过 maxSize，则不放回池中
	}
	sp.Put(b) // 将缓冲区放回池中
}

// findPool 根据给定大小查找最合适的缓冲池
func (p *Pool) findPool(size int) *sizedPool {
	if size > p.maxSize {
		return nil // 如果请求大小超过最大值，返回 nil
	}
	// 计算 size 除以 minSize 的商和余数
	div, rem := bits.Div64(0, uint64(size), uint64(p.minSize))
	// 计算商的二进制表示中的位数，用于确定池的索引
	idx := bits.Len64(div)
	// 如果整除且商是2的幂，则调整索引
	if rem == 0 && div != 0 && (div&(div-1)) == 0 {
		idx = idx - 1
	}
	return p.pools[idx]
}
