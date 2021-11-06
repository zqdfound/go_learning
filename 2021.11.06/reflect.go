package main

import (
	"fmt"
	"reflect"
)

// func Sprint(x interface{}) string{

// 	type stringer interface{
// 		String() string
// 	}

// 	switch x:= x.(type){
// 	case stringer:
// 		return x.string

// 	case string:
// 		return x

// 	default:
// 		return "???"

// }

func main() {
	t := reflect.TypeOf("3")
	fmt.Print(t)
}
