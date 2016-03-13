package queue

import "sync"

// Реализация потокобезопасной очереди на основе списка.
type ChannelBasedLimitQueue struct {
	size int
	channel chan interface{}
	mutex sync.RWMutex
}

// Создает очередь на основе списка.
func NewChannelBasedLimitQueue(limit uint32) *ChannelBasedLimitQueue {
	return &ChannelBasedLimitQueue{channel: make(chan interface{}, limit)}
}

// Добавляет элемент item в начало очереди
func (this *ChannelBasedLimitQueue) Add(item interface{}) bool {
	this.channel <- item
	this.mutex.Lock()
	this.size++
	this.mutex.Unlock()

	return true
}

// Извлекает объект из очереди
func (this *ChannelBasedLimitQueue) Remove() (interface{}, bool) {
	item := <- this.channel
	this.mutex.Lock()
	this.size--
	this.mutex.Unlock()

	return item, true
}

// Возвращает размер очереди
func (this *ChannelBasedLimitQueue) Size() int {
	this.mutex.RLock()
	size := this.size
	this.mutex.RUnlock()

	return size
}
