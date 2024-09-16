package main

import "fmt"

func main() {
	fmt.Println("Моя перваая программа Go!")

	//var hello = len("Привет")
	//word := "hello"
	slovo := "hello"

	//for i := 0; i < len(slovo); i++ {
	//	fmt.Println(slovo[i])
	//}

	//for i, symbol := range slovo {
	//	fmt.Println(i)
	//	fmt.Println(symbol)
	//}

	runes := []rune(slovo)
	bytes := []byte(slovo)

	strFromRunes := string(runes)
	strFromBytes := string(bytes)
	fmt.Println(strFromRunes)
	fmt.Println(strFromBytes)

}
