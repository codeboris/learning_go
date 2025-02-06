package main

import "fmt"

type contract struct {
	ID     int
	Number string
	Date   string
}
type contactInfo struct {
	Address string
	Phone   string
}

type user struct {
	ID   int
	Name string
	contactInfo
}

type employee struct {
	ID   int
	Name string
	contactInfo
}

func main() {
	fmt.Println("task 6.1:")
	cont := contract{
		ID:     1,
		Number: "#000A\n101",
		Date:   "2024-01-31",
	}
	fmt.Printf("%#v", cont)

	fmt.Println("task 6.2:")
	cont2 := contract{
		ID:     1,
		Number: "#000A101\t01",
		Date:   "2024-01-31",
	}
	fmt.Printf("{ID:%d Number:%s Date:%s}\n", cont2.ID, cont2.Number, cont2.Date)

	fmt.Println("task 6.3:")
	fmt.Printf("Договор № %#v от %s\n", cont.Number, cont.Date)

	fmt.Println("task 6.4:")
	u := user{
		ID:   1,
		Name: "Иван",
		contactInfo: contactInfo{
			Address: "ул. Ленина, 1",
			Phone:   "+7 900 123-45-67",
		},
	}

	e := employee{
		ID:   2,
		Name: "Петр",
		contactInfo: contactInfo{
			Address: "ул. Гагарина, 5",
			Phone:   "+7 901 987-65-43",
		},
	}

	fmt.Println(u.Address, u.Phone, e.Address, e.Phone)
}
