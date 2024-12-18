package countdownlatchx

import (
	"context"
	"sync"
)

// Group 实现了一个计数器，当计数器归零时，等待的函数将被释放。
// 它是线程安全的，可以在多个goroutine中使用。
type Group struct {
	wg sync.WaitGroup
}

// NewGroup 创建一个新的CountDownLatch实例，初始化计数器为delta。
func NewGroup(delta int) *Group {
	cdl := &Group{wg: sync.WaitGroup{}}
	cdl.wg.Add(delta)
	return cdl
}

// CountDown 将计数器减一。当计数器归零时，等待的函数将被释放。
func (cdl *Group) CountDown() {
	cdl.wg.Done()
}

// Wait 等待计数器归零。
func (cdl *Group) Wait() {
	cdl.wg.Wait()
}

// WaitContext 等待计数器归零，同时监听context的取消事件。
// 如果context被取消，将返回context的错误。
func (cdl *Group) WaitContext(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		cdl.Wait()
		close(done) // 使用通道关闭来通知主goroutine计数器已归零。
	}()

	select {
	case <-ctx.Done():
		// 如果context被取消，则返回取消错误。
		return ctx.Err()
	case <-done:
		// 如果计数器归零，正常返回。
		return nil
	}
}
