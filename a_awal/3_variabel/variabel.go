package main

import "fmt"

func main() {
	// var namaVariabel tipeData
	var nama string
	nama = "Putra Dae"
	fmt.Println(nama)

	nama = "Budi Dubi"
	fmt.Printf("%s\n\n", nama)

	// var namaVariabel = nilai
	var angkaGenap = 2
	fmt.Println(angkaGenap)
	
	angkaGenap = 8
	fmt.Printf("%d\n\n", angkaGenap)

	// namaVariabel := nilai
	affahIya := true
	fmt.Println(affahIya)
	
	affahIya = false
	fmt.Printf("%t", affahIya)

	
	/* NOTE: Format Specifier
	- %d ==> Integer
	- %s ==> String
	- %t ==> Boolean
	- %f ==> Floating Point
	- %v ==> Versatile (Any Type)
	*/
}