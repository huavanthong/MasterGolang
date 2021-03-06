package main

import (
	"fmt"
	"math"
)

/*************************************************************************************************/
/* Parent structure using interface
/*************************************************************************************************/
// Go Interface - `Shape`
type Shape interface {
	Area() float64
	Perimeter() float64
}

/*************************************************************************************************/
/* Rectangle will be the Shape'children if it implements all Shape's methods
/*************************************************************************************************/
// Struct type `Rectangle` - implements the `Shape` interface by implementing all its methods.
type Rectangle struct {
	Length, Width float64
}

/********* Methods inherit *********/
// Methods inherit from Shape
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

/********* Methods utility *********/
// Set width by using pointer
func (r *Rectangle) setWidth(width float64) {
	(*r).Width = width
}

// Set length by using pointer
func (r *Rectangle) setLength(length float64) {
	(*r).Length = length
}

// show info by using value
func (r Rectangle) show() {
	fmt.Println("Length : ", r.Length)
	fmt.Println("Width  : ", r.Width)
}

/*************************************************************************************************/
/* Circle will be the Shape'children if it implements all Shape's methods
/*************************************************************************************************/
// Struct type `Circle` - implements the `Shape` interface by implementing all its methods.
type Circle struct {
	Radius float64
}

/********* Methods inherit *********/
// Methods inherit from Shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

/********* Methods extend *********/
// Methods extend
func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

/********* Methods utility *********/
func (c *Circle) setRadius(radius float64) {
	(*c).Radius = radius
}

func (c Circle) show() {
	fmt.Println("Radius : ", c.Radius)
}

// Generic function to calculate the total area of multiple shapes of different types
func CalculateTotalArea(shapes ...Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	return totalArea
}

func main() {

	var circle Circle = Circle{4.0}
	fmt.Printf("Circle Type = %T\n", circle)
	fmt.Printf("Area = %f, Perimeter = %f, Diameter = %f\n\n", circle.Area(), circle.Perimeter(), circle.Diameter())
	circle.show()

	var s Shape = Circle{5.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())
	// We can't use show() for Shape
	// s.show()

	s = Rectangle{4.0, 6.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n", s.Area(), s.Perimeter())
	// We can't use show() for Shape was converted to Rectangle
	// s.show()

	totalArea := CalculateTotalArea(Circle{2}, Rectangle{4, 5}, Circle{10})
	fmt.Println("Total area = ", totalArea)
	fmt.Println()

	var rect Rectangle
	rect.setLength(10)
	rect.setWidth(12)
	fmt.Printf("Rect Type = %T, Rhape Value = %v\n", rect, rect)
	fmt.Printf("Area = %f, Perimeter = %f\n", rect.Area(), rect.Perimeter())
	rect.show()
}
