package main

import "fmt"

func getFullName() (string, string) {
	return " Sehan", "Alfarisi"

}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}

// Jika ingin menghiraukan cukup menggunakan tanda _(GARIS BAWAH)

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
