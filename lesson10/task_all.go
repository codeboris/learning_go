package main

import (
	"fmt"
	_ "github.com/codeboris/learning_go/lesson08/first"
	"github.com/codeboris/learning_go/lesson10/first"
	"github.com/codeboris/learning_go/lesson10/second"
)

func main() {
	fmt.Println(first.Hello())
	fmt.Println(second.Hello())
}
