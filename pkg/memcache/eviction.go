package memcache

import (
	"math/rand"
	"time"
)

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

// LFUEviction The LFU eviction policy evicts the least frequently used item from the cache when the cache is full.
type LFUEviction struct {
}

func (e *LFUEviction) Evict(c *Cache) {
	// Find the least frequently used item and remove it from the cache
	var minAccessCount int
	var minKey string

	// Find the item with the minimum access count
	for key, value := range c.data {
		if minKey == "" || value.accessCount < minAccessCount {
			minAccessCount = value.accessCount
			minKey = key
		}
	}
	// Delete the least frequently used item
	if minKey != "" {
		delete(c.data, minKey)
	}
}

// RandomEviction The Random eviction policy evicts a randomly selected item from the cache when the cache is full.
type RandomEviction struct {
}

func (e *RandomEviction) Evict(c *Cache) {
	// Select a random item from the cache and remove it

	if len(c.data) == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(c.data))

	var key string
	i := 0
	for k := range c.data {
		if i == randomIndex {
			key = k
			break
		}
		i++
	}

	delete(c.data, key)

}

// LRUKEviction The LRU-K eviction policy evicts the least recently used item among the K least frequently used items in the cache when the cache is full.
type LRUKEviction struct {
	// Fields specific to LRU-K eviction policy
	// k is a parameter that determines how many items should be kept in the cache before they are considered for eviction
	k int
}

func (e *LRUKEviction) Evict(c *Cache) {
	// Implementation of LRU-K eviction policy
	// Find the K least frequently used items in the cache and then remove the least recently used item among them
	var leastFrequentKeys []string
	for i := 0; i < e.k; i++ {
		var leastAccessCount int
		var leastAccessKey string
		for key, value := range c.data {
			if i == 0 || value.accessCount < leastAccessCount {
				leastAccessCount = value.accessCount
				leastAccessKey = key
			}
		}
		if leastAccessKey != "" {
			leastFrequentKeys = append(leastFrequentKeys, leastAccessKey)
		}
	}

	var leastRecentAccessKey string
	for key := range c.data {
		if leastRecentAccessKey == "" || c.data[key].lastAccess.Before(c.data[leastRecentAccessKey].lastAccess) {
			leastRecentAccessKey = key
		}
	}

	for _, key := range leastFrequentKeys {
		if key == leastRecentAccessKey {
			delete(c.data, key)

			return
		}
	}

	delete(c.data, leastRecentAccessKey)

}
