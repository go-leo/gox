package bigcachex

import (
	"context"
	"errors"
	"github.com/allegro/bigcache"
	"github.com/go-leo/gox/cachex"
)

var _ cachex.Store = (*Cache)(nil)

type Cache struct {
	BigCache  *bigcache.BigCache
	Marshal   func(key string, obj interface{}) ([]byte, error)
	Unmarshal func(key string, data []byte) (interface{}, error)
}

func (store *Cache) Get(ctx context.Context, key string) (any, error) {
	data, err := store.BigCache.Get(key)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil, cachex.ErrNil
	}
	if err != nil {
		return nil, err
	}
	// if Unmarshal is nil, return bytes data
	if store.Unmarshal == nil {
		return data, nil
	}
	obj, err := store.Unmarshal(key, data)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (store *Cache) Set(ctx context.Context, key string, val any) error {
	// if Marshal is nil, convert val to bytes
	if store.Marshal == nil {
		data, ok := val.([]byte)
		if !ok {
			return errors.New("bigcachex: failed to convert to bytes")
		}
		return store.BigCache.Set(key, data)
	}
	data, err := store.Marshal(key, val)
	if err != nil {
		return err
	}
	return store.BigCache.Set(key, data)
}

func (store *Cache) Delete(ctx context.Context, key string) error {
	return store.BigCache.Delete(key)
}
