package main

import (
	"fmt"
)

func main() {
	users := map[string]int{
		"Alex":       35,
		"Pet":        44,
		"Jane":       33,
		"Anastasiia": 29,
	}

	for key, value := range users {
		fmt.Println(key, value)
	}

	usersers := make([]string, 10)
	anotherMaps := make(map[string]int)
	fmt.Println(usersers)

	fmt.Println("hello")
	anotherMaps["John"] = 29

	fmt.Println(anotherMaps)
	fmt.Println(len(users))
	fmt.Println("Hello")
}
