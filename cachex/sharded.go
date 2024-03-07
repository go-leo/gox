package cachex

import (
	"context"
	"crypto/rand"
	"math"
	"math/big"
	insecurerand "math/rand"
)

var _ Store = (*shardedCache)(nil)

type shardedCache struct {
	stores []Store
	seed   uint32
}

func ShardedCache(stores []Store) Store {
	rnd, err := rand.Int(rand.Reader, big.NewInt(0).SetUint64(uint64(math.MaxUint32)))
	var seed uint32
	if err != nil {
		seed = insecurerand.Uint32()
	} else {
		seed = uint32(rnd.Uint64())
	}
	return &shardedCache{stores: stores, seed: seed}
}

func (store *shardedCache) Get(ctx context.Context, key string) (any, error) {
	return store.bucket(key).Get(ctx, key)
}

func (store *shardedCache) Set(ctx context.Context, key string, val any) error {
	return store.bucket(key).Set(ctx, key, val)
}

func (store *shardedCache) Delete(ctx context.Context, key string) error {
	return store.bucket(key).Delete(ctx, key)
}

func (store *shardedCache) bucket(k string) Store {
	return store.stores[int(store.djb33(k))%len(store.stores)]
}

// djb2 with better shuffling. 5x faster than FNV with the hash.Hash overhead.
func (store *shardedCache) djb33(k string) uint32 {
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
