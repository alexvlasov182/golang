package main

import (
	"fmt"
	"math"
)

// Compositon of Interfaces
type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

func (s Square) Perimeter() float32 {
	return s.sideLength * 4
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
	circle1 := Square{6}

	printShapeArea(square)
	// printShapeArea(circle)

	printInterface(square)
	printInterface(circle)
	printShapeArea(circle1)
	printInterface("qwe")
	printInterface(22)
	printInterface(true)
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

func printInterface(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case bool:
		fmt.Println("bool", v)
	default:
		fmt.Println("unknown type", v)
	}
}
