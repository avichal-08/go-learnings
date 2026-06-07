package main

func isPalindrome(x int) bool {
	y := 0
	for i := x; i > 0; i = i / 10 {
		y = y*10 + i%10
	}
	return y == x
}
