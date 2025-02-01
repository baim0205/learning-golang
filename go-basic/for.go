package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Question struct {
	Options []string `json:"options"`
}

type Quiz struct {
	Science map[string]Question `json:"science"`
}

type Data struct {
	Quiz Quiz `json:"quiz"`
}

func getFileJSON(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: File '%s' tidak ditemukan.\n", filename)
		return nil
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error: Tidak dapat membaca file.")
		return nil
	}

	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		fmt.Println("Error: File JSON tidak valid.")
		return nil
	}

	if len(data.Quiz.Science) == 0 {
		fmt.Println("Error: Kategori 'science' tidak ditemukan dalam file JSON.")
		return nil
	}

	var optionList [][]string
	for _, question := range data.Quiz.Science {
		optionList = append(optionList, question.Options)
	}

	return optionList
}

func main() {
	filename := "data.json"
	options := getFileJSON(filename)

	if options != nil {
		for idx, option := range options {
			fmt.Println("=-=-=-=-=")
			fmt.Printf("Question %d Options: %v\n", idx+1, option)
			fmt.Println("=-=-=-=-=")
		}
	} else {
		fmt.Println("Tidak ada data yang bisa ditampilkan.")
	}
}
