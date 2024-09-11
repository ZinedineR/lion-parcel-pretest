package main

import (
	"fmt"
	"strings"
)

// penggunaan rekursif: palindrom artinya tepi 'kiri' dan tepi 'kanan' hurufnya sama, terus check sampai tengah
// contoh radar >> palindrome("radar", 0, 4) >> rekursif terus hingga left dan right sama, dengan kondisi input[left] == input[right]
// contoh hello >> palindrome("hello", 0,4) >> langsung return false karena pada input[0] dan input[4] dia tidak sama, langsung saja return false
func palindrome(input string, left, right int) bool {
	if left >= right {
		return true
	} else if input[left] == input[right] {
		return palindrome(input, left+1, right-1)
	} else {
		return false
	}
}

func main() {
	//masukkan input di sini
	input := "kasur rusak"
	//menghapus spasi dan newline (contoh: kasur rusak)
	input = strings.TrimSpace(input)
	//membaca jumlah keseluruhan lenght
	var n = len(input) - 1
	fmt.Println("input:", input)
	fmt.Println(palindrome(input, 0, n))
}
