package lazyloadx

import (
	"errors"
	"golang.org/x/sync/singleflight"
	"sync"
)

type Group struct {
	m   sync.Map
	g   singleflight.Group
	New func() (any, error)
}

func (g *Group) Load(key string) (any, bool, error) {
	// 1. 初次检查,如果 key 已存在于 sync.Map 中，直接返回该值。
	if value, ok := g.m.Load(key); ok {
		return value, true, nil
	}
	// 2. 并发控制,g.g.Do 会确保并发控制，即如果有多个 goroutine 同时请求同一个 key，只有一个 goroutine 会执行闭包中的逻辑。
	value, err, shared := g.g.Do(key, func() (any, error) {
		// 3. 再次检查 key 是否已存在于 sync.Map 中
		// 如果存在，则直接返回该值。
		if value, ok := g.m.Load(key); ok {
			return value, nil
		}
		// 如果不存在，则调用 g.New 创建新值。
		if g.New == nil {
			return nil, errors.New("lazyloadx: New function is nil")
		}
		value, err := g.New()
		if err != nil {
			return nil, err
		}
		return value, nil
	})
	if err != nil {
		return nil, false, err
	}

	// 4. 存储结果,如果 shared 为 false，表示当前 goroutine 是唯一执行闭包的 goroutine，此时将结果存储到 sync.Map 中。
	if !shared {
		g.m.Store(key, value)
	}
	return value, false, err
}

func (g *Group) Delete(key string) {
	g.m.Delete(key)
}
