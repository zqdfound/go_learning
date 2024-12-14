package main

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestWg(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			//fmt.Printf("打印%d", i)
			fmt.Println("打印", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
