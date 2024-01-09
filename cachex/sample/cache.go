package sample

import "sync"

// Cache 简单缓存
type Cache struct {
	Cache sync.Map
}

func (store *Cache) Get(key string) (interface{}, bool) {
	return store.Cache.Load(key)
}

func (store *Cache) Set(key string, val interface{}) {
	store.Cache.Store(key, val)
}
