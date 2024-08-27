package osx

import "testing"

// TestWordSize 测试 WordSize 函数
func TestWordSize(t *testing.T) {
	// 调用 WordSize 函数
	size := WordSize()

	// 验证结果是否为预期值
	expectedSize := 64 // 假定你的计算机是64位的
	if size != expectedSize {
		t.Errorf("WordSize() = %d, want %d", size, expectedSize)
	}
}
