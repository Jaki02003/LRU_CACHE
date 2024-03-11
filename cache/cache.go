package cache

import "container/list"

type Node struct {
	Key   interface{}
	Value interface{}
}

type LRUCache struct {
	size     int
	cache    map[interface{}]*list.Element
	usedList *list.List
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		size:     size,
		cache:    make(map[interface{}]*list.Element),
		usedList: list.New(),
	}
}

func (c *LRUCache) Get(key interface{}) (interface{}, bool) {
	if element, ok := c.cache[key]; ok {
		c.usedList.MoveToFront(element)
		return element.Value.(*Node).Value, true
	}
	return nil, false
}

func (c *LRUCache) Put(key, value interface{}) {
	if element, ok := c.cache[key]; ok {
		element.Value.(*Node).Value = value
		c.usedList.MoveToFront(element)
		return
	}

	newNode := &Node{Key: key, Value: value}
	element := c.usedList.PushFront(newNode)
	c.cache[key] = element

	if c.usedList.Len() > c.size {
		c.removeOldest()
	}
}

func (c *LRUCache) removeOldest() {
	element := c.usedList.Back()
	if element == nil {
		return
	}
	c.usedList.Remove(element)
	key := element.Value.(*Node).Key
	delete(c.cache, key)
}

func (c *LRUCache) Len() int {
	return c.usedList.Len()
}
