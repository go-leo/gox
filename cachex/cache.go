package cachex

import "context"

// Store 定义接口
type Store interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
}

// ContextStore 定义接口
type ContextStore interface {
	Store
	WithContext(ctx context.Context) ContextStore
}
