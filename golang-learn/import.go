package main

import (
	"fmt"
	"golang-learn/helper"
)

func main() {
	result := helper.SayHello("Burhan")
	fmt.Println(result)
	
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidka bisa di akses
	// fmt.Println(helper.sayGoodbye("Ted")) // tidak bisa di akses
}