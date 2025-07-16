package main

import "fmt"

func main() {
	hari := "senin"

	if hari == "rabu" {
 		fmt.Println("Menggunakan seragam batik.")
 	} else if hari == "senin" {
		fmt.Println("Menggunakan seragam putih.")
 	} else if hari == "selasa" {
 		fmt.Println("Menggunakan seragam biru.")
 	} else {
 		fmt.Println("Menggunakan seragam bebas.")
 	}
}