package main

import "fmt"

func main() {
	//deklarasi panjang
	// var person map[string]string = map[string]string{
	// deklarasi singkat
	person := map[string]string{
		"name":    "Roihan",
		"address": "Sukoharjo",
	}
	// }

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
		book["title"] = "belajar GO-Lang"
		book["author"] = "Roihan"
		book["ups"] = "salah"
	
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}