package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Burhan")
	fmt.Println(result)

	fmt.Println(getHello("edward"))
	fmt.Println(getHello("joko"))
}