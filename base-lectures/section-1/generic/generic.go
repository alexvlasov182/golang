package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

type User struct {
	email string
	name  string
}

func main() {
	a := []int64{1, 2, 3}
	b := []float64{1.1, 2.2, 3.3}
	c := []string{"1", "2", "3"}

	d := []User{
		{
			email: "asd@gmail.com",
			name:  "Oleksandr",
		},
		{
			email: "lsd@gmail.com",
			name:  "Anastasiia",
		},
	}

	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(serchElement(c, "2"))
	fmt.Println(serchElement(d, User{
		email: "lsd@gmail.com",
		name:  "Lochart",
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

func serchElement[C comparable](elements []C, searchEl C) bool {
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
