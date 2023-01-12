package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("hello", firstName, lastName)
}

func main()  {
	firstName := "Budi"
	sayHelloTo(firstName, "Joko")
	sayHelloTo("Roihan", "Arrafli")
}