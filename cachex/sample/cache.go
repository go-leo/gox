package sample

import (
	"context"
	"sync"
)

// Cache 简单缓存
type Cache struct {
	Map sync.Map
}

func (store *Cache) Get(ctx context.Context, key string) (interface{}, bool) {
	return store.Map.Load(key)
}

func (store *Cache) Set(ctx context.Context, key string, val interface{}) {
	store.Map.Store(key, val)
}
