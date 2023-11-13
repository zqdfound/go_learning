package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	startTime := time.Now()
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			//Spider(strconv.Itoa(i*25), nil)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(startTime)
	fmt.Printf("waitgroup used time %s \n", elapsed)

}
