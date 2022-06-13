package main

import "fmt"

func main() {

	var name1 = "Sehan" // Jika nama beda maka hasilnya false
	var name2 = "sehan" // Jika nama sama maka hasilnya true.

	var result bool = name1 == name2
	fmt.Println(result)

	//Integer
	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
