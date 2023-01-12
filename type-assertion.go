package main

import "fmt"

func random() interface{} {
	return 100
}

func main() {
	var result interface{} = random()
	// var resultString = result.(bool)
	// fmt.Println(resultString)
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown type")
	}
}