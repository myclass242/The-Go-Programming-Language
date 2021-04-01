package popcount

var pc [256]byte

func init() {
	// for i := range pc {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount return the population counnt (number of set bits) of x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//exercise 2.3
func PopCount_2_3(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

//exercise 2.4
func PopCount_2_4(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int((x >> i) & 1)
	}
	return count
}

//exercise 2.5
func PopCount_2_5(x uint64) int {
	for count := 0; ; count++ {
		if x == 0 {
			return count
		}
		x = x & (x - 1)	// clear the rightmost non_zero bit of x
	}
}
