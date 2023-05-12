package main

import (
	"fmt"
	"math/cmplx"
)

var toBe bool = false
var maxInt uint64 = 1<<64 -1
var complex complex128 = cmplx.Sqrt(-5 + 12i)

func main() {
	stringFromat := "Type: %T Value: %v\n"
	fmt.Printf(stringFromat, toBe, toBe)
	fmt.Printf(stringFromat, maxInt, maxInt)
	fmt.Printf(stringFromat, complex, complex)
}
