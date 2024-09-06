package main

import "fmt"

// membuat const untuk pengujian pemanggilan di println
const sub, name = "Hellooo", "Abah Rohim!!"

// membuat variable untuk pengujian pemanggilan menggunakan printtf
var nilaiPositif uint8 = 90
var nilaiNegatif = -123456789

func main() {
	fmt.Println("Hellooo Dunia Abah Rohim!!")
	fmt.Println(sub, "Dunia", name)
	// String format %d pada fmt.Printf() digunakan untuk memformat data numerik non-desimal
	fmt.Printf("Bilangan Positif : %d\n", nilaiPositif)
	fmt.Printf("Bilangan Positif : %d\n", nilaiNegatif)
}
