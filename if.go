package main

import "fmt"

func main() {
	var name = "abcde"

	if name == "Roihan" {
		fmt.Println("Hello " + name)
	} else if name == "Arrafli" {
		fmt.Println("Hi " + name)
	}else {
		fmt.Println("Hi, boleh kenalan?")
	}
	
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("terlalu pendek")
	}
}