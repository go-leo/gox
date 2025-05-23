package asyncbatchx

import (
	"errors"
	"sync"
	"time"
)

type Group[Obj any] struct {
	// mu：互斥锁，保证并发安全
	mu sync.Mutex
	// buf：存储待处理的对象切片
	buf []Obj
	// submitCh：提交信号通道
	submitCh chan struct{}
	// closed：标记是否已关闭
	closed bool
	// closedCh：关闭通知通道
	closedCh chan struct{}
	// wg：用于等待任务处理完成
	wg sync.WaitGroup
	// interval：提交间隔时间
	interval time.Duration
	// size：触发提交的批大小
	size int
	// f：处理函数，用于处理一批对象
	f func(objs []Obj)
}

// New 创建并返回一个新的Group对象，用于异步批量处理对象。
// 参数size指定每次批量处理的对象数量，interval指定批量处理的间隔时间。
// 函数f是在达到批量大小或时间间隔时被调用的处理函数。
func New[Obj any](size int, interval time.Duration, f func(objs []Obj)) *Group[Obj] {
	// 校验size参数必须大于0
	if size <= 0 {
		panic("asyncbatchx: size must be greater than 0")
	}
	// 校验interval参数必须大于0
	if interval <= 0 {
		panic("asyncbatchx: interval must be greater than 0")
	}

	// 初始化Group对象
	g := &Group[Obj]{
		mu:       sync.Mutex{},
		buf:      make([]Obj, 0, size),
		submitCh: make(chan struct{}, 1),
		closed:   false,
		closedCh: make(chan struct{}),
		wg:       sync.WaitGroup{},
		interval: interval,
		size:     size,
		f:        f,
	}

	// 启动一个goroutine来处理批量逻辑
	g.wg.Add(1)
	go g.loop()
	return g
}

// Submit 向组中提交一个对象。
// 该函数在内部使用互斥锁来确保线程安全，并检查组是否已关闭。
// 如果组已关闭，将返回错误。如果组未关闭且对象被成功添加到缓冲区，
// 当缓冲区大小达到阈值时，将通过提交通道发送信号。
func (g *Group[Obj]) Submit(obj Obj) error {
	// 加锁以保护组的状态和缓冲区
	g.mu.Lock()
	// 检查组是否已关闭
	if g.closed {
		// 如果组已关闭，解锁并返回错误
		g.mu.Unlock()
		return errors.New("asyncbatchx: group is closed")
	}
	// 将对象添加到缓冲区
	g.buf = append(g.buf, obj)
	// 仅当 buf 达到 size 时，才发送信号，避免空唤醒
	if len(g.buf) >= g.size {
		// 如果缓冲区大小达到阈值，解锁并发送信号
		g.mu.Unlock()
		select {
		case g.submitCh <- struct{}{}:
		default:
		}
	} else {
		// 如果缓冲区大小未达到阈值，仅解锁
		g.mu.Unlock()
	}
	// 操作成功，返回 nil
	return nil
}

// Close 关闭一个对象组。
//
// 本函数旨在确保组内的对象不会在组关闭后被添加或访问。
// 它通过关闭组的关闭通道来通知其他协程该组已关闭，以便它们可以相应地采取行动。
//
// 注意：本函数不接受参数，也不返回任何值。
func (g *Group[Obj]) Close() {
	// 上锁以确保组的关闭状态在多协程环境下能安全地被检查和修改。
	g.mu.Lock()
	// 如果组已经关闭，则解锁并直接返回，避免重复关闭。
	if g.closed {
		g.mu.Unlock()
		return
	}
	// 将组的关闭状态设置为true，表示该组现在已关闭。
	g.closed = true
	// 解锁以释放互斥锁，允许其他协程检查组的关闭状态。
	g.mu.Unlock()
	// 关闭组的关闭通道，通知所有监听该通道的协程组已关闭。
	close(g.closedCh)
	// 等待任务处理完成
	g.wg.Wait()
}

func (g *Group[Obj]) loop() {
	defer g.wg.Done()
	ticker := time.NewTicker(g.interval)
	defer ticker.Stop()
	for {
		select {
		// 当提交通道接收到信号时，检查缓冲区是否已满，然后处理
		case <-g.submitCh:
			g.mu.Lock()
			if len(g.buf) < g.size {
				g.mu.Unlock()
				continue
			}
			batch := g.buf
			g.buf = make([]Obj, 0, g.size)
			g.mu.Unlock()
			g.f(batch)
		// 当时间间隔到达时，处理缓冲区中的对象
		case <-ticker.C:
			g.mu.Lock()
			if len(g.buf) <= 0 {
				g.mu.Unlock()
				continue
			}
			batch := g.buf
			g.buf = make([]Obj, 0, g.size)
			g.mu.Unlock()
			g.f(batch)
		// 当关闭通道接收到信号时，处理剩余对象并退出
		case <-g.closedCh:
			g.mu.Lock()
			if len(g.buf) <= 0 {
				g.mu.Unlock()
				return
			}
			batch := g.buf
			g.mu.Unlock()
			g.f(batch)
			return
		}
	}
}
