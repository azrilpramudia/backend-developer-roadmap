// zero value
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var p *int
	var slice []int
	var m map [string]int

	fmt.Println("int:", i)
	fmt.Println("float:", f)
	fmt.Println("bool:", b)
	fmt.Println("string:", s)
	fmt.Println("slice:", slice)
	fmt.Println("pointer:", p)
	fmt.Println("map:", m)
}