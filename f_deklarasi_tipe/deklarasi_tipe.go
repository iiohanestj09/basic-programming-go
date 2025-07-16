package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktpPutra NoKTP = "123456789"
	var status Married = false

	fmt.Println(ktpPutra)
	fmt.Println(status)
}