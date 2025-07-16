package main

import (
	"errors"
	"fmt"
)

func main() {
	var file string = ""
	var err error


	if file == "" {
		err = fmt.Errorf("Gagal membaca file: %w", errors.New("nama file kosong"))
	} 

	if err != nil {
		fmt.Println("Terjadi error!", err)
	} 
}