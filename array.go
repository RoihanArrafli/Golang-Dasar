package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Roihan"
	names[1] = "Arrafli"
	names[2] = "Roi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// array langsung
	var values = [3]int{
		100,
		90,
		80,
	}
	// function array
	// len(array) = untuk mendapatkan panjang array
	// array[index] = mendapat data posisi index
	// array[index] = value = mengubah data posisi index

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(names))
	fmt.Println(len(values))
	values[0] = 200
	fmt.Println(values)
}