package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulangan ke", counter)
	// 	counter++
	// }

	// for statement
	// init statement = statement sebelum for di eksekusi
	// post statement = statement yang akan selalu di eksekusi di akhir setiap perulangan
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	slice := []string{"Roihan", "Budi", "Joko"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, value := range slice{
		// fmt.Println("index", i, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Roihan"
	person["title"] = "student"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
