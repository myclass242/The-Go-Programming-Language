package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %GC\n", freezingF, ftoc(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, ftoc(boilingF))

	fmt.Printf("%v\n", f() == f())

	v := 1
	incr(&v)
	fmt.Printf("%d\n", incr(&v))
	//f, err := os.Open("a.txt")
	//f, err  = os.Create("a.txt")
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 /9
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	(*p)++
	return *p
}
