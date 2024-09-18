package main

import "fmt"

/*
* Задача 2.1
* Значения rune и byte по умолчанию имеют 0
* Так как это синонимы типов int32 и uint8
* А все числовые типы по умолчанию 0
 */

func main() {
	// Задача 2.2
	dividend := 16
	divisor := 3

	quotient := dividend / divisor
	remainder := dividend % divisor

	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %T\n", quotient, remainder, quotient)
}
