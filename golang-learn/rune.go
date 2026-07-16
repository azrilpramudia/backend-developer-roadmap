package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	teks := "Hello 😼"

	// cara yang salah
	fmt.Println("Panjang pakai len() :", len(teks))

	// cara yang benar
	fmt.Println("Panjang menggunakan rune :", utf8.RuneCountInString(teks))

	arrayRune := []rune(teks)
	fmt.Println("Panjang pakai konversi []rune :", len(arrayRune))
}