package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(prediction("Monday"))
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

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "Monday":
		return "Wish you good start :)", nil
	case "Tuesday":
		return "Good tuesday", nil
	case "Wednesday":
		return "This is middle of week, that's gooood ;)", nil
	default:
		return "I want a weekend", errors.New("the best days of the week")
	}
}
