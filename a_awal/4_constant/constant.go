package main

import "fmt"

func main() {
	// Saat membuat Constan, wajib langsung inisialisasi nilainya
	const kata = "Putra"
	const angka = 1000

	fmt.Println(kata)
	fmt.Println(angka)

	// Error:
	// kata = "Dadang"
	// angka = 123

	
	// Bisa juga Multiline
	const (
		firstName = "Putra"
		lastName = "Dae"
		umur = 20
	)

	fmt.Printf("%s %s\nUmur: %d\n", firstName, lastName, umur)
}