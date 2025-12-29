package distributed

import "math/rand"

// DeterministicSubSetting 根据给定的客户端ID生成后端列表的一个确定性子集。
// 子集的选择基于客户端的ID和子集的大小，确保相同的客户端ID选择相同的子集。
//
// 参数:
//
//	backends: 所有后端的列表。
//	clientID: 客户端的唯一标识符，用于确定选择哪个子集。
//	subsetSize: 每个子集的大小。
//
// 返回值:
//
//	包含选定子集后端的字符串切片。
//
// See: [google sre](https://landing.google.com/sre/book/chapters/load-balancing-datacenter.html)
// See: [用 subsetting 限制连接池中的连接数量](https://xargin.com/limiting-conn-wih-subset)
func DeterministicSubSetting(backends []string, clientID, subsetSize int) []string {
	// 验证参数有效性
	if subsetSize <= 0 || len(backends) == 0 {
		return nil
	}

	// 计算子集的数量，即，可以从后端总数中划分出多少个子集。
	subsetCount := len(backends) / subsetSize

	// 处理余数，确保充分利用所有后端
	if len(backends)%subsetSize != 0 {
		subsetCount++
	}

	// 处理clientID超出子集范围的情况
	if clientID >= subsetCount*subsetSize {
		return nil
	}

	// 根据clientID确定轮次，确保同一轮次使用相同的随机源。
	// 将客户端分组到轮次中；每一轮次使用相同的洗牌列表：
	round := clientID / subsetCount

	// 根据轮次初始化随机源，用于洗牌后端列表。
	r := rand.New(rand.NewSource(int64(round)))
	// 洗牌后端列表，确保每一轮次有不同的顺序。
	r.Shuffle(len(backends), func(i, j int) { backends[i], backends[j] = backends[j], backends[i] })

	// 计算当前客户端的子集ID，用于确定选定子集的起始索引。
	// 对应当前客户端的子集ID：
	subsetID := clientID % subsetCount

	// 根据子集ID和子集大小计算选定子集的起始索引。
	start := subsetID * subsetSize
	// 根据起始索引和子集大小返回选定的子集。
	return backends[start : start+subsetSize]
}

// https://colobu.com/2016/03/22/jump-consistent-hash/
func JumpHash(key uint64, buckets int, checkAlive func(int) bool) int {
	var b, j int64 = -1, 0
	if buckets <= 0 {
		buckets = 1
	}
	for j < int64(buckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}
	if checkAlive != nil && !checkAlive(int(b)) {
		return JumpHash(key+1, buckets, checkAlive) // 最好设置深度，避免key+1一直返回当掉的服务器
	}
	return int(b)
}
