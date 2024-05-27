package main

// CacheItem represents a single cache item with its value and expiration time.

type CacheItem struct {
	Value      interface{}
	Experation int64
}

// Cache defines the cache interface.
type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}
