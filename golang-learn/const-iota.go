package main

import "fmt"

func main() {
	const (
		A = iota
		B
		C
		D
	)

	fmt.Println("Nilai A:", A)
	fmt.Println("Nilai B:", B)
	fmt.Println("Nilai C:", C)
	fmt.Println("Nilai D:", D)
}