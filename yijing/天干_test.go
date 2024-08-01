package yijing

import (
	"testing"
)

// TestString tests the String method of 天干 struct.
func Test天干_String(t *testing.T) {
	tests := []struct {
		tg   天干
		want string
	}{
		{天干{名: "甲"}, "甲"},
		{天干{名: "乙"}, "乙"},
		{天干{名: "丙"}, "丙"},
		// Add more test cases here...
	}

	for _, tt := range tests {
		t.Run(tt.tg.名, func(t *testing.T) {
			if got := tt.tg.String(); got != tt.want {
				t.Errorf("天干.String() = %v, want %v", got, tt.want)
			}
		})
	}
}