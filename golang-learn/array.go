package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Edward"
	name[1]	= "Developer"
	name[2] = "Happy"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	// simple syntax
	var values = [...]int {
		90,
		80,
		95,
		100,
		110,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}