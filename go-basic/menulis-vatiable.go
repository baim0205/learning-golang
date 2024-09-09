/*
Menulis variable itu ada 2 tipe :
pertama : disebut manifest typing, tipe ini adalah kita menulis [var nama_variable, tipe_data, dan nilai variablenya]
kedua : disebut dengan type inference yang mana menuslinya secara langsung dan nanti akan di ketahui nilainya setelah di jalankan.
*/

package main

import "fmt"

func main() {

	// manifest typing
	var nama_variable string = "ini string"
	var nama_variable_3 = "ini juga tring"
	fmt.Println(nama_variable)
	fmt.Println(nama_variable_3)

	// inference type
	nama_variable_a := "ini string juga"
	nama_variable_a = "ini juga string tanpa :"
	fmt.Println(nama_variable_a)
	fmt.Println(nama_variable_a)

	// Multi variable inference type
	var angka_satu_a, angka_dua_a, angka_tiga_a string = "satu_A", "dua_B", "tiga_C"
	fmt.Println(angka_satu_a)
	fmt.Println(angka_dua_a)
	fmt.Println(angka_tiga_a)

	// Multi variable inference type
	angka_satu, angka_dua, angka_tiga := "satu", "dua", "tiga"
	fmt.Println(angka_satu)
	fmt.Println(angka_dua)
	fmt.Println(angka_tiga)

	// variable underscore
	_ = "ini testing"
	_ = "testing kedua"
	name, _ := "lala", "candil"
	fmt.Println(name)

	// keyword new
	nameDua := new(string)

	fmt.Println(nameDua)
	fmt.Println(*nameDua)

}
