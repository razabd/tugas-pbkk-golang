package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama dengan minimal 3 kata: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Memisahkan kata2nya
	words := strings.Fields(input)

	if len(words) < 3 {
		fmt.Println("Harap masukkan minimal 3 kata.")
		return
	}

	// Reversing wordnya
	for i := 0; i < len(words); i++ {
		// Pisahkan suku kata berdasarkan karakter
		runes := []rune(words[i])
		for j := len(runes) - 1; j >= 0; j-- {
			fmt.Print(string(runes[j]))
		}
		if i < len(words)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
