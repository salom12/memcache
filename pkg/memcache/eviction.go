package memcache

import "time"

type EvictionPolicy interface {
	Evict(c *Cache)
}

// SimpleEviction a simple eviction policy
type SimpleEviction struct {
}

func (e *SimpleEviction) Evict(c *Cache) {
	// implementation for a simple eviction policy
	// For example, remove the first item in the cache
	for key := range c.data {
		delete(c.data, key)
		return
	}
}

// LRUEviction The LRU eviction policy evicts the least recently used item from the cache when the cache is full.
type LRUEviction struct {
}

func (e *LRUEviction) Evict(c *Cache) {
	// Find the least recently used item and remove it from the cache
	minAccessTime := time.Now()
	var minKey string

	// Find the least recently used item by iterating over the cache
	for key, value := range c.data {
		if value.lastAccess.Before(minAccessTime) {
			minAccessTime = value.lastAccess
			minKey = key
		}
	}

	// Delete the least recently used item from the cache
	delete(c.data, minKey)
}
