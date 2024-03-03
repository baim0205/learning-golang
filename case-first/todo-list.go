package main

import (
	"fmt"
)

// Untuk menampung nilai [] yang sudah di input
// Variable itu tetap di luar fungsi 
var toDoList []string

// Tampilkan Daftar Tugas
func tampilkan() {
	if len(toDoList) == 0 {
		fmt.Println("To-do list is empty")
	} else {
		// i dibawah ini merupakan index awal
		for i, task := range toDoList {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	}
}

// Menambahkan Daftar Tugas
func menambahkan() {
	var task string
	fmt.Print("Masukkan tugas baru: ")
	fmt.Scanln(&task)
	toDoList = append(toDoList, task)
	fmt.Printf("Tugas '%s' telah ditambahkan ke To-Do list.\n", task)
}

// Menghapus Daftar Tugas
func menghapus() {
	tampilkan()
	var taskNumber int
	fmt.Print("Masukkan nomor tugas yang ingin dihapus: ")
	fmt.Scanln(&taskNumber)

	if taskNumber >= 1 && taskNumber <= len(toDoList) {
		task := toDoList[taskNumber-1]
		toDoList = append(toDoList[:taskNumber-1], toDoList[taskNumber:]...)
		fmt.Printf("Tugas '%s' telah dihapus dari To-Do list.\n", task)
	} else {
		fmt.Println("Nomor tugas tidak valid.")
	}
}

func main() {
	var choice string

	for {
		fmt.Println("\nAplikasi To-Do List")
		fmt.Println("1. Tampilkan To-Do List")
		fmt.Println("2. Tambah Tugas Baru")
		fmt.Println("3. Hapus Tugas")
		fmt.Println("4. Keluar")

		fmt.Print("Pilih menu (1/2/3/4): ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			tampilkan()
		case "2":
			menambahkan()
		case "3":
			menghapus()
		case "4":
			fmt.Println("Terima kasih! Sampai jumpa.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}


func authMiddleware(next http.Handler) http.Handler { return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  // Lakukan logika otentikasi di sini // Contoh sederhana: cek apakah ada header Authorization authToken := r.Header.Get("Authorization") if authToken != "token_yang_valid" { http.Error(w, "Unauthorized", http.StatusUnauthorized) return } // Jika otentikasi berhasil, lanjutkan ke handler selanjutnya next.ServeHTTP(w, r) }) }
