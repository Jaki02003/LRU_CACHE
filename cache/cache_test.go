package cache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cacheSize := 3
	lruCache := NewLRUCache(cacheSize)

	lruCache.Put("key1", "value1")
	lruCache.Put("key2", "value2")
	lruCache.Put("key3", "value3")

	if size := lruCache.Len(); size != cacheSize {
		t.Errorf("Expected cache size %d, but got %d", cacheSize, size)
	}

	value, exists := lruCache.Get("key2")
	if !exists || value != "value2" {
		t.Errorf("Get operation failed for key2. Expected 'value2', got '%v'", value)
	}

	_, exists = lruCache.Get("key4")
	if exists {
		t.Errorf("Get operation returned a value for a non-existing key (key4)")
	}

	lruCache.Put("key4", "value4")

	_, exists = lruCache.Get("key1")
	if exists {
		t.Errorf("Eviction failed. Expected key1 to be evicted, but it's still in the cache")
	}

	if size := lruCache.Len(); size != cacheSize {
		t.Errorf("Expected cache size %d after eviction, but got %d", cacheSize, size)
	}
}
