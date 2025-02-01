package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan angka pertama: ")
	input1, _ := reader.ReadString('\n')
	angka1, _ := strconv.ParseFloat(input1[:len(input1)-1], 64)

	fmt.Print("Masukkan angka kedua: ")
	input2, _ := reader.ReadString('\n')
	angka2, _ := strconv.ParseFloat(input2[:len(input2)-1], 64)

	// Melakukan operasi aritmatika
	penjumlahan := angka1 + angka2
	pengurangan := angka1 - angka2
	perkalian := angka1 * angka2
	pembagian := angka1 / angka2

	fmt.Println("Hasil Penjumlahan:", penjumlahan)
	fmt.Println("Hasil Pengurangan:", pengurangan)
	fmt.Println("Hasil Perkalian:", perkalian)
	fmt.Println("Hasil Pembagian:", pembagian)
}
