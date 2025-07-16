package main

import "fmt"

func main() {
	name1 := "Putra"
	name2 := "Budi"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 3
	var value2 = 4
	var value3 = 3
	var value4 = 5

	fmt.Println(value1 > value2)
	fmt.Println(value1 >= value3)
	fmt.Println(value2 < value4)
	fmt.Println(value2 <= value3)
	fmt.Println(value3 == value1)
	fmt.Println(value4 != value2)
}