package main

import (
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("need parameters")
		return
	}

	counts := make(map[string]int)
	for _, fileName := range files {
		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}
}
