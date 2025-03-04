package main

import (
	"fmt"
	"log"
)

func main() {
	a := 1
	do(a)
}

func do(v any) {
	a := 10

	v1, ok := v.(int)
	if !ok {
		log.Fatalln("неудалось декларировать тип")
	}
	a += v1

	fmt.Println(a)
}
