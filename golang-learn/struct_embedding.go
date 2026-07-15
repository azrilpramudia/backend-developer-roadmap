package main

import "fmt"

type Address struct {
	Street, City, State string
}

type Person struct {
	Name string
	Address
}

func main() {
	
	p := Person{
		Name: "Alice",
		Address: Address{"123 main st", "Bandung", "ID"},
	}

	fmt.Println(p.Street)
}