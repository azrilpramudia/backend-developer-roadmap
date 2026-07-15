package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

// Struct method
func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var edward Customer
	fmt.Println(edward)

	edward.Name = "Burhan"
	edward.Address = "USA"
	edward.Age = 20
	fmt.Println(edward)
	fmt.Println(edward.Name)
	fmt.Println(edward.Address)
	fmt.Println(edward.Age)

	joko := Customer {
		Name: "Joko",
		Address: "Indonesia",
		Age: 20,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 40}
	fmt.Println(budi)

	budi.sayHello("Agus")
	edward.sayHello("budi")
	joko.sayHello("edi")
}
