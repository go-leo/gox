package redisx

import (
	"context"
	"errors"
	"github.com/go-leo/gox/cachex"
	"github.com/redis/go-redis/v9"
	"time"
)

var _ cachex.Store = (*Cache)(nil)

type Cache struct {
	Client redis.UniversalClient
	// 过期时间
	TTL       func(key string) time.Duration
	Marshal   func(key string, obj interface{}) ([]byte, error)
	Unmarshal func(key string, data []byte) (interface{}, error)
}

func (store *Cache) Get(ctx context.Context, key string) (any, error) {
	data, err := store.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, cachex.ErrNil
	}
	if err != nil {
		return nil, err
	}
	if store.Unmarshal == nil {
		return data, nil
	}
	obj, err := store.Unmarshal(key, []byte(data))
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (store *Cache) Set(ctx context.Context, key string, val any) error {
	ttl := time.Duration(0)
	if store.TTL != nil {
		ttl = store.TTL(key)
	}
	if store.Marshal == nil {
		return store.Client.Set(ctx, key, val, ttl).Err()
	}
	data, err := store.Marshal(key, val)
	if err != nil {
		return err
	}
	return store.Client.Set(ctx, key, data, ttl).Err()
}

func (store *Cache) Delete(ctx context.Context, key string) error {
	return store.Client.Del(ctx, key).Err()
}
