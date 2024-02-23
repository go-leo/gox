package redis

import (
	"context"
	"errors"
	"github.com/go-leo/gox/cachex"
	"github.com/redis/go-redis/v9"
	"time"
)

var _ cachex.Store = (*Cache)(nil)

var (
	ErrUnmarshalNil = errors.New("unmarshal function is nil")

	ErrMarshalNil = errors.New("marshal function is nil")
)

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
		return nil, cachex.Nil
	}
	if err != nil {
		return nil, err
	}
	if store.Unmarshal == nil {
		return nil, ErrUnmarshalNil
	}
	obj, err := store.Unmarshal(key, []byte(data))
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (store *Cache) Set(ctx context.Context, key string, val any) error {
	if store.Marshal == nil {
		return ErrMarshalNil
	}
	data, err := store.Marshal(key, val)
	if err != nil {
		return err
	}
	ttl := time.Duration(0)
	if store.TTL != nil {
		ttl = store.TTL(key)
	}
	return store.Client.Set(ctx, key, data, ttl).Err()
}

func (store *Cache) Delete(ctx context.Context, key string) error {
	return store.Client.Del(ctx, key).Err()
}
