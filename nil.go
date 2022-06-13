package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	// dibawah ini adalah dengan "nil"
	var person map[string]string = NewMap("Sehan")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

	// dibawah ini adalah tanpa nill
	/**var person map[string]string

	if person["name"] == "" {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

	*/
}
