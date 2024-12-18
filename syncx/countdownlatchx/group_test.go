package countdownlatchx

import (
	"testing"
	"time"
)

func TestGroup(t *testing.T) {
	cdl := NewGroup(2)

	// 启动一个协程模拟等待完成
	go func() {
		time.Sleep(time.Millisecond * 100)
		cdl.CountDown()
	}()
	go func() {
		time.Sleep(time.Millisecond * 500)
		cdl.CountDown()
	}()
	// 等待
	cdl.Wait()
}
