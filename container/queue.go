package container

import (
	"container/list"
	"sync"
)

// 队列是一种先进先出的数据结构

type QueueBySlice struct {
	lock   sync.RWMutex
	values []interface{}
}

func (queue *QueueBySlice) Push(v interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	queue.values = append(queue.values, v)
}

func (queue *QueueBySlice) Pop() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if len(queue.values) == 0 {
		return nil
	}
	res := queue.values[0]
	queue.values = queue.values[1:]
	return res
}

func (queue *QueueBySlice) Len() int {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	return len(queue.values)
}

func (queue *QueueBySlice) Empty() bool {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	return len(queue.values) == 0
}

type QueueByList struct {
	lock sync.RWMutex
	list list.List
}

func (queue *QueueByList) Push(v interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	queue.list.PushFront(v)
}

func (queue *QueueByList) Pop() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.list.Len() == 0 {
		return nil
	}
	res := queue.list.Back()
	queue.list.Remove(res)
	return res.Value
}

func (queue *QueueByList) Len() int {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	return queue.list.Len()
}

func (queue *QueueByList) Empty() bool {
	queue.lock.RLock()
	defer queue.lock.RUnlock()
	return queue.list.Len() == 0
}
