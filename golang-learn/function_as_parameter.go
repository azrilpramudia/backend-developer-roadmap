package main

import "fmt"

// type declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "gila" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Burhan", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("gila", filter)
}