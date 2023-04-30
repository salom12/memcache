package memcache

import (
	"testing"
	"time"
)

func TestSimpleEviction(t *testing.T) {
	// Create a cache with a max size of 3
	cache := NewCache(3, &SimpleEviction{})

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

	// Check that the first item was evicted
	_, err := cache.Get("item1")
	if err == nil {
		t.Errorf("LRUEviction test failed: item3 should have been evicted")
	}
}

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
	cache.Set("item6", 5)

	// Check that the least frequently used item ("item2") has been evicted
	if _, err := cache.Get("item2"); err == nil {
		t.Errorf("Expected error due to item2 being evicted, but got nil error")
	}

	// Check that the other items are still in the cache
	if _, err := cache.Get("item1"); err != nil {
		t.Errorf("Expected item1 to still be in cache, but got error: %v", err)
	}
	if _, err := cache.Get("item6"); err != nil {
		t.Errorf("Expected item3 to still be in cache, but got error: %v", err)
	}
	if _, err := cache.Get("item5"); err != nil {
		t.Errorf("Expected item5 to be in cache, but got error: %v", err)
	}
}

func TestRandomEviction_Evict(t *testing.T) {
	c := NewCache(3, &RandomEviction{})

	// Add some values to the cache
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)

	// Cache should be full now
	if len(c.data) != 3 {
		t.Errorf("Cache size should be 3 but got %d", len(c.data))
	}

	// Evict a random item
	c.Evict()

	// Cache should have one less item now
	if len(c.data) != 2 {
		t.Errorf("Cache size should be 2 but got %d", len(c.data))
	}

	// Evict another random item
	c.Evict()

	// Cache should have one less item now
	if len(c.data) != 1 {
		t.Errorf("Cache size should be 1 but got %d", len(c.data))
	}
}

func TestLRUKEviction(t *testing.T) {
	// Create a new cache with LRU-K eviction policy
	cache := NewCache(4, &LRUKEviction{K: 2})

	// Add items to the cache
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)
	cache.Set("d", 4)

	// Access item "a" to increase its access count
	cache.Get("a")

	// Add new items to the cache
	cache.Set("e", 5)
	cache.Set("f", 6)

	// Check that item "b" was evicted due to being least frequently used
	if _, err := cache.Get("b"); err == nil {
		t.Errorf("Expected item 'b' to be evicted, but it was found in cache")
	}

	// Check that item "c" was evicted due to being least recently used among the K least frequently used items
	if _, err := cache.Get("c"); err == nil {
		t.Errorf("Expected item 'c' to be evicted, but it was found in cache")
	}
}
