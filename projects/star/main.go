package main

import (
	"fmt"
	"strings"
)

func addAsterisks(s string) string {
	// Преобразуем строку в срез символов, а затем соединяем их через '*'
	return strings.Join(strings.Split(s, ""), "*")
}

func main() {
	// Ввод строки
	var inputStr string
	fmt.Scanln(&inputStr)

	// Получаем результат
	result := addAsterisks(inputStr)

	// Выводим результат
	fmt.Println(result)
}
