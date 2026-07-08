package main

import "fmt"

func main() {
	
	type NoKTP string

	var ktpEdo NoKTP = "111111111"

	var contoh string = "22222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpEdo)
	fmt.Println(contohKtp)
}