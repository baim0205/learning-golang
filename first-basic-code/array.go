// Tanggal 3 Maret 2024 Rohim!!!
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	aray()
	aray_buah()
	aray_multidimensi()
	loop_aray()
	underscore_aray()
	aray_make()
	aray_make_terpisah()
	test_get_json()
}

func aray() {
	var car = [5]string{"Mazda", "Honda", "Toyota", "BMW", "Mercedes"}
	fmt.Println("Ini Array gaya standard", car)
}

func aray_buah() {
	var buah = [4]string{
		"Jeruk",
		"Apel",
		"Peer",
		"Manggo",
		"Watermelon",
	}
	fmt.Println("Ini Aray Gaya Vertical", buah)
}

func aray_multidimensi() {
	//number1 := [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	number2 := [2][3]int{{3, 2, 3}, {3, 4, 5}} // cara singkatnya
	//fmt.Println("Number One is", number1)
	fmt.Println("Number Two is", number2)
}

func loop_aray() {
	room := []string{"Kamar Tidur", "Kamar Mandi", "Studio", "Dapur"}
	for i := 0; i < len(room); i++ {
		fmt.Printf("Elemen %d : %s\n", i, room[i])
	}
}

func underscore_aray() {
	sytle := []string{"Cool", "Hot", "Handsome"}
	for _, sytle := range sytle { // variable i tidak dipakai maka akan error yang mana golang tidak mengijinkan ada variable yang tidak di pakai
		fmt.Printf("Modelnya : %s\n", sytle)
	}
	for i, _ := range sytle { // ini jika hanya ingin mengambil elementnya saja
		fmt.Println(i)
	}
}

func aray_make() {
	movie := make([]string, 2)
	movie[0] = "Doraemon"
	movie[1] = "Dragonball"
	fmt.Println(movie)
}

func aray_make_terpisah() {
	movie := make([]string, 2)
	movie[0] = "Doraemon"
	movie[1] = "Dragonball"
	for _, m := range movie {
		fmt.Println(m)
	}
}

// Testing Using Data Json

func test_get_json() {
	fmt.Println("\n==>INI=TEST=GET=JSON=DATA<==")
	type Data struct { // mendefinisikan struktur bernama Data
		Movies []string `json:"movies"` // untuk menyimpan data json
	}
	// Open File Json
	file, err := os.Open("movie.json")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()
	// Dekode Data Json ke struktur datanya
	var data Data                   // ini adalah deklarasi variable data yang bertype Data yang sudah di definisikan sebelumnya untuk menyimpan file json
	dekode := json.NewDecoder(file) // ini untuk membaca dan mendekode data json (kenapa decode karena sebelumnya tipe json dan di rubah ke basa golang)
	err = dekode.Decode(&data)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Daftar Film")
	for _, movie := range data.Movies {
		fmt.Println(movie)
	}
}
