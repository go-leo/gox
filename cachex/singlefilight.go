package cachex

import (
	"context"
	"golang.org/x/sync/singleflight"
)

var _ Store = (*SingleFlightCache)(nil)

type SingleFlightCache struct {
	Store       Store
	getGroup    singleflight.Group
	setGroup    singleflight.Group
	deleteGroup singleflight.Group
}

func (store *SingleFlightCache) Get(ctx context.Context, key string) (any, error) {
	val, err, _ := store.getGroup.Do(key, func() (interface{}, error) {
		return store.Store.Get(ctx, key)
	})
	return val, err
}

func (store *SingleFlightCache) Set(ctx context.Context, key string, val any) error {
	_, err, _ := store.setGroup.Do(key, func() (interface{}, error) {
		return nil, store.Store.Set(ctx, key, val)
	})
	return err
}

func (store *SingleFlightCache) Delete(ctx context.Context, key string) error {
	_, err, _ := store.deleteGroup.Do(key, func() (interface{}, error) {
		return nil, store.Store.Delete(ctx, key)
	})
	return err
}
