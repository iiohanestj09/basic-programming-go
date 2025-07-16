package main

import "fmt"

func main() {
	var names[3]string

	names[0] = "Yohanes"
	names[1] = "Putra"
	names[2] = "Dae"

	fmt.Printf("%s %s %s\n", names[0], names[1], names[2])

	var values = [3]int {
		94,
		97,
		80,
	}

	fmt.Println(values)
	fmt.Printf("%d %d %d\n", values[0], values[1], values[2])

	fmt.Println(len(names))
	fmt.Println(len(names[1]))

	fmt.Println(len(values))

	var temp [10]string
	fmt.Println(temp)
	fmt.Println(len(temp))
}