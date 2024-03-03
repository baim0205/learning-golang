// Tanggal 3 Maret 2024 Rohim!!
// Slice itu sama seperti Dictionary milik Python!!
/* Perbedaan dengan array adalah saat deklarasi variable, jika jumlah element nya tidak dituliskan maka variable tersebut adalah slice Contoh :
var buah_a = []string{"Belimbing","Ape;"} <-- ini Sliece
var buah_a = [2]string{"Belimbing","Apel;"} <-- ini Aray
var buah_a = [...]string{"Belimbing","Apel;"} <-- ini Aray
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	slice_one()
	two_index_slice()

}

func slice_one() {
	buah := []string{"Belimbing", "Watermelon", "Manggo"}
	fmt.Println(buah[0])
}

func two_index_slice() {
	buah := []string{
		"Belimbing",
		"Watermelon",
		"Manggo",
	}
	newBuah := buah[0:2]
	fmt.Println(newBuah)
}

func json_open_slice() {
	type Data struct {
		Movies []string `json="movies"`
	}
	file, err := os.Open("movie.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}
