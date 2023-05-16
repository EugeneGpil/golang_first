package main

import "fmt"

func main() {
	names := [4]string{"PES1", "PES2", "PES3", "PES4"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
