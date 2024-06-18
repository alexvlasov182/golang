package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func main() {
	a := []int64{1, 2, 3}
	b := []float64{1.1, 2.2, 3.3}
	c := []string{"1", "2", "3"}

	d := []User{
		{
			Email: "asd@gmail.com",
			Name:  "Oleksandr",
		},
		{
			Email: "lsd@gmail.com",
			Name:  "Anastasiia",
		},
	}

	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(searchElement(c, "2"))
	fmt.Println(searchElement(d, User{
		Email: "lsd@gmail.com",
		Name:  "Lochart",
	}))

	printAny(d)
	printAny(a)
	printAny(b)
}

func sum[V Number](input []V) V {
	var result V

	for _, number := range input {
		result += number
	}

	return result
}

func searchElement[C comparable](elements []C, searchEl C) bool {
	for _, el := range elements {
		if el == searchEl {
			return true
		}
	}

	return false
}

func printAny[A any](input A) {
	fmt.Println(input)
}
