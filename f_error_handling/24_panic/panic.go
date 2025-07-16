package main

import "fmt"

func hello() {
	fmt.Println("Hello Sahabat!")
}

func main() {
	defer hello()

	panic("Panic dipanggil")
	fmt.Println("Tidak akan dieksekusi")
}