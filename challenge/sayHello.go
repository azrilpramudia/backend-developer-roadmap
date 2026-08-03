package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Say Something: ")
	scanner.Scan()
	sayHello := scanner.Text()

	fmt.Print("Your Name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Your Place: ")
	scanner.Scan()
	place := scanner.Text()

	if sayHello == "Hai" {
		fmt.Printf("Hai %s dari %s\n", name, place)
	} else if sayHello == "Hello" {
		fmt.Printf("Hello %s dari %s\n", name, place)
	} else {
		fmt.Printf("Hai Apakah Boleh Kenalan?")
	}
}
