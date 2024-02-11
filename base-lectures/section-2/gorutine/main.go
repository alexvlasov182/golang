package main

import "fmt"

func main() {
	a := createPtr()

	fmt.Println(*a)

	*a = 10

	fmt.Println(*a)
}

func createPtr() *int {
	a := 5
	return &a
}
