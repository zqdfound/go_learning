package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// counts := make(map[string]int)
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	counts[input.Text()]++
	// }
	// for k, v := range counts {
	// 	fmt.Print("%s\t%d\n", k, v)
	// }

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for k, v := range counts {
		//if v > 1 {
		fmt.Printf("%s\t%d\n", k, v)
		//}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
