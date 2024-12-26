package main

import "fmt"

func fibonacci(a int, b int, n int) {
	if n == 0 {
		return
	}
	fmt.Println(a)
	fibonacci(b, a+b, n-1)
}

func main() {
	fibonacci(0, 1, 23)
}
