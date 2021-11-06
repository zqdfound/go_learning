//接口实现测试
package main

import "fmt"

//define a interface
type Personer interface {
	Run()
	Name() string
}

//implement interface
type person struct {
	name string
	age  int
}

func (person) Run() {
	fmt.Println("is Runing")
}

func (p person) Name() string {
	return p.name
}

func main() {
	var p Personer
	fmt.Println(p) // nil

	//实例化结构体并赋值给接口
	p = person{
		"zqdfound",
		18,
	}
	p.Run()
	fmt.Println(p.Name())

	//类型断言
	var p2 person = p.(person)
	fmt.Print(p2.age)
	//断言可以有另一个bool类型的返回值 表示是否成功
	if p2, ok := p.(person); ok {
		fmt.Println(ok)
		fmt.Println(p2.age)
	}

}
