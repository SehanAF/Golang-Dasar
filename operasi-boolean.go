package main

import "fmt"

func main() {

	var ujian = 80
	var absensi = 75

	// Versi panjang (step by step)
	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	// Versi singkat
	fmt.Println(ujian >= 80 && absensi >= 80)
}
