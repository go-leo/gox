package redis

import (
	"context"
	"errors"
	"github.com/go-leo/gox/cachex"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
	Cache      redis.UniversalClient
	Expiration func(key string) time.Duration
	Marshal    func(key string, obj interface{}) ([]byte, error)
	Unmarshal  func(key string, data []byte) (interface{}, error)
	ErrHandler func(err error)
	ctx        context.Context
}

func (store *Cache) Get(key string) (interface{}, bool) {
	ctx := store.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	data, err := store.Cache.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, false
	}
	if err != nil {
		store.handleErr(err)
		return nil, false
	}
	if store.Unmarshal == nil {
		return data, true
	}
	obj, err := store.Unmarshal(key, []byte(data))
	if err == nil {
		return obj, true
	}
	store.handleErr(err)
	return nil, false
}

func (store *Cache) Set(key string, val interface{}) {
	ctx := store.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	var exp time.Duration
	if store.Expiration != nil {
		exp = store.Expiration(key)
	}
	var err error
	switch value := val.(type) {
	case []byte:
		_, err = store.Cache.Set(ctx, key, value, exp).Result()
	case string:
		_, err = store.Cache.Set(ctx, key, value, exp).Result()
	default:
		if store.Unmarshal == nil {
			err = errors.New("unmarshal function is nil")
		} else {
			var data []byte
			if data, err = store.Marshal(key, val); err == nil {
				_, err = store.Cache.Set(context.Background(), key, data, exp).Result()
			}
		}
	}
	if err == nil {
		return
	}
	store.handleErr(err)
}

func (store *Cache) WithContext(ctx context.Context) cachex.ContextStore {
	if ctx == nil {
		panic("nil context")
	}
	cloned := *store
	cloned.ctx = ctx
	return &cloned
}

func (store *Cache) handleErr(err error) {
	if store.ErrHandler != nil {
		store.ErrHandler(err)
		return
	}
	panic(err)
}
