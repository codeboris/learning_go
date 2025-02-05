package main

import "fmt"

type square int

func main() {
	fmt.Println("task 5.1:")
	str := "Hello"
	ptr := &str
	fmt.Println(ptr)

	fmt.Println("task 5.2:")
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Println("task 5.3:")
	var strPtr *string
	v1 := "string"
	strPtr = &v1
	*strPtr = "string changed"
	fmt.Println(v1)

	fmt.Println("task 5.4:")
	var b int
	fmt.Println(b)
	fmt.Println(&b)

	fmt.Println("task 5.5:")
	c := 5
	change(&c)
	fmt.Println(c)

	fmt.Println("task 5.6:")
	var s square = 25
	fmt.Println(s)

	fmt.Println("task 5.7:")
	var q square = 30
	q += 15
	fmt.Println(q)

	fmt.Println("task 5.8:")
	var z square = 34
	z += 10
	fmt.Println(z, "м²")
}

func change(c *int) {
	*c++
}
