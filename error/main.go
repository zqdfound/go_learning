package main

import (
	"errors"
	"fmt"
)

func main() {
	if r, err := Divide(1, 8); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r)
	}
}

func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("can not divide Zero!!!")
	}
	return a / b, nil
}
