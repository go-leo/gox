package mutexx

import "sync"

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
