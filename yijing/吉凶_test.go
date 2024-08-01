package yijing

import "testing"


func Test吉凶_String(t *testing.T) {
    tests := []struct {
        jx 吉凶
        want string
    }{
        {吉凶{名: "大吉"}, "大吉"},
        {吉凶{名: "中吉"}, "中吉"},
        {吉凶{名: "小吉"}, "小吉"},
        {吉凶{名: "凶"}, "凶"},
    }
    for _, tt := range tests {
        if got := tt.jx.String(); got != tt.want {
            t.Errorf("吉凶.String() = %v, want %v", got, tt.want)
        }
    }
}