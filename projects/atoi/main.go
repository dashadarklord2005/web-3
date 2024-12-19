package main

import (
	"fmt"
	"strconv"
)

func squareDigits(num int) string {
	result := ""
	numStr := strconv.Itoa(num)

	for _, digit := range numStr {
		d, _ := strconv.Atoi(string(digit))
		squared := d * d
		result += strconv.Itoa(squared)
	}

	return result
}

func main() {
	var input int
	fmt.Scan(&input) // Читаем только число
	fmt.Println(squareDigits(input)) // Печатаем только результат
}
