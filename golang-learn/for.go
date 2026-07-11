package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke: ", counter)
	// 	counter++
	// }

	// for with short statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke: ", counter)
	}

	// manual for
	fmt.Println("Selesai")

	names := []string{"eko", "burhan", "subhan"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// For Range
	for index, name := range names{
		fmt.Println("index", index, "=", name)
	}

	// don't need index
	for _, name := range names{
		fmt.Println(name)
	}
}