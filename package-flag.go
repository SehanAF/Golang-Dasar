package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localHost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "Put your number")

	flag.Parse() // Parse itu penting, jika tidak menggunakan parse, si data tidak akan keluar.

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	fmt.Println("Number : ", *number)

}
