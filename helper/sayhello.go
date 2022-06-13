package helper

import "fmt"

var version = 1
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

//di satu package tidak boleh ada yang sama

func sayGoodBye(name string) {
	fmt.Println("Goodbye", name)
}
