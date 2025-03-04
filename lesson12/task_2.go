package main

import (
	"fmt"
)

type Bird interface {
	Fly()
}

type Duck struct{}
type Sparrow struct{}

func (d Duck) Fly() {
	fmt.Println("Утка - Умею летать!")
}

func (d Duck) Swim() {
	fmt.Println("Утка - Умею плавать!")
}

func (s Sparrow) Fly() {
	fmt.Println("Воробей - Умею летать!")
}

func main() {
	var d, s Bird
	d = Duck{}
	Do(d)
	s = Sparrow{}
	Do(s)
}

func Do(b Bird) {
	b.Fly()

	if d, ok := b.(Duck); ok {
		d.Swim()
	}
}
