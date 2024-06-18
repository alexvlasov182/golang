package main

import "fmt"

func main() {
	defer printMessageNew()

	fmt.Println("main()")
}

func printMessageNew() {
	fmt.Println("PrintMessageNew")
}
