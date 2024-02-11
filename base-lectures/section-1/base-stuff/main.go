package main

import (
	"fmt"
	"reflect"
)

func main() {
	message := "I'm Ninja GO"
	fmt.Println(reflect.TypeOf(message))
	fmt.Println(message)

	var messageInt int
	messageInt = 12

	fmt.Println(reflect.TypeOf(messageInt))
	fmt.Println(messageInt)

	messageByte := []byte("hi!")
	fmt.Println(messageByte)

	var a byte = 62
	fmt.Printf("%c\n", a)

	var b rune = 'A'
	fmt.Println(b)

	x, y, z := 1, 2, 3

	fmt.Println(x, y, z)
	x, y = y, x
	fmt.Println(x, y, z)

}
