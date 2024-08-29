package cachex

import (
	"context"
)

type Updater interface {
	Update(ctx context.Context, key string, val any) error
}

var _ Store = (*WriteThroughCache)(nil)

type WriteThroughCache struct {
	Store
	Updater Updater
}

func (cache *WriteThroughCache) Set(ctx context.Context, key string, val any) error {
	if err := cache.Store.Delete(ctx, key); err != nil {
		return err
	}
	if err := cache.Updater.Update(ctx, key, val); err != nil {
		return err
	}
	return cache.Store.Set(ctx, key, val)
}
