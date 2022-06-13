package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name

}

func main() {
	sehan := Man{"Sehan"}
	sehan.Married()

	fmt.Println(sehan.Name)

}
