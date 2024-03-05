package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square{5}
	circle := Circle{8}

	printShapeArea(square)
	printShapeArea(circle)

	printInterface(square)
	printInterface(circle)
	printInterface(22)
	printInterface(true)
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

func printInterface(i interface{}) {
	switch value := i.(type) {
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("bool", value)
	default:
		fmt.Println("unknown type", value)
	}
}
