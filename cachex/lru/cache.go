package lru

import (
	lru "github.com/hashicorp/golang-lru"
)

// Cache LRU缓存
type Cache struct {
	LRUCache *lru.Cache
}

func (store *Cache) Get(key string) (interface{}, bool) {
	return store.LRUCache.Get(key)
}

func (store *Cache) Set(key string, val interface{}) {
	store.LRUCache.Add(key, val)
}
