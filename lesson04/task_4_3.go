package main

import "fmt"

func main() {
	hello1(func() {
		fmt.Println("Hello, Go!")
	})
}

func hello1(f func()) {
	f()
}
