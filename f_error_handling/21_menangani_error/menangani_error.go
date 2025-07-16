package main

import "fmt"

func main() {
	var pembilang, penyebut int = 10, 0
	var hasil int
	var err error

	if penyebut == 0 {
		err = fmt.Errorf("tidak bisa membagi dengan nol")
	} else {
		hasil = pembilang / penyebut
	}

	if err != nil {
		fmt.Println("Terjadi error:", err)
	} else {
		fmt.Println("Hasil pembagian:", hasil)
	}
}