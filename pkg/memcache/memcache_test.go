package memcache

import "testing"

func TestMemCache(t *testing.T) {
	cache := NewCacheSimple(10)

	// Test Get on an empty cache
	_, err := cache.Get("1")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Test Set
	cache.Set("1", "value1")
	cache.Set("2", "value2")
	cache.Set("3", "value3")

	// Test Get
	val, err := cache.Get("1")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if val != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}

	val, err = cache.Get("2")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if val != "value2" {
		t.Errorf("Expected value2, got %v", val)
	}

	val, err = cache.Get("3")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if val != "value3" {
		t.Errorf("Expected value3, got %v", val)
	}

	// Test Evict
	cache.Evict()
	_, err = cache.Get("1")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Test Overwrite
	cache.Set("2", "new-value2")
	val, err = cache.Get("2")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if val != "new-value2" {
		t.Errorf("Expected new-value2, got %v", val)
	}
}
