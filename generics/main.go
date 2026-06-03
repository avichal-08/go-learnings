package main

import "fmt"

// SlicesIndex searches for a value inside any slice type.
//
// Before generics, we'd need:
//
//     IndexInt([]int)
//     IndexString([]string)
//     IndexFloat([]float64)
//
// Generics allow writing the algorithm once and
// reusing it for many types.
//
// Type Parameters:
//
// S -> slice type
// E -> element type
//
// Constraint:
//
// S ~[]E
//
// Means:
//   S must have an underlying type of []E.
//
// Examples:
//
//   []string
//   []int
//
// Also works for custom slice types:
//
//   type Names []string
//
// because its underlying type is []string.
//
// Constraint:
//
// E comparable
//
// Means values of type E can be compared using:
//
//   ==
//   !=
//
// Required because we perform:
//
//   v == s[i]
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {

	// Iterate through all elements.
	for i := range s {

		// Generic comparison.
		//
		// Works because E satisfies comparable.
		if v == s[i] {
			return i
		}
	}

	// Not found.
	return -1
}

// List is a generic linked list.
//
// T represents the type of data stored.
//
// Examples:
//
//   List[int]
//   List[string]
//   List[User]
//
// T can be ANY type because of:
//
//   any
//
// any is simply:
//
//   interface{}
type List[T any] struct {
	head, tail *element[T]
}

// One node of the linked list.
//
// Stores:
//
//   value
//   pointer to next node
//
// element[int]
//
// becomes:
//
//   struct {
//       next *element[int]
//       val int
//   }
//
// element[string]
//
// becomes:
//
//   struct {
//       next *element[string]
//       val string
//   }
type element[T any] struct {
	next *element[T]
	val  T
}

// Push appends a value to the end of the list.
//
// Because List is generic:
//
//   List[int]    -> Push(int)
//   List[string] -> Push(string)
//
// The compiler generates type-safe versions.
func (lst *List[T]) Push(v T) {

	// Empty list.
	if lst.tail == nil {

		// Create first node.
		lst.head = &element[T]{val: v}
		lst.tail = lst.head

	} else {

		// Create new node.
		lst.tail.next = &element[T]{val: v}

		// Advance tail pointer.
		lst.tail = lst.tail.next
	}
}

// Traverses the linked list and returns
// all values as a slice.
func (lst *List[T]) AllElements() []T {

	var elems []T

	// Standard linked-list traversal.
	//
	// Start at head.
	// Continue until nil.
	for e := lst.head; e != nil; e = e.next {

		// Collect values.
		elems = append(elems, e.val)
	}

	return elems
}

func main() {

	var s = []string{
		"foo",
		"bar",
		"zoo",
	}

	// Compiler infers:
	//
	// S = []string
	// E = string
	//
	// So we don't need to specify
	// type arguments manually.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// Explicit type arguments.
	//
	// Rarely needed because Go can
	// usually infer them automatically.
	_ = SlicesIndex[[]string, string](s, "zoo")

	// Create a linked list storing ints.
	//
	// Here:
	//
	// T = int
	lst := List[int]{}

	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	fmt.Println("list:", lst.AllElements())
}
