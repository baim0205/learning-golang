// Buat Create File hasil input yang akan di masukan ke dalam format jason

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// Membuat struktur untuk data identitas yang ada di json
	type Identitas struct {
		Nama   string `json:"nama"`
		NIK    string `json:"nik"`
		Alamat string `json:"alamat"`
	}

	// Meminta input dari pengguna
	var input_nama string
	var input_nik string
	var input_alamat string

	fmt.Println("Masukan Nama :")
	fmt.Scanln(&input_nama)

	fmt.Println("Masukan NIK :")
	fmt.Scanln(&input_nik)

	fmt.Println("Masukan Alamat")
	fmt.Scanln(&input_alamat)

	// Membuat objek Identitas dari input pengguna (penampung)
	data := Identitas{
		Nama:   input_nama,
		NIK:    input_nik,
		Alamat: input_alamat,
	}

	// Membuat file JSON
	file, err := os.Create("data_identitas.json")
	if err != nil {
		fmt.Println("Error nih!!", err)
		return
	}
	defer file.Close()

	// Mengkodekan data ke dalam format JSON dan menulisnya ke file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Data Berhasil Disimpan di data_identitas.json")
}
