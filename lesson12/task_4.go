package main

import (
	"errors"
	"fmt"
)

type BirdNext interface {
	Sing() string
}

type DuckNext struct {
	voice string
}

func (d DuckNext) Sing() string {
	return d.voice
}

// В коде нет паники, но есть ошибочная логика с проверкой на nil, примеры решения коментированы в коде
func main() {
	var d DuckNext // 1 - var d *DuckNext - можно передать указатель, он будет nil и тогда сравнение пройдет успешно
	// 2 - d := DuckNext{voice: "Quack"} // Инициализированная структура тоже поможет избежать

	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b BirdNext) (string, error) {
	// 3 - Здесь можно проверить динамеский тип duck, ok := b.(DuckNext) вместо проверки на nil
	if b != nil { // Здесь возникает паника так как b это пустая структура типа DuckNext она не может сравниваться с nil
		return b.Sing(), nil
	}
	return "", errors.New("ошибка пения")
}
