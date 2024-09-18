package main

import "fmt"

const newMyConst = 42

func main() {
	const newMyConst = 43
	fmt.Println(newMyConst)
}
