package main

import "fmt"

func main() {
	var name1 = "Edward"
	var name2 = "Edward"

	var result bool = name1 == name2

	fmt.Println(result)

	x := 5
	y := 7

	fmt.Println(x < y)
	fmt.Println(x > y)
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
}