package main

import "fmt"

func main() {
	nilai := 77
	kkm := 75
	kehadiran := 0.5
	nilaiAbsensi := 0.8

	lulusUjian := nilai >= kkm 			// true
	absensi := kehadiran > nilaiAbsensi // false

	fmt.Println(lulusUjian || absensi) 	// OR --> true OR false = true 
	fmt.Println(lulusUjian && absensi)  // AND --> true AND false = false
	fmt.Println(!absensi)				// NOT --> NOT false = true
}