package syncx

import (
	"runtime"
	"sync"
	"sync/atomic"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	maxBackoff  = 16
)

type SpinMutex struct {
	state int32
}

func (m *SpinMutex) Lock() {
	backoff := 1
	for !atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}
		if backoff < maxBackoff {
			backoff <<= 1
		}
	}
}

func (m *SpinMutex) Unlock() {
	atomic.StoreInt32(&m.state, 0)
}

type ChanMutex struct {
	state chan struct{}
	once  sync.Once
}

func NewChanMutex() *ChanMutex {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	return &ChanMutex{
		state: ch,
	}
}

func (m *ChanMutex) Lock() {
	<-m.state
}

func (m *ChanMutex) Unlock() {
	m.state <- struct{}{}
}
