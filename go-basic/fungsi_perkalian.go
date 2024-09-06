package main

import "fmt"

func perkalian(num_1 int, num_2 int) int {
	return num_1 * num_2
}

func main() {
	num_1 := 1
	num_2 := 2

	// Memanggil fungsi penjumlahan untuk di jalankan!!
	result := perkalian(num_1, num_2)
	fmt.Println("Hasilnya Adalah :", result)

}
