package main

import "fmt"

func endApp() {
	fmt.Println("End app")

	// recover
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp() // defer

	if error {
		panic("Ups error") // panic
	}
}

func main() {
	runApp(true)
	fmt.Println("Edward")
}