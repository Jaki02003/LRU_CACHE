package cache

import "container/list"

type Node struct {
	Key   interface{}
	Value interface{}
}

type LRUCache struct {
	size      int
	cache     map[interface{}]*list.Element
	usedList  *list.List
	onEvicted func(key, value interface{})
}

func NewLRUCache(size int, onEvicted func(key, value interface{})) *LRUCache {
	return &LRUCache{
		size:      size,
		cache:     make(map[interface{}]*list.Element),
		usedList:  list.New(),
		onEvicted: onEvicted,
	}
}

func (c *LRUCache) Len() int {
	return c.usedList.Len()
}
