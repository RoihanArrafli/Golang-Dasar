package main

import "fmt"

func fullName() (string, string, string) {
	return "Roihan", "Arrafli", "."
}

func main() {
	firstName, middleName,_ := fullName()
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(fullName)
	fmt.Println(firstName, middleName)
}