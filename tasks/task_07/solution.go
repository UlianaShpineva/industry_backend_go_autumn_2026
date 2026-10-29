package main

import (
	"container/list"
	"sync"
)

type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	mu       sync.Mutex
	ll       *list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		ll:       list.New(),
		items:    make(map[K]*list.Element),
	}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	el, ok := c.items[key]
	if !ok {
		return
	}
	c.ll.MoveToBack(el)
	return el.Value.(entry[K, V]).value, ok
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.capacity <= 0 {
		return
	}

	el, ok := c.items[key]
	if !ok {
		if c.ll.Len() >= c.capacity {
			delete(c.items, c.ll.Front().Value.(entry[K, V]).key)
			c.ll.Remove(c.ll.Front())
		}
		c.items[key] = c.ll.PushBack(entry[K, V]{key: key, value: value})
		return
	}
	el.Value = entry[K, V]{key: key, value: value}
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
