package sample

import (
	"context"
	"github.com/go-leo/gox/cachex"
	"sync"
)

var _ cachex.Store = (*Cache)(nil)

// Cache 简单缓存
type Cache struct {
	Map sync.Map
}

func (store *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	val, ok := store.Map.Load(key)
	if !ok {
		return nil, cachex.Nil
	}
	return val, nil
}

func (store *Cache) Set(ctx context.Context, key string, val any) error {
	store.Map.Store(key, val)
	return nil
}

func (store *Cache) Delete(ctx context.Context, key string) error {
	store.Map.Delete(key)
	return nil
}
