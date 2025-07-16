package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("Perulangan ke-%d\n", i + 1)
	}

	counter := 1
	for counter <= 5 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	nama := "Putra"
	for index, char := range nama {
		fmt.Printf("Index %d: %c\n", index, char)
	}

}
