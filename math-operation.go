package main

import "fmt"

func main() {
	var result = 130 + 120
	fmt.Println(result)

	var a = 120
	var b = 110
	var c = a * b
	fmt.Println(c)

	//Versi Simple
	var i = 220
	i += 220 // i = i + 10
	fmt.Println(i)

	//Versi Lebih Simple
	i++ // i = i + i
	fmt.Println(i)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)

}
