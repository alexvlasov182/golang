package main

import "fmt"

func main() {
	defer handlerPanic()

	messages := []string{
		"messages 1",
		"messages 2",
		"messages 3",
		"messages 4",
	}

	messages[4] = "messages 5"
	fmt.Println(messages)
}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}

	fmt.Println("handlerPanic() executed successfully")
}
