package main

import "fmt"

func main() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for key, value := range pow {
		fmt.Printf("2**%d = %d\n", key, value)

		// this is copy of slice's value
		// can't modify slice here
		value = 0

		// but can this way
		// pow[key] = 0
	}

	/*
	2**0 = 1
	2**1 = 2
	2**2 = 4
	2**3 = 8
	2**4 = 16
	2**5 = 32
	2**6 = 64
	2**7 = 128
	*/

	fmt.Println(pow) // [1 2 4 8 16 32 64 128]
}
