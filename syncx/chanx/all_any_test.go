package chanx

import (
	"context"
	"testing"
	"time"

	"golang.org/x/exp/slices"
)

// TestAll 测试 All 函数是否能正确收集所有通道的值。
func TestAll(t *testing.T) {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// 创建一些模拟通道并发送一些值
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()
	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()
	go func() {
		ch3 <- 5
		ch3 <- 6
		close(ch3)
	}()

	// 调用 All 函数收集值
	values := All(context.Background(), ch1, ch2, ch3)
	ints := <-values

	// 检查收集的值是否正确
	expected := []int{1, 2, 3, 4, 5, 6}
	if len(ints) != len(expected) {
		t.Errorf("All returned %d values, want %d", len(values), len(expected))
	}
	for i, v := range ints {
		if !slices.Contains(expected, v) {
			t.Errorf("All returned %d at index %d, want %d", v, i, expected[i])
		}
	}

	// 给一些时间让所有值都通过通道，避免测试提前结束。
	time.Sleep(100 * time.Millisecond)
}

// TestAllWithTimeout 测试 All 函数是否在 context 超时后返回。
func TestAllWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	ch := make(chan int)

	// 不向通道发送任何值，直接让它阻塞

	// 调用 All 函数，期望它会在超时后返回
	values := All(ctx, ch)

	// 检查是否没有值被收集到
	if len(<-values) != 0 {
		t.Errorf("All returned %d values, want 0", len(values))
	}
}

// TestAllWithEmptyChannels 测试 All 函数是否能处理空通道。
func TestAllWithEmptyChannels(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch1 := make(chan int)
	ch2 := make(chan int)

	// 调用 All 函数收集值，但通道是空的
	values := All(ctx, ch1, ch2)

	// 检查是否没有值被收集到
	if len(<-values) != 0 {
		t.Errorf("All returned %d values, want 0", len(values))
	}
}

func TestAny(t *testing.T) {
	type arg int // 使用别名来模拟泛型参数

	// 创建一个测试上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 创建几个输入通道，并给它们起个名字
	ch1 := make(chan arg, 1)
	ch2 := make(chan arg, 1)
	ch3 := make(chan arg, 1)

	// 待会要发送的数据
	data1 := arg(1)
	data2 := arg(2)
	data3 := arg(3)

	// 发送数据到通道的goroutine
	go func() {
		time.Sleep(time.Millisecond * 500) // 模拟延迟
		ch1 <- data1
	}()
	go func() {
		time.Sleep(time.Millisecond * 100) // 模拟延迟
		ch2 <- data2
	}()
	go func() {
		time.Sleep(time.Millisecond * 300) // 模拟延迟
		ch3 <- data3
	}()

	out := Any(ctx, ch1, ch2, ch3)

	// 设置一个时间限制，以防止在没有数据的情况下无限期地阻塞
	select {
	case res := <-out:
		t.Logf("Race won: %v", res)
	case <-ctx.Done():
		t.Error("Race context done")
	}

	// 关闭通道
	close(ch1)
	close(ch2)
	close(ch3)
}
