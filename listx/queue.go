package listx

import (
	"sync/atomic"
	"unsafe"
)

// LockFreeQueue 基于无锁队列实现
type LockFreeQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

// 通过链表实现，这个数据结构代表链表中的节点
type lockFreenode struct {
	value interface{}
	next  unsafe.Pointer
}

// NewLockFreeQueue 创建无锁队列
func NewLockFreeQueue() *LockFreeQueue {
	n := unsafe.Pointer(&lockFreenode{})
	return &LockFreeQueue{head: n, tail: n}
}

// Enqueue 入队
func (q *LockFreeQueue) Enqueue(v interface{}) {
	n := &lockFreenode{value: v}
	for {
		tail := load(&q.tail)
		next := load(&tail.next)
		if tail == load(&q.tail) { // 尾还是尾
			if next == nil { // 还没有新数据入队
				if cas(&tail.next, next, n) { //增加到队尾
					cas(&q.tail, tail, n) //入队成功，移动尾巴指针
					return
				}
			} else { // 已有新数据加到队列后面，需要移动尾指针
				cas(&q.tail, tail, next)
			}
		}
	}
}

// Dequeue 出队，没有元素则返回nil
func (q *LockFreeQueue) Dequeue() interface{} {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if head == load(&q.head) { // head还是那个head
			if head == tail { // head和tail一样
				if next == nil { // 说明是空队列
					return nil
				}
				// 只是尾指针还没有调整，尝试调整它指向下一个
				cas(&q.tail, tail, next)
			} else {
				// 读取出队的数据
				v := next.value
				// 既然要出队了，头指针移动到下一个
				if cas(&q.head, head, next) {
					return v // Dequeue is done.  return
				}
			}
		}
	}
}

// load 封装load,避免直接将*node转换成unsafe.Pointer
func load(p *unsafe.Pointer) (n *lockFreenode) {
	return (*lockFreenode)(atomic.LoadPointer(p))
}

// 封装CAS,避免直接将*node转换成unsafe.Pointer
func cas(p *unsafe.Pointer, old, new *lockFreenode) (ok bool) {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
