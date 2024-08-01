package yijing

import (
	"testing"
)


// TestString 测试阴阳的String方法
func Test阴阳_String(t *testing.T) {
	testCases := []struct {
		name     string
		yinYang  阴阳
		expected string
	}{
		{"测试案例1", 阴阳{"阴"}, "阴"},
		{"测试案例2", 阴阳{"阳"}, "阳"},
		{"测试案例3", 阴阳{"太极"}, "太极"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.yinYang.String()
			if actual != tc.expected {
				t.Errorf("预期结果为 %s, 实际结果为 %s", tc.expected, actual)
			}
		})
	}
}