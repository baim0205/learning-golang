package main

import "fmt"

// Kebutuhannya untuk di pakai di beberapa kode program nya
// Jadi dia tidak akan ada error atau di komplain oleh golang
func main() {
	const first_name string = "Mohammad"
	const last_name = "Nurrohim"
	const value = 1000

	fmt.Println(first_name, last_name, value)

	const (
		ini_nama string = "Alif"
		ini_usia        = 3
	)
	fmt.Println(ini_nama, ini_usia)
}
