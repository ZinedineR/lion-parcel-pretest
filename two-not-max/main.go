package main

import (
	"fmt"
	"math"
)

// pada dasarnya menggunakan for loop untuk pengecekkan, inisialisasi nilai paling terkecil
// lalu tiap looping di cek apakah input pada array lebih besar dari nilai paling terkecil
func maxChecker(numbers []int) int {
	//agar bisa menghandle input negatif
	maxNumber := math.MinInt
	for _, number := range numbers {
		if number > maxNumber {
			maxNumber = number
		}
	}
	return maxNumber
}

func main() {
	numbers := []int{3, 5, 1, 9, 2}
	fmt.Println("array: ", numbers)
	fmt.Println("angka paling besar: ", maxChecker(numbers))

	numbers = []int{-1, -3249234, -1213123213123213}
	fmt.Println("array: ", numbers)
	fmt.Println("angka paling besar: ", maxChecker(numbers))
}
