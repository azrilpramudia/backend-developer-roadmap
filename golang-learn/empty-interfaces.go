package main

import (
	"fmt"
	"os"
)

func EmptyInt(a ...any) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}

func main() {
	EmptyInt("Number: ", 10, "Status: ", true, 18.00)
}
