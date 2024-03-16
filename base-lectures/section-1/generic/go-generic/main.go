package main

import "fmt"

type Number interface {
	int64 | float64
}

type User struct {
	Email string
	Name  string
}

func main() {
	a := []int64{1, 2, 3}
	b := []float64{1.1, 2.2, 3.3}
	c := []string{"1", "2", "3"}

	d := []User{
		{
			Email: "asd@gmail.com",
			Name:  "Alex",
		},
		{
			Email: "123asd@gmail.com",
			Name:  "Galex",
		},
		{
			Email: "321asd@gmail.com",
			Name:  "Shmalex",
		},
	}

	fmt.Println()
	fmt.Print()
	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(searchElement(c, "6"))
	fmt.Println(searchElement(d, User{
		Email: "321asd@gmail.com",
		Name:  "Shmalex",
	}))
	printAny(d)
	printAny(a)
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
