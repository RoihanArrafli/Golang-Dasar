// closure adalah kemampuan sebuah function berinteraksi dengan data data disekitarnya dalam
// scope yang sama
// scope = lingkup kerja si variable atau function

package main

import "fmt"


func main()  {
	counter := 0
	name := "Budi"

	increment := func ()  {
		name = "Arrafli"
		name := "Roihan"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}

	fmt.Println(name)
	increment()
	increment()
	fmt.Println(counter)
}