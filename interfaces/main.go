package main

import (
	"fmt"
	"math"
)

// An interface defines BEHAVIOR.
//
// Any type that implements all methods of an interface
// automatically satisfies that interface.
//
// No explicit "implements" keyword exists in Go.
//
// geometry can represent ANY type that knows how to:
//   - calculate area
//   - calculate perimeter
type geometry interface {
	area() float64
	perim() float64
}

// rect stores rectangle dimensions.
type rect struct {
	width, height float64
}

// circle stores circle dimensions.
type circle struct {
	radius float64
}

// Because rect has area() and perim(),
// it automatically satisfies geometry.
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Because circle also has area() and perim(),
// it automatically satisfies geometry.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// measure accepts a geometry interface.
//
// This means it can work with ANY type that satisfies geometry.
//
// Examples:
//
//	rect
//	circle
//
// measure does not care about the concrete type.
// It only cares about the behavior promised by geometry.
//
// This is called programming to an interface rather than
// programming to a concrete implementation.
func measure(g geometry) {

	// Printing an interface prints the concrete value
	// stored inside it.
	fmt.Println(g)

	// Dynamic dispatch:
	//
	// If g contains rect:
	//     rect.area() is called.
	//
	// If g contains circle:
	//     circle.area() is called.
	//
	// The correct method is selected at runtime
	// based on the concrete type stored inside
	// the interface.
	fmt.Println(g.area())

	// Same dynamic dispatch happens here.
	fmt.Println(g.perim())
}

// Sometimes an interface is too generic and we need
// access to the original concrete type.
//
// A type assertion asks:
//
//	"Is the value inside this interface a circle?"
//
// If yes:
//
//	extract the circle value.
//
// If no:
//
//	return ok = false.
func detectCircle(g geometry) {

	// Safe type assertion.
	//
	// c  -> extracted circle value
	// ok -> whether assertion succeeded
	if c, ok := g.(circle); ok {

		// We can now access fields that exist
		// only on circle.
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {

	// Concrete rectangle value.
	r := rect{
		width:  3,
		height: 4,
	}

	// Concrete circle value.
	c := circle{
		radius: 5,
	}

	// Interface value created internally:
	//
	// geometry
	// ├── type: rect
	// └── value: {3,4}
	//
	// measure only sees geometry.
	measure(r)

	// Interface value created internally:
	//
	// geometry
	// ├── type: circle
	// └── value: {5}
	//
	// Same function works without modification.
	measure(c)

	// r contains a rect.
	// Assertion to circle fails.
	detectCircle(r)

	// c contains a circle.
	// Assertion succeeds.
	detectCircle(c)
}
