package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.City = "Bandung Barat"
	// fmt.Println(address1)
	// fmt.Println(address2)

	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 *Address =  &address1 // pointer

	address2.City = "Bandung Barat"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // beruabh menjadi bandung
}