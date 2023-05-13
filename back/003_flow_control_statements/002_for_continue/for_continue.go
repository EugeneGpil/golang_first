package main

import "fmt"

func main() {
	sum := 1
	// init and post statements are optional
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1
	// and at this point can drop semicolons
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
