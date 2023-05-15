package randx

import (
	"math/rand"
	"sync"
)

func SyncSource(seed int64) rand.Source {
	source := rand.NewSource(seed)
	return &syncSource{
		lk:    sync.Mutex{},
		src:   source,
		src64: source.(rand.Source64),
	}
}

type syncSource struct {
	lk    sync.Mutex
	src   rand.Source
	src64 rand.Source64
}

func (r *syncSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *syncSource) Uint64() (n uint64) {
	r.lk.Lock()
	n = r.src64.Uint64()
	r.lk.Unlock()
	return
}

func (r *syncSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}
