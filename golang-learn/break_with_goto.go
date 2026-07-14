package main

import "fmt"

func main() {
	for i := 0; ; i++{
		if i == 10 {
			goto END
		}
		fmt.Println(i)
	}
END:
	fmt.Println("END")
}