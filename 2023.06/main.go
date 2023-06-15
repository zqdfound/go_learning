package main

import "fmt"

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
func main() {
	//var arr [5]int
	//arr = [5]int{1, 2, 3, 4, 5}
	//fmt.Println(arr)

	arr := make([]int, 3, 5)
	arr[0], arr[1], arr[2] = 2, 7, 9
	brr := append(arr, 8, 11, 13)
	fmt.Println(arr)
	fmt.Println(cap(brr)) //brr 两倍扩容，并且不再与arr共享内存空间

	for i, ele := range arr {
		//fmt.Println(i, ele)
		fmt.Printf("%p %p %d %d \n", &ele, &arr[i], ele, arr[i])
	}
}
