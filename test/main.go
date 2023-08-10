package main

import "fmt"

func change(a int) {
	a = 2
}
func change1(a *int) {
	*a = 2
}
func main() {
	i := 1
	change(i)
	fmt.Println(i)
	change1(&i)
	fmt.Println(i)
}
