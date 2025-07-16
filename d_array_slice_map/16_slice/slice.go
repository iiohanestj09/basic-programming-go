package main

import "fmt"

func main() {
	// Perbedaan Array dan Slice saat Inisialisasi
	var fruits1 = []string{"Apel", "Anggur", "Alpukat"} // ini Slice
	var fruits2 = [3]string{"Bluebery", "Belimbing", "Bengkuang"} // ini Array
	var fruits3 = [...]string{"Ceri", "Cokelat", "Cranberry"} // ini Array

	fmt.Println(fruits1)
	fmt.Println(fruits2)
	fmt.Println(fruits3)

	// Hubungan Array dan Slice
	var months = [...]string {
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(months)

	var slice2 = months[9:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Neptunus") 
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	
	fmt.Println(months)

	// Slice itu Sendiri
	newSlice := make([]string, 2, 5) 	// Tipe data untuk isi Slice, Panjang, Kapasitas
	newSlice[0] = "Putra"
	newSlice[1] = "Dae"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)	// Agar Slice Baru Independen dan tidak terikat dengan Source Copy nya
	fmt.Println(copySlice)
}