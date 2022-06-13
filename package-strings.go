package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Sehan Alfarisi", "Sehan"))
	fmt.Println(strings.Contains("Sehan Alfarisi", "Joko"))

	fmt.Println(strings.Split("Sehan Alfarisi Givi", " "))

	fmt.Println(strings.ToLower("Sehan Alfarisi Givi"))
	fmt.Println(strings.ToUpper("Sehan Alfarisi Givi"))
	fmt.Println(strings.ToTitle("Sehan Alfarisi Givi"))

	fmt.Println(strings.Trim("    Sehan Alfarisi    ", " "))

	fmt.Println(strings.ReplaceAll("Sehan Sehan Sehan", "Sehan", "Givi"))

}
