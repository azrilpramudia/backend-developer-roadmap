package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Budi"
	// person["address"] = "Munich"

	person := map[string]string{
		"name" : "Budi",
		"address" : "Bandung",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book ["title"] = "Buku golang"
	book ["author"] = "Budi"
	book ["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}