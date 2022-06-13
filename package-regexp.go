package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("agero"))
	fmt.Println(regex.MatchString("saLsa"))

	search := regex.FindAllString("Eko eka edo eto eyo eki", 10)
	fmt.Println(search)
}

/** Regular Expression wajib dipelajari lebih detail karena
Sangat penting untuk membuat validasi atau apapun itu
Yang berhubung dengan text, akan pasti terpakai */
