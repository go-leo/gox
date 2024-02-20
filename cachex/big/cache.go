package big

import (
	"errors"
	"github.com/allegro/bigcache"
)

type Cache struct {
	BigCache   *bigcache.BigCache
	Marshal    func(key string, obj interface{}) ([]byte, error)
	Unmarshal  func(key string, data []byte) (interface{}, error)
	ErrHandler func(err error)
}

func (store *Cache) Get(key string) (interface{}, bool) {
	data, err := store.BigCache.Get(key)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil, false
	}
	if err != nil {
		store.handleErr(err)
		return nil, false
	}
	if store.Unmarshal == nil {
		return data, true
	}
	obj, err := store.Unmarshal(key, data)
	if err == nil {
		return obj, true
	}
	store.handleErr(err)
	return nil, false
}

func (store *Cache) Set(key string, val interface{}) {
	var err error
	switch value := val.(type) {
	case []byte:
		err = store.BigCache.Set(key, value)
	case string:
		err = store.BigCache.Set(key, []byte(value))
	default:
		if store.Unmarshal == nil {
			err = errors.New("unmarshal function is nil")
		} else {
			var data []byte
			if data, err = store.Marshal(key, val); err == nil {
				err = store.BigCache.Set(key, data)
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
