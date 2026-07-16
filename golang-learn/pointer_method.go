package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	ted := Man{"Ted"}
	ted.Married()

	fmt.Println(ted.Name)
}