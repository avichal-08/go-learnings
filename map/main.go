package main

import (
	"fmt"
	// "maps"
)

func main() {
	m := make(map[string]int)

	m["key1"] = 0
	m["key2"] = 2

	key1val, key1Present := m["key1"]
	fmt.Println("key1 value:", key1val)
	fmt.Println("key1 is present:", key1Present)

	_, key2Present := m["key2"]
	fmt.Println("key2 is present:", key2Present)
}
