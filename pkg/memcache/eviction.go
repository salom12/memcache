package memcache

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
