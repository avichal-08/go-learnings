package main

func factorial(x int) int {

	fac := 1

	for i := 2; i <= x; i++ {
		fac = fac * i
	}

	return fac

	//-----------Recursion
	/* if x == 1 {
		return x
	} else {
		return x * factorial(x-1)
	} */
}
