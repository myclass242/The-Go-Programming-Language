package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// files := os.Args[1:]
	files := []string{"echo1.go", "echo2.go", "echo3.go"}
	if len(files) == 0 {
		fmt.Println("need parameters")
		return
	}

	counts := make(map[string][]string)
	for _, fileName := range files {
		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, line := range strings.Split(string(content), "\n") {
			if counts[line] == nil {
				counts[line] = make([]string, 1, 10)
			}
			counts[line] = append(counts[line], fileName)
		}
	}

	for line, fileNames := range counts {
		if len(fileNames) > 1 {
			fmt.Printf("%s 'in' files %v\n", line, fileNames)
		}
	}
	fmt.Println(counts)
}
