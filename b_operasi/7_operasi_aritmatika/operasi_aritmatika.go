package main

import "fmt"

func main() {
	/* Operator Aritmatika di Go:
	- Penjumlaan (+)
	- Pengurangan (-)
	- Perkalian (*)
	- Pembagian (/)
	- Sisa Bagi (%)
	*/
	var a = 12
	var b = 2
	result1 := a + b
	result2 := 10 - 12
	result3 := 5.5 / 10

	fmt.Printf("%d %d %.2f\n", result1, result2, result3)

	/* Augmented Assignment
	- a = a + 10 --> a += 10
	- a = a - 10 --> a -= 10
	- a = a * 10 --> a *= 10
	- a = a / 10 --> a /= 10
	- a = a % 10 --> a %= 10
	*/

	var result4 = 1
	result4 *= 5

	fmt.Println(result4)

	/* Increment dan Decrement
	- Increment: a = a + 1 --> a++
	- Decrement: b = b - 1 --> b--
	*/

	var result5 = 0
	result5++
	result5++
	result5++
	
	fmt.Println(result5)
}