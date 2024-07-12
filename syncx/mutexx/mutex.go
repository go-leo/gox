package mutexx

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift      = iota
	starvationThresholdNs = 1e6
)

// WaiterCount 锁的等待者数量
func WaiterCount(m *sync.Mutex) int {
	// 获取state字段的值
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m)))
	v = v >> mutexWaiterShift //得到等待者的数值
	return int(v)
}

// HolderCount 锁的持有者数量 0或者1
func HolderCount(m *sync.Mutex) int {
	v := WaiterCount(m)
	return v & mutexLocked
}

// IsLocked 锁是否被持有
func IsLocked(m *sync.Mutex) bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(m)))
	return state&spinMutexLocked == spinMutexLocked
}

// IsWoken 是否有等待者被唤醒
func IsWoken(m *sync.Mutex) bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(m)))
	return state&mutexWoken == mutexWoken
}

// IsStarving 锁是否处于饥饿状态
func IsStarving(m *sync.Mutex) bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(m)))
	return state&mutexStarving == mutexStarving
}
