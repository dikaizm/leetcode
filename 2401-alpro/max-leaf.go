package main

import (
	"fmt"
)

func main() {
	var num int = 0
	var max int = 0

	fmt.Print("Masukkan daun: ")
	fmt.Scan(&num)
	max = num

	for num != -1 {
		if num > max {
			max = num
		}
		fmt.Scan(&num)
	}

	fmt.Printf("Daun terbesar %d", max)
}

// Input:
// 3 2 5 2
// 1 10
// 5 9 5 8 30 1
