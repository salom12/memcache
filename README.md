# Golang Memcache Library

The is an in-memory key value cache library for golang that has a `maxsize` limit,
it contains various implementations of the "Eviction Policy Alogrithms" which define how items in the cache are evicted when the cache is full.

The currently implemented eviction policies includes:

- [x] LRU
- [x] LFU
- [x] Random
- [x] LRU-K.

### Install
```
go get github.com/salom12/memcache
```

### Simple Example

```go
package main

import (
    "fmt"
    "github.com/salom12/memcache"
)

func main() {
    // Create a new cache with a maximum size of 100
    cache := memcache.NewCacheSimple(100)

    // Set a value in the cache
    cache.Set("key", "value")

    // Get a value from the cache
    value, err := cache.Get("key")
    if err != nil {
    	panic(err)
    }

    fmt.Println(value)

    // Delete a value from the cache
    cache.Delete("key")
}
```

### Create a new cache

To create a new cache with a specified maximum size, use the NewCache function:

```go
cache := NewSimpleCache(maxSize)
```

### Set a value in the cache

```go
cache.Set(key, value) // value is an interface{} which means it can be anything
```

### Get a value from the cache

```go
value, err := cache.Get(key)
```

### Delete a value from the cache

```go
cache.Delete(key)
```

### Eviction policies

The cache is designed to use different eviction policies to manage its size. The following eviction policies are currently supported:

- **SimpleEviction**: a simple policy that removes the first item in the cache
- **LRUEviction**: the least recently used item is removed from the cache when the cache is full
- **LFUEviction**: the least frequently used item is removed from the cache when the cache is full
- **RandomEviction**: a randomly selected item is removed from the cache when the cache is full
- **LRUKEviction**: the least recently used item among the K least frequently used items in the cache is removed when the cache is full.

To use an eviction policy, create an instance of the desired policy and pass it to the Cache object

```go
    c1 := NewCache(3, &SimpleEviction{})
    c2 := NewCache(3, &LRUEviction{})
    c3 := NewCache(3, &LFUEviction{})
    c4 := NewCache(3, &RandomEviction{})
    c5 := NewCache(3, &LRUKEviction{k: 2}) // k is a parameter that determines how many items should be kept in the cache before they are considered for eviction
```

### Run example server

```sh
go run ./cmd/memcache/main.go
```

### Run tests

```sh
go test -v ./...
```

### Run benchmark

```sh
go test -v  -bench=. -benchmem benchmark/eviction_benchmark_test.go
```
