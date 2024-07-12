package groupx

import (
	"testing"
	"time"
)

func TestCountDownLatchWait(t *testing.T) {
	cdl := NewCountDownLatch(1)

	// 启动一个协程模拟等待完成
	go func() {
		time.Sleep(time.Millisecond * 100)
		cdl.CountDown()
	}()

	// 等待
	cdl.Wait()
}
