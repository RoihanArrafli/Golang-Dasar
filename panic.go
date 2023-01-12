// panic function adalah function yang bisa digunakan untuk menghentikan program
// recover adalah function yang bisa kita gunakan untuk menangkap data panic, dgn recover panic akan berhenti, dan program akan teta berjalan 

package main 

import "fmt"

func endApp()  {
	message := recover()
	if message != nil {
		fmt.Println("aplikasi error", message)
	}
		fmt.Println("Aplikasi Selesai")
}

func runApp(error bool)  {
	defer endApp()
	if error{
		panic("ERROR")
	}
	message := recover()
	fmt.Println("Aplikasi Berjalan", message)
}

func main()  {
	runApp(true)
	fmt.Println("Roihan")
}