package main

import "fmt"

func main() {
	const (
		firstName string = "Edward"
		lastName = "Developer"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Error
	// firstName = "Budi"
	// lastName = "Joko"
}


// const firstName string = "Edward"
// const lastName = "Developer"