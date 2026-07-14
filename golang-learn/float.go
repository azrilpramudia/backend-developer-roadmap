package main

import "fmt"

func main() {
	var price float64 = 19.99
	tax := 1.5

	fmt.Println("Price: ", price)
	fmt.Println("Tax: ", tax)

	total := price + tax
	fmt.Println("Total: ", total)
}