package main

import (
	"errors"
	"fmt"
)

type myErr struct {
	status  int
	message string
}

func (e myErr) Error() string {
	return e.message
}

type myFirstError struct {
	msg string
}

func (e *myFirstError) Error() string {
	return e.msg
}

func main() {
	fmt.Println("Задача 11.1:")
	err := fmt.Errorf("%w", errors.New("ошибка1"))
	err = fmt.Errorf("ошибка2:%w", err)
	err = fmt.Errorf("ошибка3:%w", err)
	fmt.Println(err)

	fmt.Println("Задача 11.2:")
	err1 := errors.Unwrap(err)
	fmt.Println(err1)

	fmt.Println("Задача 11.3:")
	errSome1 := fmt.Errorf("%w", myErr{status: 1, message: "ошибка1"})
	errSome2 := fmt.Errorf("%w", myErr{status: 1, message: "ошибка2"})
	errSome3 := fmt.Errorf("%w", myErr{status: 1, message: "ошибка3"})

	err = fmt.Errorf("%w", errSome1)

	fmt.Println("ошибка \"ошибка1\" была:", errors.Is(err, errSome1))
	fmt.Println("ошибка \"ошибка2\" была:", errors.Is(err, errSome2))
	fmt.Println("ошибка \"ошибка3\" была:", errors.Is(err, errSome3))

	fmt.Println("Задача 11.4:")
	some1 := errors.New("ошибка1")
	some2 := fmt.Errorf("ошибка2: %w", some1)
	some3 := fmt.Errorf("ошибка3: %w", some2)

	var myErr *myFirstError
	if errors.As(some3, &myErr) {
		fmt.Println("Обнаружена ошибка типа myFirstError")
	} else {
		fmt.Println("В цепочке ошибок нет myFirstError")
	}

	fmt.Println("Цепочка ошибок:", some3)

}
