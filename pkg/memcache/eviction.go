package memcache

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type EvictionPolicy interface {
	Evict(c *Cache) error
}

// SimpleEviction a simple eviction policy
type SimpleEviction struct {
}

func (e *SimpleEviction) Evict(c *Cache) error {
	// implementation for a simple eviction policy
	// For example, remove the first item in the cache

	if len(c.keys) == 0 {
		return fmt.Errorf("cache is empty")
	}

	for _, key := range c.keys {
		delete(c.data, key)
		return nil
	}
	return nil
}

// LRUEviction The LRU eviction policy evicts the least recently used item from the cache when the cache is full.
type LRUEviction struct {
}

func (e *LRUEviction) Evict(c *Cache) error {
	// Find the least recently used item and remove it from the cache
	if len(c.keys) == 0 {
		return fmt.Errorf("cache is empty")
	}

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

	// Remove the key from the keys slice
	for i, key := range c.keys {
		if key == minKey {
			c.keys = append(c.keys[:i], c.keys[i+1:]...)
			break
		}
	}

	return nil
}

// LFUEviction The LFU eviction policy evicts the least frequently used item from the cache when the cache is full.
type LFUEviction struct {
}

func (e *LFUEviction) Evict(c *Cache) error {
	// Find the least frequently used item and remove it from the cache
	if len(c.keys) == 0 {
		return fmt.Errorf("cache is empty")
	}

	var minAccessCount int
	var minKey string

	// Find the item with the minimum access count
	for _, key := range c.keys {
		value, exists := c.data[key]
		if exists && (minKey == "" || value.accessCount < minAccessCount) {
			minAccessCount = value.accessCount
			minKey = key
		}
	}

	// Delete the least frequently used item
	if minKey != "" {
		delete(c.data, minKey)

		// Remove the key from the keys slice
		for i, k := range c.keys {
			if k == minKey {
				c.keys = append(c.keys[:i], c.keys[i+1:]...)
				break
			}
		}
	}

	return nil
}

// RandomEviction The Random eviction policy evicts a randomly selected item from the cache when the cache is full.
type RandomEviction struct {
}

func (e *RandomEviction) Evict(c *Cache) error {
	// Select a random item from the cache and remove it

	if len(c.keys) == 0 {
		return fmt.Errorf("cache is empty")
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(c.keys))

	key := c.keys[randomIndex]

	delete(c.data, key)

	// Remove the key from the keys slice
	c.keys = append(c.keys[:randomIndex], c.keys[randomIndex+1:]...)
	return nil

}

// LRUKEviction The LRU-K eviction policy evicts the least recently used item among the K least frequently used items in the cache when the cache is full.
type LRUKEviction struct {
	// Fields specific to LRU-K eviction policy
	// k is a parameter that determines how many items should be kept in the cache before they are considered for eviction
	k int
}

func (e *LRUKEviction) Evict(c *Cache) error {
	// Find the K least frequently used items
	kItems := make([]CacheValue, 0, e.k)
	for _, key := range c.keys {
		value, ok := c.data[key]
		if !ok {
			return fmt.Errorf("key %s does not exist in cache", key)
		}
		k := len(kItems)
		if k < e.k {
			// Fill up the k-items list until we have K items
			kItems = append(kItems, value)
			if len(kItems) == e.k {
				// Sort the k-items list in ascending order of access count
				sort.Slice(kItems, func(i, j int) bool {
					return kItems[i].accessCount < kItems[j].accessCount
				})
			}
		} else {
			// Replace the least frequently used item if we find an item with a lower access count
			if value.accessCount < kItems[k-1].accessCount {
				kItems[k-1] = value
				// Sort the k-items list in ascending order of access count
				sort.Slice(kItems, func(i, j int) bool {
					return kItems[i].accessCount < kItems[j].accessCount
				})
			}
		}
	}

	if len(kItems) == 0 {
		return fmt.Errorf("cache is empty")
	}

	// Evict the least recently used item among the K least frequently used items
	lru := kItems[0]
	lruKey := ""
	for key, value := range c.data {
		if value == lru {
			lruKey = key
			break
		}
	}

	// remove value from data
	delete(c.data, lruKey)

	// Remove the key from the keys slice
	for i, k := range c.keys {
		if k == lruKey {
			c.keys = append(c.keys[:i], c.keys[i+1:]...)
			break
		}
	}

	return nil
}
