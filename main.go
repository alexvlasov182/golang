package main

import (
	"fmt"

	inmemorycache "exmaple.com/home-work/home-work1/in-memory-cache"
)

func main() {
	// Create cache
	cache := inmemorycache.NewCache()

	// Set new value
	cache.Set("userId", 123)

	// Get value by key
	cache.Get("userId")

	// Delete value by key
	// cache.Delete("userId")

	fmt.Println(cache)
}
