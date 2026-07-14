package main

import "fmt"

func main() {
	i := 0
BEGIN:
	fmt.Println("Number", i)
	if i == 9 {
		goto END
	}
	i++
	goto BEGIN
END:
}