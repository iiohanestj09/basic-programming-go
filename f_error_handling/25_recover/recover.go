package main

import "fmt"

func checkSudahBelajar(sudahBelajar bool) {
	if sudahBelajar == false {
		panic("Ups, kamu belum belajar")
	}

	fmt.Println("Mantap, sudah belajar")
}

func selesai() {
	er := recover()

	if er != nil {
		fmt.Println("Ada eror:", er)
	}

	fmt.Println("Selesai")
}

func main() {
	defer selesai()
	checkSudahBelajar(false)
}
