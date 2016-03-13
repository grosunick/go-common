package queue

import "container/list"
import "sync"

// Tread safe queue, based on lists
type ListBasedConcurrentQueue struct {
	list  *list.List
	mutex sync.RWMutex
}

// ListBasedConcurrentQueue factory
func NewListBasedConcurrentQueue() *ListBasedConcurrentQueue {
	return &ListBasedConcurrentQueue{list.New(), sync.RWMutex{}}
}

// Adds element to queue
func (this *ListBasedConcurrentQueue) Add(item interface{}) bool {
	this.mutex.Lock()
	this.list.PushFront(item)
	this.mutex.Unlock()

	return true
}

// Removes element from queue
func (this *ListBasedConcurrentQueue) Remove() (interface{}, bool) {
	this.mutex.Lock()
	item := this.list.Remove(this.list.Back())
	this.mutex.Unlock()

	return item, true
}

// Returns queue size
func (this *ListBasedConcurrentQueue) Size() int {
	this.mutex.RLock()
	size := this.list.Len()
	this.mutex.RUnlock()

	return size
}