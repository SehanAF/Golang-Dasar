package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Cuyy"
	} else {
		return "Hello " + name
	}

}

func main() {
	result := getHello("Sehan")
	fmt.Println(result)

	fmt.Println(getHello(""))

}
