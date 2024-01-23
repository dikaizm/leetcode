package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a >= b {
		fmt.Print(permutasi(a, b))
	} else {
		fmt.Print(permutasi(b, a))
	}
}

func permutasi(n int, r int) int {
	return factorial(n) / factorial(n-r)
}

func factorial(n int) int {
	var result int = 1
	var i int
	for i = 1; i <= n; i++ {
		result = result * i
	}
	return result
}
