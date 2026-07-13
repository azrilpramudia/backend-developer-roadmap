package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string ) {
	firstName = "Burhan"
	middleName = "Udin"
	lastName = "Fire"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}