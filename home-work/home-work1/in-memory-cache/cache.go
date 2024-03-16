// Package cache provides an in-memory cache implementation
package cache

import (
	"sync"
)

// Cache represent an in-memory cache.
type Cache struct {
	mu    sync.RWMutex
	items map[string]interface{}
}

// New creates a new instance of Cache
func New() *Cache {
	return &Cache{items: make(map[string]interface{})}
}

// Set adds or updates a value in the cache.
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}

// Get retrieves a value from the cache by key
func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.items[key]
}

// Delete removes a value from the cache by key
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}
