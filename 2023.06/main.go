package main

import (
	"fmt"
	"time"
)

type Human struct {
	Age    int
	Height float64
	Sex    bool
}

//func main() {
//	//var a Human
//	a := Human{
//		Age:    18,
//		Height: 174.1,
//		Sex:    false,
//	}
//	fmt.Println(a)
//	fmt.Printf("%d %.2f %t \n", a.Age, a.Height, a.Sex)
//	fmt.Printf("%+v \n", a)
//}

//if 语句
//func main() {
//	//ab只在if 代码块里面生效
//	if a, b := true, 10; !a {
//		fmt.Println(a)
//	} else if b >= 10 {
//		fmt.Println(b)
//	}
//}

// 循环
//func main() {
//	//for i := 0; i < 10; i++ {
//	//	fmt.Println(i)
//	//}
//	i := 0
//	for {
//		i++
//		if i%2 == 0 {
//			continue
//		}
//		fmt.Println(i)
//		if i > 10 {
//			break
//		}
//	}
//}

// 切片
//func main() {
//	//var arr [5]int
//	//arr = [5]int{1, 2, 3, 4, 5}
//	//fmt.Println(arr)
//
//	arr := make([]int, 3, 5)
//	arr[0], arr[1], arr[2] = 2, 7, 9
//	brr := append(arr, 8, 11, 13)
//	fmt.Println(arr)
//	fmt.Println(cap(brr)) //brr 两倍扩容，并且不再与arr共享内存空间
//	//ele内存地址不会变化
//	for i, ele := range arr {
//		//fmt.Println(i, ele)
//		fmt.Printf("%p %p %d %d \n", &ele, &arr[i], ele, arr[i])
//	}
//}

// 字符串
//func main() {
//	s := "abcd"
//	for _, ele := range s {
//		//fmt.Println(ele)
//		fmt.Printf("%c \n", ele)
//	}
//}

// map
//func main() {
//	var m map[string]int
//	m = make(map[string]int, 100)
//	m["a"] = 1
//	m["b"] = 2
//	fmt.Println(m["a"])
//	delete(m, "a")
//	if v, exists := m["a"]; exists {
//		fmt.Println(v)
//	}
//}

// 函数
//func foo(a, b int) (int, error) {
//	if b == 0 {
//		return 0, errors.New("除数为0")
//	}
//	c := a / b
//	return c, nil
//}
//func main() {
//	m, n := 1, 0
//	if p, err := foo(m, n); err == nil {
//		fmt.Println(p)
//	} else {
//		fmt.Printf("发生异常", err)
//	}
//}

// 接口
// 一组行为规范的集合
//type People interface {
//	Say(int, int) int
//}
//type Tom struct {
//}
//
//// 具体实现
//func (Tom) Say(a, b int) int {
//	return a + b
//}
//
//func foo(h People) {
//	c := h.Say(1, 2)
//	fmt.Println(c)
//}
//func main() {
//	var p People
//	p = Tom{}
//	foo(p)
//}

// 时间工具
func main() {
	t0 := time.Now()
	fmt.Println(t0.Unix())
	time.Sleep(50 * time.Millisecond)
	t1 := time.Now()
	diff := t1.Sub(t0)
	fmt.Println(diff.Milliseconds())
}
