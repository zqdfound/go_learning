package main

import (
	"fmt"
	"testing"
)

func Test_print(t *testing.T) {
	a := "escape"
	fmt.Print("go", a)
	fmt.Println("go", a)
	fmt.Printf("go %s", a)
	b := fmt.Sprintf("go %s", a)
	fmt.Println(b)
}
