package atomicx

import "sync/atomic"

func SubUint32(addr *uint32, delta int32) (new uint32) {
	return atomic.AddUint32(addr, ^uint32(delta-1))
}

func SubUint64(addr *uint64, delta int64) (new uint64) {
	return atomic.AddUint64(addr, ^uint64(delta-1))
}

func DecrUint32(addr *uint32) (new uint32) {
	return SubUint32(addr, 1)
}

func DecrUint64(addr *uint64) (new uint64) {
	return SubUint64(addr, 1)
}

func IncrUint32(addr *uint32) (new uint32) {
	return atomic.AddUint32(addr, 1)
}

func IncrUint64(addr *uint64) (new uint64) {
	return atomic.AddUint64(addr, 1)
}

//func SubUintptr(addr *uintptr, delta uintptr) (new uintptr) {
//	return atomic.AddUintptr(addr, ^(delta - 1))
//}
