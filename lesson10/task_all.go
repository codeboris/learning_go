package main

import (
	"fmt"
	//v1 "github.com/codeboris/learn_go_mod"
	//v2 "github.com/codeboris/learn_go_mod/v1.1.0"
	_ "github.com/codeboris/learning_go/lesson08/first"
	"github.com/codeboris/learning_go/lesson10/first"
	"github.com/codeboris/learning_go/lesson10/second"
)

func main() {
	fmt.Println(first.Hello())
	fmt.Println(second.Hello())
	//v1.Hello()
	//v2.Hello()
}
