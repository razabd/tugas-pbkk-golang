package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("Masukkan angka: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}

	switch input {
	case 42:
		fmt.Println("hello universe")
	default:
		fmt.Println(input)
	}
}
