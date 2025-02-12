package main

import "fmt"

func main() {
	fmt.Println("task 8.1:")
	animals := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	fmt.Println(animals)

	fmt.Println("task 8.2:")
	animals2 := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	keys := []string{"слон", "бегемот", "осьминог"}

	for _, key := range keys {
		value, exists := animals2[key]
		fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", key, value, exists)
	}

	fmt.Println("task 8.3:")
	animals3 := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	delete(animals3, "бегемот")

	fmt.Println(animals3)

	fmt.Println("task 8.4:")
	animals4 := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	animals4["выдра"] = struct{}{}

	fmt.Println(animals4)

	fmt.Println("task 8.5:")
	animals5 := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	animals5["бегемот"] = 2

	fmt.Println(animals5)
}
