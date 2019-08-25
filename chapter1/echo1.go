// Echo1 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
	fmt.Println(s)
}
