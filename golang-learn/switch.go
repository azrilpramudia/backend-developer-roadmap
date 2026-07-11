package main

import "fmt"

func main() {
	name := "Edward"

	switch name {
	case "Edward":
		fmt.Println("Hello Edward")
	case "Burhan":
		fmt.Println("Hello Burhan")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// switch short statements
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch non condition
	name = "Dev"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")	
	}
}