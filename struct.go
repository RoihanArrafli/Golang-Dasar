// struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data
// struct biasanya representasi data dalam program aplikasi yang kita buat
// data di struct disimpan dalam field
// sederhananya struct adalah kumpulan field

package main 

import "fmt"

type Customer struct{
	Name, Address 	string
	Age				int
	Married			bool
}

func (customer Customer) sayHi(name string)  {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main()  {
	var Roihan Customer
	Roihan.Name = "Roihan Arrafli"
	Roihan.Address = "Indonesia"
	Roihan.Age = 17
	Roihan.Married = false

	Roihan.sayHi("Roihan")

	// fmt.Println(Roihan)
	// fmt.Println(Roihan.Name)
	// fmt.Println(Roihan.Address)
	// fmt.Println(Roihan.Age)
	// fmt.Println(Roihan.Married)

	// joko := Customer {
	// 	Name: "Joko",
	// 	Address: "Karanganyar",
	// 	Age: 25,
	// 	Married: true,
	// }
	// fmt.Println(joko)

	// budi := Customer {"Budi K", "Klaten", 30, true}
	// fmt.Println(budi)
}