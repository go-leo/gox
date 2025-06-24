package bytex

import (
	"bytes"
	"sync"
)

const (
	// DefaultBufferSize is the default size of the buffer in bytes.
	DefaultBufferSize = 1024

	// MaxBufferSize is the maximum size of the buffer in bytes.
	MaxBufferSize = 4096
)

// defaultBufferPool is a pool of bytes.Buffer
var defaultBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, DefaultBufferSize))
	},
}

// GetBuffer retrieves a bytes.Buffer from the pool.
func GetBuffer() *bytes.Buffer {
	return defaultBufferPool.Get().(*bytes.Buffer)
}

// PutBuffer returns a bytes.Buffer to the pool.
func PutBuffer(buf *bytes.Buffer) {
	if buf.Cap() > MaxBufferSize {
		return
	}
	buf.Reset()
	defaultBufferPool.Put(buf)
}
