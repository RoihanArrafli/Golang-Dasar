package main

import "fmt"

func main() {
	var name1 = "Roihan"
	var name2 = "Arrafli"

	// != tidak sama dengan
	// >= lebih dari sama dengan
	// <= kurang dari sama dengan
	// == sama dengan
	
	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 300

	fmt.Println(value1 < value2)
	fmt.Println(value1 > value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}