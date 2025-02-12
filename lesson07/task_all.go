package main

import "fmt"

func main() {
	fmt.Println("task 7.1:")
	var strArray = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(strArray)

	fmt.Println("task 7.2:")
	var strArray2 = [4]string{"яблоко", "груша", "слива", "абрикос"}
	fmt.Println(strArray2)

	fmt.Println("task 7.3:")
	var strArray3 = [4]string{"яблоко", "груша", "помидор", "абрикос"}
	strArray3[2] = "персик"
	fmt.Println(strArray3)

	fmt.Println("task 7.4:")
	var intSlice = []int{5, 2, 8, 3, 1, 9}
	fmt.Println(intSlice)

	fmt.Println("task 7.5:")
	var intSlice2 = make([]int, 10)
	fmt.Println(intSlice2)

	fmt.Println("task 7.6:")
	intSlice3 := make([]int, 10)
	intSlice3 = append(intSlice3, 4, 1, 8, 9)
	fmt.Println(intSlice3)

	fmt.Println("task 7.7:")
	intSlice4 := []int{1, 2, 3}
	intSlice5 := []int{4, 5, 6}
	result := append(intSlice4, intSlice5...)
	fmt.Println(result)

	fmt.Println("task 7.8:")
	intSlice6 := []int{1, 2, 3, 4, 5, 6}
	index := 3
	intSlice6 = append(intSlice6[:index], intSlice6[index+1:]...)

	fmt.Println(intSlice6)
}
