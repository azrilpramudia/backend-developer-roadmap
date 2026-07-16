package main

import "fmt"

func main() {
	c1 := complex(2, 3)
	c2 := 4 + 5i

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)

	// angka kompleks terdiri dari 2 bagian real dan imaginary
	fmt.Println("Real:", real(c1))
	fmt.Println("Imaginary:", imag(c1))
}