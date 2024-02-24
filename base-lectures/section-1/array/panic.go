package main

import "fmt"

func main() {
	// panic("aaaaaaaa hilfe")
	messages := []string{
		"messages 1",
		"messages 2",
		"messages 3",
		"messages 4",
	}

	messages[4] = "messages 5"
	fmt.Println(messages)

}
