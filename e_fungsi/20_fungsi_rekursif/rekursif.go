package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1 // base case	
	}
	
	return n * faktorial(n - 1) // rekursif
}

func main() {
	fmt.Println("Faktorial dari 5 adalah:", faktorial(5))
}