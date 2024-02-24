package main

import "fmt"

func main() {
	messages := []string{
		"messages 1",
		"messages 2",
		"messages 3",
		"messages 4",
	}

	for _, message := range messages {
		fmt.Println("Much Better for:", message)
	}

	for i := 0; i < len(messages); i++ {
		fmt.Println("Simple for: ", i)
	}

	counter := 0
	for {
		if counter == 100 {
			break
		}
		counter++
		fmt.Println(counter)
	}

}
