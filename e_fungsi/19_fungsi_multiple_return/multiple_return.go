package main

import "fmt"

func bagiDanSisa(a int, b int) (int, int) {
	hasil := a / b
	sisa := a % b	
	return hasil, sisa		// Mengembalikkan Lebih dari satu nilai dengan memisahnya dengan koma
}

func main() {
	hasil, sisa := bagiDanSisa(10, 3)		// Menangkap hasil pengembalian dari fungsi dengan 2 var
	fmt.Println("Hasil pembagian:", hasil)
	fmt.Println("Sisa bagi:", sisa)
}