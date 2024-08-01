package yijing

import "testing"

// 单元测试
func Test五行_克(t *testing.T) {
	tests := []struct {
		wuxing 五行
		want    五行
	}{
		{金, 木},
		{木, 土},
		{水, 火},
		{火, 金},
		{土, 水},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.wuxing.克(); got != tt.want {
				t.Errorf("%v 克() = %v, want %v", tt.wuxing, got, tt.want)
			}
		})
	}
}