package main

/*
Функции часто возвращает значения типа error, а проверяющему коду следует обрабатывать ошибки проверяя, является ли error равным nil.

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("ошибка конвертации числа: %v\n", err)
    return
}
fmt.Println("Сконвертированное целое число:", i)
error равное nil означает успех; не-nil значение error означает неудачу.
*/

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
