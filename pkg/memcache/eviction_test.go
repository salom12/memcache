package memcache

import (
	"testing"
	"time"
)

func TestLRUEviction(t *testing.T) {
	// Create a cache with a max size of 3
	cache := NewCache(3, &LRUEviction{})

	// Add some items to the cache
	cache.Set("item1", "value1")
	cache.Set("item2", "value2")

	// wait a little bit
	time.Sleep(time.Millisecond)
	cache.Set("item3", "value3")

	// Access some items to change their access times
	cache.Get("item1")
	cache.Get("item2")

	// Add another item to exceed the max size and trigger eviction
	cache.Set("item4", "value4")

	// Check that the least recently used item was evicted
	_, err := cache.Get("item3")
	if err == nil {
		t.Errorf("LRUEviction test failed: item3 should have been evicted")
	}
}
