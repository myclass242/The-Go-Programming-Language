package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s file1 file2 ...\n", os.Args[0])
		return
	}

	counts := make(map[string]int)
	for _, arg := range files {
		content, err := os.ReadFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(content), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
