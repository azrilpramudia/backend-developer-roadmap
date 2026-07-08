package main

import "fmt"

func main() {
	var name string

	name = "Edward Developer"
	fmt.Println(name)

	name = "Edward Dev"
	fmt.Println(name)

	// Declaration for all variable
	var (
		firstName = "Edward"
		lastName = "Developer"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}

// Optional
// var name = "Edward"
// fmt.Println(name)

// Simple syntax
// name := "Edward"
// fmt.Println(name)