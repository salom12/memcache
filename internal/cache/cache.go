package cache

import "github.com/salom12/memcache/pkg/memcache"

var cache *memcache.Cache

func InitMemcache(maxSize int, policy memcache.EvictionPolicy) {
	cache = memcache.NewCache(maxSize, policy)
}

func GetMemcache() *memcache.Cache {
	return cache
}
