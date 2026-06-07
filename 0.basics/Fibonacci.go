package main

func fibonacci(x int) int {
	a, b := 0, 1

	for i := 2; i <= x; i++ {
		a, b = b, a+b
	}

	return b
}
