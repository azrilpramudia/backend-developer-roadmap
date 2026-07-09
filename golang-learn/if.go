package main

import "fmt"

func main() {
	name := "Budi"

	if name == "Budi" {
		fmt.Println("Hello, Budi")
	} else if name == "Joko" {
		fmt.Println("Hello, Joko")
	} else if name == "Asep" {
		fmt.Println("Hello, Asep")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// if shorts
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}