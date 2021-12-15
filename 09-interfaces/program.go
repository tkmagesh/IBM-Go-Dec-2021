package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func main() {
	c := Circle{Radius: 10}
	//fmt.Println("Area = ", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{Height: 10, Width: 20}
	//fmt.Println("Area = ", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}

type ShapeWithArea interface {
	Area() float64
}

func PrintArea(shapeWithArea ShapeWithArea) {
	fmt.Println("Area = ", shapeWithArea.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

func PrintPerimeter(shapeWithPerimeter ShapeWithPerimeter) {
	fmt.Println("Perimeter = ", shapeWithPerimeter.Perimeter())
}

/*
type Shape interface {
	Area() float64
	Perimeter() float64
}
*/

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintShape(s Shape) {
	PrintArea(s)
	PrintPerimeter(s)
}
