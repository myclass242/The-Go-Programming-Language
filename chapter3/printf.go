package main

import (
	"fmt"
	"math"
)

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

	var f float32 = 16777216
	fmt.Println(f == f+1)

	const E = 2.71828

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d, e^x = %8.2f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)

	var cx complex128 = complex(1, 2)
	var cy complex128 = complex(3, 4)
	fmt.Println(cx * cy)
	fmt.Println(real(cx * cy))
	fmt.Println(imag(cx * cy))
	fmt.Println(1i * 1i)

	s := "Hello World"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[:5])
	t := s
	s += ", 你好世界"
	fmt.Println(s)
	fmt.Println(t)
	// s[0] = 'L'

	const GoUsage = `Go is a tool for managing Go source code.
	
	Usage:
		Go command [arguments]
	......`
	fmt.Println(GoUsage)
}
