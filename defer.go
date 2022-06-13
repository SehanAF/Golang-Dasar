package main

import "fmt"

// Defer

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runAppLication(value int) {
	defer logging()
	fmt.Println("Run appLication")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runAppLication(0) // contoh 0, nilai asli adalah 10 biar tidak erorr
}

// dalam menggunakan defer wajib di posisi atas, tidak boleh dipaling bawah.
