package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args) == 1 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range os.Args[1:] {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, num := range counts {
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
