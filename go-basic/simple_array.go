package main

import "fmt"

func main() {
	var nama_hari = [7]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	fmt.Println("ini adalah hari", nama_hari)

	for i, deret := range nama_hari {
		fmt.Println("Hari Ke -", i+1, ":", deret)
	}
}
