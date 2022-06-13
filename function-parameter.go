package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {

	sayHelloTo("Sehan", "Alfarisi")
	sayHelloTo("Givi", "Alexa")
}
