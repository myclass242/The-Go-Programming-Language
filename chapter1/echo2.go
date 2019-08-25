package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	for _, arg := range os.Args[:] {
		//fmt.Println(index, arg)
		s += sep + arg
		sep = " "
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(s)
}
