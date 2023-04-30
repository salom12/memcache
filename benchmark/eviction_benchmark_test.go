package memcache

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/salom12/memcache/pkg/memcache"
)

func BenchmarkLRUEviction(b *testing.B) {
	c := memcache.NewCache(1000, &memcache.LRUEviction{})

	for i := 0; i < b.N; i++ {
		// add some random data to the cache
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		c.Set(key, value)
	}
}

func BenchmarkLFUEviction(b *testing.B) {
	// create a new cache with LFU eviction policy
	cache := memcache.NewCache(100, &memcache.LFUEviction{})

	// add 1000 items to the cache
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		cache.Set(key, value)
	}

	// reset the benchmark timer
	b.ResetTimer()

	// perform 10000 Get operations on random cache keys
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", rand.Intn(1000))
		cache.Get(key)
	}
}

func BenchmarkLRUKEviction(b *testing.B) {
	cache := memcache.NewCache(100, &memcache.LRUKEviction{})

	for i := 0; i < b.N; i++ {
		cache.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
		cache.Get(fmt.Sprintf("key%d", rand.Intn(i+1)))
		cache.Evict()
	}
}

func BenchmarkSimpleEviction(b *testing.B) {
	c := memcache.NewCache(1000, &memcache.SimpleEviction{})

	for i := 0; i < b.N; i++ {
		// add some random data to the cache
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		c.Set(key, value)
	}
}

func BenchmarkRandomEviction(b *testing.B) {
	c := memcache.NewCache(1000, &memcache.RandomEviction{})

	for i := 0; i < b.N; i++ {
		// add some random data to the cache
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		c.Set(key, value)
	}
}
