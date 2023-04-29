package memcache

import (
	"fmt"
	"sync"
)

type Cache struct {
	Store
	data    map[string]any
	maxSize int
	mutex   sync.Mutex

	evictionPolicy EvictionPolicy
}

func NewCache(maxSize int) *Cache {
	return &Cache{
		data:           make(map[string]any),
		maxSize:        maxSize,
		evictionPolicy: &SimpleEviction{},
	}
}

func (c *Cache) Get(key string) (any, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, exists := c.data[key]
	if !exists {
		return value, fmt.Errorf("key not found in cache: %v", key)
	}

	return value, nil
}

func (c *Cache) Set(key string, value any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.data) == c.maxSize {
		c.Evict()
	}

	c.data[key] = value
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
}

func (c *Cache) Evict() {
	c.evictionPolicy.Evict(c)
}
