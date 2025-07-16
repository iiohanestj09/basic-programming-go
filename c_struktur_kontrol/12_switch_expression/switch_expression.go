package main

import "fmt"

func main() {
	matkul := "Matematika"

	switch matkul {
		case "Pemweb":
			fmt.Println("Memakai ruangan digedung B11")
		case "Matkom":
			fmt.Println("Memakai ruangan digedung A20")
		case "Strukdat":
			fmt.Println("Memakai ruangan digedung B12")
		case "Basdat", "SO":
			fmt.Println("Menggunakan ruangan digedung A19")
		default:
			fmt.Println("Matakuliah tidak terdaftar")
	}
}