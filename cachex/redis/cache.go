package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
	Client     redis.UniversalClient
	Expiration func(key string) time.Duration
	Marshal    func(key string, obj interface{}) ([]byte, error)
	Unmarshal  func(key string, data []byte) (interface{}, error)
	ErrHandler func(err error)
}

func (store *Cache) Get(ctx context.Context, key string) (interface{}, bool) {
	data, err := store.Client.Get(ctx, key).Result()
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

func (store *Cache) Set(ctx context.Context, key string, val interface{}) {
	var exp time.Duration
	if store.Expiration != nil {
		exp = store.Expiration(key)
	}
	var err error
	switch value := val.(type) {
	case []byte:
		_, err = store.Client.Set(ctx, key, value, exp).Result()
	case string:
		_, err = store.Client.Set(ctx, key, value, exp).Result()
	default:
		if store.Unmarshal == nil {
			err = errors.New("unmarshal function is nil")
		} else {
			var data []byte
			if data, err = store.Marshal(key, val); err == nil {
				_, err = store.Client.Set(context.Background(), key, data, exp).Result()
			}
		}
	}
	if err == nil {
		return
	}
	store.handleErr(err)
}

func (store *Cache) handleErr(err error) {
	if store.ErrHandler != nil {
		store.ErrHandler(err)
		return
	}
	panic(err)
}
