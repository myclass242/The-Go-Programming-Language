package main

import (
	"fmt"
	"flag"
	"strings"
)


var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
func main() {
	flag.Parse()
	fmt.Printf(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	p := new([0]int)
	q := new(struct{})
	fmt.Printf("%p %v\n", p, q)
	fmt.Printf("%p %v\n", q, q)
	//fmt.Println(p == q)

	//_, ok := q.(int)
	var f1 float64
	var f2 float32
	f1 = 3.1415
	f2 = 3.1416
	f1 = float64(f2)
	fmt.Printf("%f\n", f1)
}