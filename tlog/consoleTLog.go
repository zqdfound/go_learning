package main

import "fmt"

// 控制台日志
type ConsoleLog struct {
}

func (c ConsoleLog) Log(msg string) {
	fmt.Println("[console]", msg)
}
