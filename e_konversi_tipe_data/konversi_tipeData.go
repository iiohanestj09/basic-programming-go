package main

import "fmt"

func main() {
	// Integer 
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)   // Overflow, nilai akan melintir balik dari batas tipe datanya.

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// Int dan Float
	bilDesimal := 2.7
	bilBulat := int(bilDesimal)

	fmt.Println(bilDesimal)
	fmt.Println(bilBulat)

	// String dan Char
	name := "Putra"
	t := name[2]
	eString := string(t)

	fmt.Println(name)
	fmt.Println(t)
	fmt.Println(eString)
}