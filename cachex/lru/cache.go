package lru

import (
	lru "github.com/hashicorp/golang-lru"
)

// Cache LRU缓存
type Cache struct {
	Cache *lru.Cache
}

func (store *Cache) Get(key string) (interface{}, bool) {
	return store.Cache.Get(key)
}

func (store *Cache) Set(key string, val interface{}) {
	store.Cache.Add(key, val)
}
