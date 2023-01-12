// defer function = function yang bisa kita jadwalkan untuk di eksekusi setelah sebuah function
// selesai di eksekusi
// defer function akan tetap dieksekusi walaupun terjadi error di function yang dieksekusi

package main

import "fmt"

func logging()  {
	fmt.Println("selesai memanggil function")
}

func runApllication(value int)  {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("result", result)
}

func main() {
	runApllication(0)
}