package main

import "fmt"

func main() {
	var a uint = 5 // 0101 dalam biner
	var b uint = 3 // 0011 dalam biner

	fmt.Println("a & b =", a & b) //0101 & 0011 = 0001 → 1 
	fmt.Println("a | b =", a | b) //0101 | 0011 = 0111 → 7
	fmt.Println("a ^ b =", a ^ b) //0101 ^ 0011 = 0110 → 6
	fmt.Println("a &^ b =", a &^ b)//0101 &^ 0011 = 0100 → 4
	fmt.Println("a << 1 =", a << 1) // 0101 << 1 = 1010 → 10
	fmt.Println("a >> 1 =", a >> 1) // 0101 >> 1 = 0010 → 2
}
