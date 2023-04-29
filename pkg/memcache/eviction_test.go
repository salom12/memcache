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

func TestLFUEviction(t *testing.T) {
	// Create a cache with LFU eviction policy and maximum size of 3
	cache := NewCache(3, &LFUEviction{})

	// Add 3 items to the cache
	cache.Set("item1", 1)
	cache.Set("item2", 2)
	cache.Set("item3", 3)

	// Retrieve the first item to increase its access count
	cache.Get("item1")

	// Add 2 more items to the cache, exceeding its maximum size
	cache.Set("item4", 4)
	cache.Set("item5", 5)

	// Check that the least frequently used item ("item2") has been evicted
	if _, err := cache.Get("item2"); err == nil {
		t.Errorf("Expected error due to item2 being evicted, but got nil error")
	}

	// Check that the other items are still in the cache
	if _, err := cache.Get("item1"); err != nil {
		t.Errorf("Expected item1 to still be in cache, but got error: %v", err)
	}
	if _, err := cache.Get("item3"); err != nil {
		t.Errorf("Expected item3 to still be in cache, but got error: %v", err)
	}
	if _, err := cache.Get("item4"); err != nil {
		t.Errorf("Expected item4 to be in cache, but got error: %v", err)
	}
	if _, err := cache.Get("item5"); err != nil {
		t.Errorf("Expected item5 to be in cache, but got error: %v", err)
	}
}
