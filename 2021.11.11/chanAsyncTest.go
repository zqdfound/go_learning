package main

//channel 使用
import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	fmt.Println("len = ", len(c), " cap = ", cap(c))

	go func() {
		defer fmt.Println("sub thread is end")
		for i := 0; i <= 4; i++ {
			c <- i
			fmt.Println("input ", i, " to channel,len = ", len(c), " and cap = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
	fmt.Printf("main End")
}
