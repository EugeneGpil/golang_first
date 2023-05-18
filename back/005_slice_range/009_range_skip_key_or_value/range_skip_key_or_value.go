package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for key := range pow {
		pow[key] = 1 << uint(key) // == 2**key
	}

	for _, value := range pow {
		fmt.Println(value)
	}

	/*
	1
	2
	4
	8
	16
	32
	64
	128
	256
	512
	*/
}
