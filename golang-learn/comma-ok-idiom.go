package main

import "fmt"

func main() {
	myMap := map[string]int{"Apple": 5, "Banana": 10}
	
	value, ok := myMap["Apple"]

	if ok {
		fmt.Println("Value found: ", value)
	} else {
		fmt.Println("Key not found")
	}

	value, ok = myMap["cherry"]

	if ok {
		fmt.Println("Value Found: ", value)
	} else {
		fmt.Println("Key not found")
	}
}

