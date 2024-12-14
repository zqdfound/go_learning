package main

import (
	"flag"
	"fmt"
	"testing"
)

func Test_Flag(t *testing.T) {
	// 定义命令参数
	// flag.xxx序列函数，xxx指的是数据类型，例如String、Int、Bool

	// 定义个字符串参数，参数名是name, 参数默认值是tizi, 参数提示信息是：姓名
	namePtr := flag.String("name", "tizi", "姓名")
	// 定义个int参数，参数名：age, 默认值：18，提示信息是：年龄
	agePtr := flag.Int("age", 18, "年龄")
	vipPtr := flag.Bool("vip", true, "是否vip")

	// 自定义命令行 -h 参数，输出什么内容。
	flag.Usage = func() {
		// -h 参数通常用于输出帮助信息，命令有什么参数，怎么用之类的信息
		// 先打印下怎么使用命令
		fmt.Println("Usage: flagdemo [-h] [-name string] [-vip] [-age int]")
		// 调用PrintDefaults打印前面定义的参数列表。
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	// 因为前面flag.xxx函数返回的是指针，所以需要在变量前面加上*号读取参数内容
	fmt.Println("name:", *namePtr)
	fmt.Println("age:", *agePtr)
	fmt.Println("vpi:", *vipPtr)
}
