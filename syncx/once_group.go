package syncx

import (
	"sync"
)

type onceGroup struct {
	m sync.Map
}

func (o *onceGroup) Do(key string, f func()) {
	actual, loaded := o.m.LoadOrStore(key, &sync.Once{})
	if !loaded {
		return
	}
	once := actual.(*sync.Once)
	once.Do(f)
}
