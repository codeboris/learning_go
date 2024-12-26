package main

import "fmt"

func hello2() func() {
	return func() {
		fmt.Println("Hello, Go!")
	}
}

func main() {
	anonFunc := hello2()
	anonFunc()
}
