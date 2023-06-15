package main

import "fmt"

func add(p *int) int {
	*p++ //改变指针指向的值
	return *p
}

// 元祖赋值，最大公约数
func gcd(x, y int) int {
	if y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	//i := 0
	//add(&i)
	//fmt.Println(add(&i))
	//fmt.Print(gcd(18, 9))
	defer func() {
		fmt.Println("is all end")
	}()
	fmt.Println("continue")
}
