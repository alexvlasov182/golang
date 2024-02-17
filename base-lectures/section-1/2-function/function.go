package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	message, err := checkAge(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

}
func printMessage(message string) {
	fmt.Println(message)
}
func sayHello(name string) string {
	return "Hello, " + name
}

func checkAge(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Enter, but be careful", nil
	} else if age >= 45 && age < 65 {
		return "You sure that this music will like you", nil
	} else if age >= 65 {
		return "This is too much for you", errors.New("you are too old")
	}

	return "Decline enter", errors.New("you are too young")
}
