package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) { //Ini adalah Pointer di Function
	address.Country = "Indonesia"
}

func main() {
	//ini adalah deklarasi variabel singkat
	/**address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 */
	// (&) adalah pointer jadi address 2 adalah pointer ke address1

	// *address adalah pointer(*)
	//contoh deklarasi variabel panjang
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	//var address4 *Address = &Address{"Subang", "Jawa Barat", "Indonesia"} // ini adalah contoh pointer
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	// cara kerja operator adalah menarik address1 ke address2 (operator*)
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // ini adalah operator*

	fmt.Println(address1) // address1 tidak berubah jika address2 tidak memakai pointer(&)
	fmt.Println(address2) // address 1 akan berubah dan mengikuti address2 jika address2 memakai pointer(&)
	fmt.Println(address3)

	//ini adalah contoh untuk function New
	var address4 *Address = new(Address)
	//ini adlaah contoh jika ingin menambahkan city.
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{ //ini adalah Pointer di Function
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	var alamatPointer *Address = &alamat //Pointer di Function
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)

}
