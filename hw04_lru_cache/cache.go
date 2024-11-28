package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheValue struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.Lock()
	defer c.Unlock()

	v, exist := c.items[key]
	if exist {
		v.Value = cacheValue{
			key:   key,
			value: value,
		}
		c.queue.MoveToFront(v)
		return true
	}
	if c.queue.Len() == c.capacity {
		item := c.queue.Back()
		c.queue.Remove(item)
		delete(c.items, item.Value.(cacheValue).key)
	}
	cValue := cacheValue{
		key:   key,
		value: value,
	}
	v = c.queue.PushFront(cValue)
	c.items[key] = v

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	v, exist := c.items[key]
	if exist {
		c.queue.MoveToFront(v)
		return v.Value.(cacheValue).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.Lock()
	defer c.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
