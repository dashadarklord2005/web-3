package main

import (
	"fmt"
	"math"
)

func main() {
	// Ввод значений катетов a и b
	var a, b float64
	fmt.Scan(&a, &b)

	// Вычисление гипотенузы по теореме Пифагора
	c := math.Sqrt(a*a + b*b)

	// Вывод результата
	fmt.Println(c)
}
