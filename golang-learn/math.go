package main

import "fmt"

// math operators
func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d*e

	fmt.Println(c)

	f := 10
	g := 5

	fmt.Println(f + g)
	fmt.Println(f - g)
	fmt.Println(f * g)
	fmt.Println(f / g)
	fmt.Println(f % g)

	
	// augmented assignment
	var i = 10
	i += 10 // i = i + 10

	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)


	// Unary Operator
	var j = 1
	j++ // j = j + 1
	fmt.Println(j)
	j++ // j = j + 1
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)
}