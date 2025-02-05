package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hello, Go!")
	}
	hello1(f)
}

func hello1(f func()) {
	f()
}
