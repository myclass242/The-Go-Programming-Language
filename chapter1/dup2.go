package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fileName := range files {
			f, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, count map[string]int) {
	var content map[string]int = make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
		content[input.Text()]++
	}
	for _, n := range content {
		if n > 1 {
			fmt.Printf("%s has repeat lines\n", f.Name())
			break
		}
	}
}