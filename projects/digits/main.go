package main

import (
	"fmt"
)

func findMaxDigit(input string) rune {
	maxDigit := '0'

	for _, digit := range input {
		if digit > maxDigit {
			maxDigit = digit
		}
	}

	return maxDigit
}

func main() {
	var input string
	fmt.Scan(&input) // Чтение строки

	fmt.Println(string(findMaxDigit(input)))
}
