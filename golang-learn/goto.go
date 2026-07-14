package main

import "fmt"

func main() {
	fmt.Println(1)
	goto NEXT
	fmt.Println(2)
NEXT:
	fmt.Println(3)
}