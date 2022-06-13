package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	helper.SayHello("Sehan")
	// helper.sayGoodBye("Sehan") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error

}
