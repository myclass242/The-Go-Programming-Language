package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Println(input.Err())
	fmt.Println(counts)

	fmt.Println("Prints the text of each line that appears more than once in the standard input")
	for line, num := range counts {
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, line)
		}
	}
}
