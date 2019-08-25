package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[:], " "))
	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(os.Args)
}
