package main

import "fmt"

func main() {

	// harus menyebutkan tipe data
	var name string
	name = "Roihan"
	fmt.Println(name)
	name = "Arrafli"
	fmt.Println(name)

	// langsung mendeklarasikan isi variablenya
	var friendName = "Budi"
	fmt.Println(friendName)
	var age = 17
	fmt.Println(age)

	// langsung diinisialisasi variablenya
	country := "Indonesia"
	fmt.Println(country)

	// multiple variable
	var (
		firstName = "Roi"
		lastName = "Han"
	)
	fmt.Print(firstName)
	fmt.Print(lastName)
}