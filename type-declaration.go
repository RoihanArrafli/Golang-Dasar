package main

import "fmt"

func main() {
	type noKtp string
	type Married bool

	var noEktp noKtp = "331174627864"
	var marriedstatus Married = false

	fmt.Println(noEktp)
	fmt.Println(marriedstatus)
}