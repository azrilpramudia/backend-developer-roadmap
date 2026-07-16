package main

import (
	"fmt"
	"golang-learn/database"
	_ "golang-learn/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}