package cyclicbarrierx

import (
	"fmt"
	"sync"
	"testing"
)

func TestGroup(t *testing.T) {
	// 定义需要到达屏障的线程数量
	parties := 5
	// 定义屏障动作
	barrierAction := func() {
		println("所有线程已到达屏障")
	}
	// 创建 CyclicBarrier 实例
	cb := NewGroup(parties, barrierAction)

	wg := sync.WaitGroup{}
	wg.Add(parties)
	// 模拟多个线程到达屏障
	for i := 0; i < parties; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("线程 %d 已到达屏障\n", i)
			cb.Wait()
			fmt.Printf("线程 %d 已执行屏障动作\n", i)
		}(i)
	}

	wg.Wait()
}
