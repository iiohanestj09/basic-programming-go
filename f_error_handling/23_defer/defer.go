package main

import "fmt"

func study() {
	fmt.Println("Semangat Belajar Golang!")
}

func main() {
	defer study() 
	div := 0
	result := 3 / div
	fmt.Println(result) 
}