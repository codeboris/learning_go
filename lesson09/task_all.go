package main

import "fmt"

func fruitMarket(fruitName string) string {
	fruits := map[string]int{
		"апельсин": 5,
		"яблоки":   3,
		"сливы":    3,
		"груши":    0,
	}
	if count, exists := fruits[fruitName]; exists {
		return fmt.Sprintf("%s - %d", fruitName, count)
	}

	return "Фрукта нет в наличии"
}

func checkFood(food string) {
	switch food {
	case "груша", "яблоко", "апельсин":
		fmt.Println("Это фрукт")
	case "тыква", "огурец", "помидор":
		fmt.Println("Это овощ")
	default:
		fmt.Println("что-тостранное...")
	}
}

func main() {
	fmt.Println("task 9.1:")
	fmt.Println(fruitMarket("апельсин"))

	fmt.Println("task 9.2:")
	slace := []int{1, 2, 3}

OUT:
	for i := 0; i < len(slace); i++ {
		fmt.Printf("v1:%d\n", slace[i])
		for j := 0; j < len(slace); j++ {
			fmt.Printf("\tv2:%d\n", slace[j])
			for k := 0; k < len(slace); k++ {
				fmt.Printf("\t\tv3:%d\n", slace[k])
				for l := 0; l < len(slace); l++ {
					fmt.Printf("\t\t\tv4:%d\n", slace[l])
					if slace[l] == slace[1] {
						continue OUT
					}
				}
			}
		}
	}

	fmt.Println("task 9.3:")
	checkFood("тыква")
	checkFood("груша")
}
