package main

import "fmt"

func main() {
	var name = "Alexa"

	if name == "Sehan" {
		fmt.Println("Hello Sehan")
	} else if name == "Alfarisi" {
		fmt.Println("Hello Alfarisi")
	} else if name == "Alexa" {
		fmt.Println("Hello Alexa")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	var length = len(name)
	if length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
