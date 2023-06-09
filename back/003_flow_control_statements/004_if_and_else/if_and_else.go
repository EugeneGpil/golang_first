package main

import "fmt"
import "math"

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {

		// variable v here is available
		
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// variable v here is already undefined

	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
