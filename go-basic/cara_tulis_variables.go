package main

import "fmt"

func main() {
	// Ini Variable dengan deklarasi tipe data
	var name string

	name = "Mohammad Nurrohim"
	fmt.Println(name)

	name = "Khalifah Mohammad El Mutaqien"
	fmt.Println(name)

	// Variable langsung tanpa harus deklarasikan tipe datanya
	var friend_name = "Farah"
	fmt.Println(friend_name)

	var age int8 = 30
	fmt.Println(age)

	// Variable tanpa harus menulis var ini tidak wajib apabila kita sudah menggantinya dengan :=
	my_name := "Nurrohim"
	fmt.Println(my_name)
	my_name = "Mohammad"
	fmt.Println(my_name)
	// jadi kalau dengan := itu langsung tanpa harus var nama_variable

	// Variable metode ini cara untuk membuat variable secara banyak dan mempermudah dalam membacanya
	var (
		first_name = "Mohammad"
		last_name  = "Nurrohim"
	)
	fmt.Println(first_name, last_name)
}
