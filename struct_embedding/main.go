package main

import "fmt"

// base is a normal struct.
//
// It contains data and methods.
//
// Unlike OOP languages, Go does not support
// class inheritance.
type base struct {
	num int
}

// Methods can be attached to structs.
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container EMBEDS base.
//
// Notice:
//
//     base
//
// instead of:
//
//     b base
//
// This is called embedding.
//
// Embedding provides:
//
//   - field promotion
//   - method promotion
//
// IMPORTANT:
//
// This is NOT inheritance.
//
// container does not become a subtype of base.
//
// Instead:
//
//     container HAS A base
//
// not:
//
//     container IS A base
type container struct {

	// Embedded field.
	base

	// Normal field.
	str string
}

func main() {

	// Construct container.
	//
	// Since base is a field,
	// we initialize it like any other field.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Even though num belongs to base,
	// Go promotes the field.
	//
	// These are equivalent:
	//
	//     co.num
	//     co.base.num
	//
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Explicit access.
	fmt.Println("also num:", co.base.num)

	// container does NOT define describe().
	//
	// However, base does.
	//
	// Since base is embedded,
	// its methods are promoted to container.
	//
	// So Go allows:
	//
	//     co.describe()
	//
	// Internally:
	//
	//     co.base.describe()
	//
	fmt.Println("describe:", co.describe())

	// Interface requiring describe().
	type describer interface {
		describe() string
	}

	// container does not explicitly implement
	// describe().
	//
	// But because describe() is promoted from
	// embedded base, container satisfies
	// describer automatically.
	var d describer = co

	fmt.Println("describer:", d.describe())
}
