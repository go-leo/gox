package knuthx

import (
	"strconv"
	"testing"
)

// TestKnuth 测试 Knuth 函数的预期行为。
func TestKnuth(t *testing.T) {
	// 定义一组测试用例
	testCases := []struct {
		input    uint
		expected uint
	}{
		{0, 0},                // 测试边界条件
		{1, 1481765933},       // 测试一个小于32位的非零值
		{42, 2104627054},      // 测试一个典型的非零值
		{1 << 32, 1284865837}, // 测试一个32位的值
		{1 << 33, 2569731674}, // 测试一个大于32位的值
	}

	for _, tc := range testCases {
		t.Run("Input_"+strconv.FormatUint(uint64(tc.input), 10), func(t *testing.T) {
			actual := Knuth(tc.input)
			if actual != tc.expected {
				t.Errorf("Knuth(%d) = %d; expected %d", tc.input, actual, tc.expected)
			}
		})
	}
}
