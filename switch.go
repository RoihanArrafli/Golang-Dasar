package main

import "fmt"

func main() {
	name := "Roihan"

	switch name {
	case "Roihan":
		fmt.Println("Hello " + name)
	case "Budi": 
		fmt.Println("Hello " + name)
	default:
		fmt.Println("Hi, boleh kenalan?")
	}

	// switch dengan short statement
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama terlalu pendek")
	// }

	// switch tanpa kondisi
	length := len(name)

	switch  {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length < 5:
		fmt.Println("nama terlalu pendek")
	default:
		fmt.Println("nama sudah benar")
	}
}