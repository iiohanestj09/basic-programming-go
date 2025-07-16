package main

import "fmt"

func main() {

	// Number --> Integer dan Floating Point
	fmt.Println("Satu =", 1)
	fmt.Println("Negatif Dua =", -2)
	fmt.Println("Tiga Koma Empat =", 3.4)

	// Boolean
	fmt.Println("Benar =", true)
	fmt.Println("Salah =", false)

	// String
	fmt.Println("Putra")
	fmt.Println("Putra Daeee")

	fmt.Println(len("Eko"))
	fmt.Println("Eko Kurniawan"[0])
	fmt.Println("Eland'orr"[4])
	
	fmt.Println(string("Eko Kurniawan"[0]))
	fmt.Println(string("Eland'orr"[4]))
}