package main

import "fmt"

func hello3() {
	defer fmt.Println("Завершение работы!")
	fmt.Println("Hello, Go!")
	return
}

func main() {
	hello3()
}
