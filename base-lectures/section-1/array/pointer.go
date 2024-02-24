package main

import "fmt"

func main() {
	message := "I will be really good Software Engineer Golang"
	fmt.Println(message)
	changeMessage(&message)
	fmt.Println(message)

	number := 5
	var p *int

	p = &number

	fmt.Println(p)
	fmt.Println(&number)

	*p = 10

	fmt.Println(number)
	fmt.Println(p)
}

func changeMessage(message *string) {
	*message += " ( from the function changeMessage())"
}
