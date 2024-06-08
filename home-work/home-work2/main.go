package main

import (
	"log"
	"sync"
	"time"
)

// Task defines the interface for a task.
type Task interface {
	Execute(cahce Cache)
}

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

// TTLCache is an implementation of a cache with TTL.
type TTLCache struct {
	items sync.Map
	ttl   time.Duration
}

// NewTTLCache creates a new TTLCache with the specified TTL duration.
func NewTTLCache(ttl time.Duration) *TTLCache {
	cache := &TTLCache{
		ttl: ttl,
	}
	go cache.cleanup()
	return cache
}

// Set adds a new item to the cache.
func (c *TTLCache) Set(key string, value interface{}) {
	experation := time.Now().Add(c.ttl).UnixNano()
	c.items.Store(key, CacheItem{
		Value:      value,
		Experation: experation,
	})
}

// Get retrieves an item from the cache
func (c *TTLCache) Get(key string) (interface{}, bool) {
	item, found := c.items.Load(key)
	if !found {
		return nil, false
	}

	cacheItem := item.(CacheItem)
	if time.Now().UnixNano() > cacheItem.Experation {
		c.items.Delete(key)
		return nil, false
	}

	return cacheItem.Value, true
}

// cleanup periodically removes expired items from the cache.
func (c *TTLCache) cleanup() {
	for {
		time.Sleep(c.ttl)
		now := time.Now().UnixNano()

		c.items.Range(func(key, value interface{}) bool {
			cacheItem := value.(CacheItem)
			if now > cacheItem.Experation {
				c.items.Delete(key)
			}

			return true
		})
	}
}

// CheckTask is a task that sets a value in the cache
type CacheTask struct {
	key   string
	value interface{}
}

// Execute sets the value in the cache
func (t *CacheTask) Execute(cache Cache) {
	log.Printf("Setting value for key: %s", t.key)
	cache.Set(t.key, t.value)
}

// Worker processes task from the tasks channel and uses the cache.
func Worker(id int, cache Cache, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		log.Printf("Worker %d processing task", id)
		task.Execute(cache)
	}
}

func main() {
	cache := NewTTLCache(5 * time.Second)
	tasks := make(chan Task, 10)
	var wg sync.WaitGroup

	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, cache, tasks, &wg)
	}

	taskList := []Task{
		&CacheTask{"foo", "bar"},
		&CacheTask{"baz", "qux"},
		&CacheTask{"quux", "corge"},
	}

	for _, task := range taskList {
		tasks <- task
	}

	close(tasks)
	wg.Wait()

	time.Sleep(1 * time.Second)

	value, found := cache.Get("foo")
	if found {
		log.Println("Found value:", value)
	} else {
		log.Println("Value not found or expired")
	}

	time.Sleep(6 * time.Second)

	value, found = cache.Get("foo")
	if found {
		log.Println("Found value:", value)
	} else {
		log.Println("Value not found or expired")
	}
}
