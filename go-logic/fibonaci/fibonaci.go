package main

import "fmt"

// fungsi untuk menghasilkan deret Fibonacci hingga i elemen
func fibo(i int) []int {
	// Membuat slice untuk menyimpan deret Fibonacci dengan panjang i
	fib := make([]int, i) // make itu ganti dari dictionary

	// Inisialisasi dua angka pertama dalam deret Fibonacci
	fib[0], fib[1] = 0, 1

	// Mengisi slice dengan deret Fibonacci
	for f := 2; f < i; f++ {
		fmt.Print(f)
		fib[f] = fib[f-1] + fib[f-2]
	}

	return fib // Mengembalikan slice yang berisi deret Fibonacci
}

func main() {
	var i int

	// Meminta pengguna untuk memasukkan jumlah elemen deret Fibonacci yang diinginkan
	fmt.Print("Masukkan Jumlah Elemen Fibonacci : ")
	fmt.Scanln(&i)

	// Memanggil fungsi fibo untuk menghasilkan deret Fibonacci hingga i elemen
	result := fibo(i)

	// Mencetak hasil deret Fibonacci
	fmt.Println("Deret Fibonacci:")
	fmt.Println(result)
}
