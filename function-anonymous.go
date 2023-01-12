package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main()  {
	blacklist := func (name string)bool  {
		return name == "Roihan"
	}

	registerUser("Roihan", blacklist)
	registerUser("Arrafli", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("Roi", func(name string) bool {
		return name == "root"
	})
}