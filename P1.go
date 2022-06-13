package main

import (
	"fmt"
	"time"
)

//Latihan Project Daftar Film

func Film() {
	var day int
	fmt.Print("Pilih hari: ")
	fmt.Scan(&day)
	fmt.Println("Wait....")
	time.Sleep(1 * time.Second)
	fmt.Println("Tersedia film: ")
	time.Sleep(1 * time.Second)

	switch day {
	case 1:
		fmt.Println("One Piece")
	case 2:
		fmt.Println("Naruto Shippuden")
	case 3:
		fmt.Println("Ultraman")
	case 4:
		fmt.Println("Jujutsu Kaisen")
	case 5:
		fmt.Println("Viking")
	case 6:
		fmt.Println("Peaky Blinders")
	case 7:
		fmt.Println("GOTY")
	}

}

func main() {
	fmt.Println("Daftar Film")
	fmt.Println("---//---")
	time.Sleep(1 * time.Second)

	daftarFilm := []string{
		"1. Senin\n",
		"2. Selasa\n",
		"3. Rabu\n",
		"4. Kamis\n",
		"5. Jumat\n",
		"6. Sabtu\n",
		"7. Minggu\n",
	}
	fmt.Println(daftarFilm)
	fmt.Println("\n")

	Film()
}
