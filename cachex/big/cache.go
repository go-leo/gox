package big

import (
	"context"
	"errors"
	"github.com/allegro/bigcache"
	"github.com/go-leo/gox/cachex"
)

var _ cachex.Store = (*Cache)(nil)

var (
	ErrUnmarshalNil = errors.New("unmarshal function is nil")

	ErrMarshalNil = errors.New("marshal function is nil")
)

type Cache struct {
	BigCache  *bigcache.BigCache
	Marshal   func(key string, obj interface{}) ([]byte, error)
	Unmarshal func(key string, data []byte) (interface{}, error)
}

func (store *Cache) Get(ctx context.Context, key string) (any, error) {
	data, err := store.BigCache.Get(key)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil, cachex.Nil
	}
	if err != nil {
		return nil, err
	}
	if store.Unmarshal == nil {
		return nil, ErrUnmarshalNil
	}
	obj, err := store.Unmarshal(key, data)
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
	return store.BigCache.Set(key, data)
}

func (store *Cache) Delete(ctx context.Context, key string) error {
	return store.BigCache.Delete(key)
}
