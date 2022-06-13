package main

import (
	"fmt"
	"golang-dasar/database" // tinggal menambah kan _ dan hapus fmt
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
