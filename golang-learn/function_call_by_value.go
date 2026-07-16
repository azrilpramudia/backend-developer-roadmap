package main

import "fmt"

func changeName(name string) {
	name = "meow"
}

func main() {
	namaAsli := "Edward"
	changeName(namaAsli)

	fmt.Println("Nama Asli:", namaAsli)
}