package main

import (
	"fmt"
	"math"
)

func trianglePrint(input int) {
	//handle positif int, basically hanya perlu print * sejumlah angka terkecil setelah 0 dari input, print * lalu print newline
	//terdapat dua loop, loop pertama adalah loop dari angka terkecil setelah 0 ke jumlah input
	//loop kedua berada di dalam loop pertama, print * sejumlah angka loop pertama
	if input > 0 {
		for i := 1; i <= input; i++ {
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
		//handle negatif int, kebalikan dari input positif, bedanya * hanay di print di angka terakhir jumlah angka terbesar setelah input
		//kuncinya terletak pada penggunaan math.Abs, untuk mengubah menjadi angka positif sehingga bisa dilakukan looping
		//terdapat 3 loop, loop pertama seperti penggunaan input positif
		//loop kedua melakukan print spasi hingga sebelum jumlah inputan - loop pertama
		//loop ketiga melakukan print * sejumlah loop pertama
	} else if input < 0 {
		// Print reversed triangle (right aligned)
		absLength := int(math.Abs(float64(input)))
		for i := 1; i <= absLength; i++ {
			for j := 0; j < absLength-i; j++ {
				fmt.Print(" ")
			}
			for k := 0; k < i; k++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}

func main() {
	//ubah input di sini
	input := 6
	fmt.Println("print triangle:", input)
	trianglePrint(input)

	input = -5
	fmt.Println("print triangle:", input)
	trianglePrint(input)
}
