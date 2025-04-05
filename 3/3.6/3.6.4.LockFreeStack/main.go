package main

import (
	"sync/atomic"
	"unsafe"
)

type Node struct {
	value int
	next  *Node
}

type LockFreeStack struct {
	top unsafe.Pointer
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (s *LockFreeStack) Push(value int) {
	newTop := &Node{value: value}
	for {
		oldTop := atomic.LoadPointer(&s.top)
		newTop.next = (*Node)(oldTop)
		if atomic.CompareAndSwapPointer(
			&s.top,
			oldTop,
			unsafe.Pointer(newTop),
		) {
			break
		}
	}
}

func (s *LockFreeStack) Pop() (int, bool) {
	for {
		oldTop := atomic.LoadPointer(&s.top)
		if oldTop == nil {
			return 0, false
		}

		newTop := (*Node)(oldTop).next
		if atomic.CompareAndSwapPointer(
			&s.top,
			oldTop,
			unsafe.Pointer(newTop),
		) {
			return (*Node)(oldTop).value, true
		}
	}
}

func main() {

}
