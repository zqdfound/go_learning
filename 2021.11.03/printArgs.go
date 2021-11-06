package main

import (
	"fmt"
	"os"
)

func main() {
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	fmt.Print(len(os.Args))
	s := 0
	for s < len(os.Args) {
		fmt.Println(os.Args[s])
		s++
	}
}
