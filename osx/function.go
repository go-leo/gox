package osx

import (
	"unsafe"
)

func WordSize() int {
	// 获取指针的大小，即计算机的字长
	wordSize := unsafe.Sizeof(new(interface{}))
	// 根据字节大小转换为位数
	return int(wordSize * 8)
}
