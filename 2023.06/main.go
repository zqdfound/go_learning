package main

import "fmt"

type Human struct {
	Age    int
	Height float64
	Sex    bool
}

func main() {
	var a Human
	a = Human{
		Age:    18,
		Height: 174.1,
		Sex:    false,
	}
	fmt.Println(a)
	fmt.Printf("%d %.2f %t \n", a.Age, a.Height, a.Sex)
	fmt.Printf("%+v \n", a)
}
