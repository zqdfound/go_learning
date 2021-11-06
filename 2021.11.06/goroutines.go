//线程
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const u = 20
	fibSum := fib(u)
	fmt.Print(u, fibSum)

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Println(r)
			time.Sleep(delay)
		}
	}
}

//斐波那契数求和
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
