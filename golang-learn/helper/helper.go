package helper

import "fmt"

var version = "1.0.0"
var Application = "golang"

func sayGoodbye(name string) string {
	return "GoodBye" + name
}

func SayHello(name string) string {
	return "Hello " + name
}

// Access modifier
func Contoh() {
	sayGoodbye("Ted")
	fmt.Println(version)
}