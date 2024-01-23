package main

import (
	"fmt"
)

func main() {
	var num int = 0
	var digit int = 0
	var currDigit int = 0
	var found bool = false

	fmt.Scan(&digit)
	fmt.Scan(&num)

	for num > 0 && !found {
		currDigit = num % 10
		num = num / 10

		if currDigit == digit {
			found = true
		}
	}

	fmt.Print(found)
}
