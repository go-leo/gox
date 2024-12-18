package cyclicbarrierx

import (
	"context"
	"sync"
)

// Group 实现了一个可重用的屏障，等待一组线程到达指定点后一起继续执行。
// 当所有等待的线程都到达屏障时，会执行一个指定的屏障动作。
type Group struct {
	barrierWg     sync.WaitGroup // 用于同步屏障动作的等待组
	partiesWg     sync.WaitGroup // 用于同步到达屏障的线程的等待组
	parties       int            // 需要到达屏障的线程数量
	barrierAction func()         // 在所有线程到达屏障后执行的动作
}

// NewGroup 创建并返回一个新的 CyclicBarrier 实例。
// 参数 parties 指定需要到达屏障的线程数量。
// 参数 barrierAction 是一个函数，当所有线程都到达屏障时会被执行。
func NewGroup(parties int, barrierAction func()) *Group {
	if parties < 1 {
		panic("groupx: parties must be greater than 0")
	}
	cb := &Group{
		barrierWg:     sync.WaitGroup{},
		partiesWg:     sync.WaitGroup{},
		parties:       parties,
		barrierAction: barrierAction,
	}
	cb.init()
	return cb
}

func (cb *Group) init() {
	// 初始化 barrierWg，表示屏障动作开始前的一个等待单位
	cb.barrierWg.Add(1)
	// 初始化 partiesWg，并为每个参与线程增加计数
	cb.partiesWg.Add(cb.parties)
	// 启动一个 goroutine，在所有线程到达屏障后执行屏障动作
	go func() {
		defer cb.barrierWg.Done()
		cb.partiesWg.Wait()
		if cb.barrierAction != nil {
			cb.barrierAction()
		}
	}()
}

// Wait 方法用于表示一个线程到达了屏障。
// 线程调用此方法后，会等待直到所有线程都到达屏障，然后一起继续执行。
func (cb *Group) Wait() {
	cb.partiesWg.Done()
	cb.barrierWg.Wait()
}

// WaitContext 方法与 Wait 类似，但支持通过 context 上下文取消等待。
// 如果上下文被取消，方法会返回相应的错误；否则，它会等待直到所有线程到达屏障。
func (cb *Group) WaitContext(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		cb.Wait()
		close(done)
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}
}
