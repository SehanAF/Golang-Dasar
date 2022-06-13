package main

import "fmt"

func main() {

	name := "Sehan raka"

	switch name {
	case "Sehan":
		fmt.Println("Hello Sehan")
		fmt.Println("Hello Sehan")
	case "Raka":
		fmt.Println("Hello Raka")
		fmt.Println("Hello Raka")
	default:
		fmt.Println("Kenalan dong")
		fmt.Println("Kenalan dong")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 18:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")

	}
}
