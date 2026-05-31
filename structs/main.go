package main

import "fmt"

// Struct = collection of related fields.
//
// Think of it as a custom type that groups data together.
//
// Similar to:
// - class (without inheritance) in Java
// - object shape in JavaScript
// - struct in C
type User struct {
	Name string
	Age  int
}

func main() {

	// field order doesn't matter.
	u1 := User{
		Name: "Avichal",
		Age:  20,
	}

	fmt.Println(u1)

	// Access fields using dot notation.
	fmt.Println(u1.Name)
	fmt.Println(u1.Age)

	// Modify a field.
	u1.Age = 21

	fmt.Println(u1)

	// Positional initialization.
	// Avoid in production because field order matters.
	u2 := User{"John", 25}

	fmt.Println(u2)

	// Zero-value struct.
	// Every field gets its zero value automatically.
	var u3 User

	fmt.Println(u3)

	// Structs are VALUE TYPES.
	// Assignment creates a copy.
	u4 := u1

	u4.Name = "Copied User"

	fmt.Println("Original:", u1.Name)
	fmt.Println("Copy:", u4.Name)

	// Original remains unchanged.
}
