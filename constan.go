package main

import "fmt"

func main() {
	// const firstName string = "Roihan "
	// const lastName = "Arrafli"
	// const value = 1000

	// multiple variable
	const (
		firstName string 	= "Roihan "
		lastName 			= "Arrafli"
		value 				= 1000
	)

	// error
	// value = 2000;

	fmt.Print(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}