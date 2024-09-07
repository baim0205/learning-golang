/* ======
Nama : Mohammad Nurrohim
Test : Logic Sederhana

Buatlah :
1. Form Pilihan untuk memilih (a. square, b. pyramid, c. pyramid setengah kiri, d. pyramid setengan kanan)
2. Setelah memilih bentuk, buatkan form masukan nilai nya

Flow :
Munculin Pilihan Cetak => Pilih yang mau di cetak => Munculkan Input Nilai yang mau di masukan => Logic Looping dan Case

Yang Dapat Dipelajari :
1. Belajar kondisi dengan menggunakan case
2. Belajar metode forloop
3. Belajar Algoritma pyramid

====== */

package main

import "fmt"

func main() {
	var pilihan string
	var rows int
	fmt.Println("====== Selamat Datang ====== \n")
	fmt.Println("a. square")
	fmt.Println("b. pyramid")
	fmt.Println("c. pyramid setengah kiri")
	fmt.Println("d. pyramid setengah kanan\n")
	fmt.Print("Pilihlah bentuk yang ingin di cetak : ")
	fmt.Scan(&pilihan)

	fmt.Print("Masukan Nilai : ")
	fmt.Scan(&rows)

	switch pilihan {
	case "a":
		for i := 1; i <= rows; i++ {
			for j := 1; j <= rows; j++ {
				fmt.Print("* ")
			}
			fmt.Println()
		}
	case "b":
		for i := 1; i <= rows; i++ {
			for j := i; j <= rows; j++ {
				fmt.Print(" ")
			}
			for k := 1; k <= i; k++ {
				fmt.Print("* ")
			}
			fmt.Println()
		}
	case "c":
		for i := 1; i <= rows; i++ {
			for j := i; j <= rows; j++ {
				fmt.Print("* ")
			}
			fmt.Println()
		}
	case "d":
		for i := 1; i <= rows; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print("* ")
			}
			fmt.Println()
		}

	}

}
