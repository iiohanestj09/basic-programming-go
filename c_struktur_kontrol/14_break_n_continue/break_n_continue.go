package main

import "fmt"

func main() {
	for a := 0; a < 10; a++ {
		if a == 6 {
			break
		}

		fmt.Println("a =", a)
	}
	fmt.Println("")

	for b := 0; b < 10; b++ {
		if b % 2 == 0 {
			continue
		}

		fmt.Println("b =", b)
	}

}