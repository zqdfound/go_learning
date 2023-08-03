package main

import (
	"fmt"
	"g6/activationlock/bypassUtils"
)

func main() {
	//key, _ := hex.DecodeString("cb84798c3ca85a674194550a2e96aed8")
	for i := 0; i < 30; i++ {
		code, err := bypassUtils.Create(nil)
		if err == nil {
			hash := code.Hash()
			fmt.Print("bypassMap.put(\"")
			fmt.Print(code)
			fmt.Print("\"")
			fmt.Print(",")
			fmt.Print("\"")
			fmt.Print(hash)
			fmt.Print("\");")
			fmt.Println()
		}
	}

}
