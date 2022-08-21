package container

// 栈是一种先进后出的数据结构
//   本类分别用slice、container/list（双向链表）以及自定义的结构实现

import (
	"container/list"
	"sync"
)

type StackBySlice struct {
	lock   sync.RWMutex
	values []interface{}
}

func (stack *StackBySlice) Push(v interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.values = append(stack.values, v)
}

func (stack *StackBySlice) Peek() interface{} {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	return stack.values[len(stack.values)-1]
}

// TODO 在一些文章里（Reddit）提到，使用slice作为Stack的底层支持，速度更快。
// 但是，slice最大的问题其实是存在一个共用底层数组的的问题，哪怕你不断的Pop，但可能对于Golang来说，其占用的内存，并不一定减少。
func (stack *StackBySlice) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if len(stack.values) == 0 {
		return nil
	}
	res := stack.values[len(stack.values)-1]
	stack.values = stack.values[:len(stack.values)-1]
	return res
}

func (stack *StackBySlice) Len() int {
	return len(stack.values)
}

func (stack *StackBySlice) Empty() bool {
	return len(stack.values) == 0
}

type StackByList struct {
	lock sync.RWMutex
	list list.List
}

func (stack *StackByList) Push(v interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.PushBack(v)
}

func (stack *StackByList) Peek() interface{} {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	if stack.list.Len() == 0 {
		return nil
	}
	return stack.list.Back().Value
}

func (stack *StackByList) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.list.Len() == 0 {
		return nil
	}
	res := stack.list.Back()
	stack.list.Remove(res)
	return res.Value
}

func (stack *StackByList) Len() int {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	return stack.list.Len()
}

func (stack *StackByList) Empty() bool {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	return stack.list.Len() == 0
}

type (
	CustomizeStack struct {
		len  int
		top  *node // 栈顶
		lock sync.RWMutex
	}
	node struct {
		prev  *node
		value interface{}
	}
)

func (stack *CustomizeStack) Push(v interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.len++
	newValue := node{
		prev:  stack.top,
		value: v,
	}
	stack.top = &newValue
}

func (stack *CustomizeStack) Peek() interface{} {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	if stack.len == 0 {
		return nil
	}
	return stack.top.value
}

func (stack *CustomizeStack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.len == 0 {
		return nil
	}
	res := stack.top.value
	stack.top = stack.top.prev
	stack.len--
	return res
}

func (stack *CustomizeStack) Len() int {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	return stack.len
}

func (stack *CustomizeStack) Empty() bool {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	return stack.len == 0
}
