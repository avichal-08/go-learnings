package main

import "fmt"

func main() {
	a := 4
	fmt.Println("add of a", &a)
	square(&a)
	fmt.Println("value of a after square", a)

	ptr := &a
	fmt.Println(*ptr)

}

func square(n *int) {
	fmt.Println("add of n", *n)
	*n *= *n
}

/* if a func is returning pointer of a variable(ex: square returnin pointer to main) then
the compiler will use heap to store that returning memory for main func to access using pointer
because stack gets eliminated automatically after function life-time */
