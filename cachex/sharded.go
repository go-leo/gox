package cachex

import (
	"context"
)

var _ Store = (*ShardedCache)(nil)

type ShardedCache struct {
	stores []Store
	shards uint32
	seed   uint32
}

func (store *ShardedCache) Get(ctx context.Context, key string) (any, error) {
	return store.bucket(key).Get(ctx, key)
}

func (store *ShardedCache) Set(ctx context.Context, key string, val any) error {
	return store.bucket(key).Set(ctx, key, val)
}

func (store *ShardedCache) Delete(ctx context.Context, key string) error {
	return store.bucket(key).Delete(ctx, key)
}

func (store *ShardedCache) bucket(k string) Store {
	return store.stores[store.djb33(k)%store.shards]
}

// djb2 with better shuffling. 5x faster than FNV with the hash.Hash overhead.
func (store *ShardedCache) djb33(k string) uint32 {
	var (
		l = uint32(len(k))
		d = 5381 + store.seed + l
		i = uint32(0)
	)
	// Why is all this 5x faster than a for loop?
	if l >= 4 {
		for i < l-4 {
			d = (d * 33) ^ uint32(k[i])
			d = (d * 33) ^ uint32(k[i+1])
			d = (d * 33) ^ uint32(k[i+2])
			d = (d * 33) ^ uint32(k[i+3])
			i += 4
		}
	}
	switch l - i {
	case 1:
	case 2:
		d = (d * 33) ^ uint32(k[i])
	case 3:
		d = (d * 33) ^ uint32(k[i])
		d = (d * 33) ^ uint32(k[i+1])
	case 4:
		d = (d * 33) ^ uint32(k[i])
		d = (d * 33) ^ uint32(k[i+1])
		d = (d * 33) ^ uint32(k[i+2])
	}
	return d ^ (d >> 16)
}
