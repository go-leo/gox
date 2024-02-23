package lru

import (
	"context"
	"errors"
	"github.com/go-leo/gox/cachex"
	lru "github.com/hashicorp/golang-lru"
)

var (
	ErrEvicted = errors.New("eviction occurred")
)

var _ cachex.Store = (*Cache)(nil)

// Cache LRU缓存
type Cache struct {
	LRUCache *lru.Cache
}

func (store *Cache) Get(ctx context.Context, key string) (any, error) {
	val, ok := store.LRUCache.Get(key)
	if !ok {
		return nil, cachex.Nil
	}
	return val, nil
}

func (store *Cache) Set(ctx context.Context, key string, val any) error {
	if store.LRUCache.Add(key, val) {
		return ErrEvicted
	}
	return nil
}

func (store *Cache) Delete(ctx context.Context, key string) error {
	_ = store.LRUCache.Remove(key)
	return nil
}
