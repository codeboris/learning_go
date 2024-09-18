package main

import "fmt"

const (
	myConst1 = iota
	myConst2
	myConst3
	myConst4
	myConst5
)

func main() {
	fmt.Println(myConst1, myConst2, myConst3, myConst4, myConst5)
}
