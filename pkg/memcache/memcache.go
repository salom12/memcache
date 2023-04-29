package memcache

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Cache struct {
	Store
	data    map[string]CacheValue
	maxSize int
	mutex   sync.Mutex

	evictionPolicy EvictionPolicy
}

type CacheValue struct {
	data        any
	accessCount int
	lastAccess  time.Time
}

func NewCacheSimple(maxSize int) *Cache {
	return &Cache{
		data:           make(map[string]CacheValue),
		maxSize:        maxSize,
		evictionPolicy: &SimpleEviction{},
	}
}

func NewCache(maxSize int, policy EvictionPolicy) *Cache {
	return &Cache{
		data:           make(map[string]CacheValue),
		maxSize:        maxSize,
		evictionPolicy: policy,
	}
}

func (c *Cache) Get(key string) (any, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, exists := c.data[key]
	if !exists {
		return value.data, fmt.Errorf("key not found in cache: %v", key)
	}

	// update access count
	value.accessCount++

	// upadte last access
	value.lastAccess = time.Now()

	// Update cache
	c.data[key] = value

	return value.data, nil
}

func (c *Cache) GetString(key string) (string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, err := c.Get(key)
	if err != nil {
		return "", err
	}

	if reflect.TypeOf(value).Kind() != reflect.String {
		return "", fmt.Errorf("value is not a string")
	}

	return value.(string), nil
}

func (c *Cache) Set(key string, value any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.data) == c.maxSize {
		c.Evict()
	}

	now := time.Now()
	c.data[key] = CacheValue{
		data:        value,
		accessCount: 1,
		lastAccess:  now,
	}
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
}

func (c *Cache) Evict() {
	c.evictionPolicy.Evict(c)
}
