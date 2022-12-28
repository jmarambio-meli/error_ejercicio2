package main

import (
	"errors"
	"fmt"
)

type customError struct {
	msg string
}

func (c customError) Error() string {
	return "Error: El salario es menor a 10000"
}

func NewCustom(message string) error {
	return &customError{msg: message}
}

var (
	ErrOther = NewCustom("other error from custom")
)

func main() {
	var salary int = 15000
	err := SendError(salary)
	if err != nil {
		if errors.Is(err, ErrOther) {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Debe pagar Impuesto")
	}
	fmt.Println(errors.Is(err, ErrOther))

}

func SendError(salary int) error {
	if salary < 10000 {
		return ErrOther
	}

	return nil
}
