package main

import (
	"fmt"
)

//func main() {
//	s := gocron.NewScheduler(time.UTC)
//	s.Every(1).Seconds().Do(printTime)
//	s.StartAsync()
//	select {}
//}
//func printTime() {
//	fmt.Println(time.Now())
//}

func fullName(firstName *string, lastName *string) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic has been recovered")
		}
	}()
	if firstName == nil {
		panic("Firsr Name can't be null")
	}
	if lastName == nil {
		panic("Last Name can't be null")
	}

	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

//func main() {
//firName := "paul"
//fullName(&firName, nil)
//fmt.Println("returned normally from main")
//var sb strings.Builder
//sb.WriteString("a")
//var a int
//var b *int
//fmt.Printf("%T", a)
//fmt.Printf("%T", b)
//var num int = 42
//ptr := &num       // ptr 存储 num 的地址
//fmt.Println(*ptr) // 输出 42（解引用，获取 ptr 指向的值）
//fmt.Println(*ptr) // 输出 42（解引用，获取 ptr 指向的值）
//
//*ptr = 100       // 修改 ptr 指向的内存的值
//fmt.Println(num) // 输出 100（num 的值被修改）
//
//var p *int
//var p1 int
//fmt.Println(&p) // 输出 100（num 的值被修改）
//fmt.Println(p1) // 输出 100（num 的值被修改）
//}

// 日志接口
type Logger interface {
	Log(msg string)
}

// 控制台日志
type ConsoleLogger struct {
}
