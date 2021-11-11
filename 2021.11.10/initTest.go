package main

import "fmt"

func main() {
	defer fmt.Print("D1")
	defer fmt.Print("D2")

	func1()
	func2()
}
func func1() {
	fmt.Println("f1")
}
func func2() {
	fmt.Println("f2")
}
