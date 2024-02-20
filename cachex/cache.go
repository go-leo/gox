package cachex

import "context"

// Store 定义接口
type Store interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	Set(ctx context.Context, key string, val interface{})
}
