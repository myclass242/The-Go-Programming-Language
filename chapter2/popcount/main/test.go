package main

import (
	"fmt"
	"learngo/chapter2/popcount"
	"time"
)

func main() {
	var loopNumber uint64 = 9999999999
	start := time.Now()
	for i := uint64(0); i < loopNumber; i++ {
		popcount.PopCount(i)
	}
	fmt.Println(time.Since(start).Nanoseconds())

	start = time.Now()
	for i := uint64(0); i < loopNumber; i++ {
		popcount.PopCount_2_3(i)
	}
	fmt.Println(time.Since(start).Nanoseconds())

	start = time.Now()
	for i := uint64(0); i < loopNumber; i++ {
		popcount.PopCount_2_4(i)
	}
	fmt.Println(time.Since(start).Nanoseconds())

	start = time.Now()
	for i := uint64(0); i < loopNumber; i++ {
		popcount.PopCount_2_5(i)
	}
	fmt.Println(time.Since(start).Nanoseconds())
}
