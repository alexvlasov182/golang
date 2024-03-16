package main

import (
	"fmt"

	cache "exmaple.com/home-work/home-work1/in-memory-cache"
)

func main() {
	// Create cache
	cache := cache.New()

	// Set new value
	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId") // Reassigning userId variable

	fmt.Println(userId)

}
