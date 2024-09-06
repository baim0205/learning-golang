package main

import "fmt"

var a, b float64
var result float64

func main() {
	fmt.Print("Masukan Angka Pertama : ")
	fmt.Scan(&a)
	fmt.Print("Masukan Angka Kedua : ")
	fmt.Scan(&b)

	result = a + b
	fmt.Printf("Hasilnya Adalah : %.2f\n", result)
}
