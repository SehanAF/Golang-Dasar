package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func sayHi(customer Customer, name string) { // Ini adalah Struct Method
	fmt.Println("Hello", name, "My Name Is", customer.Name)
}

func main() { // ini adalah struct
	var sehan Customer //< adalah cara ke 1
	sehan.Name = "Sehan Alfarisi"
	sehan.Address = "Indonesia"
	sehan.Age = 30

	sayHi(sehan, "Agung")

	fmt.Println(sehan)
	fmt.Println(sehan.Name)
	fmt.Println(sehan.Address)
	fmt.Println(sehan.Age)

	alexa := Customer{ //< adalah cara ke 2
		Name:    "alexa",
		Address: "semarang",
		Age:     35,
	}
	fmt.Println(alexa)

	jaka := Customer{"Budi", "Jakarta", 35} //< adalah cara ke 3
	fmt.Println(jaka)
}
