package main

import "fmt"

func luasPersegi(sisi int) int {		
	return sisi * sisi					// Hanya mengembalikan satu nilai hasil dari sisi * sisi
}

func main() {
	hasil := luasPersegi(5)
	fmt.Println("Luas persegi adalah:", hasil)
}