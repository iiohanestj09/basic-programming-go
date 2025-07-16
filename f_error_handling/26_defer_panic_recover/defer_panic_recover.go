package main

import "fmt"

func tanganiPanic() {
	if pesan := recover(); pesan != nil {
		fmt.Println("Terjadi kesalahan:", pesan)
	}
}

func bagi(a, b int) {
	defer tanganiPanic() 
	defer fmt.Println("Fungsi selesai dijalankan.") 

	fmt.Printf("Membagi %d dengan %d\n", a, b)
	if b == 0 {
		panic("Tidak bisa dibagi dengan nol!")
	}

	hasil := a / b
	fmt.Println("Hasil:", hasil)
}

func main() {
	bagi(10, 2)
	fmt.Println()

	bagi(5, 0)
	fmt.Println("\nProgram selesai")
}