// tipe data slice adalah potongan dari sebuah aray
// slice mirip array tapi ukuran slice bisa berubah
// tipe data slice ada 3 = pointer = penunjuk data pertamadi array pada slice
// ,length = panjang dari slice
// , capacity = kapasitas dari slice, length tidak boleh lebih dari capacity

package main

import "fmt"

func main()  {
	var months = [...]string {
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "diubah"
	// fmt.Println(slice1)

	// slice1[0] = "mei lagi"
	// fmt.Println(months)

	var slice2 = months[2:4]
	fmt.Println(slice2)

	// append(slice, data) = membuat slice baru dengan menambah data ke posisi terakhir slice ,
	// jika kapasitas sudah penuh, maka akan membuat array baru
	var slice3 = append(slice2, "roi")
	fmt.Println(slice3)
	slice3[1] = "Bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// make([] typedata, length, capacity) = membuat slice baru
	newSlice := make([]string, 2, 5)

	newSlice[0] = "roi"
	newSlice[1] = "han"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}