package gocachex

import (
	"context"
	"github.com/go-leo/gox/cachex"
	"github.com/patrickmn/go-cache"
	"time"
)

var _ cachex.Store = (*Cache)(nil)

// Cache TTL缓存
type Cache struct {
	Cache *cache.Cache
	// 过期时间
	TTL func(ctx context.Context, key string) time.Duration
}

func (store *Cache) Get(ctx context.Context, key string) (any, error) {
	val, ok := store.Cache.Get(key)
	if !ok {
		return nil, cachex.ErrNil
	}
	return val, nil
}

func (store *Cache) Set(ctx context.Context, key string, val any) error {
	ttl := cache.DefaultExpiration
	if store.TTL != nil {
		ttl = store.TTL(ctx, key)
	}
	store.Cache.Set(key, val, ttl)
	return nil
}

func (store *Cache) Delete(ctx context.Context, key string) error {
	store.Cache.Delete(key)
	return nil
}
