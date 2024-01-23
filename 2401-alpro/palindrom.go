package main

import "fmt"

const ARR_SIZE int = 256

func main() {
	var arr [ARR_SIZE]int
	var n int

	fmt.Print("Masukkan jumlah array: ")
	fmt.Scan(&n)

	// input array
	inputArr(n, &arr)
	displayArr(n, arr)

	// reverse isi array
	reverseArr(n, &arr)
	displayArr(n, arr)

	// cek palindrom
	if palindrom(n, arr) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func inputArr(n int, arr *[ARR_SIZE]int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
}

func reverseArr(n int, arr *[ARR_SIZE]int) {
	var temp, i int
	for i = 0; i < n/2; i++ {
		temp = arr[i]
		arr[i] = arr[n-i-1]
		arr[n-i-1] = temp
	}
}

func palindrom(n int, arr [ARR_SIZE]int) bool {
	var i int
	var match bool = true

	for i = 0; i < n/2 && match; i++ {
		if arr[i] != arr[n-i-1] {
			match = false
		}
	}

	return match
}

func displayArr(n int, arr [ARR_SIZE]int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println("")
}
