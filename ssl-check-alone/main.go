package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func checkSSLCertificate(URL string) {
	// Buat klien HTTP baru dengan transport TLS
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Kirim permintaan HTTP ke URL yang ditentukan
	resp, err := client.Get(URL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Dapatkan sertifikat dari koneksi TLS
	cert := resp.TLS.PeerCertificates[0]

	// Periksa tanggal kedaluwarsa sertifikat
	expirationDate := cert.NotAfter
	currentTime := time.Now()

	// Menghitung sisa hari atau menampilkan tanggal kedaluwarsa
	remainingDays := expirationDate.Sub(currentTime).Hours() / 24
	if expirationDate.Before(currentTime) {
		fmt.Println("Sertifikat SSL telah kedaluwarsa!")
	} else {
		fmt.Printf("Sertifikat SSL masih berlaku. Tanggal kedaluwarsa: %s\n", expirationDate.Format("2006-01-02"))
		fmt.Printf("Sisa hari: %.0f\n", remainingDays)
	}
}

func main() {
	// Panggil fungsi checkSSLCertificate dengan URL situs web yang ingin diperiksa
	URL := "https://nicepay.com"
	checkSSLCertificate(URL)
}
