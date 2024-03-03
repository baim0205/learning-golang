package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

func main() {
	// Ganti alamat situs web dan port yang sesuai
	hostname := "iki-international.tw"
	port := 443

	// Buat alamat URL dengan format "hostname:port"
	address := fmt.Sprintf("%s:%d", hostname, port)

	// Buat waktu tenggat kadaluarsa untuk sertifikat (misalnya, 30 hari dari sekarang)
	expirationThreshold := time.Now().AddDate(0, 0, 30)

	// Buat konfigurasi TLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Nonaktifkan verifikasi sertifikat SSL
	}

	conn, err := tls.Dial("tcp", address, tlsConfig)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer conn.Close()

	// Dapatkan sertifikat dari koneksi TLS
	state := conn.ConnectionState()
	cert := state.PeerCertificates[0]

	// Dapatkan waktu kedaluwarsa dari sertifikat
	expiration := cert.NotAfter

	// Bandingkan waktu kedaluwarsa dengan waktu tenggat kadaluarsa
	if expiration.Before(expirationThreshold) {
		fmt.Printf("Sertifikat SSL untuk %s telah kadaluarsa pada %v\n", hostname, expiration)
	} else {
		fmt.Printf("Sertifikat SSL untuk %s masih berlaku hingga %v\n", hostname, expiration)
	}
}
