package ttl

import (
	"context"
	"github.com/go-leo/cache"
	"time"
)

// Cache TTL缓存
type Cache struct {
	Cache cache.Cache
	// 过期时间
	TTL func(key string) time.Duration
}

func (store *Cache) Get(ctx context.Context, key string) (interface{}, bool) {
	return store.Cache.Get(key)
}

func (store *Cache) Set(ctx context.Context, key string, val interface{}) {
	ttl := cache.DefaultExpiration
	if store.TTL != nil {
		ttl = store.TTL(key)
	}
	store.Cache.Set(key, val, ttl)
}
