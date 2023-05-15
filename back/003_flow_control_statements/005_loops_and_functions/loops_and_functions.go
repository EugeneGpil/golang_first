package main

import "fmt"
import "math"

func Sqrt(x float64) float64 {
	var z float64 = 1

	var zIn2SimilarToXTimes = 0

	var zIn2In10Int int64
	var xIn10Int int64 = int64(x * 10000000000)
	for zIn2SimilarToXTimes <= 10 {
		zIn2In10Int = int64(z * z * 10000000000)
		if xIn10Int == zIn2In10Int {
			zIn2SimilarToXTimes++
		}
		z -= (z*z - x) / (2*z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(4))
	fmt.Println(math.Sqrt(4))
}
