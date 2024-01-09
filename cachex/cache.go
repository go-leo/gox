package cachex

// Store 定义接口
type Store interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
}
