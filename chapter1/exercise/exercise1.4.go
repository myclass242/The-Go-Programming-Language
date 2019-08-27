package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range(files) {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "os.Open: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, filenameCount := range counts {
		for filename, n := range filenameCount {
			if n > 1 {
				fmt.Printf("file:%s line:%s count:%d\n", filename, line, n)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f.Name()]++
	}
}
