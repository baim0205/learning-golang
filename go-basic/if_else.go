package main

import "fmt"

/*
Buatkan form input nilai yang nanti nilai nya akan di gunakan untuk sebuah kondisi
Lalu Buatlah kondisi :
1. Jika Input Nilainya 100 Maka lulus dengan sempurna
2. Jika Input Nilainya Diatas 80 maka Lulus
3. Jika Nilai Nya = 70 - 79 maka belum lulus tapi dapat remedial perbaikan (ini masih belum jelas apakah bisa menggunakan lebih besar dari 70 kurang dari 80)
4. Jika nilai nya di bawah 70 maka sudah pasti tidak lulus
*/

func main() {
	var nilai string
	// fmt.Println(nilai)
	fmt.Print("Masukan Nilai nya : ")
	fmt.Scan(&nilai)
	if nilai == "100" {
		fmt.Println(nilai, "Nilai abang untochable!! Ngeriihhh!!! Sudah Pasti Lulus")
	} else if nilai > "80" {
		fmt.Println(nilai, "Anda Lulus!!!")
	} else if nilai == "70" {
		fmt.Println(nilai, "Abang Nyaris Lulus bang!!! Tapi Abang Dapat Remedia!!!")
	} else {
		fmt.Println(nilai, "Maaf Abangkuh!!! Anda Tidak Lulus Coba Lagi Next Time yah1!!")
	}
}
