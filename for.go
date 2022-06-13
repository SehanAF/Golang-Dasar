package main

import "fmt"

func main() {
	counter := 1

	for counter < 9 {
		fmt.Println("perulagan ke", counter)
		counter++
	}

	//Versi simple

	for counter := 1; counter < 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Sehan", "Alfarisi", "Kasep"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])

	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)

		//Bila tidak mau memakai Index ( i ganti dengan _ )
		//fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Sehan"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
