package main

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
)

func Test_Signal(t *testing.T) {
	// 创建一个信号通道，用于接收信号，通道容量为 1
	c := make(chan os.Signal, 1)

	// 调用Notify注册，通过通道c接收Interrupt和Kill两种信号
	signal.Notify(c, os.Interrupt, os.Kill)

	// 从通道中读取收到的信号，没有信号则阻塞协程
	switch <-c {
	case os.Interrupt:
		fmt.Println("收到：Interrupt信号")
	case os.Kill:
		fmt.Println("收到：Kill信号")
	default:
		fmt.Println("收到其他信号")

	}
}
