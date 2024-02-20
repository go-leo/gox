package sample

import "sync"

// Cache 简单缓存
type Cache struct {
	Map sync.Map
}

func (store *Cache) Get(key string) (interface{}, bool) {
	return store.Map.Load(key)
}

func (store *Cache) Set(key string, val interface{}) {
	store.Map.Store(key, val)
}
