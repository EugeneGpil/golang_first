package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Negative values is not supported, %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var z float64 = 1

	var zIn2SimilarToXTimes = 0

	var zIn2In10Int int64
	var xIn10Int int64 = int64(x * 10000000000)
	for zIn2SimilarToXTimes <= 10 {
		zIn2In10Int = int64(z * z * 10000000000)
		if xIn10Int == zIn2In10Int {
			zIn2SimilarToXTimes++
		}
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
