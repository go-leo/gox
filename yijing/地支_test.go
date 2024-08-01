package yijing

import (
	"testing"
)

// TestString方法，测试String方法的输出是否正确
func Test地支_String(t *testing.T) {
	tests := []struct {
		地支  地支
		期望 string
	}{
		{地支{名: "子"}, "子"},
		{地支{名: "丑"}, "丑"},
		{地支{名: "寅"}, "寅"},
		// ... 添加更多测试用例
	}

	for _, tt := range tests {
		if got := tt.地支.String(); got != tt.期望 {
			t.Errorf("地支.String() = %v, 期望 %v", got, tt.期望)
		}
	}
}