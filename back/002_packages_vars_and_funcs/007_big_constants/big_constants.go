package main

import "fmt"

// Create a huge number by shifting a 1 bit left 100 places.
// In other words, the binary number that is 1 followed by 100 zeroes
const big = 1 << 100

// Shift it right again 99 places, so we end up with 1<<1, or 2.
const small = big >> 99

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))

	// cannot use big (untyped int constant 1267650600228229401496703205376) as int value
	// in argument (overflows)
	// fmt.Println(needInt(big))

	fmt.Println(needFloat(big))
}
