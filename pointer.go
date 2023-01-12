//pointer adalah kemampuan membuat reference kemlokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
//  sederhannaya , dengan kemampuan pointer , kita bisa membuat pass by value

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address)  {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Sukoharjo", "Jawa Tengah", "Indonesia"}
	// var address2 *Address = &address1
	address2 := &address1
	address3 := &address1
	address2.City = "Solo"

	*address2 = Address{"Semarang", "Jawa Tengah", "Indonesia"}
	
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Surakarta"
	fmt.Println(address4)

	var alamat = Address{
		City: "Karanganyar",
		Province: "Jawa Tengah",
		Country: "",
	}
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Print(alamat)
}