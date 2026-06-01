package main

import "fmt"

// rect is a struct type.
//
// Methods can be attached to named types.
// In Go, methods are just functions with a special receiver.
type rect struct {
	width, height int
}

// POINTER RECEIVER
//
// Receiver:
//     *rect
//
// This method operates on a pointer to rect.
//
// Use pointer receivers when:
//   - the method needs to modify the struct
//   - the struct is large and copying is expensive
//   - you want consistency across methods
//
// Even though this method doesn't modify the struct,
// it's still common to use pointer receivers.
func (r *rect) area() int {
	return r.width * r.height
}

// VALUE RECEIVER
//
// Receiver:
//     rect
//
// The struct is copied when this method is called.
//
// Best for:
//   - small structs
//   - read-only operations
//
// Changes made inside this method would only affect
// the copy, not the original struct.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	// Create a rect value.
	r := rect{
		width:  10,
		height: 5,
	}

	// Calling a POINTER receiver on a VALUE.
	//
	// Go automatically takes the address:
	//
	//     (&r).area()
	//
	// This works because r is addressable.
	fmt.Println("area:", r.area())

	// Normal value receiver call.
	fmt.Println("perim:", r.perim())

	// Explicit pointer.
	rp := &r

	// Calling a POINTER receiver on a POINTER.
	//
	// Exact match.
	fmt.Println("area:", rp.area())

	// Calling a VALUE receiver on a POINTER.
	//
	// Go automatically dereferences:
	//
	//     (*rp).perim()
	//
	// This is another convenience provided by Go.
	fmt.Println("perim:", rp.perim())
}
