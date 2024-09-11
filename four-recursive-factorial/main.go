package main

import (
	"fmt"
)

// rumus dasar factorial adalah n! = n. (n-1) !
// lakukan rekursif terus di mana n hingga n = 1, di mana dia akan mereturn 1
// contoh 5! = 5 * factorial(5-1) >> di mana factorial (5-1) = 4 * factorial (4-1)... dst
func factorial(input int) int {
	if input == 1 {
		return input
	} else {
		return input * factorial(input-1)
	}
}

func main() {
	//input di sini
	//pastikan int adalah bukan negatif, karena faktorial pada angka negatif tidak terdefinisi
	input := 5
	fmt.Println("length:", input)
	fmt.Println("jumlah factorial: ", factorial(input))
}
