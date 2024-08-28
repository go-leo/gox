package listx

import (
	"sync/atomic"
	"unsafe"
)

// LockFreeQueue 基于无锁队列实现,通过链表实现
type LockFreeQueue struct {
	// headPtr 队列头指针
	headPtr unsafe.Pointer
	// tailPtr 队列尾指针
	tailPtr unsafe.Pointer
}

// lockFreeElement 代表链表中的节点
type lockFreeElement struct {
	value   any
	nextPtr unsafe.Pointer
}

// NewLockFreeQueue 创建无锁队列
// Reference: https://www.sobyte.net/post/2021-07/implementing-lock-free-queues-with-go/
func NewLockFreeQueue() *LockFreeQueue {
	elemPtr := unsafe.Pointer(&lockFreeElement{})
	return &LockFreeQueue{headPtr: elemPtr, tailPtr: elemPtr}
}

// Enqueue 入队
func (q *LockFreeQueue) Enqueue(v any) {
	n := &lockFreeElement{value: v}
	for {
		tail := load(&q.tailPtr)
		next := load(&tail.nextPtr)
		// 尾还是尾
		if tail == load(&q.tailPtr) {
			// 还没有新数据入队
			if next == nil {
				//增加到队尾
				if cas(&tail.nextPtr, next, n) {
					//入队成功，移动尾巴指针
					_ = cas(&q.tailPtr, tail, n)
					return
				}
			} else {
				// 已有新数据加到队列后面，需要移动尾指针
				_ = cas(&q.tailPtr, tail, next)
			}
		}
	}
}

// Dequeue 出队，没有元素则返回nil
func (q *LockFreeQueue) Dequeue() any {
	for {
		head := load(&q.headPtr)
		tail := load(&q.tailPtr)
		next := load(&head.nextPtr)
		// head还是那个head
		if head == load(&q.headPtr) {
			// head和tail一样
			if head == tail {
				// 说明是空队列
				if next == nil {
					return nil
				}
				// 只是尾指针还没有调整，尝试调整它指向下一个
				_ = cas(&q.tailPtr, tail, next)
			} else {
				// 读取出队的数据
				v := next.value
				// 既然要出队了，头指针移动到下一个
				if cas(&q.headPtr, head, next) {
					// Dequeue is done.
					return v
				}
			}
		}
	}
}

func load(p *unsafe.Pointer) *lockFreeElement {
	return (*lockFreeElement)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *lockFreeElement) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
