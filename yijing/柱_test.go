package yijing

import (
	"testing"
)

func Test柱_String(t *testing.T) {
	tests := []struct {
		z     柱
		want  string
	}{
		{柱{甲, 子}, "天干: 甲, 地支: 子"},
		// 添加更多测试用例
	}

	for _, tt := range tests {
		if got := tt.z.String(); got != tt.want {
			t.Errorf("柱.String() = %v, want %v", got, tt.want)
		}
	}
}