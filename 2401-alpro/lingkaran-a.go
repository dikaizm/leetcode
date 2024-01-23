package main

import (
	"fmt"
	"math"
)

func main() {
	// Membaca koordinat pusat dan radius lingkaran 1
	fmt.Print("Lingkaran 1: ")
	var x1, y1, r1 int
	fmt.Scanln(&x1, &y1, &r1)

	// Membaca koordinat pusat dan radius lingkaran 2
	fmt.Print("Lingkaran 2: ")
	var x2, y2, r2 int
	fmt.Scanln(&x2, &y2, &r2)

	// Membaca koordinat titik sembarang
	fmt.Print("Titik sembarang: ")
	var x, y int
	fmt.Scanln(&x, &y)

	// Menghitung jarak titik ke pusat lingkaran 1
	distanceToCircle1 := distanceToCircle(x, y, x1, y1)

	// Menghitung jarak titik ke pusat lingkaran 2
	distanceToCircle2 := distanceToCircle(x, y, x2, y2)

	// Memeriksa posisi titik
	checkPosition(distanceToCircle1, distanceToCircle2, r1, r2)
}

// Euclidean formula
// d(p, q)^2 = (q1-p1)^2 + (q2-p2)^2
// d(p, q) = sqrt((q1-p1)^2 + (q2-p2)^2)
func distanceToCircle(x int, y int, cx int, cy int) float64 {
	return math.Sqrt(float64((x-cx)*(x-cx) + (y-cy)*(y-cy)))
}

func checkPosition(dCircle1 float64, dCircle2 float64, r1 int, r2 int) {
	if dCircle1 <= float64(r1) && dCircle2 <= float64(r2) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dCircle1 <= float64(r1) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dCircle2 <= float64(r2) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
