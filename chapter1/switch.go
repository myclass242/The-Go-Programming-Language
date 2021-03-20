package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	x, _ := strconv.Atoi(os.Args[1])

	switch {
	case x > 0:
		{
			fmt.Printf("x > 0")
		}
	default:
		{
			fmt.Printf("x = 0")
		}
	case x < 0:
		{
			fmt.Printf("x < 0")
		}
	}
}
