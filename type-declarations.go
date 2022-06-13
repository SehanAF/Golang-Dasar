package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpSehan NoKTP = "34241048164175"
	var marriedStatus Married = true
	fmt.Println(noKtpSehan)
	fmt.Println(marriedStatus)
}
