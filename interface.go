// interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung

package main

import "fmt"

type HasName interface {
	getName()string
}

func sayHello(hasName HasName)  {
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string  {
	return person.Name
}

type Animal struct{
	Name string
}

func (animal Animal) getName() string  {
	return animal.Name
}

func main()  {
	var Roihan Person
	Roihan.Name = "Roihan"

	cat := Animal{
		Name: "Push",
	}

	sayHello(cat)
	sayHello(Roihan)
}