package main

import "fmt"

func main() {
	i := 42
	f := float64(i)
	var uint uint = uint(f)

	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(uint)
}
